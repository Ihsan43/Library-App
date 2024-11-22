package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

// Configuration structs
type ApiConfig struct {
	ApiPort string
}

type DbConfig struct {
	Host     string
	Port     string
	Name     string
	Password string
	User     string
	DbDrive  string
}

type TokenConfig struct {
	ApplicationName     string
	JwtSignatureKey     []byte
	JwtSigningMethod    *jwt.SigningMethodHMAC
	AccessTokenLifeTime time.Duration
}

type Config struct {
	ApiConfig
	DbConfig
	TokenConfig
}

func (c *Config) readConfig() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	c.ApiConfig.ApiPort = os.Getenv("API_PORT")

	c.DbConfig.Host = os.Getenv("DB_HOST")
	c.DbConfig.Port = os.Getenv("DB_PORT")
	c.DbConfig.Name = os.Getenv("DB_NAME")
	c.DbConfig.Password = os.Getenv("DB_PASSWORD")
	c.DbConfig.User = os.Getenv("DB_USER")
	c.DbConfig.DbDrive = os.Getenv("DB_DRIVE")

	appTokenExpire, err := strconv.Atoi(os.Getenv("APP_TOKEN_EXPIRE"))
	if err != nil {
		return err
	}

	accessTokenLifeTime := time.Duration(appTokenExpire) * time.Minute
	
	c.TokenConfig = TokenConfig{
		JwtSigningMethod:    jwt.SigningMethodHS256,
		AccessTokenLifeTime: accessTokenLifeTime,
	}

	c.TokenConfig.ApplicationName = os.Getenv("APP_NAME")
	c.TokenConfig.JwtSignatureKey = []byte(os.Getenv("SIGNATURE_KEY"))

	fmt.Println(c.TokenConfig.JwtSignatureKey)
	fmt.Println(c.TokenConfig.ApplicationName)

	requiredVars := map[string]any{
		"API_PORT":         c.ApiConfig.ApiPort,
		"DB_HOST":          c.DbConfig.Host,
		"DB_PORT":          c.DbConfig.Port,
		"DB_NAME":          c.DbConfig.Name,
		"DB_PASSWORD":      c.DbConfig.Password,
		"DB_USER":          c.DbConfig.User,
		"DB_DRIVE":         c.DbConfig.DbDrive,
		"APP_TOKEN_EXPIRE": c.TokenConfig.AccessTokenLifeTime,
		"APP_NAME":         c.TokenConfig.ApplicationName,
		"SIGNATURE_KEY":    c.TokenConfig.JwtSignatureKey,
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
