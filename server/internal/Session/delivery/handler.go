package delivery

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	sessionUsecase "server/internal/Session/usecase"
	userUsecase "server/internal/User/usecase"
	"server/internal/domain/entity"
	"time"
	"server/internal/domain/dto"
)

const allowedOrigin = "http://84.23.53.216"

type Result struct {
	Body interface{}
}

type Error struct {
	Err string
}

type SessionHandler struct {
	sessions sessionUsecase.UsecaseI
	users userUsecase.UsecaseI	
}

func NewSessionHandler(sessions sessionUsecase.UsecaseI, users userUsecase.UsecaseI) *SessionHandler{
	return &SessionHandler{
		sessions: sessions,
		users: users,
	}
}

// SignUp godoc
// @Summary      Signing up a user
// @Description  Signing up a user
// @Tags        users
// @Accept     application/json
// @Produce  application/json
// @Param 	user	 body	 store.User	 true	 "User object for signing up"
// @Success  200 {object}  integer "success create User return id"
// @Failure 400 {object} error "bad request"
// @Failure 500 {object} error "internal server error"
// @Router   /api/users [post]
func (handler *SessionHandler) SignUp(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")

	jsonbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(&Error{Err: entity.ErrProblemsReadingData.Error()})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	reqUser := dto.ReqCreateUser{}
	err = json.Unmarshal(jsonbody, &reqUser)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(&Error{Err: entity.ErrUnmarshalingJson.Error()})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	id, err := handler.users.CreateUser(dto.ToEntityCreateUser(&reqUser))
	if err != nil {
		if err == entity.ErrInternalServerError {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(&Error{Err: err.Error()})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	
		return
	}

	body := map[string]interface{}{
		"ID": id,
	}

	err = json.NewEncoder(w).Encode(&Result{Body: body})
	if err != nil {
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

	w.Header().Add("Access-Control-Allow-Origin", allowedOrigin)
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Set("content-type", "application/json")

	jsonbody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(&Error{Err: entity.ErrProblemsReadingData.Error()})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	reqUser := dto.ReqLoginUser{}
	err = json.Unmarshal(jsonbody, &reqUser)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(&Error{Err: entity.ErrProblemsReadingData.Error()})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	cookieUC, err := handler.sessions.Login(dto.ToEntityLoginUser(&reqUser))
	if err != nil {
		if err == entity.ErrInternalServerError {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}

		return
	}

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    cookieUC.SessionToken,
		Expires:  time.Now().Add(cookieUC.MaxAge),
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)
	body := map[string]interface{}{
		"Username": reqUser.Username,
	}

	err = json.NewEncoder(w).Encode(&Result{Body: body})
	if err != nil {
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

	w.Header().Add("Access-Control-Allow-Origin", allowedOrigin)
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Set("content-type", "application/json")
	cookie, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		w.WriteHeader(http.StatusUnauthorized)
		err = json.NewEncoder(w).Encode(&Error{Err: entity.ErrUnauthorized.Error()})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = handler.sessions.Logout(&entity.Cookie{
		SessionToken: cookie.Value,
	})	
	if err != nil {
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

	w.Header().Add("Access-Control-Allow-Origin", allowedOrigin)
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Set("content-type", "application/json")

	cookie, err := r.Cookie("session_id")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		err = json.NewEncoder(w).Encode(&Error{Err:entity.ErrUnauthorized.Error()})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	
	username, err := handler.sessions.Check(cookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if username == nil {
		w.WriteHeader(http.StatusUnauthorized)
		err := json.NewEncoder(w).Encode(&Error{Err: entity.ErrUnauthorized.Error()})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	http.SetCookie(w, cookie)
	body := map[string]interface{}{
		"Username": username,
	}
	err = json.NewEncoder(w).Encode(&Result{Body: body})
	
	if err != nil {
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
	w.Header().Add("Access-Control-Allow-Origin", allowedOrigin)
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Set("content-type", "application/json")

	cookie, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		w.WriteHeader(http.StatusUnauthorized)
		err = json.NewEncoder(w).Encode(&Error{Err: entity.ErrUnauthorized.Error()})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	user, err := handler.sessions.GetUserProfile(cookie.Value)
	if err == entity.ErrInternalServerError{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if user == nil {
		w.WriteHeader(http.StatusUnauthorized)
		err = json.NewEncoder(w).Encode(&Error{Err: entity.ErrUnauthorized.Error()})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	err = json.NewEncoder(w).Encode(&Result{Body: user})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}


func (handler *SessionHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", allowedOrigin)
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Set("content-type", "application/json")

	cookie, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		w.WriteHeader(http.StatusUnauthorized)
		err = json.NewEncoder(w).Encode(&Error{Err: entity.ErrUnauthorized.Error()})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	id, err := handler.sessions.GetIdByCookie(cookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if id == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		err := json.NewEncoder(w).Encode(&Error{Err: entity.ErrUnauthorized.Error()})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	jsonbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(&Error{Err: entity.ErrProblemsReadingData.Error()})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	updatedUser := &dto.ReqUpdateUser{}
	err = json.Unmarshal(jsonbody, &updatedUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = handler.users.UpdateUser(dto.ToEntityUpdateUser(updatedUser, id))
	if err != nil {
		if err == entity.ErrInternalServerError {
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else if err == entity.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(&Error{Err: err.Error()})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	
		return
	}

}