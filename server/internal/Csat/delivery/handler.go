package delivery

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	csatUsecase "server/internal/Csat/usecase"
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

type CsatHandler struct {
	csatUsecase csatUsecase.UsecaseI
	logger      *mw.ACLog
}

func NewCsatHandler(csatUsecase csatUsecase.UsecaseI, logger *mw.ACLog) *CsatHandler {
	return &CsatHandler{
		csatUsecase: csatUsecase,
		logger:      logger,
	}
}

func (handler *CsatHandler) RegisterHandler(router *mux.Router) {
	router.HandleFunc("/api/csat/quizzes/{id}/config", handler.GetQuestionnaire).Methods(http.MethodGet)
	router.HandleFunc("/api/csat/quizzes", handler.AddAnswer).Methods(http.MethodPost)
	router.HandleFunc("/api/csat/quizzes/answers/{id}", handler.GetAnswerList).Methods(http.MethodGet)
}

func (handler *CsatHandler) GetQuestionnaire(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	strid, ok := vars["id"]
	if !ok {
		handler.logger.LogError("problems with parameters", errors.New("id is missing in parameters"), w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id64, err := strconv.ParseUint(strid, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		handler.logger.LogError("problems with parameters", errors.New("id is not number"), w.Header().Get("request-id"), r.URL.Path)
		err = json.NewEncoder(w).Encode(&RespError{Err: "id is not a number"})
		return
	}

	id := uint(id64)

	questionnaire, err := handler.csatUsecase.GetQuestionnaireByID(id)
	if err != nil {
		handler.logger.LogError("problems with getting cart", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body := questionnaire

	encoder := json.NewEncoder(w)
	err = encoder.Encode(&Result{Body: body})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		handler.logger.LogError("problems while marshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		return
	}
}

func (handler *CsatHandler) AddAnswer(w http.ResponseWriter, r *http.Request) {
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

	reqAnswer := entity.Answer{}
	err = json.Unmarshal(jsonbody, &reqAnswer)

	if err != nil {
		handler.logger.LogError("problems with unmarshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.csatUsecase.AddAnswer(&reqAnswer)
	if err != nil {
		handler.logger.LogError("problems with getting cart", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

func (handler *CsatHandler) GetAnswerList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	strid, ok := vars["id"]
	if !ok {
		handler.logger.LogError("problems with parameters", errors.New("id is missing in parameters"), w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id64, err := strconv.ParseUint(strid, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		handler.logger.LogError("problems with parameters", errors.New("id is not number"), w.Header().Get("request-id"), r.URL.Path)
		err = json.NewEncoder(w).Encode(&RespError{Err: "id is not a number"})
		return
	}

	id := uint(id64)

	answers, err := handler.csatUsecase.GetAnswersByQuestionId(id)

	if err != nil {
		handler.logger.LogError("problems with getting restauratns", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body := answers

	encoder := json.NewEncoder(w)
	err = encoder.Encode(&Result{Body: body})

	if err != nil {
		handler.logger.LogError("problems with marshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
