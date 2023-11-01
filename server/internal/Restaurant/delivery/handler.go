package delivery

import (
	"encoding/json"
	"fmt"
	"net/http"
	restaurantUsecase "server/internal/Restaurant/usecase"
	"strconv"

	"github.com/gorilla/mux"
)

const allowedOrigin = "http://84.23.53.216"

type Result struct {
	Body interface{}
}

type Error struct {
	Err string
}

type RestaurantHandler struct {
	restaurants restaurantUsecase.UsecaseI
}

func NewRestaurantHandler(restaurants restaurantUsecase.UsecaseI) *RestaurantHandler {
	return &RestaurantHandler{restaurants: restaurants}
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
func (handler *RestaurantHandler) GetRestaurantList(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", allowedOrigin)
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Set("content-type", "application/json")

	rests, err := handler.restaurants.GetRestaurants()

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

func (handler *RestaurantHandler) GetRestaurantById(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", allowedOrigin)
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Set("content-type", "application/json")

	vars := mux.Vars(r)
	strid, ok := vars["id"]
	if !ok {
		fmt.Println("id is missing in parameters")
	}

	id64, err := strconv.ParseUint(strid, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(&Error{Err: "id is not a number"})
		return
	}

	id := uint(id64)

	rest, err := handler.restaurants.GetRestaurantById(id)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(&Error{Err: "data base error"})
		return
	}

	body := map[string]interface{}{
		"RestaurantWithProducts": rest,
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(&Result{Body: body})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(&Error{Err: "error while marshalling JSON"})
		return
	}
}
