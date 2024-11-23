package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"primaryKey;default:uuid_generate_v4()"`
	RoleID    int       `gorm:"column:role_id"`
	Username  string    `gorm:"unique"`
	Name      string
	Email     string `gorm:"unique"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type Role struct {
	ID   int `gorm:"datatype:serial;primaryKey"`
	Name string
}

func (User) TableName() string {
	return "authentication.users"
}

func (Role) TableName() string {
	return "authentication.roles"
}
