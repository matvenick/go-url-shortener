// Package handlers реализует обработчики запросов.
package handlers

import (
	"net/http"
)

func ExpandHandler(_ http.ResponseWriter, _ *http.Request) {
	// Обработка запроса ExpandHandler
	// ...
}

// func ShortenHandler(w http.ResponseWriter, r *http.Request) {
// 	var requestBody ShortenHandlerRequestBody

// 	// Декодируем JSON-тело запроса в структуру RequestBody.
// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&requestBody); err != nil {
// 		http.Error(w, "Failed to decode JSON request body", http.StatusBadRequest)
// 		return
// 	}

// 	// Вызываем функцию для генерации короткой ссылки.
// 	shortURL := generateShortURL()

// 	// Формируем JSON-ответ.
// 	responseBody := ShortenHandlerResponseBody{Result: shortURL}
// 	responseJSON, err := json.Marshal(responseBody)
// 	if err != nil {
// 		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
// 		return
// 	}

// 	// Отправляем ответ клиенту с поддержкой сжатия.
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusCreated)

// 	// Записываем JSON-ответ в ResponseWriter с обработкой возможной ошибки.
// 	if _, err := w.Write(responseJSON); err != nil {
// 		http.Error(w, "Failed to write JSON response", http.StatusInternalServerError)
// 		return
// 	}
// }

// // generateShortURL - возвращаю функцию generateShortURL.
// func generateShortURL() string {
// 	// Реализация GenerateShortURL.
// 	return fmt.Sprintf("http://localhost:8080/%v", genRandomString())
// }
