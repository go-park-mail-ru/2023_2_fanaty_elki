package middleware

import (
	"net/http"
	sessionUsecase "server/internal/Session/usecase"
)

type SessionMiddleware struct {
	sessionUC sessionUsecase.UsecaseI
}

func NewSessionMiddleware(sessionUC sessionUsecase.UsecaseI) *SessionMiddleware {
	return &SessionMiddleware{
		sessionUC: sessionUC,
	}
}
func (mw *SessionMiddleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		cookie, err := r.Cookie("session_id")
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		userId, err := mw.sessionUC.GetIdByCookie(cookie.Value)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if userId == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
