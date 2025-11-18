package middleware

import (
	"log"
	"net/http"
	"time"

	"scalable-backend-service/logger"

	"go.uber.org/zap"
)

func WithTimer(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		log.Default().Printf("Timing traffic on route %s", r.URL.Path)
		handler.ServeHTTP(w, r)
		defer logger.GetLoggerInstance().Info("Request time", zap.Int64("nanoseconds", time.Since(startTime).Nanoseconds()))
	})
}
