package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			if _, err := fmt.Fprint(w, "http://localhost:8080/EwHXdJfB"); err != nil {
				http.Error(w, "Failed to write the response", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)
			return
		}

		http.Error(w, "Invalid request", http.StatusBadRequest)
	})

	http.HandleFunc("/EwHXdJfB", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.Header().Set("Location", "https://practicum.yandex.ru/")
			w.WriteHeader(http.StatusTemporaryRedirect)
			return
		}

		http.Error(w, "Invalid request", http.StatusBadRequest)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
