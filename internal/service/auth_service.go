package service

import (
	"context"
	"errors"

	"github.com/revandpratama/reflect/auth-service/internal/dto"
	"github.com/revandpratama/reflect/auth-service/internal/entity"
	"github.com/revandpratama/reflect/auth-service/internal/repository"
	"github.com/revandpratama/reflect/auth-service/pkg/auth"
)

type authService struct {
	repository repository.AuthRepository
}

type AuthService interface {
	Login(context context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error)
	Register(context context.Context, req *dto.RegisterRequest) error
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
	if err := auth.ValidatePassword(user.Password, req.Password); err != nil {
		return nil, err
	}

	token, err := auth.CreateToken(user)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		AccessToken: token,
	}, nil
}

func (s *authService) Register(context context.Context, req *dto.RegisterRequest) error {

	encryptedPassword, err := auth.EncryptPassword(req.Password)
	if err != nil {
		return err
	}

	if s.repository.IsEmailExists(context, req.Email) || s.repository.IsUsernameExists(context, req.Username) {
		return errors.New("email or username already exists")
	}

	newUser := entity.User{
		Name:     req.Name,
		Username: req.Username,
		Email:    req.Email,
		Password: encryptedPassword,
	}

	if err := s.repository.CreateUser(context, &newUser); err != nil {
		return err
	}

	return nil
}
