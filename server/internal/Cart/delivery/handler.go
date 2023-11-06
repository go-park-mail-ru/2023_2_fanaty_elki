package delivery

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	cartUsecase "server/internal/Cart/usecase"
	"server/internal/domain/dto"
	"server/internal/domain/entity"
)

const allowedOrigin = "http://84.23.53.216"

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

// GetCart godoc
// @Summary      giving information about cart and its products
// @Description  giving an array of products in cart
// @Tags        Cart
// @Accept     */*
// @Produce  application/json
// @Success  200 {object} []*dto.CartProduct "success returning information about cart"
// @Failure 401 {object} error "unauthorized"
// @Failure 500 {object} error "internal server error"
// @Router   /cart [get]
func (handler *CartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", allowedOrigin)
	w.Header().Add("Access-Control-Allow-Credentials", "true")
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

// AddProductToCart
// @Summary      adding product to cart
// @Description  adding product to cart by its id
// @Tags        Cart
// @Accept     application/json
// @Produce  application/json
// @Success  200  "success adding product"
// @Failure 400 {object} error "bad request"
// @Failure 401 {object} error "unauthorized"
// @Failure 500 {object} error "internal server error"
// @Router   /cart/add [post]
func (handler *CartHandler) AddProductToCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", allowedOrigin)
	w.Header().Add("Access-Control-Allow-Credentials", "true")
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

// DeleteProductFromCart
// @Summary      deleting product from cart
// @Description  deleting product from cart by its id
// @Tags        Cart
// @Accept     application/json
// @Produce  application/json
// @Success  200  "success deleting product"
// @Failure 400 {object} error "bad request"
// @Failure 401 {object} error "unauthorized"
// @Failure 500 {object} error "internal server error"
// @Router   /cart/delete [post]
func (handler *CartHandler) DeleteProductFromCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", allowedOrigin)
	w.Header().Add("Access-Control-Allow-Credentials", "true")
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

// UpdateItemCountUp
// @Summary      updating product count up
// @Description  updating product count up in cart by its id
// @Tags        Cart
// @Accept     application/json
// @Produce  application/json
// @Success  200  "updating product count up"
// @Failure 400 {object} error "bad request"
// @Failure 401 {object} error "unauthorized"
// @Failure 500 {object} error "internal server error"
// @Router   /cart/update/up [patch]
func (handler *CartHandler) UpdateItemCountUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", allowedOrigin)
	w.Header().Add("Access-Control-Allow-Credentials", "true")
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

// UpdateItemCountDown
// @Summary      updating product count down
// @Description  updating product count down in cart by its id
// @Tags        Cart
// @Accept     application/json
// @Produce  application/json
// @Success  200  "updating product count down"
// @Failure 400 {object} error "bad request"
// @Failure 401 {object} error "unauthorized"
// @Failure 500 {object} error "internal server error"
// @Router   /cart/update/up [patch]
func (handler *CartHandler) UpdateItemCountDown(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", allowedOrigin)
	w.Header().Add("Access-Control-Allow-Credentials", "true")
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
