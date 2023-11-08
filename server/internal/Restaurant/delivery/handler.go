package delivery

import (
	"encoding/json"
	"net/http"
	restaurantUsecase "server/internal/Restaurant/usecase"
	"strconv"
	"errors"
	"github.com/gorilla/mux"
	mw "server/internal/middleware"
)

type Result struct {
	Body interface{}
}

type RespError struct {
	Err string
}

type RestaurantHandler struct {
	restaurants restaurantUsecase.UsecaseI
	logger *mw.ACLog
}

func NewRestaurantHandler(restaurants restaurantUsecase.UsecaseI, logger *mw.ACLog) *RestaurantHandler {
	return &RestaurantHandler{
		restaurants: restaurants,
		logger: logger,
	}
}

func (handler *RestaurantHandler) RegisterHandler(router *mux.Router) {
	router.HandleFunc("/api/restaurants", handler.GetRestaurantList).Methods(http.MethodGet)
	router.HandleFunc("/api/restaurants/{id}", handler.GetRestaurantById).Methods(http.MethodGet)
}

// GetRestaurantsList godoc
// @Summary      giving restaurats
// @Description  giving array of restaurants
// @Tags        Restaurants
// @Accept     */*
// @Produce  application/json
// @Success  200 {object}  []entity.Restaurant "success returning array of restaurants"
// @Failure 500 {object} error "internal server error"
// @Router   /restaurants [get]
func (handler *RestaurantHandler) GetRestaurantList(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")

	rests, err := handler.restaurants.GetRestaurants()

	if err != nil {
		handler.logger.LogError("problems with getting restauratns", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body := map[string]interface{}{
		"restaurants": rests,
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(&Result{Body: body})

	if err != nil {
		handler.logger.LogError("problems with marshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// GetRestaurantById godoc
// @Summary      giving information about restaurant and its products
// @Description  giving restaurant object and array of menu types with array of products in each menu type
// @Tags        Restaurants
// @Accept     */{id}
// @Produce  application/json
// @Success  200 {object}  dto.RestaurantWithProducts "success returning information about restaurant"
// @Failure 500 {object} error "internal server error"
// @Router   /restaurants/{id} [get]
func (handler *RestaurantHandler) GetRestaurantById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	vars := mux.Vars(r)
	strid, ok := vars["id"]
	if !ok {
		handler.logger.LogError("problems with parameters", errors.New("id is missing in parameters"), w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id64, err := strconv.ParseUint(strid, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		handler.logger.LogError("problems with parameters", errors.New("id is not number"), w.Header().Get("request-id"), r.URL.Path)
		err = json.NewEncoder(w).Encode(&RespError{Err: "id is not a number"})
		return
	}

	id := uint(id64)

	rest, err := handler.restaurants.GetRestaurantById(id)

	if err != nil {
		handler.logger.LogError("problems restaurants id", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body := map[string]interface{}{
		"RestaurantWithProducts": rest,
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(&Result{Body: body})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		handler.logger.LogError("problems with marshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		return
	}
}

