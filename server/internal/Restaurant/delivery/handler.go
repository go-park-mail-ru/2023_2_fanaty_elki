package delivery

import (
	"encoding/json"
	"errors"
	"net/http"
	restaurantUsecase "server/internal/Restaurant/usecase"
	"server/internal/domain/entity"
	mw "server/internal/middleware"
	"strconv"

	"github.com/gorilla/mux"
)

type Result struct {
	Body interface{}
}

type RespError struct {
	Err string
}

type RestaurantHandler struct {
	restaurants restaurantUsecase.UsecaseI
	logger      *mw.ACLog
}

func NewRestaurantHandler(restaurants restaurantUsecase.UsecaseI, logger *mw.ACLog) *RestaurantHandler {
	return &RestaurantHandler{
		restaurants: restaurants,
		logger:      logger,
	}
}

func (handler *RestaurantHandler) RegisterHandler(router *mux.Router) {
	router.HandleFunc("/api/restaurants", handler.GetRestaurantList).Methods(http.MethodGet)
	router.HandleFunc("/api/restaurants/tips", handler.GetRestaurantTipList).Methods(http.MethodGet)
	router.HandleFunc("/api/restaurants/{id:[0-9]+}", handler.GetRestaurantById).Methods(http.MethodGet)
	router.HandleFunc("/api/restaurants/{id}/products", handler.GetRestaurantProducts).Methods(http.MethodGet)
	router.HandleFunc("/api/restaurants/{category}", handler.GetRestaurantListByCategory).Methods(http.MethodGet)
	router.HandleFunc("/api/categories", handler.GetCategoryList).Methods(http.MethodGet)
	router.HandleFunc("/api/restaurants/", handler.Search).Methods(http.MethodGet)
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

	w.Header().Set("Content-Type", "application/json")

	rests, err := handler.restaurants.GetRestaurants()

	if err != nil {
		handler.logger.LogError("problems with getting restauratns", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body := rests

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
	w.Header().Set("Content-Type", "application/json")

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
		if err == entity.ErrNotFound {
			handler.logger.LogError("problems restaurants id", err, w.Header().Get("request-id"), r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		handler.logger.LogError("problems restaurants id", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body := rest

	encoder := json.NewEncoder(w)
	err = encoder.Encode(&Result{Body: body})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		handler.logger.LogError("problems with marshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		return
	}
}

func (handler *RestaurantHandler) GetRestaurantProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

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

	menu, err := handler.restaurants.GetRestaurantProducts(id)

	if err != nil {
		if err == entity.ErrNotFound {
			handler.logger.LogError("problems restaurants id", err, w.Header().Get("request-id"), r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		handler.logger.LogError("problems restaurants id", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body := menu

	encoder := json.NewEncoder(w)
	err = encoder.Encode(&Result{Body: body})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		handler.logger.LogError("problems with marshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		return
	}
}

func (handler *RestaurantHandler) GetRestaurantListByCategory(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	categoryname, ok := vars["category"]
	if !ok {
		handler.logger.LogError("problems with parameters", errors.New("category is missing in parameters"), w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rests, err := handler.restaurants.GetRestaurantsByCategory(categoryname)

	if err != nil {
		if err == entity.ErrNotFound {
			handler.logger.LogError("problems restaurants id", err, w.Header().Get("request-id"), r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		handler.logger.LogError("problems with getting restauratns by category", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body := rests

	encoder := json.NewEncoder(w)
	err = encoder.Encode(&Result{Body: body})

	if err != nil {
		handler.logger.LogError("problems with marshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *RestaurantHandler) GetCategoryList(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	cats, err := handler.restaurants.GetCategories()

	if err != nil {
		handler.logger.LogError("problems with getting categories", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body := cats

	encoder := json.NewEncoder(w)
	err = encoder.Encode(&Result{Body: body})

	if err != nil {
		handler.logger.LogError("problems with marshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *RestaurantHandler) Search(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	search := r.URL.Query().Get("search")
	if search == "" {
		handler.logger.LogError("problems with parameters", errors.New("missing parameters"), w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rests, err := handler.restaurants.Search(search)
	if err != nil {
		if err == entity.ErrNotFound {
			handler.logger.LogError("problems restaurants id", err, w.Header().Get("request-id"), r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		handler.logger.LogError("problems with searching", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body := rests

	encoder := json.NewEncoder(w)
	err = encoder.Encode(&Result{Body: body})

	if err != nil {
		handler.logger.LogError("problems with marshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *RestaurantHandler) GetRestaurantTipList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cookie, _ := r.Cookie("session_id")
	rests, err := handler.restaurants.GetRestaurantTips(cookie.Value)

	if err != nil {
		handler.logger.LogError("problems with getting restaurant tips", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body := rests

	encoder := json.NewEncoder(w)
	err = encoder.Encode(&Result{Body: body})

	if err != nil {
		handler.logger.LogError("problems with marshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
