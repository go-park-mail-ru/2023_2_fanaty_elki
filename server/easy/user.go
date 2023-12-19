package dto

import (
	"database/sql"
	"server/internal/domain/entity"
	//proto "server/proto/user"
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
	Username    string            `json:"Username"`
	Birthday    string            `json:"Birthday"`
	PhoneNumber string            `json:"PhoneNumber"`
	Email       string            `json:"Email"`
	Icon        string            `json:"Icon"`
	Addresses   []*RespGetAddress `json:"Addresses"`
	Current     uint              `json:"CurrentAddressId"`
}

type ReqLoginUser struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

type RespID struct {
	ID uint `json:"ID"`
}

type ReqUpdateUser struct {
	Username    string `json:"Username"`
	Password    string `json:"Password"`
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

type DBUpdateUser struct {
	ID          uint
	Username    string
	Password    string
	Birthday    sql.NullString
	PhoneNumber string
	Email       string
	Icon        sql.NullString
}

type DBGetUser struct {
	ID          uint
	Username    string
	Password    string
	Birthday    sql.NullString
	PhoneNumber string
	Email       string
	Icon        sql.NullString
}

// func ToRespGetUser(user *DBGetUser) *proto.DBGetUser {
// 	return &proto.DBGetUser{
// 		ID: uint64(user.ID),
// 		Username: user.Username,
// 		Password: user.Password,
// 		Birthday:    transformSqlStringToString(user.Birthday),
// 		PhoneNumber: user.PhoneNumber,
// 		Email:       user.Email,
// 		Icon:        transformSqlStringToString(user.Icon),
// 	}
// }

// func ToDBGetUser(user *proto.DBGetUser) *DBGetUser {
// 	return &DBGetUser{
// 		ID: uint(user.ID),
// 		Password: user.Password,
// 		Username: user.Username,
// 		Birthday:    *transformStringToSqlString(user.Birthday),
// 		PhoneNumber: user.PhoneNumber,
// 		Email:       user.Email,
// 		Icon:        *transformStringToSqlString(user.Icon),
// 	}
// }

// func ToRespCreateUser(user *DBCreateUser) *proto.DBCreateUser {
// 	return &proto.DBCreateUser{
// 		ID:			 uint64(user.ID),
// 		Username: 	 user.Username,
// 		Password: 	 user.Password,
// 		Birthday:    transformSqlStringToString(user.Birthday),
// 		PhoneNumber: user.PhoneNumber,
// 		Email:       user.Email,
// 		Icon:        transformSqlStringToString(user.Icon),
// 	}
// }

// func ToDBCreateUser(user *proto.DBCreateUser) *DBCreateUser {
// 	return &DBCreateUser{
// 		ID:			 uint(user.ID),
// 		Username: 	 user.Username,
// 		Password: 	 user.Password,
// 		Birthday:    *transformStringToSqlString(user.Birthday),
// 		PhoneNumber: user.PhoneNumber,
// 		Email:       user.Email,
// 		Icon:        *transformStringToSqlString(user.Icon),
// 	}
// }

// func ToRespUpdateUser(user *DBUpdateUser) *proto.DBUpdateUser {
// 	return &proto.DBUpdateUser{
// 		ID:			 uint64(user.ID),
// 		Username: 	 user.Username,
// 		Password:    user.Password,
// 		Birthday:    transformSqlStringToString(user.Birthday),
// 		PhoneNumber: user.PhoneNumber,
// 		Email:       user.Email,
// 		Icon:        transformSqlStringToString(user.Icon),
// 	}
// }

// func ToDBUpdateUser(user *proto.DBUpdateUser) *DBUpdateUser {
// 	return &DBUpdateUser{
// 		ID:		     uint(user.ID),
// 		Username: 	 user.Username,
// 		Password:    user.Password,
// 		Birthday:    *transformStringToSqlString(user.Birthday),
// 		PhoneNumber: user.PhoneNumber,
// 		Email:       user.Email,
// 		Icon:        *transformStringToSqlString(user.Icon),
// 	}
// }

// func ToEntityGetUser(reqUser *DBGetUser) *entity.User {
// 	if reqUser == nil{
// 		return nil
// 	}
// 	return &entity.User{
// 		ID: reqUser.ID,
// 		Username: reqUser.Username,
// 		Password: reqUser.Password,
// 		Birthday:    transformSqlStringToString(reqUser.Birthday),
// 		PhoneNumber: reqUser.PhoneNumber,
// 		Email:       reqUser.Email,
// 		Icon:        transformSqlStringToString(reqUser.Icon),
// 	}
// }

// func ToEntityCreateUser(reqUser *ReqCreateUser) *entity.User {
// 	if reqUser == nil{
// 		return nil
// 	}
// 	return &entity.User{
// 		ID: reqUser.ID,
// 		Username: reqUser.Username,
// 		Password: reqUser.Password,
// 		Birthday:    reqUser.Birthday,
// 		PhoneNumber: reqUser.PhoneNumber,
// 		Email:       reqUser.Email,
// 		Icon:        reqUser.Icon,
// 	}
// }

// func ToEntityUpdateUser(reqUser *ReqUpdateUser, id uint) *entity.User {
// 	if reqUser == nil{
// 		return nil
// 	}
// 	return &entity.User{
// 		ID: id,
// 		Username: reqUser.Username,
// 		Password: reqUser.Password,
// 		Birthday:    reqUser.Birthday,
// 		PhoneNumber: reqUser.PhoneNumber,
// 		Email:       reqUser.Email,
// 		Icon:        reqUser.Icon,
// 	}
// }

func ToEntityLoginUser(reqUser *ReqLoginUser) *entity.User {
	if reqUser == nil {
		return nil
	}
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
		Birthday:    *transformStringToSqlString(user.Birthday),
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Icon:        *transformStringToSqlString(user.Icon),
	}
}

func ToRepoUpdateUser(user *entity.User) *DBUpdateUser {
	return &DBUpdateUser{
		ID:          user.ID,
		Username:    user.Username,
		Password:    user.Password,
		Birthday:    *transformStringToSqlString(user.Birthday),
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Icon:        *transformStringToSqlString(user.Icon),
	}
}

func ToReqGetUserProfile(user *DBGetUser) *ReqGetUserProfile {
	return &ReqGetUserProfile{
		Username:    user.Username,
		Birthday:    transformSqlStringToString(user.Birthday),
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Icon:        transformSqlStringToString(user.Icon),
	}
}

func transformStringToSqlString(str string) *sql.NullString {
	if str != "" {
		return &sql.NullString{String: str, Valid: true}
	}
	return &sql.NullString{Valid: false}
}

func transformSqlStringToString(str sql.NullString) string {
	if str.Valid {
		return str.String
	}
	return ""
}
