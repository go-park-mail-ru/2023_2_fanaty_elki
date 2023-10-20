package delivery

import (
	"server/internal/usecases"
	"net/http"
	"encoding/json"
)

type RestaurantHandler struct {
	restaurants usecases.RestaurantUsecase
}

func NewRestaurantHandler(restaurants *usecases.RestaurantUsecase) *RestaurantHandler{
	return &RestaurantHandler{restaurants: *restaurants}
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
func (api *RestaurantHandler) GetRestaurantList(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", allowedOrigin)
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Set("content-type", "application/json")

	rests, err := api.restaurants.GetRestaurants()

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