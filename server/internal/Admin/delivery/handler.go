package delivery

import (
	"encoding/json"
	//"fmt"
	"io/ioutil"
	"net/http"
	adminUsecase "server/internal/Admin/usecase"
	userUsecase "server/internal/User/usecase"
	"server/internal/domain/dto"
	"server/internal/domain/entity"
	mw "server/internal/middleware"
	"time"

	"github.com/gorilla/mux"
)

type Result struct {
	Body interface{}
}

type RespError struct {
	Err string
}

type AdminHandler struct {
	admins adminUsecase.UsecaseI
	users    userUsecase.UsecaseI
	logger   *mw.ACLog
}

func NewAdminHandler(admins adminUsecase.UsecaseI, users userUsecase.UsecaseI, logger *mw.ACLog) *AdminHandler {
	return &AdminHandler{
		admins: admins,
		users:    users,
		logger:   logger,
	}
}

func (handler *AdminHandler) RegisterAdminHandler(router *mux.Router) {
	router.HandleFunc("/api/csat/admin/logout", handler.Logout).Methods(http.MethodDelete)
	router.HandleFunc("/api/csat/admin/auth", handler.Auth).Methods(http.MethodGet)
}

func (handler *AdminHandler) RegisterCorsHandler(router *mux.Router) {
	router.HandleFunc("/api/csat/admin/login", handler.Login).Methods(http.MethodPost)
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
func (handler *AdminHandler) Login(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	w.Header().Set("Content-Type", "application/json")
	if r.Header.Get("Content-Type") != "application/json" {
		handler.logger.LogError("bad content-type", entity.ErrBadContentType, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	jsonbody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		handler.logger.LogError("problems with reading json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reqUser := dto.ReqAdmin{}
	err = json.Unmarshal(jsonbody, &reqUser)

	if err != nil {
		handler.logger.LogError("problems with unmarshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cookieUC, err := handler.admins.Login(dto.ToEntityLoginAdmin(&reqUser))

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
		Name:     "admin_id",
		Value:    cookieUC.SessionToken,
		Expires:  time.Now().Add(cookieUC.MaxAge),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, cookie)

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
func (handler *AdminHandler) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("admin_id")
	err := handler.admins.Logout(&entity.Cookie{
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
func (handler *AdminHandler) Auth(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	oldCookie, _ := r.Cookie("admin_id")
	userId, err := handler.admins.Check(oldCookie.Value)
	if userId == 0 {
		handler.logger.LogError("unauthorized", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusUnauthorized)
		oldCookie.Expires = time.Now().AddDate(0, 0, -1)
		http.SetCookie(w, oldCookie)
		return
	}

	cookie := &http.Cookie{
		Name:     "admin_id",
		Value:    oldCookie.Value,
		Expires:  time.Now().Add(150 * time.Hour),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, cookie)
	admin, err := handler.admins.CreateCookieAuth(&entity.Cookie{
		UserID:       userId,
		SessionToken: cookie.Value,
	})
	if err != nil {
		handler.logger.LogError("problems with auth cookie", err, w.Header().Get("request-id"), r.URL.Path)
	}
	err = json.NewEncoder(w).Encode(&Result{Body: admin})
	if err != nil {
		handler.logger.LogError("problems with marshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
