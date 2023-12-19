package delivery

import (
	"encoding/json"

	"io/ioutil"
	"net/http"
	orderUsecase "server/internal/Order/usecase"
	sessionUsecase "server/internal/Session/usecase"
	"server/internal/domain/dto"
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

type OrderHandler struct {
	orderUC   orderUsecase.UsecaseI
	sessionUC sessionUsecase.SessionUsecaseI
	logger    *mw.ACLog
}

func NewOrderHandler(orderUC orderUsecase.UsecaseI, sessionUC sessionUsecase.SessionUsecaseI, logger *mw.ACLog) *OrderHandler {
	return &OrderHandler{
		orderUC:   orderUC,
		sessionUC: sessionUC,
		logger:    logger,
	}
}

func (handler *OrderHandler) RegisterHandler(router *mux.Router) {
	router.HandleFunc("/api/orders", handler.CreateOrder).Methods(http.MethodPost)
	router.HandleFunc("/api/orders", handler.UpdateOrder).Methods(http.MethodPatch)
	router.HandleFunc("/api/orders", handler.GetOrders).Methods(http.MethodGet)
	router.HandleFunc("/api/orders/{id}", handler.GetOrder).Methods(http.MethodGet)
}

func (handler *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Header.Get("Content-Type") != "application/json" {
		handler.logger.LogError("bad content-type", entity.ErrBadContentType, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cookie, _ := r.Cookie("session_id")

	UserID, _ := handler.sessionUC.GetIDByCookie(cookie.Value)

	jsonbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handler.logger.LogError("problems with reading json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reqOrder := dto.ReqCreateOrder{UserID: UserID}
	err = json.Unmarshal(jsonbody, &reqOrder)
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
	err = json.NewEncoder(w).Encode(&Result{Body: respOrder})
	if err != nil {
		handler.logger.LogError("problems with marshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *OrderHandler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")

	jsonbody, err := ioutil.ReadAll(r.Body)
	if err != nil {

		handler.logger.LogError("problems with reading json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reqOrder := dto.ReqUpdateOrder{}
	err = json.Unmarshal(jsonbody, &reqOrder)
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

func (handler *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {

	cookie, _ := r.Cookie("session_id")
	UserID, _ := handler.sessionUC.GetIDByCookie(cookie.Value)

	respOrders, err := handler.orderUC.GetOrders(UserID)
	if err != nil {
		handler.logger.LogError("order: problems while getting orders", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(&Result{Body: respOrders})
	if err != nil {
		handler.logger.LogError("order: problems while marshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	strid, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	orderId, err := strconv.ParseUint(strid, 10, 64)
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
		OrderID: uint(orderId),
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

	err = json.NewEncoder(w).Encode(&Result{Body: respOrder})
	if err != nil {
		handler.logger.LogError("order: problems with marshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
