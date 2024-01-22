package main

import (
	"flag"
	"go-url-shortener/internal/app/config"
	"go-url-shortener/internal/app/server"
)

func main() {
	serverAddress := flag.String("a", "localhost:8080", "HTTP-сервер адрес")
	baseURL := flag.String("b", "http://localhost:8080", "Базовый адрес для сокращения URL")
	flag.Parse()

	conf := config.NewConfig()
	configureFromFlags(conf, *serverAddress, *baseURL)

	srv := server.NewServer()
	srv.Start(*serverAddress)
}

func configureFromFlags(conf *config.Config, serverAddress, baseURL string) {
	conf.ServerAddress = serverAddress
	conf.BaseURL = baseURL
}
