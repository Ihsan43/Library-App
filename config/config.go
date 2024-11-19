package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Configuration structs
type apiConfig struct {
	ApiPort string
}

type dbConfig struct {
	Host     string
	Port     string
	Name     string
	Password string
	User     string
	DbDrive  string
}

type Config struct {
	apiConfig
	dbConfig
}

func (c *Config) readConfig() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	c.apiConfig.ApiPort = os.Getenv("API_PORT")

	c.dbConfig.Host = os.Getenv("DB_HOST")
	c.dbConfig.Port = os.Getenv("DB_PORT")
	c.dbConfig.Name = os.Getenv("DB_NAME")
	c.dbConfig.Password = os.Getenv("DB_PASSWORD")
	c.dbConfig.User = os.Getenv("DB_USER")
	c.dbConfig.DbDrive = os.Getenv("DB_DRIVE")

	requiredVars := map[string]any{
		"API_PORT":    c.apiConfig.ApiPort,
		"DB_HOST":     c.dbConfig.Host,
		"DB_PORT":     c.dbConfig.Port,
		"DB_NAME":     c.dbConfig.Name,
		"DB_PASSWORD": c.dbConfig.Password,
		"DB_USER":     c.dbConfig.User,
		"DB_DRIVE":    c.dbConfig.DbDrive,
	}

	if err := validateEnvVars(requiredVars); err != nil {
		return err
	}

	fmt.Println("All required environment variables are set!")
	return nil
}

func validateEnvVars(envVars map[string]any) error {
	for key, value := range envVars {
		if value == nil || value == "" {
			return fmt.Errorf("missing required environment variable: %s", key)
		}
	}
	return nil
}

func NewConfig() (*Config, error) {
	cfg := Config{}

	if err := cfg.readConfig(); err != nil {
		return nil, err
	}

	return &cfg, nil
}
