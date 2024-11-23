package service

import (
	"context"

	"github.com/revandpratama/reflect/auth-service/internal/dto"
	"github.com/revandpratama/reflect/auth-service/internal/repository"
	"github.com/revandpratama/reflect/auth-service/pkg/auth"
)

type authService struct {
	repository repository.AuthRepository
}

type AuthService interface {
	Login(context context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error)
}

func NewAuthService(repository repository.AuthRepository) *authService {
	return &authService{
		repository: repository,
	}
}

func (s *authService) Login(context context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error) {

	user, err := s.repository.GetUserByUsername(context, req.Username)

	if err != nil {
		return nil, err
	}

	//verify password

	token, err := auth.CreateToken(user)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		AccessToken: token,
	}, nil
}
