package config

import "os"

// Config структура для хранения конфигурационных переменных.
type Config struct {
	ServerAddress string
	BaseURL       string
}

// NewConfig создает новую конфигурацию, используя переменные окружения или значения по умолчанию.
func NewConfig() *Config {
	return &Config{
		ServerAddress: getEnv("SERVER_ADDRESS", "localhost:8080"),
		BaseURL:       getEnv("BASE_URL", "http://localhost:8080"),
	}
}

// getEnv возвращает значение переменной окружения или значение по умолчанию, если переменная не задана.
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// LoadFromFlags загружает значения из флагов командной строки, перезаписывая переменные окружения.
func (c *Config) LoadFromFlags(serverAddress, baseURL string) {
	c.ServerAddress = serverAddress
	c.BaseURL = baseURL
}
