package delivery

import (
	"errors"
	"io/ioutil"
	"net/http"
	promoUsecase "server/internal/Promo/usecase"
	"server/internal/domain/entity"
	mw "server/internal/middleware"

	"github.com/gorilla/mux"
	easyjson "github.com/mailru/easyjson"
	easyjsonopt "github.com/mailru/easyjson/opt"
)

//Result struct
type Result struct {
	Body interface{}
}

//RespError struct
type RespError struct {
	Err string
}

//PromoHandler struct
type PromoHandler struct {
	promoUsecase promoUsecase.PromoUsecaseI
	logger       *mw.ACLog
}

//NewPromoHandler creates new promo handler object
func NewPromoHandler(promoUsecase promoUsecase.PromoUsecaseI, logger *mw.ACLog) *PromoHandler {
	return &PromoHandler{
		promoUsecase: promoUsecase,
		logger:       logger,
	}
}

//RegisterHandler regisers promocode api
func (handler *PromoHandler) RegisterHandler(router *mux.Router) {
	router.HandleFunc("/api/promo", handler.UsePromo).Methods(http.MethodPost)
	router.HandleFunc("/api/promo/{promocode}", handler.DeletePromo).Methods(http.MethodDelete)
}

//UsePromo handles use promocode request
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

	var promocode easyjsonopt.String
	err = promocode.UnmarshalJSON(jsonbody)
	if err != nil {
		handler.logger.LogError("problems with unmarshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	promo, err := handler.promoUsecase.UsePromo(cookie.Value, promocode.V)

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
		return
	}

	body := promo

	_, err = easyjson.MarshalToWriter(body, w)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

//DeletePromo handles deletes promocode request
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
		return
	}

}
