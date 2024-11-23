package repository

import (
	"context"

	"github.com/revandpratama/reflect/auth-service/internal/entity"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

type AuthRepository interface {
	GetUserByUsername(context context.Context, username string) (*entity.User, error)
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{
		db: db,
	}
}


func (r *authRepository) GetUserByUsername(context context.Context, username string) (*entity.User, error) {
	var user entity.User
	err := r.db.WithContext(context).First(&user, "username = ?", username).Error

	return &user, err
}