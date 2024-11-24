package server

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/revandpratama/reflect/auth-service/adapter"
	"github.com/revandpratama/reflect/auth-service/client"
	"github.com/revandpratama/reflect/auth-service/internal/controller"
	"github.com/revandpratama/reflect/auth-service/internal/repository"
	"github.com/revandpratama/reflect/auth-service/internal/service"
	"github.com/revandpratama/reflect/auth-service/pkg/logger"
)

type Server struct {
	shutdown chan os.Signal
	done     chan bool
}

func NewServer() *Server {
	return &Server{
		shutdown: make(chan os.Signal, 1),
		done:     make(chan bool, 1),
	}
}

func (s *Server) Start() {
	signal.Notify(s.shutdown, os.Interrupt, syscall.SIGTERM)
	//initialize app
	adapter.LoadConfig()
	adapter.ConnectDB()

	authRepository := repository.NewAuthRepository(adapter.DB)
	authService := service.NewAuthService(authRepository)
	authController := controller.NewAuthController(authService)

	logger.MakeLog(logger.Logger{
		Level:   logger.LEVEL_INFO,
		Message: "starting grpc server...",
	})

	go func() {

		if err := adapter.StartGRPCServer(authController); err != nil {
			log.Printf("Error starting gRPC server: %v", err)
			s.done <- true
		}
	}()

	logger.MakeLog(logger.Logger{
		Level:   logger.LEVEL_INFO,
		Message: "starting grpc server started",
	})

	client.TestClient()

	logger.MakeLog(logger.Logger{
		Level:   logger.LEVEL_INFO,
		Message: "server started",
	})

	select {
	case sig := <-s.shutdown:
		logger.MakeLog(logger.Logger{
			Level:   logger.LEVEL_INFO,
			Message: fmt.Sprintf("shutting down... :%v", sig),
		})
	case <-s.done:
		logger.MakeLog(logger.Logger{
			Level:   logger.LEVEL_FATAL,
			Message: "server stopped due to error",
		})
	}

	s.cleanup()

}

func (s *Server) cleanup() {
	// Add cleanup logic here
	logger.MakeLog(logger.Logger{
		Level:   logger.LEVEL_WARN,
		Message: "cleaning up resources...",
	})

	sqlDB, err := adapter.DB.DB()
	if err != nil {
		logger.MakeLog(logger.Logger{
			Level:   logger.LEVEL_ERROR,
			Message: err.Error(),
		})
	}
	// Close database connection
	if sqlDB != nil {
		if err := sqlDB.Close(); err != nil {
			logger.MakeLog(logger.Logger{
				Level:   logger.LEVEL_ERROR,
				Message: fmt.Sprintf("Error closing database connection: %v", err),
			})

		}
	}

	// Add any other cleanup tasks here

	logger.MakeLog(logger.Logger{
		Level:   logger.LEVEL_INFO,
		Message: "server shutdown complete",
	})
}
