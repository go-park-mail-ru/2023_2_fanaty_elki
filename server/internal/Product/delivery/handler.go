package delivery

import (
	"fmt"
	"github.com/gorilla/mux"
	easyjson "github.com/mailru/easyjson"
	"net/http"
	productUsecase "server/internal/Product/usecase"
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

//ProductHandler struct
type ProductHandler struct {
	productUsecase productUsecase.ProductUsecaseI
	logger         *mw.ACLog
}

//NewProductHandler creates product handler
func NewProductHandler(productUsecase productUsecase.ProductUsecaseI, logger *mw.ACLog) *ProductHandler {
	return &ProductHandler{
		productUsecase: productUsecase,
		logger:         logger,
	}
}

//RegisterHandler registers product handler api
func (handler *ProductHandler) RegisterHandler(router *mux.Router) {
	router.HandleFunc("/api/products/{id}", handler.GetProduct).Methods(http.MethodGet)
}

//GetProduct handles get product request
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
		return
	}

	body := product

	_, err = easyjson.MarshalToWriter(body, w)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
