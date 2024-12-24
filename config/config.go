package config

import (
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

	KafkaHost     string
	KafkaPort     string
	KafkaClientID string
	KafkaTopic    string
}

var ENV *Config

func LoadConfig() error {

	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigName(".env")

	if err := viper.ReadInConfig(); err != nil {

		return err
	}

	if err := viper.Unmarshal(&ENV); err != nil {
		return err
	}

	return nil
}
