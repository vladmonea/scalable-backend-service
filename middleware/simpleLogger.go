package middleware

import (
	"net/http"

	"scalable-backend-service/logger"

	"go.uber.org/zap"
)

func WithSimpleLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.GetLoggerInstance().Info("Incoming traffic on route", zap.String("path", r.URL.Path))
		handler.ServeHTTP(w, r)
	})
}
