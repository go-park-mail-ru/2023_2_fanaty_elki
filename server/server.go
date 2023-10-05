package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"regexp"
	"server/store"
	"time"
)

// @title Prinesi-Poday API
// @version 1.0
// @license.name Apache 2.0
// @host http://84.23.53.216:8001/
const keyServerAddr = "serverAddr"
const allowedOrigin = "http://84.23.53.216"

type Result struct {
	Body interface{}
}

type Error struct {
	Err string
}

type Handler struct {
	restaurantstore *store.RestaurantStore
	userstore       *store.UserStore
	sessions        map[string]uint
}

var (
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// GetRestaurants godoc
// @Summary      giving restaurats
// @Description  giving array of restaurants
// @Tags        Restaurants
// @Accept     */*
// @Produce  application/json
// @Success  200 {object}  []store.Restaurant "success returning array of restaurants"
// @Failure 500 {object} error "internal server error"
// @Router   /restaurants [get]
func (api *Handler) GetRestaurantList(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", allowedOrigin)
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Set("content-type", "application/json")

	rests, err := api.restaurantstore.GetRestaurants()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(&Error{Err: "data base error"})
		return
	}
	body := map[string]interface{}{
		"restaurants": rests,
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(&Result{Body: body})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(&Error{Err: "error while marshalling JSON"})
		return
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
// @Router   /users [post]
func (api *Handler) SignUp(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		return
	}

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

	username := keyVal["username"]
	password := keyVal["password"]
	birthday := keyVal["birthday"]
	email := keyVal["email"]

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

	re := regexp.MustCompile(`\d{2}-\d{2}-\d{4}`)
	if birthday != "" && !re.MatchString(birthday) {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(&Error{Err: "incorrect birthday"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	re = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !re.MatchString(email) {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(&Error{Err: "incorrect email"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	user := api.userstore.FindUserBy("username", keyVal["username"])
	if user != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(&Error{Err: "username already exists"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	user = api.userstore.FindUserBy("email", keyVal["email"])
	if user != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(&Error{Err: "email already exists"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	in := &store.User{
		Username: username,
		Password: password,
		Birthday: birthday,
		Email:    email,
	}

	id := api.userstore.SignUpUser(in)

	body := map[string]interface{}{
		"id": id,
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
func (api *Handler) Login(w http.ResponseWriter, r *http.Request) {

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

	user := api.userstore.FindUserBy("username", keyVal["username"])
	if user == nil {
		w.WriteHeader(http.StatusNotFound)
		err = json.NewEncoder(w).Encode(&Error{Err: "user not found"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	if user.Password != keyVal["password"] {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(&Error{Err: "incorrect password"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	SID := randStringRunes(32)

	api.sessions[SID] = user.ID

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    SID,
		Expires:  time.Now().Add(10 * time.Hour),
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	body := map[string]interface{}{
		"cookie":   cookie.Value,
		"username": user.Username,
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
func (api *Handler) Logout(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", allowedOrigin)
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	session, err := r.Cookie("session_id")
	w.Header().Set("content-type", "application/json")
	if err == http.ErrNoCookie {
		w.WriteHeader(http.StatusUnauthorized)
		err = json.NewEncoder(w).Encode(&Error{Err: "unauthorized"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	if _, ok := api.sessions[session.Value]; !ok {
		w.WriteHeader(http.StatusUnauthorized)
		err = json.NewEncoder(w).Encode(&Error{Err: "unauthorized"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	delete(api.sessions, session.Value)

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
func (api *Handler) Auth(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", allowedOrigin)
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	session, err := r.Cookie("session_id")
	w.Header().Set("content-type", "application/json")
	if err == http.ErrNoCookie {
		w.WriteHeader(http.StatusUnauthorized)
		err = json.NewEncoder(w).Encode(&Error{Err: "unauthorized"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	id, ok := api.sessions[session.Value]
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		err = json.NewEncoder(w).Encode(&Error{Err: "unauthorized"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	user := api.userstore.GetUserById(id - 1)
	http.SetCookie(w, session)
	body := map[string]interface{}{
		"username": user.Username,
	}
	err = json.NewEncoder(w).Encode(&Result{Body: body})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

const PORT = ":3333"

func main() {
	mux := http.NewServeMux()
	api := &Handler{
		restaurantstore: store.NewRestaurantStore(),
		userstore:       store.NewUserStore(),
		sessions:        make(map[string]uint, 10),
	}
	mux.HandleFunc("/restaurants", api.GetRestaurantList)
	mux.HandleFunc("/users", api.SignUp)
	mux.HandleFunc("/login", api.Login)
	mux.HandleFunc("/logout", api.Logout)
	mux.HandleFunc("/auth", api.Auth)

	server := &http.Server{
		Addr:    PORT,
		Handler: mux,
	}

	fmt.Println("Server start")
	err := server.ListenAndServe()

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")

	} else if err != nil {
		fmt.Printf("error listening for server: %s\n", err)
	}
}
