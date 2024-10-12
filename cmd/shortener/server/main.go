package main

import (
	"flag"
	"go-url-shortener/internal/app/config"
	"go-url-shortener/internal/app/server"
)

func main() {
	serverAddress := flag.String("a", "localhost:8080", "HTTP-сервер адрес")
	baseURL := flag.String("b", "http://localhost:8080", "Базовый адрес для сокращения URL")
	filePath := flag.String("f", "C:/tmp/short-url-db.json", "Путь к файлу для сохранения данных")
	flag.Parse()

	conf := config.NewConfig()
	configureFromFlags(conf, *serverAddress, *baseURL, *filePath)

	srv := server.NewServer(conf)
	srv.Start(*serverAddress)
}

func configureFromFlags(conf *config.Config, serverAddress, baseURL, filePath string) {
	conf.ServerAddress = serverAddress
	conf.BaseURL = baseURL
	conf.FilePath = filePath
}
