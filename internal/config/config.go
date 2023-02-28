package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Config struct {
	Database DatabaseConfig `mapstructure:"db"`
	OAuth    struct {
		Google struct {
			ClientID     string `mapstructure:"client_id"`
			ClientSecret string `mapstructure:"client_secret"`
			RedirectURI  string `mapstructure:"redirect_uri"`
		} `mapstructure:"google"`
	} `mapstructure:"oauth"`
	Server ServerConfig `mapstructure:"server"`
}

type DatabaseConfig struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Name     string `mapstructure:"name"`
}

type OAuthGoogleConfig struct {
	ClientID     string `mapstructure:"client_id"`
	ClientSecret string `mapstructure:"client_secret"`
	RedirectURI  string `mapstructure:"redirect_uri"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
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

func GetOAuthConfig(cfg *Config) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     cfg.OAuth.Google.ClientID,
		ClientSecret: cfg.OAuth.Google.ClientSecret,
		Endpoint:     google.Endpoint,
		RedirectURL:  cfg.OAuth.Google.RedirectURI,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
	}
}
