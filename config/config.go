package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type dbConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

type authConfig struct {
	TimeLimitAccessToken  string
	TimeLimitRefreshToken  string
	HmacSecretAccessToken string
	HmacSecretRefreshToken string
	IsTesting  string
}
type AppConfig struct {
	DBConfig   dbConfig
	AuthConfig authConfig
}

func initConfig() AppConfig {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")
		fmt.Println("[Warning] env file not found/ failed to be loaded, proceed to use default env value...")
		fmt.Println(">> here's the error details:", err)
		fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")
	}
	return AppConfig{
		DBConfig: dbConfig{
			Host:     getENV("DB_HOST", "localhost"),
			User:     getENV("DB_USER", "postgres"),
			Password: getENV("DB_PASSWORD", "postgres"),
			DBName:   getENV("DB_NAME", "db_restaurant"),
			Port:     getENV("DB_PORT", "5432"),
		},
		AuthConfig: authConfig{
			TimeLimitAccessToken:  getENV("ACCESS_TOKEN_EXPIRATION", "900"),
			TimeLimitRefreshToken:  getENV("REFRESH_TOKEN_EXPIRATION", "86400"),
			HmacSecretAccessToken: getENV("HMAC_SECRET_ACCESS_TOKEN", "very-secret"),
			HmacSecretRefreshToken: getENV("HMAC_SECRET_REFRESH_TOKEN", "super-secret"),
			IsTesting:  getENV("IS_TESTING", "true"),
		},
	}
}

func getENV(key, defaultVal string) string {
	env := os.Getenv(key)
	if env == "" {
		return defaultVal
	}
	return env
}

var Config = initConfig()
