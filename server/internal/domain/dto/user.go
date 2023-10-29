package dto

import (
	"server/internal/domain/entity"
	"database/sql"
)

type ReqCreateUser struct {
	ID          uint           `json:"ID"`
	Username    string         `json:"Username"`
	Password    string         `json:"Password"`
	Birthday    string		   `json:"Birthday"`
	PhoneNumber string         `json:"PhoneNumber"`
	Email       string         `json:"Email"`
	Icon        string		   `json:"Icon"`
}

type DBCreateUser struct {
	ID          uint           
	Username    string         
	Password    string         
	Birthday    sql.NullString
	PhoneNumber string         
	Email       string         
	Icon        sql.NullString 
}

type ReqLoginUser struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

func ToEntityCreateUser(reqUser *ReqCreateUser) *entity.User {
	return &entity.User{
		ID: reqUser.ID,
		Username: reqUser.Username,         
		Password: reqUser.Password,         
		Birthday:    reqUser.Birthday,
		PhoneNumber: reqUser.PhoneNumber,         
		Email:       reqUser.Email,         
		Icon:        reqUser.Icon, 
	}
} 

func ToEntityLoginUser(reqUser *ReqLoginUser) *entity.User {
	return &entity.User{
		Username: reqUser.Username,         
		Password: reqUser.Password,         
	}
} 

func ToRepoUser (user *entity.User) *DBCreateUser{
	return &DBCreateUser{
		ID:          user.ID,           
		Username:    user.Username,         
		Password:    user.Password,         
		Birthday:    *transformStringToSqlString(user.Birthday),
		PhoneNumber: user.PhoneNumber,         
		Email:       user.Email,         
		Icon:        *transformStringToSqlString(user.Icon), 
	}
}

func transformStringToSqlString(str string) *sql.NullString {
	if str != "" {
		return &sql.NullString{String: str, Valid: true}
	}
	return &sql.NullString{Valid: false}
}