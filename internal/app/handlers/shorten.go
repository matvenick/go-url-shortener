// Package handlers реализует обработчики запросов.
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"math/rand/v2"
)

func (h *Handlers) ShortenHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody RequestBody

	// Декодируем JSON-тело запроса в структуру RequestBody.
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestBody); err != nil {
		http.Error(w, "Failed to decode JSON request body", http.StatusBadRequest)
		return
	}

	// Вызываем функцию для генерации короткой ссылки.
	randString := genRandomString()

	err := h.store.SaveURL(randString, requestBody.URL)
	if err != nil {
		http.Error(w, "Failed to decode JSON request body", http.StatusInternalServerError)
		return
	}

	// Формируем JSON-ответ.
	responseBody := ResponseBody{
		Result: fmt.Sprintf("http://localhost:8080/%v", randString),
	}
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

// RequestBody - определение структуры тела запроса.
type RequestBody struct {
	URL string `json:"url"`
}

// ResponseBody - определение структуры тела ответа.
type ResponseBody struct {
	Result string `json:"result"`
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
