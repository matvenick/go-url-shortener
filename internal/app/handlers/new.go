package handlers

import "go-url-shortener/internal/app/storage"

type Handlers struct {
	store *storage.Storage
}

func NewHandlers(store *storage.Storage) *Handlers {
	return &Handlers{
		store: store,
	}
}
