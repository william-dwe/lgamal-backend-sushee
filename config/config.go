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
	ExpiresAt  string
	HmacSecret string
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
			DBName:   getENV("DB_NAME", "wallet_db_william"),
			Port:     getENV("DB_PORT", "5432"),
		},
		AuthConfig: authConfig{
			ExpiresAt:  getENV("AUTH_EXPIRATION", "900"),
			HmacSecret: getENV("HMAC_SECRET", "very-secret"),
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
