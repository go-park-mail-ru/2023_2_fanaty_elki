package dto

import (
	"database/sql"
	"server/internal/domain/entity"
	"time"
)

type ReqCreateComment struct {
	Text         string `json:"Text"`
	Rating       uint8  `json:"Rating"`
	RestaurantId uint
	UserId       uint
}

type DBReqCreateComment struct {
	Text         sql.NullString
	RestaurantId uint
	UserId       uint
	Rating       uint8
}

type RespCreateComment struct {
	Username string
	Icon     string
	Text     string
	Rating   uint8
	Date     time.Time
}

type DBRespCreateComment struct {
	Text         sql.NullString
	UserId       uint
	Rating       uint8
	RestaurantId uint
	Date         time.Time
}

type RespGetComment struct {
	Username string    `json:"Username"`
	Icon     string    `json:"Icon"`
	Text     string    `json:"Text"`
	Rating   uint8     `json:"Rating"`
	Date     time.Time `json:"Date"`
}

type DBRespGetComment struct {
	Username string
	Icon     string
	Text     sql.NullString
	Rating   uint8
	Date     time.Time
}

//easyjson:json
type RespComments []*RespGetComment

// func (c ReqCreateComment) FromReqToEntCreateComment() *entity.Comment {
// 	return &entity.Comment{
// 		Text:         c.Text,
// 		Rating:       c.Rating,
// 		RestaurantId: c.RestaurantId,
// 		UserId:       c.UserId,
// 	}
// }

// func (c DBRespCreateComment) FromDBRespToEntCreateComment() *entity.Comment {
// 	return &entity.Comment{
// 		Text:   transformSqlStringToString(c.Text),
// 		Rating: c.Rating,
// 		Date:   c.Date,
// 		UserId: c.UserId,
// 	}
// }

// func FromEntToDBReqCreateComment(comment *entity.Comment) *DBReqCreateComment {
// 	return &DBReqCreateComment{
// 		Text:         *transformStringToSqlString(comment.Text),
// 		Rating:       comment.Rating,
// 		RestaurantId: comment.RestaurantId,
// 		UserId:       comment.UserId,
// 	}
// }

func FromEntToRespCreateComment(comment *entity.Comment) *RespCreateComment {
	return &RespCreateComment{
		Text:   comment.Text,
		Rating: comment.Rating,
		Date:   comment.Date,
	}
}

func (c DBRespGetComment) FromDBtoDel() *RespGetComment {
	return &RespGetComment{
		Text:     transformSqlStringToString(c.Text),
		Username: c.Username,
		Icon:     c.Icon,
		Rating:   c.Rating,
		Date:     c.Date,
	}
}
