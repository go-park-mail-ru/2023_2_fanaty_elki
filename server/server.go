package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"

	"./store"
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
}

func (api *Handler) getRestaurantList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Printf("%s: got / request. \n",
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

func (api *Handler) getUserList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Printf("%s: got / request. \n",
		ctx.Value(keyServerAddr),
	)

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
	}
	mux.HandleFunc("/restaurants", api.getRestaurantList)
	mux.HandleFunc("/users", api.getUserList)
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
	err := server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error listening for server: %s\n", err)
	}

}
