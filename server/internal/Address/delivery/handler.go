package delivery

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	addressUsecase "server/internal/Address/usecase"
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

type AddressHandler struct {
	addressUC   addressUsecase.UsecaseI
	sessionUC sessionUsecase.SessionUsecaseI
	logger    *mw.ACLog
}

func NewAddressHandler(addressUC addressUsecase.UsecaseI, sessionUC sessionUsecase.SessionUsecaseI, logger *mw.ACLog) *AddressHandler {
	return &AddressHandler{
		addressUC:   addressUC,
		sessionUC: sessionUC,
		logger:    logger,
	}
}

func (handler *AddressHandler) RegisterHandler(router *mux.Router) {
	router.HandleFunc("/api/users/me/addresses", handler.CreateAddress).Methods(http.MethodPost)
	router.HandleFunc("/api/users/me/addresses/{id}", handler.DeleteAddress).Methods(http.MethodDelete)
	router.HandleFunc("/api/users/me/addresses/{id}", handler.SetAddress).Methods(http.MethodPatch)
	// router.HandleFunc("/api/addresss/{id}", handler.GetAddress).Methods(http.MethodGet)
}

func (handler *AddressHandler) CreateAddress(w http.ResponseWriter, r *http.Request) {
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
	reqAddress := dto.ReqCreateAddress{Cookie: cookie.Value}
	err = json.Unmarshal(jsonbody, &reqAddress)
	if err != nil {
		handler.logger.LogError("problems with unmarshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.addressUC.CreateAddress(&reqAddress)
	switch err {
	case entity.ErrInternalServerError:
		handler.logger.LogError("problems with creating address", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	case entity.ErrBadRequest:
		handler.logger.LogError("problems with address", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	case entity.ErrNotFound:
		handler.logger.LogError("no cookie", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusNotFound)
		return
	case entity.ErrAddressAlreadyExist:
		handler.logger.LogError("address exist", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(entity.StatusAddressAlreadyExist)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (handler *AddressHandler) DeleteAddress(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")

	cookie, _ := r.Cookie("session_id")

	vars := mux.Vars(r)
	strid, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	addressId, err := strconv.ParseUint(strid, 10, 64)
	if err != nil {
		handler.logger.LogError("problems while parsing addresss json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.addressUC.DeleteAddress(uint(addressId), cookie.Value)
	switch err {
	case entity.ErrInternalServerError:
		handler.logger.LogError("problems while deleting address", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	case entity.ErrNotFound:
		handler.logger.LogError("deleting address not found", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

func (handler *AddressHandler) SetAddress(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")

	cookie, _ := r.Cookie("session_id")

	vars := mux.Vars(r)
	strid, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	addressId, err := strconv.ParseUint(strid, 10, 64)
	if err != nil {
		handler.logger.LogError("problems while parsing addresss json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.addressUC.SetAddress(uint(addressId), cookie.Value)
	switch err {
	case entity.ErrInternalServerError:
		handler.logger.LogError("problems while deleting address", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	case entity.ErrNotFound:
		handler.logger.LogError("deleting address not found", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusNotFound)
		return
	}
}