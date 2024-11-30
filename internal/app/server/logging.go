package server

import "go.uber.org/zap"

func RetrieveLogger() *zap.Logger {
	return logger
}
