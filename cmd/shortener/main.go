package main

import (
	"errors"
	"flag"
	"go-url-shortener/internal/app/config"
	"go-url-shortener/internal/app/server"
	"log"
	"os"
	"path/filepath"
)

func main() {
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)

	serverAddress := flagOrEnv("a", "SERVER_ADDRESS", "localhost:8080", "HTTP-сервер адрес")
	baseURL := flagOrEnv("b", "BASE_URL", "http://localhost:8080", "Базовый адрес для сокращения URL")

	// Используем кроссплатформенный путь для временного файла
	defaultFilePath := filepath.Join(os.TempDir(), "short-url-db.json")
	filePath := flagOrEnv("f", "FILE_STORAGE_PATH", defaultFilePath, "Путь к файлу для сохранения данных")

	flag.Parse()

	if _, er := os.Stat(*filePath); errors.Is(er, os.ErrNotExist) {
		er = os.WriteFile(*filePath, []byte("[]"), 0644)
		if er != nil {
			log.Fatalf("Failed to create file for urls: %v", er)
			return
		}
	}

	if *filePath == defaultFilePath {
		defer os.Remove(defaultFilePath)
	}

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
