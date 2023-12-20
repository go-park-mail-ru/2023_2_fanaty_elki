package delivery

import (
	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
	"net/http"
	sessionUsecase "server/internal/Session/usecase"
	userUsecase "server/internal/User/usecase"
	"server/internal/domain/dto"
	"server/internal/domain/entity"
	mw "server/internal/middleware"
	"time"
)

//Result struct
type Result struct {
	Body interface{}
}

//RespError struct
type RespError struct {
	Err string
}

//SessionHandler struct
type SessionHandler struct {
	sessions sessionUsecase.SessionUsecaseI
	users    userUsecase.Iusecase
	logger   *mw.ACLog
}

//NewSessionHandler creates session handler
func NewSessionHandler(sessions sessionUsecase.SessionUsecaseI, users userUsecase.Iusecase, logger *mw.ACLog) *SessionHandler {
	return &SessionHandler{
		sessions: sessions,
		users:    users,
		logger:   logger,
	}
}

//RegisterAuthHandler registers cors handler api
func (handler *SessionHandler) RegisterAuthHandler(router *mux.Router) {
	router.HandleFunc("/api/logout", handler.Logout).Methods(http.MethodDelete)
	router.HandleFunc("/api/auth", handler.Auth).Methods(http.MethodGet)
	router.HandleFunc("/api/users/me", handler.Profile).Methods(http.MethodGet)
	router.HandleFunc("/api/users/me", handler.UpdateProfile).Methods(http.MethodPatch)
	router.HandleFunc("/api/users/me/icon", handler.UpdateAvatar).Methods(http.MethodPatch)
	router.HandleFunc("/api/csrf", handler.CreateCsrf).Methods(http.MethodPost)
}

//RegisterCorsHandler registers cors handler api
func (handler *SessionHandler) RegisterCorsHandler(router *mux.Router) {
	router.HandleFunc("/api/login", handler.Login).Methods(http.MethodPost)
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

	w.Header().Set("Content-Type", "application/json")
	if r.Header.Get("Content-Type") != "application/json" {
		handler.logger.LogError("bad content-type", entity.ErrBadContentType, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reqUser := dto.ReqCreateUser{}

	err := easyjson.UnmarshalFromReader(r.Body, &reqUser)

	if err != nil {
		handler.logger.LogError("problems with unmarshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := handler.users.CreateUser(dto.ToEntityCreateUser(&reqUser))
	switch err {
	case entity.ErrInternalServerError:
		handler.logger.LogError("problems with creating user", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	case entity.ErrInvalidBirthday, entity.ErrInvalidPassword, entity.ErrInvalidEmail, entity.ErrInvalidUsername, entity.ErrInvalidPhoneNumber:
		handler.logger.LogError("invalid field", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	case entity.ErrConflictEmail:
		handler.logger.LogError("conflcit", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(entity.StatusConflicEmail)
		return
	case entity.ErrConflictUsername:
		handler.logger.LogError("conflcit", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(entity.StatusConflicUsername)
		return
	case entity.ErrConflictPhoneNumber:
		handler.logger.LogError("conflcit", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(entity.StatusConflicPhoneNumber)
		return
	}

	w.WriteHeader(http.StatusCreated)

	body := &dto.RespID{ID: id}

	_, err = easyjson.MarshalToWriter(body, w)
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

	w.Header().Set("Content-Type", "application/json")

	w.Header().Set("Content-Type", "application/json")
	if r.Header.Get("Content-Type") != "application/json" {
		handler.logger.LogError("bad content-type", entity.ErrBadContentType, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reqUser := dto.ReqLoginUser{}

	err := easyjson.UnmarshalFromReader(r.Body, &reqUser)

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
			w.WriteHeader(http.StatusUnauthorized)
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

	user, err := handler.sessions.GetUserProfile(cookie.Value)
	if err == entity.ErrInternalServerError {
		handler.logger.LogError("problems with getting profile", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = easyjson.MarshalToWriter(user, w)
	if err != nil {
		handler.logger.LogError("problems with marshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
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

	w.Header().Set("Content-Type", "application/json")

	oldCookie, _ := r.Cookie("session_id")
	UserID, err := handler.sessions.Check(oldCookie.Value)
	if UserID == 0 {
		handler.logger.LogError("unauthorized", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusUnauthorized)
		oldCookie.Expires = time.Now().AddDate(0, 0, -1)
		http.SetCookie(w, oldCookie)
		return
	}

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    oldCookie.Value,
		Expires:  time.Now().Add(150 * time.Hour),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, cookie)
	user, err := handler.sessions.CreateCookieAuth(&entity.Cookie{
		UserID:       UserID,
		SessionToken: cookie.Value,
	})
	if err != nil {
		handler.logger.LogError("problems with auth cookie", err, w.Header().Get("request-id"), r.URL.Path)
	}

	_, err = easyjson.MarshalToWriter(user, w)
	if err != nil {
		handler.logger.LogError("problems with marshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
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

	w.Header().Set("Content-Type", "application/json")
	cookie, _ := r.Cookie("session_id")
	user, err := handler.sessions.GetUserProfile(cookie.Value)
	if err == entity.ErrInternalServerError {
		handler.logger.LogError("problems with getting profile", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = easyjson.MarshalToWriter(user, w)
	if err != nil {
		handler.logger.LogError("problems with marshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

//UpdateProfile handles update profile request
func (handler *SessionHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	if r.Header.Get("Content-Type") != "application/json" {
		handler.logger.LogError("bad content-type", entity.ErrBadContentType, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cookie, _ := r.Cookie("session_id")
	id, _ := handler.sessions.GetIDByCookie(cookie.Value)

	updatedUser := &dto.ReqUpdateUser{}

	err := easyjson.UnmarshalFromReader(r.Body, updatedUser)
	if err != nil {
		handler.logger.LogError("prbolems with unmarshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = handler.users.UpdateUser(dto.ToEntityUpdateUser(updatedUser, id))
	switch err {
	case entity.ErrInternalServerError:
		handler.logger.LogError("problems with updating user", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	case entity.ErrInvalidEmail, entity.ErrInvalidUsername, entity.ErrInvalidPhoneNumber:
		handler.logger.LogError("invalid field", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	case entity.ErrConflictEmail:
		handler.logger.LogError("conflcit", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(entity.StatusConflicEmail)
		return
	case entity.ErrConflictUsername:
		handler.logger.LogError("conflcit", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(entity.StatusConflicUsername)
		return
	case entity.ErrConflictPhoneNumber:
		handler.logger.LogError("conflcit", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(entity.StatusConflicPhoneNumber)
		return
	case entity.ErrNotFound:
		handler.logger.LogError("user not found", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

//UpdateAvatar handles update avatar request
func (handler *SessionHandler) UpdateAvatar(w http.ResponseWriter, r *http.Request) {

	cookie, _ := r.Cookie("session_id")
	id, _ := handler.sessions.GetIDByCookie(cookie.Value)

	r.ParseMultipartForm(10 << 20)
	file, filehandler, err := r.FormFile("image")
	if err != nil {
		handler.logger.LogError("prbolems receving image", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()

	err = handler.users.UpdateAvatar(file, filehandler, id)
	if err != nil {
		handler.logger.LogError("prbolems creating foto", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

//CreateCsrf handles create csrf request
func (handler *SessionHandler) CreateCsrf(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("session_id")
	token, err := handler.sessions.CreateCsrf(cookie.Value)
	if err != nil {
		handler.logger.LogError("problems with creating csrf-token", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("X-CSRF-Token", token)
	w.Header().Set("Access-Control-Expose-Headers", "X-CSRF-Token")
	w.WriteHeader(http.StatusCreated)
}
