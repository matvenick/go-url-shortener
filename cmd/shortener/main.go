package main

import (
	"flag"
	"go-url-shortener/internal/app/config"
	"go-url-shortener/internal/app/server"
	"log"
	"os"
)

func main() {
	serverAddress := flagOrEnv("a", "SERVER_ADDRESS", "localhost:8080", "HTTP-сервер адрес")
	baseURL := flagOrEnv("b", "BASE_URL", "http://localhost:8080", "Базовый адрес для сокращения URL")
	filePath := flagOrEnv("f", "FILE_STORAGE_PATH", "C:/tmp/short-url-db.json", "Путь к файлу для сохранения данных")
	flag.Parse()

	srv, err := server.NewServer(&config.Config{
		ServerAddress: *serverAddress,
		BaseURL:       *baseURL,
		UrlsPath:      *filePath,
	})
	if err != nil {
		log.Fatalf("Failed to create shortener: %v", err)
		return
	}
	srv.Start(*serverAddress)
}

func flagOrEnv(flagName, envVarName, fallbackValue, description string) *string {
	if value, ok := os.LookupEnv(envVarName); ok {
		v := value
		return &v
	}
	return flag.String(flagName, fallbackValue, description)
}

//тест коммент
