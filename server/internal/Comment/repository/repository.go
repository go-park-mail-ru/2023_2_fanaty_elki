package repository

import (
	"server/internal/domain/dto"
	"server/internal/domain/entity"
)

type CommentRepositoryI interface {
	Create(comment *dto.DBReqCreateComment) (*entity.Comment, error) 
	Get(id uint) ([]*dto.RespGetComment, error)
}