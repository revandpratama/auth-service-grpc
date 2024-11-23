package server

import (
	"log"

	"github.com/revandpratama/reflect/auth-service/adapter"
	"github.com/revandpratama/reflect/auth-service/client"
	"github.com/revandpratama/reflect/auth-service/internal/controller"
	"github.com/revandpratama/reflect/auth-service/internal/repository"
	"github.com/revandpratama/reflect/auth-service/internal/service"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start() {

	//initialize app
	adapter.LoadConfig()
	adapter.ConnectDB()
	log.Println("1")
	authRepository := repository.NewAuthRepository(adapter.DB)
	authService := service.NewAuthService(authRepository)
	authController := controller.NewAuthController(authService)
	
	
	adapter.StartGRPCServer(authController)
	
	client.TestClient()
}
