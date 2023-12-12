// Package handlers реализует обработчики запросов.
package handlers

import (
	"net/http"
)

// ShortenHandler обрабатывает запросы на сокращение ссылок.
func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	// Реализация обработчика сокращения ссылок
}

// ExpandHandler обрабатывает запросы на раскрытие сокращенных ссылок.
func ExpandHandler(w http.ResponseWriter, r *http.Request) {
	// Реализация обработчика раскрытия сокращенной ссылки
}
