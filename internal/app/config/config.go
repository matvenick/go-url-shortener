package config

import "os"

type Config struct {
	ServerAddress string `json:"server_address,omitempty"`
	BaseURL       string `json:"base_url,omitempty"`
	FilePath      string `json:"file_path,omitempty"`
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
