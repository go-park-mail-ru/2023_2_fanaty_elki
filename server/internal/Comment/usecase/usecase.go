package usecase

import (
	commentRep "server/internal/Comment/repository"
	restRep "server/internal/Restaurant/repository"
	userRep "server/internal/User/repository"
	"server/internal/domain/dto"
	"server/internal/domain/entity"
)

type UsecaseI interface {
	CreateComment(comment *dto.ReqCreateComment) (*dto.RespCreateComment, error)
	GetComments(id uint) ([]*dto.RespGetComment, error)
}

type commentUsecase struct {
	commentRepo commentRep.CommentRepositoryI
	userRepo    userRep.UserRepositoryI
	restRepo    restRep.RestaurantRepositoryI
}

func NewCommentUsecase(commentRepI commentRep.CommentRepositoryI, userRepI userRep.UserRepositoryI,
	restRepI restRep.RestaurantRepositoryI) *commentUsecase {
	return &commentUsecase{
		commentRepo: commentRepI,
		userRepo:    userRepI,
		restRepo:    restRepI,
	}
}

func (c *commentUsecase) CreateComment(comment *dto.ReqCreateComment) (*dto.RespCreateComment, error) {
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
		respComment.Icon = "img/defaultIcon.png"
	}

	err = c.restRepo.UpdateComments(comment)
	if err != nil {
		return nil, err
	}

	return respComment, nil
}

func (c *commentUsecase) GetComments(id uint) ([]*dto.RespGetComment, error) {
	resp, err := c.commentRepo.Get(id)
	if err == entity.ErrNotFound {
		return []*dto.RespGetComment{}, nil
	}
	return resp, err
}
