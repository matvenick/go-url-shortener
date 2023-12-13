package server

import (
	"go-url-shortener/internal/app/config"
	"go-url-shortener/internal/app/handlers"
	"net/http"
)

func StartServer() {
	conf := config.NewConfig()

	// ...

	SetupRoutes()

	http.ListenAndServe(conf.ServerAddress, nil)
}

func SetupRoutes() {
	router := http.NewServeMux()
	router.HandleFunc("/shorten", handlers.ShortenHandler)
	router.HandleFunc("/expand", handlers.ExpandHandler)

	http.Handle("/", router)
}
