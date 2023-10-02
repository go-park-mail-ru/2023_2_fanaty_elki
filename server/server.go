package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"regexp"
	"server/store"
	"time"
)

const keyServerAddr = "serverAddr"

type Result struct {
	Body interface{}
	Err  string
}

type Handler struct {
	restaurantstore *store.RestaurantStore
	userstore       *store.UserStore
	sessions        map[string]uint
}

var (
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func (api *Handler) getRestaurantList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Printf("%s: got /restaurants request. \n",
		ctx.Value(keyServerAddr),
	)

	w.Header().Set("content-type", "application/json")

	rests, err := api.restaurantstore.GetRestaurants()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&Result{Err: "data base error"})
		return
	}
	body := map[string]interface{}{
		"restaurants": rests,
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(&Result{Body: body})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&Result{Err: "error while marshalling JSON"})
		return
	}
}

func (api *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Printf("%s: got /users request. \n",
		ctx.Value(keyServerAddr),
	)

	if r.Method == "POST" {

		jsonbody, err := ioutil.ReadAll(r.Body)

		w.Header().Set("content-type", "application/json")

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(&Result{Err: "problems with reading data"})
			return
		}

		keyVal := make(map[string]string)
		err = json.Unmarshal(jsonbody, &keyVal)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(&Result{Err: "problems with unmarshaling json"})
			return
		}

		username := keyVal["username"]
		password := keyVal["password"]
		birthday := keyVal["birthday"]
		phoneNumber := keyVal["phone_number"]
		email := keyVal["email"]

		if len(username) < 3 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&Result{Err: "username is too short"})
			return
		}

		if len(username) > 30 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&Result{Err: "username is too long"})
			return
		}

		if len(password) < 3 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&Result{Err: "password is too short"})
			return
		}

		if len(password) > 20 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&Result{Err: "password is too long"})
			return
		}

		re := regexp.MustCompile(`\d{2}-\d{2}-\d{4}`)
		if birthday != "" && !re.MatchString(birthday) {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&Result{Err: "incorrect birthday"})
			return
		}

		re = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		if !re.MatchString(email) {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&Result{Err: "incorrect email"})
			return
		}

		re = regexp.MustCompile(`^[0-9\-\+]{9,15}$`)
		if !re.MatchString(phoneNumber) {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&Result{Err: "incorrect phone"})
			return
		}

		user := api.userstore.FindUserBy("username", keyVal["username"])
		if user != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&Result{Err: "username already exists"})
			return
		}

		user = api.userstore.FindUserBy("phone_number", keyVal["phone_number"])
		if user != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&Result{Err: "phone number already exists"})
			return
		}

		user = api.userstore.FindUserBy("email", keyVal["email"])
		if user != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&Result{Err: "email already exists"})
			return
		}

		in := &store.User{
			Username:    username,
			Password:    password,
			Birthday:    birthday,
			PhoneNumber: phoneNumber,
			Email:       email,
		}

		id, err := api.userstore.SignUpUser(in)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&Result{Err: "error while adding user"})
			return
		}

		body := map[string]interface{}{
			"id": id,
		}
		json.NewEncoder(w).Encode(&Result{Body: body})
	}

}

func (api *Handler) Login(w http.ResponseWriter, r *http.Request) {

	jsonbody, err := ioutil.ReadAll(r.Body)

	w.Header().Set("content-type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&Result{Err: "problems with reading data"})
		return
	}

	keyVal := make(map[string]string)
	err = json.Unmarshal(jsonbody, &keyVal)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&Result{Err: "problems with unmarshaling json"})
		return
	}

	user := api.userstore.FindUserBy("username", keyVal["username"])
	if user == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(&Result{Err: "user not found"})
		return
	}

	if user.Password != keyVal["password"] {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&Result{Err: "incorrect password"})
		return
	}

	SID := RandStringRunes(32)

	api.sessions[SID] = user.ID

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    SID,
		Expires:  time.Now().Add(10 * time.Hour),
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	body := map[string]interface{}{
		"cookie": cookie.Value,
	}
	json.NewEncoder(w).Encode(&Result{Body: body})

}

func (api *Handler) Logout(w http.ResponseWriter, r *http.Request) {

	session, err := r.Cookie("session_id")
	w.Header().Set("content-type", "application/json")
	if err == http.ErrNoCookie {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&Result{Err: "unauthorized"})
		return
	}
	if _, ok := api.sessions[session.Value]; !ok {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&Result{Err: "unauthorized"})
		return
	}

	delete(api.sessions, session.Value)

	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)
}

func (api *Handler) Auth(w http.ResponseWriter, r *http.Request) {

	session, err := r.Cookie("session_id")
	w.Header().Set("content-type", "application/json")
	if err == http.ErrNoCookie {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&Result{Err: "unauthorized"})
		return
	}

	id, ok := api.sessions[session.Value]
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&Result{Err: "unauthorized"})
		return
	}

	http.SetCookie(w, session)
	body := map[string]interface{}{
		"id": id,
	}
	json.NewEncoder(w).Encode(&Result{Body: body})
}

const PORT = ":3333"

func main() {
	mux := http.NewServeMux()
	api := &Handler{
		restaurantstore: store.NewRestaurantStore(),
		userstore:       store.NewUserStore(),
		sessions:        make(map[string]uint, 10),
	}
	mux.HandleFunc("/restaurants", api.getRestaurantList)
	mux.HandleFunc("/users", api.SignUp)
	mux.HandleFunc("/login", api.Login)
	mux.HandleFunc("/logout", api.Logout)
	mux.HandleFunc("/auth", api.Auth)
	ctx := context.Background()

	server := &http.Server{
		Addr:    PORT,
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
			return ctx
		},
	}

	fmt.Println("Server start")
	err := server.ListenAndServe()

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")

	} else if err != nil {
		fmt.Printf("error listening for server: %s\n", err)
	}

}
