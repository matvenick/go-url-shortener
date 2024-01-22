// Package server предоставляет основные компоненты HTTP-сервера.
package server

import (
	"compress/gzip"
	"fmt"
	"go-url-shortener/internal/app/handlers"
	"log"
	"net/http"
	"strings"
)

// Server представляет собой HTTP-сервер.
type Server struct {
	router http.Handler
}

// NewServer создает новый экземпляр HTTP-сервера.
func NewServer() *Server {
	return &Server{
		router: SetupRoutes(),
	}
}

// Start запускает HTTP-сервер с заданным адресом.
func (s *Server) Start(address string) {
	fmt.Printf("Server is running on %s...\n", address)
	log.Fatal(http.ListenAndServe(address, s.router))
}

// SetupRoutes настраивает роуты HTTP сервера. Включаем GzipMiddleware в цепочку middleware.
func SetupRoutes() *http.ServeMux {
	router := http.NewServeMux()
	router.Handle("/shorten", LoggerMiddleware(GzipMiddleware(http.HandlerFunc(handlers.ShortenHandler))))
	router.Handle("/expand", LoggerMiddleware(GzipMiddleware(http.HandlerFunc(handlers.ExpandHandler))))

	return router
}

// GenerateShortURL - возвращаю функцию GenerateShortURL.
func GenerateShortURL() string {
	// Реализация GenerateShortURL.
	return "http://localhost:8080/EwHXdJfB"
}

// gzipResponseWriter - новая структура для обновленного ResponseWriter.
type gzipResponseWriter struct {
	http.ResponseWriter
	*gzip.Writer
}

func (w *gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

// GzipMiddleware обеспечивает сжатие ответов.
func GzipMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Проверяем заголовок Accept-Encoding клиента.
		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			// Если клиент поддерживает gzip, применяем middleware.
			w.Header().Set("Content-Encoding", "gzip")
			gz := gzip.NewWriter(w)

			// Обновляем ResponseWriter с gzip.Writer.
			gzw := &gzipResponseWriter{ResponseWriter: w, Writer: gz}

			// Продолжаем обработку запроса с обновленным ResponseWriter.
			next.ServeHTTP(gzw, r)

			// После завершения обработки запроса проверяем ошибку при закрытии gzip.Writer.
			if err := gz.Close(); err != nil {
				// Обработка ошибки, например, логирование.
				fmt.Println("Failed to close gzip.Writer:", err)
			}

			return
		}

		// Продолжаем обработку запроса с оригинальным ResponseWriter.
		next.ServeHTTP(w, r)
	})
}
