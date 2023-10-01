package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"time"
	"server/store"
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

	rests, err := api.restaurantstore.GetRestaurants()

	if err != nil {
		http.Error(w, `{"error":"db"}`, 500)
		return
	}

	body := map[string]interface{}{
		"restaurants": rests,
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(&Result{Body: body})
	if err != nil {
		log.Printf("error while marshalling JSON: %s", err)
		http.Error(w, `{"error":"error while marshalling JSON"}`, 500)
		return
	}
}

func (api *Handler) User(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Printf("%s: got /users request. \n",
		ctx.Value(keyServerAddr),
	)

	if r.Method == "GET" {
		users, err := api.userstore.GetUsers()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(&Result{Err: "problems with reading data"})
			return
		}

		body := map[string]interface{}{
			"users": users,
		}

		encoder := json.NewEncoder(w)
		err = encoder.Encode(&Result{Body: body})
		if err != nil {
			log.Printf("error while marshalling JSON: %s", err)
			w.Write([]byte("{}"))
			return
		}
	}

	if r.Method == "POST" {

		jsonbody, err := ioutil.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(&Result{Err: "problems with reading data"})
			return
		}

		keyVal := make(map[string]string)
		err = json.Unmarshal(jsonbody, &keyVal)

		if err != nil {
			//http.Error(w, `{"error":"problems with unmarshaling json"}`, 500)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(&Result{Err: "problems with unmarshaling json"})
			return
		}

		username := keyVal["username"]
		password := keyVal["password"]
		birthday := keyVal["birthday"]
		phoneNumber := keyVal["phone_number"]
		email := keyVal["email"]
		icon := keyVal["icon"]

		user, err := api.userstore.FindUserBy("username", keyVal["username"])

		if user != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&Result{Err: "username already exists"})
			return
		}

		user, err = api.userstore.FindUserBy("phone_number", keyVal["phone_number"])
		if user != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&Result{Err: "phone number already exists"})
			return
		}

		user, err = api.userstore.FindUserBy("email", keyVal["email"])
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
			Icon:        icon,
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

	jsonbody, err := ioutil.ReadAll(r.Body) // check for errors

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&Result{Err: "problems with reading data"})
		return
	}

	keyVal := make(map[string]string)
	json.Unmarshal(jsonbody, &keyVal) // check for errors

	user, err := api.userstore.FindUser(keyVal["username"])
	if err != nil {
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
		Name:    "session_id",
		Value:   SID,
		Expires: time.Now().Add(10 * time.Hour),
	}
	http.SetCookie(w, cookie)
	w.Write([]byte(SID))
}

func (api *Handler) Logout(w http.ResponseWriter, r *http.Request) {

	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&Result{Err: "problems with cookies"})
		return
	}
	// Здесь не уверен
	if _, ok := api.sessions[session.Value]; !ok {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&Result{Err: "problems with cookies"})
		return
	}

	delete(api.sessions, session.Value)

	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)
}

func (api *Handler) Auth(w http.ResponseWriter, r *http.Request) {

	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&Result{Err: "problems with authorizing"})
		return
	}

	id, ok := api.sessions[session.Value]
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&Result{Err: "problems with authorizing"})
		return
	}

	http.SetCookie(w, session)
	body := map[string]interface{}{
		"id": id,
	}
	json.NewEncoder(w).Encode(&Result{Body: body})
}

func main() {
	mux := http.NewServeMux()
	api := &Handler{
		restaurantstore: store.NewRestaurantStore(),
		userstore:       store.NewUserStore(),
		sessions:        make(map[string]uint, 10),
	}
	mux.HandleFunc("/restaurants", api.getRestaurantList)
	mux.HandleFunc("/users", api.User)
	mux.HandleFunc("/login", api.Login)
	mux.HandleFunc("/logout", api.Logout)
	mux.HandleFunc("/auth", api.Auth)
	ctx := context.Background()

	server := &http.Server{
		Addr:    ":3333",
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
