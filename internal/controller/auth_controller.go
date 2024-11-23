package controller

import (
	"context"

	"github.com/revandpratama/reflect/auth-service/internal/dto"
	"github.com/revandpratama/reflect/auth-service/internal/service"
	pb "github.com/revandpratama/reflect/auth-service/proto/generated/auth"
)

type AuthController struct {
	pb.UnimplementedAuthServiceServer
	service service.AuthService
}

func NewAuthController(service service.AuthService) *AuthController {
	return &AuthController{
		service: service,
	}
}

func (c *AuthController) Login(context context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

	request := dto.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	}

	response, err := c.service.Login(context, &request)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		AccessToken: response.AccessToken,
	}, nil
}
