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
	CreateUser(context context.Context, newUser *entity.User) error
	IsEmailExists(context context.Context, email string) bool 
	IsUsernameExists(context context.Context, username string) bool
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

func (r *authRepository) CreateUser(context context.Context, newUser *entity.User) error {
	return r.db.WithContext(context).Create(&newUser).Error
}

func (r *authRepository) IsEmailExists(context context.Context, email string) bool {
	var user []entity.User

	r.db.WithContext(context).Find(&user, "email = ?", email)
	
	return len(user) > 0
}
func (r *authRepository) IsUsernameExists(context context.Context, username string) bool {
	var user []entity.User

	r.db.WithContext(context).Find(&user, "username = ?", username)
	
	return len(user) > 0
}