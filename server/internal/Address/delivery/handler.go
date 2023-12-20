package delivery

import (
	"github.com/gorilla/mux"
	easyjson "github.com/mailru/easyjson"
	"net/http"
	addressUsecase "server/internal/Address/usecase"
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

//AddressHandler struct
type AddressHandler struct {
	addressUC addressUsecase.AddressUsecaseI
	sessionUC sessionUsecase.SessionUsecaseI
	logger    *mw.ACLog
}

//NewAddressHandler creates address handler
func NewAddressHandler(addressUC addressUsecase.AddressUsecaseI, sessionUC sessionUsecase.SessionUsecaseI, logger *mw.ACLog) *AddressHandler {
	return &AddressHandler{
		addressUC: addressUC,
		sessionUC: sessionUC,
		logger:    logger,
	}
}

//RegisterHandler registers address handler api
func (handler *AddressHandler) RegisterHandler(router *mux.Router) {
	router.HandleFunc("/api/users/me/addresses", handler.CreateAddress).Methods(http.MethodPost)
	router.HandleFunc("/api/users/me/addresses/{id}", handler.DeleteAddress).Methods(http.MethodDelete)
	router.HandleFunc("/api/users/me/addresses/{id}", handler.SetAddress).Methods(http.MethodPatch)
	// router.HandleFunc("/api/addresss/{id}", handler.GetAddress).Methods(http.MethodGet)
}

//CreateAddress handles create address request
func (handler *AddressHandler) CreateAddress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Header.Get("Content-Type") != "application/json" {
		handler.logger.LogError("bad content-type", entity.ErrBadContentType, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cookie, _ := r.Cookie("session_id")
	reqAddress := dto.ReqCreateAddress{Cookie: cookie.Value}

	err := easyjson.UnmarshalFromReader(r.Body, &reqAddress)
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

//DeleteAddress handles delete address request
func (handler *AddressHandler) DeleteAddress(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")

	cookie, _ := r.Cookie("session_id")

	vars := mux.Vars(r)
	strid, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	addressID, err := strconv.ParseUint(strid, 10, 64)
	if err != nil {
		handler.logger.LogError("problems while parsing addresss json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.addressUC.DeleteAddress(uint(addressID), cookie.Value)
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

//SetAddress handles set address request
func (handler *AddressHandler) SetAddress(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")

	cookie, _ := r.Cookie("session_id")

	vars := mux.Vars(r)
	strid, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	addressID, err := strconv.ParseUint(strid, 10, 64)
	if err != nil {
		handler.logger.LogError("problems while parsing addresss json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.addressUC.SetAddress(uint(addressID), cookie.Value)
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
