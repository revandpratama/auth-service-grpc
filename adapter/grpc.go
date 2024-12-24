package adapter

import (
	"google.golang.org/grpc"
)

func GRPC() Option {
	return &grpcserver{}
}

type grpcserver struct {
	adapter *Adapter
}

func (g *grpcserver) Start(a *Adapter) error {

	grpcServer := grpc.NewServer()

	a.GRPCServer = grpcServer
	g.adapter = a

	return nil
}

func (g *grpcserver) Stop() error {

	Adapters.GRPCServer.Stop()

	return nil
}

// func StartGRPCServer(srv *controller.AuthController) error {
// 	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", ENV.GRPCServerPort))
// 	if err != nil {
// 		// logger.MakeLog(logger.Logger{
// 		// 	Level:   logger.LEVEL_FATAL,
// 		// 	Message: fmt.Sprintf("Failed to listen: %v", err),
// 		// })
// 		return err
// 	}

// 	grpcServer := grpc.NewServer()

// 	pb.RegisterAuthServiceServer(grpcServer, srv)

// 	// logger.MakeLog(logger.Logger{
// 	// 	Level:   logger.LEVEL_INFO,
// 	// 	Message: "grpc server connected...",
// 	// })

// 	if err := grpcServer.Serve(listener); err != nil {
// 		// logger.MakeLog(logger.Logger{
// 		// 	Level:   logger.LEVEL_FATAL,
// 		// 	Message: fmt.Sprintf("Failed to serve: %v", err),
// 		// })
// 		return err
// 	}

// 	// logger.MakeLog(logger.Logger{
// 	// 	Level:   logger.LEVEL_INFO,
// 	// 	Message: "grpc server connected...",
// 	// })

// 	return nil
// }
