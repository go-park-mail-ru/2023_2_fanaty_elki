package delivery

import (
	"encoding/json"
	"fmt"
	"net/http"
	productUsecase "server/internal/Product/usecase"
	"strconv"

	"github.com/gorilla/mux"
)

type Result struct {
	Body interface{}
}

type RespError struct {
	Err string
}

type ProductHandler struct {
	productUsecase productUsecase.UsecaseI
}

func NewProductHandler(productUsecase productUsecase.UsecaseI) *ProductHandler {
	return &ProductHandler{productUsecase: productUsecase}
}

func (handler *ProductHandler) RegisterHandler(router *mux.Router) {
	router.HandleFunc("/api/product/{id}", handler.GetProduct).Methods(http.MethodGet)
}

func (handler *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

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
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(&RespError{Err: "data base error"})
		return
	}

	body := map[string]interface{}{
		"Product": product,
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(&Result{Body: body})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(&RespError{Err: "error while marshalling JSON"})
		return
	}

}
