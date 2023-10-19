package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"server/store"
	"time"

	"github.com/gomodule/redigo/redis"
	_ "github.com/lib/pq"
)

// @title Prinesi-Poday API
// @version 1.0
// @license.name Apache 2.0
// @host http://84.23.53.216:8001/
const keyServerAddr = "serverAddr"
const allowedOrigin = "http://84.23.53.216"

var (
	redisAddr = flag.String("addr", "redis://user:@localhost:6379/0", "redis addr")
)

type Result struct {
	Body interface{}
}

type Error struct {
	Err string
}

type Handler struct {
	restaurantstore *store.RestaurantRepo
	userstore       *store.UserRepo
	sessManager     *store.SessionManager
}

// var (
// 	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
// )

// func randStringRunes(n int) string {
// 	b := make([]rune, n)
// 	for i := range b {
// 		b[i] = letterRunes[rand.Intn(len(letterRunes))]
// 	}
// 	return string(b)
// }

func (api *Handler) checkSession(r *http.Request) *store.Session {
	cookieSessionID, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		return nil
	} else if err != nil {
		return nil
	}

	sess := api.sessManager.Check(&store.SessionID{
		ID: cookieSessionID.Value,
	})
	return sess
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

	username := keyVal["Username"]
	password := keyVal["Password"]
	birthday := keyVal["Birthday"]
	phoneNumber := keyVal["PhoneNumber"]
	email := keyVal["Email"]
	icon := keyVal["Icon"]

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

	user := api.userstore.FindUserBy("Username", keyVal["Username"])
	if user != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(&Error{Err: "username already exists"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	user = api.userstore.FindUserBy("Email", keyVal["Email"])
	if user != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(&Error{Err: "email already exists"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	user = api.userstore.FindUserBy("PhoneNumber", keyVal["PhoneNumber"])
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

	in := &store.User{
		Username:    username,
		Password:    password,
		Birthday:    birthdayString,
		PhoneNumber: phoneNumber,
		Email:       email,
		Icon:        iconString,
	}

	id := api.userstore.SignUpUser(in)

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

	user := api.userstore.FindUserBy("Username", keyVal["Username"])
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

	//SID := randStringRunes(32)

	//api.sessions[SID] = user.ID
	sess, err := api.sessManager.Create(&store.Session{
		UserID:    user.ID,
		Useragent: r.UserAgent(),
	})
	if err != nil {
		fmt.Println("cant create session:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    sess.ID,
		Expires:  time.Now().Add(10 * time.Hour),
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
func (api *Handler) Logout(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", allowedOrigin)
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Set("content-type", "application/json")
	session, err := r.Cookie("session_id")
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// }
	// if session == nil {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	err = json.NewEncoder(w).Encode(&Error{Err: "unauthorized"})
	// 	if err != nil {
	// 		w.WriteHeader(http.StatusInternalServerError)
	// 	}
	// 	return
	// }
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

	api.sessManager.Delete(&store.SessionID{
		ID: session.Value,
	})
	// if _, ok := api.sessions[session.Value]; !ok {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	err = json.NewEncoder(w).Encode(&Error{Err: "unauthorized"})
	// 	if err != nil {
	// 		w.WriteHeader(http.StatusInternalServerError)
	// 	}
	// 	return
	// }

	//delete(api.sessions, session.Value)
	//api.sessManager.Delete(session)
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
	//session, err := r.Cookie("session_id")
	w.Header().Set("content-type", "application/json")
	// if err == http.ErrNoCookie {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	err = json.NewEncoder(w).Encode(&Error{Err: "unauthorized"})
	// 	if err != nil {
	// 		w.WriteHeader(http.StatusInternalServerError)
	// 	}
	// 	return
	// }

	// id, ok := api.sessions[session.Value]
	// if !ok {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	err = json.NewEncoder(w).Encode(&Error{Err: "unauthorized"})
	// 	if err != nil {
	// 		w.WriteHeader(http.StatusInternalServerError)
	// 	}
	// 	return
	// }
	sess, err := r.Cookie("session_id")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		err = json.NewEncoder(w).Encode(&Error{Err: "unauthorized"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	session := api.checkSession(r)
	if session == nil {
		w.WriteHeader(http.StatusUnauthorized)
		err := json.NewEncoder(w).Encode(&Error{Err: "unauthorized"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	user := api.userstore.GetUserById(session.UserID)
	http.SetCookie(w, sess)
	body := map[string]interface{}{
		"Username": user.Username,
	}
	err = json.NewEncoder(w).Encode(&Result{Body: body})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

const PORT = ":3333"

func main() {
	flag.Parse()

	var err error
	redisConn, err := redis.DialURL(*redisAddr)
	if err != nil {
		log.Fatalf("cant connect to redis")
	}

	mux := http.NewServeMux()

	db := store.GetPostgres()
	defer db.Close()
	api := &Handler{
		restaurantstore: store.NewRestaurantRepo(db),
		userstore:       store.NewUserRepo(db),
		sessManager:     store.NewSessionManager(redisConn),
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
	err = server.ListenAndServe()

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")

	} else if err != nil {
		fmt.Printf("error listening for server: %s\n", err)
	}
}
