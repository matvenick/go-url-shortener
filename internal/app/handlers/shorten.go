// Package handlers реализует обработчики запросов.
package handlers

import (
	"fmt"
	"io"
	"net/http"

	"math/rand/v2"
)

func (h *Handlers) ShortenHandler(w http.ResponseWriter, r *http.Request) {
	url, _ := io.ReadAll(r.Body)
	// log.Println("data")
	// log.Println(string(data))

	// var requestBody RequestBody

	// Декодируем JSON-тело запроса в структуру RequestBody.
	// decoder := json.NewDecoder(r.Body)
	// if err := decoder.Decode(&requestBody); err != nil {
	// http.Error(w, "Failed to decode JSON request body", http.StatusBadRequest)
	// return
	// }

	// Вызываем функцию для генерации короткой ссылки.
	randString := genRandomString()

	err := h.store.SaveURL(randString, string(url))
	if err != nil {
		http.Error(w, "Failed to decode JSON request body", http.StatusInternalServerError)
		return
	}

	resultURL := fmt.Sprintf("http://localhost:8080/%v", randString)

	// Формируем JSON-ответ.
	// responseBody := ResponseBody{
	// 	Result: resultURL,
	// }
	// responseJSON, err := json.Marshal(responseBody)
	// if err != nil {
	// 	http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
	// 	return
	// }

	// Отправляем ответ клиенту с поддержкой сжатия.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// Записываем JSON-ответ в ResponseWriter с обработкой возможной ошибки.
	if _, err := w.Write([]byte(resultURL)); err != nil {
		http.Error(w, "Failed to write JSON response", http.StatusInternalServerError)
		return
	}
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
