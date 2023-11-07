package middleware

import (
	//"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"

	//"go.uber.org/zap/zapcore"
	"github.com/google/uuid"
	
)

type ACLog struct {
	logger *zap.SugaredLogger	
}

type responseRecorder struct {
	http.ResponseWriter
	status int
}

func (r *responseRecorder) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}

func NewLoggingResponseWriter(w http.ResponseWriter) *responseRecorder {
	return &responseRecorder{w, http.StatusOK}
}

func NewACLog(logger *zap.SugaredLogger) *ACLog {
	return &ACLog{
		logger: logger,
	}
}

func (ac *ACLog) ACLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		rec := NewLoggingResponseWriter(w)
		requestID := uuid.New().String()
		w.Header().Set("request-id", requestID)
		start := time.Now()
		next.ServeHTTP(rec, r)
		status := rec.status
		ac.logger.Infow("Access log info",
			zap.String("method", r.Method),
			zap.String("remote addr", r.RemoteAddr),
			zap.String("url", r.URL.Path),
			zap.String("request-id", requestID),
			zap.Int("status", status),
			zap.Duration("work time", time.Duration(time.Since(start).Microseconds())),
		)
	})
}
