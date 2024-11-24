package adapter

import (
	"fmt"

	"github.com/revandpratama/reflect/auth-service/internal/entity"
	"github.com/revandpratama/reflect/auth-service/pkg/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Host:     ENV.DBHost,
		Port:     ENV.DBPort,
		User:     ENV.DBUser,
		Password: ENV.DBPassword,
		DBName:   ENV.DBName,
		SSLMode:  ENV.DBSSLMode,
	}
}

func (c *DatabaseConfig) FormatDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		c.Host, c.User, c.Password, c.DBName, c.Port, c.SSLMode)
}
func ConnectDB() {
	config := NewDatabaseConfig()

	dsn := config.FormatDSN()

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		logger.MakeLog(logger.Logger{
			Level:   logger.LEVEL_FATAL,
			Message: err.Error(),
		})
		panic(err)
	}

	DB = db

	// if err := dropTable(db, "users", "roles"); err != nil {
	// 	logger.MakeLog(logger.Logger{
	// 		Level:   logger.LEVEL_ERROR,
	// 		Message: fmt.Sprintf("migrate error: %v", err.Error()),
	// 	})
	// }

	if err := autoMigrate(db); err != nil {
		logger.MakeLog(logger.Logger{
			Level:   logger.LEVEL_ERROR,
			Message: fmt.Sprintf("migrate error: %v", err.Error()),
		})
	}

	logger.MakeLog(logger.Logger{
		Level:   logger.LEVEL_INFO,
		Message: "database connected...",
	})
}

func autoMigrate(db *gorm.DB) error {

	err := db.AutoMigrate(&entity.User{}, &entity.Role{})

	return err
}

// func dropTable(db *gorm.DB, tables ...string) error {
// 	for _, tablename := range tables {
// 		err := db.Exec(fmt.Sprintf("DROP TABLE authentication.%s", tablename)).Error
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }
