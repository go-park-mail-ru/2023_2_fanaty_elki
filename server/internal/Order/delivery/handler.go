package delivery

import (
	"encoding/json"
	//	"fmt"
	"io/ioutil"
	"net/http"
	orderUsecase "server/internal/Order/usecase"
	sessionUsecase "server/internal/Session/usecase"
	"server/internal/domain/dto"
	"server/internal/domain/entity"

	//	"server/internal/domain/entity"
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
	sessionUC sessionUsecase.UsecaseI
}

func NewOrderHandler(orderUC orderUsecase.UsecaseI, sessionUC sessionUsecase.UsecaseI) *OrderHandler {
	return &OrderHandler{
		orderUC:   orderUC,
		sessionUC: sessionUC,
	}
}

func (handler *OrderHandler) RegisterHandler(router *mux.Router) {
	router.HandleFunc("/api/orders", handler.CreateOrder).Methods(http.MethodPost)
	router.HandleFunc("/api/orders", handler.UpdateOrder).Methods(http.MethodPatch)
	router.HandleFunc("/api/orders", handler.GetOrders).Methods(http.MethodGet)
	router.HandleFunc("/api/orders/{id}", handler.GetOrder).Methods(http.MethodGet)
}

func (handler *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	cookie, err := r.Cookie("session_id")
	// if err == http.ErrNoCookie {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	return
	// } else if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// }

	userId, err := handler.sessionUC.GetIdByCookie(cookie.Value)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	// if userId == 0 {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	return
	// }

	jsonbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reqOrder := dto.ReqCreateOrder{UserId: userId}
	err = json.Unmarshal(jsonbody, &reqOrder)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	respOrder, err := handler.orderUC.CreateOrder(&reqOrder)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(&Result{Body: respOrder})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *OrderHandler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	// cookie, err := r.Cookie("session_id")
	// // if err == http.ErrNoCookie {
	// // 	w.WriteHeader(http.StatusUnauthorized)
	// // 	return
	// // } else if err != nil {
	// // 	w.WriteHeader(http.StatusInternalServerError)
	// // }

	// userId, err := handler.sessionUC.GetIdByCookie(cookie.Value)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	// if userId == 0 {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	return
	// }

	jsonbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reqOrder := dto.ReqUpdateOrder{}
	err = json.Unmarshal(jsonbody, &reqOrder)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.orderUC.UpdateOrder(&reqOrder)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	cookie, err := r.Cookie("session_id")
	// if err == http.ErrNoCookie {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	return
	// } else if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// }

	userId, err := handler.sessionUC.GetIdByCookie(cookie.Value)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	// if userId == 0 {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	return
	// }

	respOrders, err := handler.orderUC.GetOrders(userId)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(&Result{Body: respOrders})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	vars := mux.Vars(r)
	strid, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	orderId, err := strconv.ParseUint(strid, 10, 64)
	if err != nil {
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

	userId, err := handler.sessionUC.GetIdByCookie(cookie.Value)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	// if userId == 0 {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	return
	// }

	reqOrder := dto.ReqGetOneOrder{
		UserId:  userId,
		OrderId: uint(orderId),
	}

	respOrder, err := handler.orderUC.GetOrder(&reqOrder)
	if err != nil {
		if err == entity.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(&Result{Body: respOrder})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
