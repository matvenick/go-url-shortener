package handlers

import (
	"go-url-shortener/internal/app/config"
	"go-url-shortener/internal/app/storage"
)

type Handlers struct {
	store *storage.Storage
	conf  *config.Config
}

func NewHandlers(store *storage.Storage, conf *config.Config) *Handlers {
	return &Handlers{
		store: store,
		conf:  conf,
	}
}
