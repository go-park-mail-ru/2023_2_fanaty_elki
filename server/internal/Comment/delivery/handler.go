package delivery

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	commentUsecase "server/internal/Comment/usecase"
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

type CommentHandler struct {
	commentUC   commentUsecase.UsecaseI
	sessionUC sessionUsecase.UsecaseI
	logger    *mw.ACLog
}

func NewCommentHandler(commentUC commentUsecase.UsecaseI, sessionUC sessionUsecase.UsecaseI, logger *mw.ACLog) *CommentHandler {
	return &CommentHandler{
		commentUC:   commentUC,
		sessionUC: sessionUC,
		logger:    logger,
	}
}

func (handler *CommentHandler) RegisterPostHandler(router *mux.Router) {
	router.HandleFunc("/api/comments/{RestaurantId}", handler.CreateComment).Methods(http.MethodPost)
}

func (handler *CommentHandler) RegisterGetHandler(router *mux.Router) {
	router.HandleFunc("/api/comments/{RestaurantId}", handler.GetComments).Methods(http.MethodGet)
}

func (handler *CommentHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	strid, ok := vars["RestaurantId"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	restaurantId, err := strconv.ParseUint(strid, 10, 64)
	if err != nil {
		handler.logger.LogError("problems while parsing comments get parameters", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if r.Header.Get("Content-type") != "application/json" {
		handler.logger.LogError("bad content-type", entity.ErrBadContentType, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cookie, _ := r.Cookie("session_id")
	userId, _ := handler.sessionUC.GetIdByCookie(cookie.Value)

	reqComment := &dto.ReqCreateComment{
		UserId: userId,
		RestaurantId: uint(restaurantId),
	}

	jsonbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handler.logger.LogError("problems with reading json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(jsonbody, &reqComment)
	if err != nil {
		handler.logger.LogError("problems with unmarshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	respComment, err := handler.commentUC.CreateComment(reqComment)
	switch err {
	case entity.ErrInvalidRating:
		handler.logger.LogError("problems with rating", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	case entity.ErrInternalServerError:
		handler.logger.LogError("problems with creating comment", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(&Result{Body:respComment})
	if err != nil {
		handler.logger.LogError("problems with marshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *CommentHandler) GetComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	strid, ok := vars["RestaurantId"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	restaurantId, err := strconv.ParseUint(strid, 10, 64)
	if err != nil {
		handler.logger.LogError("problems while parsing comments get parameters", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	respComment, err := handler.commentUC.GetComments(uint(restaurantId))
	if err != nil {
		handler.logger.LogError("problems with getting comments", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(&Result{Body:respComment})
	if err != nil {
		handler.logger.LogError("problems with marshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}