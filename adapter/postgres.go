package adapter

import (
	"database/sql"
	"fmt"

	"github.com/revandpratama/reflect/auth-service/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type postgresql struct {
	adapter *Adapter

	sql *sql.DB
}

func Postgres() Option {
	return &postgresql{}
}

func (p *postgresql) Start(a *Adapter) error {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.ENV.DBHost, config.ENV.DBUser, config.ENV.DBPassword, config.ENV.DBName, config.ENV.DBPort, config.ENV.DBSSLMode)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		return err
	}

	sql, err := p.adapter.Postgres.DB()
	if err != nil {
		return err
	}

	a.Postgres = db
	p.adapter = a
	p.sql = sql
	return nil
}

func (p *postgresql) Stop() error {

	if err := p.sql.Close(); err != nil {
		return err
	}

	return nil
}

// func (c *PostgresConfig) FormatDSN() string {
// 	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
// 		c.Host, c.User, c.Password, c.DBName, c.Port, c.SSLMode)
// }
// func ConnectDB() error {
// 	config := NewPostgresConfig()

// 	dsn := config.FormatDSN()

// 	db, err := gorm.Open(postgres.New(postgres.Config{
// 		DSN:                  dsn,
// 		PreferSimpleProtocol: true,
// 	}), &gorm.Config{})
// 	if err != nil {
// 		// logger.MakeLog(logger.Logger{
// 		// 	Level:   logger.LEVEL_FATAL,
// 		// 	Message: err.Error(),
// 		// })
// 		// panic(err)
// 		return err
// 	}

// 	DB = db

// 	// if err := dropTable(db, "users", "roles"); err != nil {
// 	// 	logger.MakeLog(logger.Logger{
// 	// 		Level:   logger.LEVEL_ERROR,
// 	// 		Message: fmt.Sprintf("migrate error: %v", err.Error()),
// 	// 	})
// 	// }

// 	if err := autoMigrate(db); err != nil {
// 		// logger.MakeLog(logger.Logger{
// 		// 	Level:   logger.LEVEL_ERROR,
// 		// 	Message: fmt.Sprintf("migrate error: %v", err.Error()),
// 		// })
// 		return err
// 	}

// 	// logger.MakeLog(logger.Logger{
// 	// 	Level:   logger.LEVEL_INFO,
// 	// 	Message: "database connected...",
// 	// })

// 	return nil
// }

// func autoMigrate(db *gorm.DB) error {

// 	err := db.AutoMigrate(&entity.User{}, &entity.Role{})

// 	return err
// }

// func dropTable(db *gorm.DB, tables ...string) error {
// 	for _, tablename := range tables {
// 		err := db.Exec(fmt.Sprintf("DROP TABLE authentication.%s", tablename)).Error
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }
