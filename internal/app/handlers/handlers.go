// handlers.go
package handlers

import (
	"encoding/json"
	"go-url-shortener/internal/app/server"
	"net/http"
)

// RequestBody представляет входные данные для эндпоинта /api/shorten.
type RequestBody struct {
	URL string `json:"url"`
}

// ResponseBody представляет данные для ответа от эндпоинта /api/shorten.
type ResponseBody struct {
	Result string `json:"result"`
}

// ShortenHandler обрабатывает запросы к эндпоинту /api/shorten.
func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody RequestBody

	// Декодируем JSON-тело запроса в структуру RequestBody.
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestBody); err != nil {
		http.Error(w, "Failed to decode JSON request body", http.StatusBadRequest)
		return
	}

	// Вызываем функцию для генерации короткой ссылки.
	shortURL := server.GetShortURL(requestBody.URL)

	// Формируем JSON-ответ.
	responseBody := ResponseBody{Result: shortURL}
	responseJSON, err := json.Marshal(responseBody)
	if err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
		return
	}

	// Отправляем ответ клиенту.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(responseJSON)
	if err != nil {
		http.Error(w, "Failed to write JSON response", http.StatusInternalServerError)
		return
	}
}

// ExpandHandler обрабатывает запросы к эндпоинту /expand.
func ExpandHandler(w http.ResponseWriter, r *http.Request) {
	// Здесь можно добавить логику для обработки запросов к /expand.
	w.WriteHeader(http.StatusOK)
}
