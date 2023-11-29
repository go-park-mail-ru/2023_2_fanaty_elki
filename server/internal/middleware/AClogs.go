package middleware

import (
	"net/http"
	"server/internal/domain/entity"
	"strconv"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type ACLog struct {
	logger      *zap.SugaredLogger
	errorLogger *zap.SugaredLogger
	hitcounter  entity.HitStats
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

func NewACLog(logger *zap.SugaredLogger, errorLogger *zap.SugaredLogger, hc entity.HitStats) *ACLog {
	return &ACLog{
		logger:      logger,
		errorLogger: errorLogger,
		hitcounter:  hc,
	}
}

func (ac *ACLog) ACLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

		switch status {
		case 200:
			ac.hitcounter.Ok.Inc()
		case 404:
			ac.hitcounter.NotFoundError.Inc()
		case 500:
			ac.hitcounter.InternalServerError.Inc()
		}

		ac.hitcounter.UrlMetric.WithLabelValues(strconv.Itoa(status), r.URL.Path).Inc()
	})
}

func (ac *ACLog) LogError(message string, err error, requestID string, url string) {
	ac.errorLogger.Errorw(message,
		zap.Error(err),
		zap.String("request-id", requestID),
		zap.String("url", url),
	)
}
