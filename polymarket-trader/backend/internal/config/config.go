package config

import (
	"os"
)

type Config struct {
	ServerPort    string
	DatabaseURL   string
	RedisAddr     string
	PolymarketAPI string
	PrivateKey    string
	ApiKey        string
	ApiSecret     string
	ApiPassphrase string
}

func LoadConfig() *Config {
	return &Config{
		ServerPort:    getEnv("SERVER_PORT", "8080"),
		DatabaseURL:   getEnv("DATABASE_URL", "host=localhost user=user password=password dbname=polymarket port=5432 sslmode=disable"),
		RedisAddr:     getEnv("REDIS_ADDR", "localhost:6379"),
		PolymarketAPI: getEnv("POLYMARKET_API_URL", "https://clob.polymarket.com"),
		PrivateKey:    getEnv("PRIVATE_KEY", ""),
		ApiKey:        getEnv("API_KEY", ""),
		ApiSecret:     getEnv("API_SECRET", ""),
		ApiPassphrase: getEnv("API_PASSPHRASE", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
