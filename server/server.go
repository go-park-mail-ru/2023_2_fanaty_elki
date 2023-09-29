package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"time"

	"server/store"
)

const keyServerAddr = "serverAddr"

type Restaurant struct {
	ID            int
	Name          string
	Rating        float32
	CommentsCount int
}

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

	//store := store.NewRestaurantStore()

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
		w.Write([]byte("{}"))
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
			http.Error(w, `{"error":"db"}`, 500)
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

		jsonbody, err := ioutil.ReadAll(r.Body) // check for errors

		keyVal := make(map[string]string)
		json.Unmarshal(jsonbody, &keyVal) // check for errors

		username := keyVal["username"]
		password := keyVal["password"]
		birthday := keyVal["birthday"]
		phoneNumber := keyVal["phone_number"]
		email := keyVal["email"]
		icon := keyVal["icon"]

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
			http.Error(w, `{"error":"db"}`, 400)
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

	keyVal := make(map[string]string)
	json.Unmarshal(jsonbody, &keyVal) // check for errors

	user, err := api.userstore.FindUser(keyVal["username"])
	if err != nil {
		http.Error(w, `no user`, 404)
		return
	}

	if user.Password != keyVal["password"] {
		http.Error(w, `bad pass`, 400)
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

func getRestaurantById(w http.ResponseWriter, r *http.Request) {

	hasid := r.URL.Query().Has("id")
	id := r.URL.Query().Get("id")
	if hasid {
		io.WriteString(w, "id = "+id)
	} else {
		io.WriteString(w, "no id")
	}

}

func getHello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Printf("%s: got /hello request\n", ctx.Value(keyServerAddr))

	myName := r.PostFormValue("myName")
	if myName == "" {
		w.Header().Set("x-missing-field", "myName")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	io.WriteString(w, fmt.Sprintf("Hello, %s!\n", myName))
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
	mux.HandleFunc("/hello", getHello)
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
