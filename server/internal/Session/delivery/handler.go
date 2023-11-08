package delivery

import (
	"encoding/json"
	// "fmt"
	"io/ioutil"
	"net/http"
	sessionUsecase "server/internal/Session/usecase"
	userUsecase "server/internal/User/usecase"
	"server/internal/domain/dto"
	"server/internal/domain/entity"
	"time"
	"github.com/gorilla/mux"
	mw "server/internal/middleware"
)


type Result struct {
	Body interface{}
}

type RespError struct {
	Err string
}

type SessionHandler struct {
	sessions sessionUsecase.UsecaseI
	users    userUsecase.UsecaseI
	logger *mw.ACLog
}

func NewSessionHandler(sessions sessionUsecase.UsecaseI, users userUsecase.UsecaseI, logger *mw.ACLog) *SessionHandler {
	return &SessionHandler{
		sessions: sessions,
		users:    users,
		logger: logger,
	}
}

func (handler *SessionHandler) RegisterAuthHandler(router *mux.Router) {
	router.HandleFunc("/api/logout", handler.Logout).Methods(http.MethodDelete)
	router.HandleFunc("/api/auth", handler.Auth).Methods(http.MethodGet)
	router.HandleFunc("/api/me", handler.Profile).Methods(http.MethodGet)
	router.HandleFunc("/api/me", handler.UpdateProfile).Methods(http.MethodPatch)
}

func (handler *SessionHandler) RegisterCorsHandler(router *mux.Router) {
	router.HandleFunc("/api/login", handler.Login).Methods(http.MethodPost)
}

func (handler *SessionHandler) RegisterHandler(router *mux.Router) {
	router.HandleFunc("/api/users", handler.SignUp).Methods(http.MethodPost)
}


// SignUp godoc
// @Summary      Signing up a user
// @Description  Signing up a user
// @Tags        users
// @Accept     application/json
// @Produce  application/json
// @Param 	user	 body	 store.User	 true	 "User object for signing up"
// @Success  201 {object}  integer "success create User return id"
// @Failure 400 {object} error "bad request"
// @Failure 500 {object} error "internal server error"
// @Router   /api/users [post]
func (handler *SessionHandler) SignUp(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")

	jsonbody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		handler.logger.LogError("problems with reading json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reqUser := dto.ReqCreateUser{}
	err = json.Unmarshal(jsonbody, &reqUser)
	
	if err != nil {
		handler.logger.LogError("problems with unmarshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	id, err := handler.users.CreateUser(dto.ToEntityCreateUser(&reqUser))
	
	if err != nil {
		if err == entity.ErrInternalServerError {
			handler.logger.LogError("problems with creating user", err, w.Header().Get("request-id"), r.URL.Path)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		handler.logger.LogError("bad request", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	body := map[string]interface{}{
		"ID": id,
	}

	err = json.NewEncoder(w).Encode(&Result{Body: body})
	if err != nil {
		handler.logger.LogError("problems marshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// Login godoc
// @Summary      Log in user
// @Description  Log in user
// @Tags        users
// @Accept     application/json
// @Produce  application/json
// @Param    user body store.User true "user object for login"
// @Success  200 {object}  string "success login User return cookie"
// @Failure 400 {object} error "bad request"
// @Failure 404 {object} error "not found"
// @Failure 500 {object} error "internal server error"
// @Router   /api/login [post]
func (handler *SessionHandler) Login(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")

	jsonbody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		handler.logger.LogError("problems with reading json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reqUser := dto.ReqLoginUser{}
	err = json.Unmarshal(jsonbody, &reqUser)

	if err != nil {
		handler.logger.LogError("problems with unmarshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cookieUC, err := handler.sessions.Login(dto.ToEntityLoginUser(&reqUser))
	if err != nil {
		if err == entity.ErrInternalServerError {
			handler.logger.LogError("problems with creating cookie", err, w.Header().Get("request-id"), r.URL.Path)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			handler.logger.LogError("incorrect data", err, w.Header().Get("request-id"), r.URL.Path)
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    cookieUC.SessionToken,
		Expires:  time.Now().Add(cookieUC.MaxAge),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, cookie)
	body := map[string]interface{}{
		"Username": reqUser.Username,
	}

	err = json.NewEncoder(w).Encode(&Result{Body: body})
	if err != nil {
		handler.logger.LogError("problems with marshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
	}

}

// Logout godoc
// @Summary      Log out user
// @Description  Log out user
// @Tags        users
// @Accept     application/json
// @Produce  application/json
// @Param    cookie header string true "Log out user"
// @Success 200 "void" "success log out"
// @Failure 400 {object} error "bad request"
// @Failure 401 {object} error "unauthorized"
// @Router   /api/logout [delete]
func (handler *SessionHandler) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("session_id")
	err := handler.sessions.Logout(&entity.Cookie{
		SessionToken: cookie.Value,
	})

	if err != nil {
		handler.logger.LogError("problems with deleting cookie", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	cookie.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, cookie)
}

// Auth godoc
// @Summary      checking auth
// @Description  checking auth
// @Tags        users
// @Accept     application/json
// @Produce  application/json
// @Param    cookie header string true "Checking user authentication"
// @Success  200 {object} integer "success authenticate return id"
// @Failure 401 {object} error "unauthorized"
// @Router   /api/auth [get]
func (handler *SessionHandler) Auth(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")

	cookie, _ := r.Cookie("session_id")
	username, err := handler.sessions.Check(cookie.Value)
	if username == nil {
		handler.logger.LogError("unauthorized", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusUnauthorized)
		cookie.Expires = time.Now().AddDate(0, 0, -1)
		http.SetCookie(w, cookie)
		return
	}

	http.SetCookie(w, cookie)
	body := map[string]interface{}{
		"Username": username,
	}
	err = json.NewEncoder(w).Encode(&Result{Body: body})

	if err != nil {
		handler.logger.LogError("problems with marshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// Profile godoc
// @Summary      getting profile
// @Description  getting profile
// @Tags        users
// @Accept     application/json
// @Produce  application/json
// @Param    cookie header string true "Checking user authentication"
// @Success  200 {object} dto.ReqGetUserProfile "success getting profile return User"
// @Failure 401 {object} error "unauthorized"
// @Router   /api/me [get]
func (handler *SessionHandler) Profile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	cookie, _ := r.Cookie("session_id")
	user, err := handler.sessions.GetUserProfile(cookie.Value)
	if err == entity.ErrInternalServerError{
		handler.logger.LogError("problems with getting profile", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(&Result{Body: user})
	if err != nil {
		handler.logger.LogError("problems with marshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}


func (handler *SessionHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {

	cookie, _ := r.Cookie("session_id")
	id, _ := handler.sessions.GetIdByCookie(cookie.Value)

	jsonbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handler.logger.LogError("problems with reading json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	updatedUser := &dto.ReqUpdateUser{}
	err = json.Unmarshal(jsonbody, &updatedUser)
	if err != nil {
		handler.logger.LogError("prbolems with unmarshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = handler.users.UpdateUser(dto.ToEntityUpdateUser(updatedUser, id))
	if err != nil {
		if err == entity.ErrInternalServerError {
			handler.logger.LogError("problems with updating user", err, w.Header().Get("request-id"), r.URL.Path)
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else if err == entity.ErrNotFound {
			handler.logger.LogError("user not found", err, w.Header().Get("request-id"), r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		handler.logger.LogError("incorrect data", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
