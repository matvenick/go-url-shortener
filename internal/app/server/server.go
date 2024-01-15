// server.go
package server

import (
	"fmt"
	"go-url-shortener/internal/app/config"
	"go-url-shortener/internal/app/handlers"
	"go.uber.org/zap"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var logger *zap.Logger

// InitializeLogger инициализирует логгер.
func InitializeLogger() {
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		fmt.Println("Failed to initialize logger:", err)
	}
	defer func() {
		_ = logger.Sync()
	}()
}

// StartServer запускает сервер.
func StartServer() {
	InitializeLogger()

	conf := config.NewConfig()

	SetupRoutes()

	if err := http.ListenAndServe(conf.ServerAddress, LoggerMiddleware(http.DefaultServeMux)); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}

// SetupRoutes настраивает маршруты сервера.
func SetupRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/api/shorten", handlers.ShortenHandler).Methods("POST")
	router.HandleFunc("/expand", handlers.ExpandHandler).Methods("GET")

	http.Handle("/", LoggerMiddleware(router))
}

// LoggerMiddleware предоставляет middleware для логирования запросов.
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		if logger == nil {
			log.Fatal("Logger is not initialized")
		}

		logger.Info("Incoming request",
			zap.String("URI", r.RequestURI),
			zap.String("Method", r.Method),
		)

		// Замените w.(interface{ StatusCode() sudo dscl . -create /Users/nimatveev UserShell /bin/bashint }).StatusCode() на w.WriteHeader(http.StatusOK)
		next.ServeHTTP(w, r)

		// Замените w.(interface{ StatusCode() int }).StatusCode() на http.StatusOK
		logger.Info("Outgoing response",
			zap.Int("StatusCode", http.StatusOK),
		)

		logger.Info("Request processed in",
			zap.Duration("Duration", time.Since(start)),
		)
	})
}

// GetShortURL генерирует короткую ссылку.
func GetShortURL(originalURL string) string {
	// Вставьте ваш код для генерации короткой ссылки здесь.
	return "http://example.com/shortURL"
}
