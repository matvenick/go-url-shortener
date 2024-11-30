// Package server предоставляет основные компоненты HTTP-сервера.
package server

import (
	"go.uber.org/zap"
	"log"
	"net/http"
	"time"
)

var logger *zap.Logger

func init() {
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
}

// LoggerMiddleware регистрирует входящие запросы и исходящие ответы.
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		logger.Info("Incoming request",
			zap.String("URI", r.RequestURI),
			zap.String("Method", r.Method),
		)

		// Замените w.(interface{ StatusCode() int }).StatusCode() на http.StatusOK
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

// RetrieveLogger возвращает экземпляр логгера.
func _() *zap.Logger {
	return logger
}
