package delivery

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	cartUsecase "server/internal/Cart/usecase"
	"server/internal/domain/dto"
	"server/internal/domain/entity"
	"github.com/gorilla/mux"
)

type Result struct {
	Body interface{}
}

type Error struct {
	Err string
}

type CartHandler struct {
	cartUsecase cartUsecase.UsecaseI
}

func NewCartHandler(cartUsecase cartUsecase.UsecaseI) *CartHandler {
	return &CartHandler{cartUsecase: cartUsecase}
}

func (handler *CartHandler) RegisterHandler(router *mux.Router) {
	router.HandleFunc("/api/cart", handler.GetCart).Methods(http.MethodGet)
	router.HandleFunc("/api/cart/add", handler.AddProductToCart).Methods(http.MethodPost)
	router.HandleFunc("/api/cart/delete", handler.DeleteProductFromCart).Methods(http.MethodPost)
	router.HandleFunc("/api/cart/update/up", handler.UpdateItemCountUp).Methods(http.MethodPatch)
	router.HandleFunc("/api/cart/update/down", handler.UpdateItemCountDown).Methods(http.MethodPatch)
}

func (handler *CartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	cookie, err := r.Cookie("session_id")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		err = json.NewEncoder(w).Encode(&Error{Err: entity.ErrUnauthorized.Error()})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	cart, err := handler.cartUsecase.GetUserCart(cookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(&Error{Err: "data base error"})
		return
	}

	body := map[string]interface{}{
		"Cart": cart,
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(&Result{Body: body})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(&Error{Err: "error while marshalling JSON"})
		return
	}
}

func (handler *CartHandler) AddProductToCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	jsonbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reqProduct := dto.ReqProductID{}
	err = json.Unmarshal(jsonbody, &reqProduct)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cookie, err := r.Cookie("session_id")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = handler.cartUsecase.AddProductToCart(cookie.Value, reqProduct.ProductID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (handler *CartHandler) DeleteProductFromCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	jsonbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reqProduct := dto.ReqProductID{}
	err = json.Unmarshal(jsonbody, &reqProduct)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cookie, err := r.Cookie("session_id")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = handler.cartUsecase.DeleteProductFromCart(cookie.Value, reqProduct.ProductID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *CartHandler) UpdateItemCountUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	jsonbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reqProduct := dto.ReqProductID{}
	err = json.Unmarshal(jsonbody, &reqProduct)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cookie, err := r.Cookie("session_id")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = handler.cartUsecase.UpdateItemCountUp(cookie.Value, reqProduct.ProductID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *CartHandler) UpdateItemCountDown(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	jsonbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reqProduct := dto.ReqProductID{}
	err = json.Unmarshal(jsonbody, &reqProduct)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cookie, err := r.Cookie("session_id")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = handler.cartUsecase.UpdateItemCountDown(cookie.Value, reqProduct.ProductID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
