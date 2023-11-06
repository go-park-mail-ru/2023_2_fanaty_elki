package dto

import (
	"database/sql"
	"server/internal/domain/entity"
)

type ReqCreateUser struct {
	ID          uint   `json:"ID"`
	Username    string `json:"Username"`
	Password    string `json:"Password"`
	Birthday    string `json:"Birthday"`
	PhoneNumber string `json:"PhoneNumber"`
	Email       string `json:"Email"`
	Icon        string `json:"Icon"`
}

type ReqGetUserProfile struct {
	Username    string `json:"Username"`
	Birthday    string `json:"Birthday"`
	PhoneNumber string `json:"PhoneNumber"`
	Email       string `json:"Email"`
	Icon        string `json:"Icon"`
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

type ReqUpdateUser struct {
	Username    string `json:"Username"`
	Password    string `json:"Password"`
	Birthday    string `json:"Birthday"`
	PhoneNumber string `json:"PhoneNumber"`
	Email       string `json:"Email"`
	Icon        string `json:"Icon"`
}

type DBUpdateUser struct {
	ID          uint
	Username    string
	Password    string
	Birthday    sql.NullString
	PhoneNumber string
	Email       string
	Icon        sql.NullString
}

func ToEntityCreateUser(reqUser *ReqCreateUser) *entity.User {
	return &entity.User{
		ID:          reqUser.ID,
		Username:    reqUser.Username,
		Password:    reqUser.Password,
		Birthday:    *transformStringToSqlString(reqUser.Birthday),
		PhoneNumber: reqUser.PhoneNumber,
		Email:       reqUser.Email,
		Icon:        *transformStringToSqlString(reqUser.Icon),
	}
}

func ToEntityUpdateUser(reqUser *ReqUpdateUser, id uint) *entity.User {
	return &entity.User{
		ID:          id,
		Username:    reqUser.Username,
		Password:    reqUser.Password,
		Birthday:    *transformStringToSqlString(reqUser.Birthday),
		PhoneNumber: reqUser.PhoneNumber,
		Email:       reqUser.Email,
		Icon:        *transformStringToSqlString(reqUser.Icon),
	}
}

func ToEntityLoginUser(reqUser *ReqLoginUser) *entity.User {
	return &entity.User{
		Username: reqUser.Username,
		Password: reqUser.Password,
	}
}

func ToRepoCreateUser(user *entity.User) *DBCreateUser {
	return &DBCreateUser{
		ID:          user.ID,
		Username:    user.Username,
		Password:    user.Password,
		Birthday:    user.Birthday,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Icon:        user.Icon,
	}
}

func ToRepoUpdateUser(user *entity.User) *DBUpdateUser {
	return &DBUpdateUser{
		ID:          user.ID,
		Username:    user.Username,
		Password:    user.Password,
		Birthday:    user.Birthday,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Icon:        user.Icon,
	}
}

func ToReqGetUserProfile(user *entity.User) *ReqGetUserProfile {
	return &ReqGetUserProfile{
		Username:    user.Username,
		Birthday:    tranformSqlStringToString(user.Birthday),
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Icon:        tranformSqlStringToString(user.Icon),
	}
}

func transformStringToSqlString(str string) *sql.NullString {
	if str != "" {
		return &sql.NullString{String: str, Valid: true}
	}
	return &sql.NullString{Valid: false}
}

func tranformSqlStringToString(sqlstr sql.NullString) string {
	if sqlstr.Valid == false {
		return ""
	}
	return sqlstr.String
}
