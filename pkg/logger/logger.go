package logger

import (
	"fmt"
	"go.uber.org/zap"
)

func LoggerCreate() *zap.Logger {
	config := zap.NewProductionConfig()

	config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)

	logger, err := config.Build()
	if err != nil {
		fmt.Printf("Error logger create: %v", err)
	}
	return logger
}
