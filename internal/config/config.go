package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DatabaseURL string
	Env         string
}

func NewConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("dotenv")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("Failed to read config file: %v", err)
	}

	config := &Config{
		DatabaseURL: viper.GetString("DATABASE_URL"),
		Env:         viper.GetString("ENV"),
	}

	return config, nil
}
