// Package handlers реализует обработчики запросов.
package handlers

import (
	"encoding/json"
	"go-url-shortener/internal/app/server"
	"net/http"
)

func ExpandHandler(_ http.ResponseWriter, _ *http.Request) {
	// Обработка запроса ExpandHandler
	// ...
}

// RequestBody - определение структуры тела запроса.
type RequestBody struct {
	URL string `json:"url"`
}

// ResponseBody - определение структуры тела ответа.
type ResponseBody struct {
	Result string `json:"result"`
}

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody RequestBody

	// Декодируем JSON-тело запроса в структуру RequestBody.
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestBody); err != nil {
		http.Error(w, "Failed to decode JSON request body", http.StatusBadRequest)
		return
	}

	// Вызываем функцию для генерации короткой ссылки.
	shortURL := server.GenerateShortURL()

	// Формируем JSON-ответ.
	responseBody := ResponseBody{Result: shortURL}
	responseJSON, err := json.Marshal(responseBody)
	if err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
		return
	}

	// Отправляем ответ клиенту с поддержкой сжатия.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// Записываем JSON-ответ в ResponseWriter с обработкой возможной ошибки.
	if _, err := w.Write(responseJSON); err != nil {
		http.Error(w, "Failed to write JSON response", http.StatusInternalServerError)
		return
	}
}
