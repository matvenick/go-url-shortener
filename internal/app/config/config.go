package config

import "os"

type Config struct {
	ServerAddress string
	BaseURL       string
}

func NewConfig() *Config {
	return &Config{
		ServerAddress: getEnv("SERVER_ADDRESS", "localhost:8080"),
		BaseURL:       getEnv("BASE_URL", "http://localhost:8080"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
