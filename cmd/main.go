package main

import (
	"github.com/revandpratama/reflect/auth-service/cmd/server"
	"github.com/revandpratama/reflect/auth-service/pkg/logger"
)

func main() {

	server := server.NewServer()

	logger.MakeLog(logger.Logger{
		Level: logger.LEVEL_INFO,
		Message: "attempting to start the server...",
	})
	server.Start()
}
