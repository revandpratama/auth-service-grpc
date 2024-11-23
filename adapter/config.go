package adapter

import (
	"github.com/revandpratama/reflect/auth-service/pkg/logger"
	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	GRPCServerPort string
}

var ENV *Config

func LoadConfig() {

	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigName(".env")

	if err := viper.ReadInConfig(); err != nil {
		logger.MakeLog(logger.Logger{
			Level:   logger.LEVEL_ERROR,
			Message: err.Error(),
		})
		panic(err)
	}

	if err := viper.Unmarshal(&ENV); err != nil {
		logger.MakeLog(logger.Logger{
			Level:   logger.LEVEL_ERROR,
			Message: err.Error(),
		})
		panic(err)
	}

	logger.MakeLog(logger.Logger{
		Level:   logger.LEVEL_INFO,
		Message: "config running...",
	})

}
