package usecase

import (
	commentRep "server/internal/Comment/repository"
	restRep "server/internal/Restaurant/repository"
	userRep "server/internal/User/repository"
	"server/internal/domain/dto"
	"server/internal/domain/entity"
)

// CommentUsecaseI interface
type CommentUsecaseI interface {
	CreateComment(comment *dto.ReqCreateComment) (*dto.RespCreateComment, error)
	GetComments(id uint) (*dto.RespComments, error)
}

// CommentUsecase struct
type CommentUsecase struct {
	commentRepo commentRep.CommentRepositoryI
	userRepo    userRep.UserRepositoryI
	restRepo    restRep.RestaurantRepositoryI
}

// NewCommentUsecase crate comment usecase
func NewCommentUsecase(commentRepI commentRep.CommentRepositoryI, userRepI userRep.UserRepositoryI,
	restRepI restRep.RestaurantRepositoryI) *CommentUsecase {
	return &CommentUsecase{
		commentRepo: commentRepI,
		userRepo:    userRepI,
		restRepo:    restRepI,
	}
}

// CreateComment creates comment
func (c *CommentUsecase) CreateComment(comment *dto.ReqCreateComment) (*dto.RespCreateComment, error) {
	if comment.Rating < 1 || comment.Rating > 5 {
		return nil, entity.ErrInvalidRating
	}

	enComment := comment.FromReqToEntCreateComment()
	enComment, err := c.commentRepo.Create(dto.FromEntToDBReqCreateComment(enComment))

	if err != nil {
		return nil, err
	}

	respComment := dto.FromEntToRespCreateComment(enComment)
	user, err := c.userRepo.FindUserByID(enComment.UserID)
	if err != nil {
		return nil, err
	}
	respComment.Username = user.Username

	if user.Icon.Valid {
		respComment.Icon = user.Icon.String
	} else {
		respComment.Icon = "img/defaultIcon.webp"
	}

	err = c.restRepo.UpdateComments(comment)
	if err != nil {
		return nil, err
	}

	return respComment, nil
}

// GetComments gets comments
func (c *CommentUsecase) GetComments(id uint) (*dto.RespComments, error) {
	resp, err := c.commentRepo.Get(id)

	var respComments dto.RespComments

	for _, comment := range resp {
		respComments = append(respComments, comment)
	}

	if err == entity.ErrNotFound {
		return &dto.RespComments{}, nil
	}
	return &respComments, err
}
