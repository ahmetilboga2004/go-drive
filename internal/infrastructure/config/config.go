package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type appConfig struct {
	Port    string
	Mode    string
	BaseURL string
}

type dbConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
	TimeZone string
	SSLMode  string
}

type jwtConfig struct {
	AccessSecretKey        string
	RefreshSecretKey       string
	AccessTokenExpiration  time.Duration
	RefreshTokenExpiration time.Duration
}

type smtpConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	From     string
}

var (
	APP  *appConfig
	DB   *dbConfig
	JWT  *jwtConfig
	SMTP *smtpConfig
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Error().Err(err).Msg(".env file loading failed")
	}

	APP = &appConfig{
		Port:    getEnv("APP_PORT"),
		Mode:    getEnv("APP_MODE"),
		BaseURL: getEnv("APP_BASE_URL"),
	}

	DB = &dbConfig{
		Host:     getEnv("DB_HOST"),
		Port:     getEnv("DB_PORT"),
		Username: getEnv("DB_USERNAME"),
		Password: getEnv("DB_PASSWORD"),
		Name:     getEnv("DB_NAME"),
		TimeZone: getEnv("DB_TIME_ZONE"),
		SSLMode:  getEnv("DB_SSL_MODE"),
	}

	JWT = &jwtConfig{
		AccessSecretKey:        getEnv("JWT_ACCESS_SECRET_KEY"),
		RefreshSecretKey:       getEnv("JWT_REFRESH_SECRET_KEY"),
		AccessTokenExpiration:  getEnvAsDuration("JWT_ACCESS_TOKEN_EXPIRATION", "15m"),
		RefreshTokenExpiration: getEnvAsDuration("JWT_REFRESH_TOKEN_EXPIRATION", "60m"),
	}

	SMTP = &smtpConfig{
		Host:     getEnv("SMTP_HOST"),
		Port:     getEnv("SMTP_PORT"),
		Username: getEnv("SMTP_USERNAME"),
		Password: getEnv("SMTP_PASSWORD"),
		From:     getEnv("SMTP_FROM"),
	}
}

func getEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	log.Error().Str("key", key).Msg("enviremont variable required")
	return ""
}

/* func getEnvAsInt(key string) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	log.Fatalf("Environment variable %s must be an integer", key)
	return 0
} */

func getEnvAsDuration(key, defaultVal string) time.Duration {
	if value, exists := os.LookupEnv(key); exists {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	duration, _ := time.ParseDuration(defaultVal)
	return duration
}
