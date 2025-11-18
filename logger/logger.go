package logger

import (
	"log"
	"sync"

	"go.uber.org/zap"
)

var (
	loggerInstance *zap.Logger
	once           sync.Once
)

func InitLogger() *zap.Logger {
	once.Do(func() {
		GetLoggerInstance()
	})
	return loggerInstance
}

func GetLoggerInstance() *zap.Logger {
	if loggerInstance == nil {
		l, err := zap.NewDevelopment()
		if err != nil {
			log.Fatalf("Failed to initialize logger - %v", err)
			return nil
		}
		loggerInstance = l
	}
	return loggerInstance
}

func Close() {
	if loggerInstance != nil {
		loggerInstance.Sync()
	}
}
