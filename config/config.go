package config

import (
	"github.com/spf13/viper"
)

// Config represents the application configuration.
type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	// Add other configuration fields as needed
}

// NewConfig initializes and returns a new Config instance.
func NewConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := &Config{
		DBHost:     viper.GetString("db.host"),
		DBPort:     viper.GetString("db.port"),
		DBUser:     viper.GetString("db.user"),
		DBPassword: viper.GetString("db.password"),
		DBName:     viper.GetString("db.name"),
		// Initialize other configuration fields
	}

	return cfg, nil
}
