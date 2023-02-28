package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Database DatabaseConfig `mapstructure:"db"`
}

type DatabaseConfig struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Name     string `mapstructure:"name"`
}

func Load() (*Config, error) {
	// Set the configuration file name and paths
	viper.SetConfigFile("configs/.config.yaml")

	// Load the configuration
	if err := viper.ReadInConfig(); err != nil {
		// If there's an error loading the configuration file, log the error and return the default configuration
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Println(fmt.Errorf("failed to read configuration file: %w", err))
		}
	}

	cfg := Config{}
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal configuration: %w", err)
	}
	return &cfg, nil
}
