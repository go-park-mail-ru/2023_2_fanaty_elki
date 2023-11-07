package middleware

import (
	"net/http"
	"time"
	"go.uber.org/zap"
	//"go.uber.org/zap/zapcore"
)

type ACLog struct {
	logger *zap.SugaredLogger	
}



func NewACLog(logger *zap.SugaredLogger) *ACLog {
	return &ACLog{
		logger: logger,
	}
}

func (ac *ACLog) ACLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		start := time.Now()
		next.ServeHTTP(w, r)
		ac.logger.Info(r.URL.Path,
			zap.String("method", r.Method),
			zap.String("remote_addr", r.RemoteAddr),
			zap.String("url", r.URL.Path),
			zap.Duration("work_time", time.Since(start)),
		)
	})
}
