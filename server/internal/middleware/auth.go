package middleware

import (
	"net/http"
	sessionUsecase "server/internal/Session/usecase"
	"server/internal/domain/entity"
	"time"
	"regexp"
)

type SessionMiddleware struct {
	sessionUC sessionUsecase.UsecaseI
	logger *ACLog
}

func NewSessionMiddleware(sessionUC sessionUsecase.UsecaseI, logger *ACLog) *SessionMiddleware {
	return &SessionMiddleware{
		sessionUC: sessionUC,
		logger: logger,
	}
}

func (mw *SessionMiddleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		cookie, err := r.Cookie("session_id")
		
		if err == http.ErrNoCookie {
			mw.logger.LogError("no cookie", err, w.Header().Get("request-id"), r.URL.Path)
			w.WriteHeader(http.StatusUnauthorized)
			return
		} else if err != nil {
			mw.logger.LogError("problems with getting cookie", err, w.Header().Get("request-id"), r.URL.Path)
			w.WriteHeader(http.StatusInternalServerError)
		}

		userId, err := mw.sessionUC.Check(cookie.Value)
		if err != nil {
			mw.logger.LogError("problems with getting user by cookie", err, w.Header().Get("request-id"), r.URL.Path)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if userId == 0 {
			mw.logger.LogError("user not found", err, w.Header().Get("request-id"), r.URL.Path)
			cookie.Expires = time.Now().AddDate(0, 0, -1)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if r.Method != http.MethodGet && r.URL.Path != "/api/csrf" && !checkCsat(r.URL.Path){
			csrfToken := r.Header.Get("X-CSRF-Token")
			err = mw.sessionUC.CheckCsrf(cookie.Value, csrfToken)
			if err != nil {
				if err == entity.ErrFailCSRF {
					mw.logger.LogError("fail csrf", err, w.Header().Get("request-id"), r.URL.Path)
					w.WriteHeader(entity.StatusFailCSRF)
					return
				}
			
				mw.logger.LogError("problems with checking csrf on server", err, w.Header().Get("request-id"), r.URL.Path)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

func checkCsat(url string) bool {
	pattern := `/api/csat`

	match, err := regexp.MatchString(pattern, url)
	if err != nil {
		return false
	}

	return match
}

