package adapter

import (
	"fmt"
	"log"
	"net"

	"github.com/revandpratama/reflect/auth-service/internal/controller"
	"github.com/revandpratama/reflect/auth-service/pkg/logger"
	pb "github.com/revandpratama/reflect/auth-service/proto/generated/auth"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	pb.UnimplementedAuthServiceServer
}

func StartGRPCServer(srv *controller.AuthController) error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", ENV.GRPCServerPort))
	if err != nil {
		logger.MakeLog(logger.Logger{
			Level:   logger.LEVEL_FATAL,
			Message: fmt.Sprintf("Failed to listen: %v", err),
		})
		return err
	}
	log.Println("2")
	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, srv)
	log.Println("3")

	logger.MakeLog(logger.Logger{
		Level:   logger.LEVEL_INFO,
		Message: "grpc server connected...",
	})

	if err := grpcServer.Serve(listener); err != nil {
		logger.MakeLog(logger.Logger{
			Level:   logger.LEVEL_FATAL,
			Message: fmt.Sprintf("Failed to serve: %v", err),
		})
		return err
	}

	logger.MakeLog(logger.Logger{
		Level:   logger.LEVEL_INFO,
		Message: "grpc server connected...",
	})

	return nil
}
