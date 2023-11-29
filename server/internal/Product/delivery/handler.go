package delivery

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	productUsecase "server/internal/Product/usecase"
	"server/internal/domain/entity"
	mw "server/internal/middleware"
	"strconv"
)

type Result struct {
	Body interface{}
}

type RespError struct {
	Err string
}

type ProductHandler struct {
	productUsecase productUsecase.UsecaseI
	logger         *mw.ACLog
}

func NewProductHandler(productUsecase productUsecase.UsecaseI, logger *mw.ACLog) *ProductHandler {
	return &ProductHandler{
		productUsecase: productUsecase,
		logger:         logger,
	}
}

func (handler *ProductHandler) RegisterHandler(router *mux.Router) {
	router.HandleFunc("/api/products/{id}", handler.GetProduct).Methods(http.MethodGet)
}

func (handler *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	strid, ok := vars["id"]
	if !ok {
		fmt.Println("id is missing in parameters")
	}

	id64, err := strconv.ParseUint(strid, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(&RespError{Err: "id is not a number"})
		return
	}

	id := uint(id64)

	product, err := handler.productUsecase.GetProductByID(id)

	if err != nil {
		if err == entity.ErrNotFound {
			handler.logger.LogError("product not found", err, w.Header().Get("request-id"), r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(&RespError{Err: "data base error"})
		return
	}

	body := product

	encoder := json.NewEncoder(w)
	err = encoder.Encode(&Result{Body: body})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(&RespError{Err: "error while marshalling JSON"})
		return
	}

}
