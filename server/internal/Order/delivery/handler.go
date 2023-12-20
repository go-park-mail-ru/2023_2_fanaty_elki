package delivery

import (
	"github.com/gorilla/mux"
	easyjson "github.com/mailru/easyjson"
	"net/http"
	orderUsecase "server/internal/Order/usecase"
	sessionUsecase "server/internal/Session/usecase"
	"server/internal/domain/dto"
	"server/internal/domain/entity"
	mw "server/internal/middleware"
	"strconv"
)

//Result struct
type Result struct {
	Body interface{}
}

//RespError struct
type RespError struct {
	Err string
}

//OrderHandler struct
type OrderHandler struct {
	orderUC   orderUsecase.OrderUsecaseI
	sessionUC sessionUsecase.SessionUsecaseI
	logger    *mw.ACLog
}

//NewOrderHandler creates order handler
func NewOrderHandler(orderUC orderUsecase.OrderUsecaseI, sessionUC sessionUsecase.SessionUsecaseI, logger *mw.ACLog) *OrderHandler {
	return &OrderHandler{
		orderUC:   orderUC,
		sessionUC: sessionUC,
		logger:    logger,
	}
}

//RegisterHandler registers order handler api
func (handler *OrderHandler) RegisterHandler(router *mux.Router) {
	router.HandleFunc("/api/orders", handler.CreateOrder).Methods(http.MethodPost)
	router.HandleFunc("/api/orders", handler.UpdateOrder).Methods(http.MethodPatch)
	router.HandleFunc("/api/orders", handler.GetOrders).Methods(http.MethodGet)
	router.HandleFunc("/api/orders/{id}", handler.GetOrder).Methods(http.MethodGet)
}

//CreateOrder handles create order request
func (handler *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Header.Get("Content-Type") != "application/json" {
		handler.logger.LogError("bad content-type", entity.ErrBadContentType, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cookie, _ := r.Cookie("session_id")

	UserID, _ := handler.sessionUC.GetIDByCookie(cookie.Value)

	reqOrder := dto.ReqCreateOrder{UserID: UserID}

	err := easyjson.UnmarshalFromReader(r.Body, &reqOrder)
	if err != nil {
		handler.logger.LogError("problems with unmarshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	respOrder, err := handler.orderUC.CreateOrder(&reqOrder)
	switch err {
	case entity.ErrInternalServerError:
		handler.logger.LogError("problems with creating order", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	case entity.ErrBadRequest:
		handler.logger.LogError("problems with address", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	case entity.ErrNotFound:
		handler.logger.LogError("Cart is empty", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusCreated)

	_, err = easyjson.MarshalToWriter(respOrder, w)
	if err != nil {
		handler.logger.LogError("problems with marshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

//UpdateOrder handles update order request
func (handler *OrderHandler) UpdateOrder(w http.ResponseWriter, r *http.Request) {

	reqOrder := dto.ReqUpdateOrder{}

	err := easyjson.UnmarshalFromReader(r.Body, &reqOrder)
	if err != nil {
		handler.logger.LogError("problems with unmarshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.orderUC.UpdateOrder(&reqOrder)
	if err != nil {
		handler.logger.LogError("problems while updating json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

//GetOrders handles get order request
func (handler *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {

	cookie, _ := r.Cookie("session_id")
	UserID, _ := handler.sessionUC.GetIDByCookie(cookie.Value)

	respOrders, err := handler.orderUC.GetOrders(UserID)
	if err != nil {
		handler.logger.LogError("order: problems while getting orders", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = easyjson.MarshalToWriter(respOrders, w)
	if err != nil {
		handler.logger.LogError("order: problems while marshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

//GetOrder handles get order request
func (handler *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	strid, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	orderID, err := strconv.ParseUint(strid, 10, 64)
	if err != nil {
		handler.logger.LogError("problems while parsing orders json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cookie, err := r.Cookie("session_id")
	// if err == http.ErrNoCookie {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	return
	// } else if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// }

	UserID, err := handler.sessionUC.GetIDByCookie(cookie.Value)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	// if UserID == 0 {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	return
	// }

	reqOrder := dto.ReqGetOneOrder{
		UserID:  UserID,
		OrderID: uint(orderID),
	}

	respOrder, err := handler.orderUC.GetOrder(&reqOrder)

	if err != nil {
		if err == entity.ErrNotFound {
			handler.logger.LogError("order: not found order", err, w.Header().Get("request-id"), r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = easyjson.MarshalToWriter(respOrder, w)
	if err != nil {
		handler.logger.LogError("order: problems with marshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
