// Package handlers реализует обработчики запросов.
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"math/rand/v2"
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
	shortURL := generateShortURL()

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

// generateShortURL - возвращаю функцию generateShortURL.
func generateShortURL() string {
	// Реализация GenerateShortURL.
	return fmt.Sprintf("http://localhost:8080/%v", genRandomString())
}

func genRandomString() string {
	res := ""
	for i := 0; i < 8; i++ {
		res += string(genRandomRune())
	}
	return res
}

func genRandomRune() rune {
	isUpper := rand.IntN(2) == 0
	symbol := rand.IntN(26)
	if isUpper {
		return rune(symbol + 'A')
	}
	return rune(symbol + 'a')
}
