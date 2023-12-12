package server

import (
	"github.com/gorilla/mux"
	"go-url-shortener/internal/app/handlers" // Замените на ваш реальный путь
	"net/http"
)

func SetupRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/shorten", handlers.ShortenHandler).Methods("POST")
	router.HandleFunc("/expand", handlers.ExpandHandler).Methods("GET")

	http.Handle("/", router)
}
