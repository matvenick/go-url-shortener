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

type flagsData struct {
	serverAddress *string
	baseURL       *string
	filePath      *string
}

func main() {
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)

	// Используем кроссплатформенный путь для временного файла
	defaultFilePath := filepath.Join(os.TempDir(), "short-url-db.json")

	var flags flagsData
	flags.serverAddress = flag.String("a", "localhost:8080", "HTTP-сервер адрес")
	flags.baseURL = flag.String("b", "http://localhost:8080", "Базовый адрес для сокращения URL")
	flags.filePath = flag.String("f", defaultFilePath, "Путь к файлу для сохранения данных")
	flag.Parse()

	cfg := &config.Config{
		ServerAddress: *flags.serverAddress,
		BaseURL:       *flags.baseURL,
		UrlsPath:      *flags.filePath,
	}
	if value, ok := os.LookupEnv("SERVER_ADDRESS"); ok {
		cfg.ServerAddress = value
	}
	if value, ok := os.LookupEnv("BASE_URL"); ok {
		cfg.BaseURL = value
	}
	if value, ok := os.LookupEnv("FILE_STORAGE_PATH"); ok {
		cfg.UrlsPath = value
	}

	if _, er := os.Stat(cfg.UrlsPath); errors.Is(er, os.ErrNotExist) {
		er = os.WriteFile(cfg.UrlsPath, []byte("[]"), 0644)
		if er != nil {
			log.Fatalf("Failed to create file for urls: %v", er)
			return
		}
	}

	if cfg.UrlsPath == defaultFilePath {
		defer os.Remove(defaultFilePath)
	}

	srv, err := server.NewServer(cfg)
	if err != nil {
		log.Fatalf("Failed to create shortener: %v", err)
		return
	}
	srv.Start(cfg.ServerAddress)
}

func flagOrEnv(flagName, envVarName, fallbackValue, description string) *string {
	if value, ok := os.LookupEnv(envVarName); ok {
		v := value
		return &v
	}
	return flag.String(flagName, fallbackValue, description)
}
