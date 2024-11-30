package server

import (
	"fmt"
	"go-url-shortener/internal/app/config"
	"go-url-shortener/internal/app/handlers"
	"go.uber.org/zap"
	"log"
	"net/http"
	"time"
)

var logger *zap.Logger

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

func StartServer() {
	InitializeLogger()

	conf := config.NewConfig()

	SetupRoutes()

	if err := http.ListenAndServe(conf.ServerAddress, LoggerMiddleware(http.DefaultServeMux)); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}

func SetupRoutes() {
	router := http.NewServeMux()
	router.HandleFunc("/shorten", handlers.ShortenHandler)
	router.HandleFunc("/expand", handlers.ExpandHandler)

	http.Handle("/", LoggerMiddleware(router))
}

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
