package delivery

import (
	//"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	cartUsecase "server/internal/Cart/usecase"
	//"server/internal/domain/dto"
	"server/internal/domain/entity"
	mw "server/internal/middleware"
	"strconv"

	"github.com/gorilla/mux"
	easyjson "github.com/mailru/easyjson"
	easyjsonopt "github.com/mailru/easyjson/opt"
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
	router.HandleFunc("/api/cart", handler.AddProductToCart).Methods(http.MethodPost)
	router.HandleFunc("/api/cart/{id}", handler.DeleteProductFromCart).Methods(http.MethodDelete)
	router.HandleFunc("/api/cart/clear", handler.CleanCart).Methods(http.MethodPost)
	router.HandleFunc("/api/cart/tips", handler.GetCartTips).Methods(http.MethodGet)
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

	// encoder := json.NewEncoder(w)
	// err = encoder.Encode(&Result{Body: body})

	_, err = easyjson.MarshalToWriter(body, w)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		handler.logger.LogError("problems while marshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		return
	}
}

func (handler *CartHandler) AddProductToCart(w http.ResponseWriter, r *http.Request) {
	jsonbody, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	handler.logger.LogError("problems with reading json", err, w.Header().Get("request-id"), r.URL.Path)
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	//err = json.Unmarshal(jsonbody, &id)

	//id := &dto.ReqProductID{}

	var id easyjsonopt.Int

	id.UnmarshalJSON(jsonbody)

	if err != nil {
		handler.logger.LogError("problems with unmarshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cookie, _ := r.Cookie("session_id")

	err = handler.cartUsecase.AddProductToCart(cookie.Value, uint(id.V))
	if err != nil {
		if err == entity.ErrNotFound {
			handler.logger.LogError("problems no product", err, w.Header().Get("request-id"), r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		handler.logger.LogError("problems with adding product to cart", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (handler *CartHandler) DeleteProductFromCart(w http.ResponseWriter, r *http.Request) {
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
		return
	}

	id := uint(id64)

	cookie, _ := r.Cookie("session_id")

	err = handler.cartUsecase.DeleteProductFromCart(cookie.Value, id)
	if err != nil {
		if err == entity.ErrNotFound {
			handler.logger.LogError("problems no product or no product in cart", err, w.Header().Get("request-id"), r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		handler.logger.LogError("problems deleting product from cart", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *CartHandler) CleanCart(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = handler.cartUsecase.CleanCart(cookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *CartHandler) GetCartTips(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cookie, _ := r.Cookie("session_id")
	tips, err := handler.cartUsecase.GetCartTips(cookie.Value)
	if err != nil {
		if err == entity.ErrNotFound {
			handler.logger.LogError("problems no products in cart", err, w.Header().Get("request-id"), r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		handler.logger.LogError("problems with getting cart", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body := tips

	// encoder := json.NewEncoder(w)
	// err = encoder.Encode(&Result{Body: body})

	_, err = easyjson.MarshalToWriter(body, w)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		handler.logger.LogError("problems while marshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		return
	}
}
