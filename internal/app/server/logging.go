// Package shortener предоставляет основные компоненты HTTP-сервера.
package server

import (
	"go.uber.org/zap"
	"log"
)

func init() {
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
}

// RetrieveLogger возвращает экземпляр логгера.
func _() *zap.Logger {
	return logger
}
