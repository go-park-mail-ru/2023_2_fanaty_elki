package delivery

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	promoUsecase "server/internal/Promo/usecase"
	"server/internal/domain/entity"
	mw "server/internal/middleware"

	"github.com/gorilla/mux"
)

type Result struct {
	Body interface{}
}

type RespError struct {
	Err string
}

type PromoHandler struct {
	promoUsecase promoUsecase.UsecaseI
	logger       *mw.ACLog
}

func NewPromoHandler(promoUsecase promoUsecase.UsecaseI, logger *mw.ACLog) *PromoHandler {
	return &PromoHandler{
		promoUsecase: promoUsecase,
		logger:       logger,
	}
}

func (handler *PromoHandler) RegisterHandler(router *mux.Router) {
	router.HandleFunc("/api/promo", handler.UsePromo).Methods(http.MethodPost)
	router.HandleFunc("/api/promo/{promocode}", handler.DeletePromo).Methods(http.MethodDelete)
}

func (handler *PromoHandler) UsePromo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Header.Get("Content-Type") != "application/json" {
		handler.logger.LogError("bad content-type", entity.ErrBadContentType, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	jsonbody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		handler.logger.LogError("problems with reading json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cookie, _ := r.Cookie("session_id")

	var promocode string
	err = json.Unmarshal(jsonbody, &promocode)
	if err != nil {
		handler.logger.LogError("problems with unmarshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	promo, err := handler.promoUsecase.UsePromo(cookie.Value, promocode)

	if err != nil {
		if err == entity.ErrNotFound {
			handler.logger.LogError("promo not found", err, w.Header().Get("request-id"), r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if err == entity.ErrActionConditionsNotMet {
			handler.logger.LogError("promo conditions no met", err, w.Header().Get("request-id"), r.URL.Path)
			w.WriteHeader(entity.StatusActionConditionsNotMet)
			return
		}

		if err == entity.ErrPromoIsAlreadyApplied {
			handler.logger.LogError("promo is already applied", err, w.Header().Get("request-id"), r.URL.Path)
			w.WriteHeader(entity.StatusPromoIsAlreadyApplied)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(&RespError{Err: "data base error"})
		return
	}

	body := promo

	encoder := json.NewEncoder(w)
	err = encoder.Encode(&Result{Body: body})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(&RespError{Err: "error while marshalling JSON"})
		return
	}

}

func (handler *PromoHandler) DeletePromo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	promocode, ok := vars["promocode"]
	if !ok {
		handler.logger.LogError("problems with parameters", errors.New("category is missing in parameters"), w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cookie, _ := r.Cookie("session_id")

	err := handler.promoUsecase.DeletePromo(cookie.Value, promocode)

	if err != nil {
		if err == entity.ErrNotFound {
			handler.logger.LogError("promo not found", err, w.Header().Get("request-id"), r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(&RespError{Err: "data base error"})
		return
	}

}
