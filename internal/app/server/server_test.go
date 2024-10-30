// Package shortener предоставляет тесты для основных компонентов HTTP-сервера.
package server

import (
	"bytes"
	"encoding/json"
	"go-url-shortener/internal/app/handlers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func TestShortenHandler(t *testing.T) {
	// Тест для эндпоинта /api/shorten.
	t.Run("/api/shorten", func(t *testing.T) {
		// Подготавливаем тестовый запрос с JSON-телом.
		requestBody := handlers.ShortenHandlerRequestBody{URL: "https://practicum.yandex.ru"}
		jsonBody, err := json.Marshal(requestBody)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/api/shorten", bytes.NewBuffer(jsonBody))
		if err != nil {
			t.Fatal(err)
		}

		// Используем mux.Router для настройки роутинга.
		router := mux.NewRouter()
		router.HandleFunc("/api/shorten", handlers.ShortenHandler).Methods("POST")

		// Тестируем запрос.
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		// Проверяем код состояния.
		if status := rr.Code; status != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusCreated)
		}

		// Проверяем формат ответа.
		var responseBody handlers.ShortenHandlerResponseBody
		err = json.Unmarshal(rr.Body.Bytes(), &responseBody)
		if err != nil {
			t.Errorf("failed to unmarshal JSON response body: %v", err)
		}

		// Проверяем только префикс ответа.
		expectedPrefix := "http://localhost:8080/"
		if !strings.HasPrefix(responseBody.Result, expectedPrefix) {
			t.Errorf("handler returned unexpected result: got %v, want prefix %v", responseBody.Result, expectedPrefix)
		}
	})
}

func TestExpandHandler(t *testing.T) {
	// Тест для эндпоинта /expand.
	t.Run("/expand", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/expand", nil)
		if err != nil {
			t.Fatal(err)
		}

		// Используем mux.Router для настройки роутинга.
		router := mux.NewRouter()
		router.HandleFunc(
			"/expand",
			handlers.ExpandHandler,
		).Methods("GET")

		// Тестируем запрос.
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		// Проверяем код состояния.
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
		}
	})
}
