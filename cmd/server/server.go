package server

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/revandpratama/reflect/auth-service/adapter"
	"github.com/revandpratama/reflect/auth-service/client"
	"github.com/revandpratama/reflect/auth-service/config"
	"github.com/revandpratama/reflect/auth-service/internal/controller"
	"github.com/revandpratama/reflect/auth-service/internal/repository"
	"github.com/revandpratama/reflect/auth-service/internal/service"
	"github.com/revandpratama/reflect/auth-service/pkg/logger"
	pb "github.com/revandpratama/reflect/auth-service/proto/generated/auth"
	"google.golang.org/grpc"
)

type Server struct {
	shutdown     chan os.Signal
	errorOccured chan error
}

func NewServer() *Server {
	return &Server{
		shutdown:     make(chan os.Signal, 1),
		errorOccured: make(chan error, 1),
	}
}

func (s *Server) Start() {
	signal.Notify(s.shutdown, os.Interrupt, syscall.SIGTERM)

	if err := config.LoadConfig(); err != nil {
		logger.MakeLog(logger.Logger{
			Level:   logger.LEVEL_FATAL,
			Message: fmt.Sprintf("failed to initialize config : %v", err),
		})
		s.errorOccured <- err
	}

	logger.MakeLog(logger.Logger{
		Level:   logger.LEVEL_INFO,
		Message: "config running...",
	})

	err := adapter.Adapters.Sync(
		adapter.Postgres(),
		adapter.GRPC(),
	)
	if err != nil {
		logger.MakeLog(logger.Logger{
			Level:   logger.LEVEL_FATAL,
			Message: fmt.Sprintf("failed to start adapter : %v", err),
		})
		s.errorOccured <- err
	}

	authRepository := repository.NewAuthRepository(adapter.DB)
	authService := service.NewAuthService(authRepository)
	authController := controller.NewAuthController(authService)

	logger.MakeLog(logger.Logger{
		Level:   logger.LEVEL_INFO,
		Message: "starting grpc server...",
	})

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", config.ENV.GRPCServerPort))
	if err != nil {
		logger.MakeLog(logger.Logger{
			Level:   logger.LEVEL_FATAL,
			Message: fmt.Sprintf("failed to listen: %v", err),
		})
		s.errorOccured <- err
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, authController)

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			s.errorOccured <- err
		}
	}()

	logger.MakeLog(logger.Logger{
		Level:   logger.LEVEL_INFO,
		Message: "grpc server running...",
	})

	// ! Testing only
	client.TestClient()

	logger.MakeLog(logger.Logger{
		Level:   logger.LEVEL_INFO,
		Message: "server is running",
	})

	select {
	case sig := <-s.shutdown:
		logger.MakeLog(logger.Logger{
			Level:   logger.LEVEL_INFO,
			Message: fmt.Sprintf("shutting down... , cause: %v", sig),
		})
	case err := <-s.errorOccured:
		logger.MakeLog(logger.Logger{
			Level:   logger.LEVEL_FATAL,
			Message: fmt.Sprintf("Server stopped due to error: %v", err),
		})
	}

	s.cleanup()

}

func (s *Server) cleanup() {
	// Add cleanup logic here
	logger.MakeLog(logger.Logger{
		Level:   logger.LEVEL_INFO,
		Message: "cleaning up resources...",
	})

	err := adapter.Adapters.Unsync()
	if err != nil {
		logger.MakeLog(logger.Logger{
			Level:   logger.LEVEL_ERROR,
			Message: fmt.Sprintf("Error cleaning up adapters: %v", err),
		})
	}

	logger.MakeLog(logger.Logger{
		Level:   logger.LEVEL_INFO,
		Message: "server shutdown complete",
	})

}
