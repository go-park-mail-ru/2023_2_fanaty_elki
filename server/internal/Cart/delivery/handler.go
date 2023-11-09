package delivery

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	cartUsecase "server/internal/Cart/usecase"
	"server/internal/domain/dto"
	mw "server/internal/middleware"
)

type Result struct {
	Body interface{}
}

type RespError struct {
	Err string
}

type CartHandler struct {
	cartUsecase cartUsecase.UsecaseI
	logger      *mw.ACLog
}

func NewCartHandler(cartUsecase cartUsecase.UsecaseI, logger *mw.ACLog) *CartHandler {
	return &CartHandler{
		cartUsecase: cartUsecase,
		logger:      logger,
	}
}

func (handler *CartHandler) RegisterHandler(router *mux.Router) {
	router.HandleFunc("/api/cart", handler.GetCart).Methods(http.MethodGet)
	router.HandleFunc("/api/cart/add", handler.AddProductToCart).Methods(http.MethodPost)
	router.HandleFunc("/api/cart/delete", handler.DeleteProductFromCart).Methods(http.MethodPost)
	router.HandleFunc("/api/cart/update/up", handler.UpdateItemCountUp).Methods(http.MethodPatch)
	router.HandleFunc("/api/cart/update/down", handler.UpdateItemCountDown).Methods(http.MethodPatch)
}

func (handler *CartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cookie, _ := r.Cookie("session_id")
	cart, err := handler.cartUsecase.GetUserCart(cookie.Value)
	if err != nil {
		handler.logger.LogError("problems with getting cart", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body := cart

	encoder := json.NewEncoder(w)
	err = encoder.Encode(&Result{Body: body})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		handler.logger.LogError("problems while marshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		return
	}
}

func (handler *CartHandler) AddProductToCart(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("content-type", "application/json")

	jsonbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handler.logger.LogError("problems with reading json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reqProduct := dto.ReqProductID{}
	err = json.Unmarshal(jsonbody, &reqProduct)

	if err != nil {
		handler.logger.LogError("problems with unmarshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cookie, _ := r.Cookie("session_id")

	err = handler.cartUsecase.AddProductToCart(cookie.Value, reqProduct.ProductID)
	if err != nil {
		handler.logger.LogError("problems with adding product to cart", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (handler *CartHandler) DeleteProductFromCart(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("content-type", "application/json")

	jsonbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handler.logger.LogError("problems with reading json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reqProduct := dto.ReqProductID{}
	err = json.Unmarshal(jsonbody, &reqProduct)

	if err != nil {
		handler.logger.LogError("problems with unmarshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cookie, _ := r.Cookie("session_id")

	err = handler.cartUsecase.DeleteProductFromCart(cookie.Value, reqProduct.ProductID)
	if err != nil {
		handler.logger.LogError("problems deleting product from cart", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *CartHandler) UpdateItemCountUp(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("content-type", "application/json")

	jsonbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handler.logger.LogError("problems with reading json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reqProduct := dto.ReqProductID{}
	err = json.Unmarshal(jsonbody, &reqProduct)

	if err != nil {
		handler.logger.LogError("problems with unmarshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cookie, _ := r.Cookie("session_id")

	err = handler.cartUsecase.UpdateItemCountUp(cookie.Value, reqProduct.ProductID)
	if err != nil {
		handler.logger.LogError("problems updating item count up from cart", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *CartHandler) UpdateItemCountDown(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("content-type", "application/json")

	jsonbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handler.logger.LogError("problems with reading json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reqProduct := dto.ReqProductID{}
	err = json.Unmarshal(jsonbody, &reqProduct)

	if err != nil {
		handler.logger.LogError("problems with unmarshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cookie, _ := r.Cookie("session_id")

	err = handler.cartUsecase.UpdateItemCountDown(cookie.Value, reqProduct.ProductID)
	if err != nil {
		handler.logger.LogError("problems updating item count down from cart", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
