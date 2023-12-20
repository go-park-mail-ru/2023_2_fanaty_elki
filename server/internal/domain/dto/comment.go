package dto

import (
	"database/sql"
	"server/internal/domain/entity"
	"time"
)

//ReqCreateComment dto
type ReqCreateComment struct {
	Text         string `json:"Text"`
	Rating       uint8  `json:"Rating"`
	RestaurantID uint
	UserID       uint
}

//DBReqCreateComment dto
type DBReqCreateComment struct {
	Text         sql.NullString
	RestaurantID uint
	UserID       uint
	Rating       uint8
}

//RespCreateComment dto
type RespCreateComment struct {
	Username string
	Icon     string
	Text     string
	Rating   uint8
	Date     time.Time
}

//DBRespCreateComment dto
type DBRespCreateComment struct {
	Text         sql.NullString
	UserID       uint
	Rating       uint8
	RestaurantID uint
	Date         time.Time
}

//RespGetComment dto
type RespGetComment struct {
	Username string    `json:"Username"`
	Icon     string    `json:"Icon"`
	Text     string    `json:"Text"`
	Rating   uint8     `json:"Rating"`
	Date     time.Time `json:"Date"`
}

//DBRespGetComment dto
type DBRespGetComment struct {
	Username string
	Icon     string
	Text     sql.NullString
	Rating   uint8
	Date     time.Time
}

//easyjson:json
type RespComments []*RespGetComment

//FromReqToEntCreateComment transforms ReqCreateComment to Comment
func (c ReqCreateComment) FromReqToEntCreateComment() *entity.Comment {
	return &entity.Comment{
		Text:         c.Text,
		Rating:       c.Rating,
		RestaurantID: c.RestaurantID,
		UserID:       c.UserID,
	}
}

//FromDBRespToEntCreateComment transforms DBRespCreateComment to Comment
func (c DBRespCreateComment) FromDBRespToEntCreateComment() *entity.Comment {
	return &entity.Comment{
		Text:   transformSQLStringToString(c.Text),
		Rating: c.Rating,
		Date:   c.Date,
		UserID: c.UserID,
	}
}

//FromEntToDBReqCreateComment transforms Comment to DBReqCreateComment
func FromEntToDBReqCreateComment(comment *entity.Comment) *DBReqCreateComment {
	return &DBReqCreateComment{
		Text:         *transformStringToSQLString(comment.Text),
		Rating:       comment.Rating,
		RestaurantID: comment.RestaurantID,
		UserID:       comment.UserID,
	}
}

//FromEntToRespCreateComment transforms Comment to RespCreateComment
func FromEntToRespCreateComment(comment *entity.Comment) *RespCreateComment {
	return &RespCreateComment{
		Text:   comment.Text,
		Rating: comment.Rating,
		Date:   comment.Date,
	}
}

//FromDBtoDel transforms DBRespGetComment to RespGetComment
func (c DBRespGetComment) FromDBtoDel() *RespGetComment {
	return &RespGetComment{
		Text:     transformSQLStringToString(c.Text),
		Username: c.Username,
		Icon:     c.Icon,
		Rating:   c.Rating,
		Date:     c.Date,
	}
}
