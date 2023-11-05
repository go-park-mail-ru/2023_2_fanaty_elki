package delivery

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	orderUsecase "server/internal/Order/usecase"
	sessionUsecase "server/internal/Session/usecase"
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

type OrderHandler struct {
	orderUC orderUsecase.UsecaseI
	sessionUC sessionUsecase.UsecaseI
}

func NewOrderHandler(orderUC orderUsecase.UsecaseI, sessionUC sessionUsecase.UsecaseI) *OrderHandler {
	return &OrderHandler{
		orderUC: orderUC,
		sessionUC: sessionUC,
	}
}

func (handler *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Access-Control-Allow-Origin", allowedOrigin)
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Set("content-type", "application/json")

	cookie, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		w.WriteHeader(http.StatusUnauthorized)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	userId, err := handler.sessionUC.GetIdByCookie(cookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if userId == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	jsonbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(&Error{Err: entity.ErrProblemsReadingData.Error()})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
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
	err = json.NewEncoder(w).Encode(&Result{Body:respOrder})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}