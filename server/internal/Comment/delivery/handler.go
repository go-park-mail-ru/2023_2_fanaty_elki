package delivery

import (
	"net/http"
	commentUsecase "server/internal/Comment/usecase"
	sessionUsecase "server/internal/Session/usecase"
	"server/internal/domain/dto"
	"server/internal/domain/entity"
	mw "server/internal/middleware"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
)

//Result struct
type Result struct {
	Body interface{}
}

//RespError struct
type RespError struct {
	Err string
}

//CommentHandler struct
type CommentHandler struct {
	commentUC commentUsecase.CommentUsecaseI
	sessionUC sessionUsecase.SessionUsecaseI
	logger    *mw.ACLog
}

//NewCommentHandler create comment handler
func NewCommentHandler(commentUC commentUsecase.CommentUsecaseI, sessionUC sessionUsecase.SessionUsecaseI, logger *mw.ACLog) *CommentHandler {
	return &CommentHandler{
		commentUC: commentUC,
		sessionUC: sessionUC,
		logger:    logger,
	}
}

//RegisterPostHandler registers comment handler api
func (handler *CommentHandler) RegisterPostHandler(router *mux.Router) {
	router.HandleFunc("/api/comments/{RestaurantId}", handler.CreateComment).Methods(http.MethodPost)
}

//RegisterGetHandler registers comment handler api
func (handler *CommentHandler) RegisterGetHandler(router *mux.Router) {
	router.HandleFunc("/api/comments/{RestaurantId}", handler.GetComments).Methods(http.MethodGet)
}

//CreateComment handles create comment request
func (handler *CommentHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	strid, ok := vars["RestaurantId"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	restaurantID, err := strconv.ParseUint(strid, 10, 64)
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
	UserID, _ := handler.sessionUC.GetIDByCookie(cookie.Value)

	reqComment := &dto.ReqCreateComment{
		UserID:       UserID,
		RestaurantID: uint(restaurantID),
	}

	err = easyjson.UnmarshalFromReader(r.Body, reqComment)
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

	_, err = easyjson.MarshalToWriter(respComment, w)

	if err != nil {
		handler.logger.LogError("problems with marshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

//GetComments handles get comments request
func (handler *CommentHandler) GetComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	strid, ok := vars["RestaurantId"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	restaurantID, err := strconv.ParseUint(strid, 10, 64)
	if err != nil {
		handler.logger.LogError("problems while parsing comments get parameters", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	respComment, err := handler.commentUC.GetComments(uint(restaurantID))
	if err != nil {
		handler.logger.LogError("problems with getting comments", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = easyjson.MarshalToWriter(respComment, w)

	if err != nil {
		handler.logger.LogError("problems with marshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
