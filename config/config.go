package config

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"mnc/merchant-bank-api/Utils/common"
	"os"
	"strconv"
	"time"
)

type ApiConfig struct {
	ApiHost string
	ApiPort string
}

type DbConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	Driver   string
}

type Config struct {
	ApiConfig
	DbConfig
	TokenConfig
}

type TokenConfig struct {
	ApplicationName     string
	JwtSignatureKey     []byte
	JwtSigningMethod    *jwt.SigningMethodHMAC
	AccessTokenLifeTime time.Duration
}

func (c *Config) ReadConfig() error {
	err := common.LoadEnv()
	if err != nil {
		return err
	}
	c.DbConfig = DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Driver:   os.Getenv("DB_DRIVER"),
	}

	c.ApiConfig = ApiConfig{
		ApiHost: os.Getenv("API_HOST"),
		ApiPort: os.Getenv("API_PORT"),
	}

	appTokenExpire, err := strconv.Atoi(os.Getenv("APP_TOKEN_EXPIRE"))

	if err != nil {
		return err
	}

	accessTokenLifeTime := time.Duration(appTokenExpire) * time.Minute

	c.TokenConfig = TokenConfig{
		ApplicationName:     os.Getenv("APP_TOKEN_NAME"),
		JwtSignatureKey:     []byte(os.Getenv("APP_TOKEN_KEY")),
		JwtSigningMethod:    jwt.SigningMethodHS256,
		AccessTokenLifeTime: accessTokenLifeTime,
	}

	if c.DbConfig.Host == "" || c.DbConfig.Port == "" || c.DbConfig.Name == "" || c.DbConfig.User == "" || c.DbConfig.Driver == "" || c.ApiConfig.ApiHost == "" || c.ApiConfig.ApiPort == "" {
		return fmt.Errorf("Missing Required environment variables")
	}
	return nil

}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := cfg.ReadConfig()
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
