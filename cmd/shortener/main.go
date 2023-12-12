package main

import (
	"fmt"
	"go-url-shortener/internal/app/config" // Замените на ваш путь к пакету config
)

func main() {
	cfg := config.InitConfig()

	// Дальнейшая логика вашего приложения, использующая cfg
	fmt.Println("Server address:", cfg.Address)
	fmt.Println("Base URL:", cfg.BaseURL)
	fmt.Println("Environment:", cfg.Environment)

	// Запуск сервера и т.д.
}
