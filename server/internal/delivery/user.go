package delivery

import (
	"server/internal/domain/entity"
	"server/internal/usecases"
	//"server/repository"
	"database/sql"
	"encoding/json"
//	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"

	//"github.com/gomodule/redigo/redis"
	_ "github.com/lib/pq"
)

type UserHandler struct {
	users usecases.UserUsecase
	sessManager  usecases.SessionUsecase
}

func NewUserHandler(users *usecases.UserUsecase) *UserHandler{
	return &UserHandler{users: *users}
}



// type Handler struct {
// 	restaurantstore *repository.RestaurantRepo
// 	userstore       *repository.UserRepo
// 	sessManager     *repository.SessionManager
// }

func (api *UserHandler) checkSession(r *http.Request) (*entity.Session, error) {
	cookieSessionID, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	sess, err := api.sessManager.Check(&entity.SessionID{
		ID: cookieSessionID.Value,
	})
	if err != nil {
		return nil, err
	}
	return sess, nil
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
// @Router   /users [post]
func (api *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {

	jsonbody, err := ioutil.ReadAll(r.Body)

	w.Header().Set("content-type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(&Error{Err: "problems with reading data"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	keyVal := make(map[string]string)
	err = json.Unmarshal(jsonbody, &keyVal)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(&Error{Err: "problems with unmarshaling json"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	username := keyVal["Username"]
	password := keyVal["Password"]
	birthday := keyVal["Birthday"]
	email := keyVal["Email"]
	icon := keyVal["Icon"]
	phoneNumber := keyVal["PhoneNumber"]

	if len(username) < 3 {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(&Error{Err: "username is too short"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	if len(username) > 30 {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(&Error{Err: "username is too long"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	if len(password) < 3 {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(&Error{Err: "password is too short"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	if len(password) > 20 {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(&Error{Err: "password is too long"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	re := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
	if birthday != "" && !re.MatchString(birthday) {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(&Error{Err: "incorrect birthday"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	re = regexp.MustCompile(`^[+]?[0-9]{3,25}$`)
	if phoneNumber != "" && !re.MatchString(phoneNumber) {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(&Error{Err: "incorrect phone number"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	re = regexp.MustCompile(`\S*@\S*`)
	if !re.MatchString(email) {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(&Error{Err: "incorrect email"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	user, err := api.users.FindUserBy("Username", keyVal["Username"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if user != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(&Error{Err: "username already exists"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	user, err = api.users.FindUserBy("Email", keyVal["Email"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if user != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(&Error{Err: "email already exists"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	user, err = api.users.FindUserBy("PhoneNumber", keyVal["PhoneNumber"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if user != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(&Error{Err: "phone number already exists"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	var birthdayString sql.NullString
	if birthday != "" {
		birthdayString = sql.NullString{String: birthday, Valid: true}
	} else {
		birthdayString = sql.NullString{Valid: false}
	}

	var iconString sql.NullString
	if icon != "" {
		iconString = sql.NullString{String: icon, Valid: true}
	} else {
		iconString = sql.NullString{Valid: false}
	}


	in := &entity.User{
		Username:    username,
		Password:    password,
		Birthday:    birthdayString,
		PhoneNumber: phoneNumber,
		Email:       email,
		Icon:        iconString,
	}

	id, err := api.users.CreateUser(in)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body := map[string]interface{}{
		"ID": id,
	}

	err = json.NewEncoder(w).Encode(&Result{Body: body})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
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
// @Router   /login [post]
func (api *UserHandler) Login(w http.ResponseWriter, r *http.Request) {

	jsonbody, err := ioutil.ReadAll(r.Body)

	w.Header().Add("Access-Control-Allow-Origin", allowedOrigin)
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Set("content-type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(&Error{Err: "problems with reading data"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	keyVal := make(map[string]string)
	err = json.Unmarshal(jsonbody, &keyVal)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(&Error{Err: "problems with unmarshaling json"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	user, err := api.users.FindUserBy("Username", keyVal["Username"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if user == nil {
		w.WriteHeader(http.StatusNotFound)
		err = json.NewEncoder(w).Encode(&Error{Err: "user not found"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	if user.Password != keyVal["Password"] {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(&Error{Err: "incorrect password"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	sess, err := api.sessManager.Create(&entity.Session{
		UserID:    user.ID,
		Useragent: r.UserAgent(),
	})
	if err != nil {
		fmt.Println("can't create session:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    sess.ID,
		Expires:  time.Now().Add(50 * time.Hour),
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)
	body := map[string]interface{}{
		"Username": user.Username,
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
// @Router   /logout [get]
func (api *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", allowedOrigin)
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Set("content-type", "application/json")
	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		w.WriteHeader(http.StatusUnauthorized)
		err = json.NewEncoder(w).Encode(&Error{Err: "unauthorized"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	api.sessManager.Delete(&entity.SessionID{
		ID: session.Value,
	})

	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)
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
// @Router   /auth [get]
func (api *UserHandler) Auth(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", allowedOrigin)
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Set("content-type", "application/json")

	sess, err := r.Cookie("session_id")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		err = json.NewEncoder(w).Encode(&Error{Err: "unauthorized"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	session, err := api.checkSession(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if session == nil {
		w.WriteHeader(http.StatusUnauthorized)
		err := json.NewEncoder(w).Encode(&Error{Err: "unauthorized"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	user, err := api.users.GetUserById(session.UserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, sess)
	body := map[string]interface{}{
		"Username": user.Username,
	}
	err = json.NewEncoder(w).Encode(&Result{Body: body})
	
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

