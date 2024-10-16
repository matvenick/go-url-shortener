package server

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"go-url-shortener/internal/app/config"
	"go-url-shortener/internal/app/handlers"
	"go-url-shortener/internal/app/storage"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// SetupRoutes настраивает роуты HTTP сервера. Включаем GzipMiddleware и LoggerMiddleware в цепочку middleware.
func SetupRoutes(h *handlers.Handlers) http.Handler {
	// "github.com/gorilla/mux"
	router := mux.NewRouter()
	// router := http.NewServeMux()
	router.Handle("/", LoggerMiddleware(GzipMiddleware(http.HandlerFunc(h.ShortenHandler)))).Methods("POST")
	router.Handle("/{shortURL}", LoggerMiddleware(GzipMiddleware(http.HandlerFunc(h.ExpandHandler)))).Methods("GET")

	return router
}

var (
	logger *zap.Logger
	_      *config.Config
)

func init() {
	zapConfig := zap.NewProductionConfig()
	zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	var err error
	logger, err = zapConfig.Build()
	if err != nil {
		log.Fatal(err)
	}
}

// URLData представляет собой структуру для хранения данных URL.
type URLData struct {
	UUID        string `json:"uuid"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

// Server представляет собой HTTP-сервер.
type Server struct {
	router       http.Handler
	urlDataStore []URLData
	mu           sync.Mutex
	config       *config.Config
	handlers     *handlers.Handlers
}

// NewServer создает новый экземпляр HTTP-сервера.
func NewServer(conf *config.Config) (*Server, error) {
	store := storage.NewStorage(conf.UrlsPath)
	err := store.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to Load storage: %v", err)
	}
	h := handlers.NewHandlers(store)
	s := &Server{
		router:       SetupRoutes(h),
		config:       conf,
		urlDataStore: make([]URLData, 0),
		handlers:     h,
	}

	// Загрузка данных из файла при старте сервера.
	if s.config.UrlsPath != "" {
		if err := s.loadURLDataFromFile(); err != nil {
			logger.Error("Failed to load URL data from file", zap.Error(err))
		}
	}

	return s, nil
}

// Start запускает HTTP-сервер с заданным адресом.
func (s *Server) Start(address string) {
	fmt.Printf("Server is running on %s...\n", address)
	log.Fatal(http.ListenAndServe(address, s.router))
}

// gzipResponseWriter - новая структура для обновленного ResponseWriter.
type gzipResponseWriter struct {
	http.ResponseWriter
	*gzip.Writer
}

func (w *gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func (w *gzipResponseWriter) Header() http.Header {
	return w.ResponseWriter.Header()
}

func (w *gzipResponseWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
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
				logger.Error("Failed to close gzip.Writer", zap.Error(err))
			}

			return
		}

		// Продолжаем обработку запроса с оригинальным ResponseWriter.
		next.ServeHTTP(w, r)
	})
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
			zap.Duration("Duration", time.Since(start)),
		)
	})
}

// saveURLDataToFile сохраняет данные из внутреннего хранилища в файл.
func (s *Server) saveURLDataToFile() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	data, err := json.Marshal(s.urlDataStore)
	if err != nil {
		logger.Error("Failed to marshal data", zap.Error(err))
		return err
	}

	err = os.WriteFile(s.config.UrlsPath, data, 0644)
	if err != nil {
		logger.Error("Failed to write data to file", zap.Error(err))
		return err
	}

	return nil
}

// loadURLDataFromFile загружает данные из файла и добавляет их во внутреннее хранилище.
func (s *Server) loadURLDataFromFile() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	fileContent, err := os.ReadFile(s.config.UrlsPath)
	if err != nil {
		logger.Error("Failed to read data from file", zap.Error(err))
		return err
	}

	var loadedData []URLData
	if err := json.Unmarshal(fileContent, &loadedData); err != nil {
		logger.Error("Failed to unmarshal data", zap.Error(err))
		return err
	}

	s.urlDataStore = loadedData
	return nil
}
