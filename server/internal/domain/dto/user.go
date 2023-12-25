package dto

import (
	"database/sql"
	"server/internal/domain/entity"
	proto "server/proto/user"
)

//ReqCreateUser dto
type ReqCreateUser struct {
	ID          uint   `json:"ID"`
	Username    string `json:"Username"`
	Password    string `json:"Password"`
	Birthday    string `json:"Birthday"`
	PhoneNumber string `json:"PhoneNumber"`
	Email       string `json:"Email"`
	Icon        string `json:"Icon"`
}

//ReqGetUserProfile dto
type ReqGetUserProfile struct {
	Username    string            `json:"Username"`
	Birthday    string            `json:"Birthday"`
	PhoneNumber string            `json:"PhoneNumber"`
	Email       string            `json:"Email"`
	Icon        string            `json:"Icon"`
	Addresses   []*RespGetAddress `json:"Addresses"`
	Current     uint              `json:"CurrentAddressId"`
}

//ReqLoginUser dto
type ReqLoginUser struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

//RespID struct
type RespID struct {
	ID uint `json:"ID"`
}

//ReqUpdateUser dto
type ReqUpdateUser struct {
	Username    string `json:"Username"`
	Password    string `json:"Password"`
	Birthday    string `json:"Birthday"`
	PhoneNumber string `json:"PhoneNumber"`
	Email       string `json:"Email"`
	Icon        string `json:"Icon"`
}

//DBCreateUser dto
type DBCreateUser struct {
	ID          uint
	Username    string
	Password    string
	Birthday    sql.NullString
	PhoneNumber string
	Email       string
	Icon        sql.NullString
}

//DBUpdateUser dto
type DBUpdateUser struct {
	ID          uint
	Username    string
	Password    string
	Birthday    sql.NullString
	PhoneNumber string
	Email       string
	Icon        sql.NullString
}

//DBGetUser dto
type DBGetUser struct {
	ID          uint
	Username    string
	Password    string
	Birthday    sql.NullString
	PhoneNumber string
	Email       string
	Icon        sql.NullString
}

//ToRespGetUser transforms DBGetUser to proto.DBGetUser
func ToRespGetUser(user *DBGetUser) *proto.DBGetUser {
	return &proto.DBGetUser{
		ID:          uint64(user.ID),
		Username:    user.Username,
		Password:    user.Password,
		Birthday:    transformSQLStringToString(user.Birthday),
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Icon:        transformSQLStringToString(user.Icon),
	}
}

//ToDBGetUser transforms proto.DBGetUser to DBGetUser
func ToDBGetUser(user *proto.DBGetUser) *DBGetUser {
	return &DBGetUser{
		ID:          uint(user.ID),
		Password:    user.Password,
		Username:    user.Username,
		Birthday:    *transformStringToSQLString(user.Birthday),
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Icon:        *transformStringToSQLString(user.Icon),
	}
}

//ToRespCreateUser transforms DBCreateUser to proto.DBCreateUser
func ToRespCreateUser(user *DBCreateUser) *proto.DBCreateUser {
	return &proto.DBCreateUser{
		ID:          uint64(user.ID),
		Username:    user.Username,
		Password:    user.Password,
		Birthday:    transformSQLStringToString(user.Birthday),
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Icon:        transformSQLStringToString(user.Icon),
	}
}

//ToDBCreateUser transforms proto.DBCreateUser to DBCreateUser
func ToDBCreateUser(user *proto.DBCreateUser) *DBCreateUser {
	return &DBCreateUser{
		ID:          uint(user.ID),
		Username:    user.Username,
		Password:    user.Password,
		Birthday:    *transformStringToSQLString(user.Birthday),
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Icon:        *transformStringToSQLString(user.Icon),
	}
}

//ToRespUpdateUser transforms DBUpdateUser to proto.DBUpdateUser
func ToRespUpdateUser(user *DBUpdateUser) *proto.DBUpdateUser {
	return &proto.DBUpdateUser{
		ID:          uint64(user.ID),
		Username:    user.Username,
		Password:    user.Password,
		Birthday:    transformSQLStringToString(user.Birthday),
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Icon:        transformSQLStringToString(user.Icon),
	}
}

//ToDBUpdateUser transforms proto.DBUpdateUser to DBUpdateUser
func ToDBUpdateUser(user *proto.DBUpdateUser) *DBUpdateUser {
	return &DBUpdateUser{
		ID:          uint(user.ID),
		Username:    user.Username,
		Password:    user.Password,
		Birthday:    *transformStringToSQLString(user.Birthday),
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Icon:        *transformStringToSQLString(user.Icon),
	}
}

//ToEntityGetUser transforms DBGetUser to User
func ToEntityGetUser(reqUser *DBGetUser) *entity.User {
	if reqUser == nil {
		return nil
	}
	return &entity.User{
		ID:          reqUser.ID,
		Username:    reqUser.Username,
		Password:    reqUser.Password,
		Birthday:    transformSQLStringToString(reqUser.Birthday),
		PhoneNumber: reqUser.PhoneNumber,
		Email:       reqUser.Email,
		Icon:        transformSQLStringToString(reqUser.Icon),
	}
}

//ToEntityCreateUser transforms ReqCreateUser to User
func ToEntityCreateUser(reqUser *ReqCreateUser) *entity.User {
	if reqUser == nil {
		return nil
	}
	return &entity.User{
		ID:          reqUser.ID,
		Username:    reqUser.Username,
		Password:    reqUser.Password,
		Birthday:    reqUser.Birthday,
		PhoneNumber: reqUser.PhoneNumber,
		Email:       reqUser.Email,
		Icon:        reqUser.Icon,
	}
}

//ToEntityUpdateUser transforms ReqUpdateUser to User
func ToEntityUpdateUser(reqUser *ReqUpdateUser, id uint) *entity.User {
	if reqUser == nil {
		return nil
	}
	return &entity.User{
		ID:          id,
		Username:    reqUser.Username,
		Password:    reqUser.Password,
		Birthday:    reqUser.Birthday,
		PhoneNumber: reqUser.PhoneNumber,
		Email:       reqUser.Email,
		Icon:        reqUser.Icon,
	}
}

//ToEntityLoginUser transforms ReqLoginUser to User
func ToEntityLoginUser(reqUser *ReqLoginUser) *entity.User {
	if reqUser == nil {
		return nil
	}
	return &entity.User{
		Username: reqUser.Username,
		Password: reqUser.Password,
	}
}

//ToRepoCreateUser transforms User to DBCreateUser
func ToRepoCreateUser(user *entity.User) *DBCreateUser {
	return &DBCreateUser{
		ID:          user.ID,
		Username:    user.Username,
		Password:    user.Password,
		Birthday:    *transformStringToSQLString(user.Birthday),
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Icon:        *transformStringToSQLString(user.Icon),
	}
}

//ToRepoUpdateUser transforms User to DBUpdateUser
func ToRepoUpdateUser(user *entity.User) *DBUpdateUser {
	return &DBUpdateUser{
		ID:          user.ID,
		Username:    user.Username,
		Password:    user.Password,
		Birthday:    *transformStringToSQLString(user.Birthday),
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Icon:        *transformStringToSQLString(user.Icon),
	}
}

//ToReqGetUserProfile transforms DBGetUset to ReqGetUserProfile
func ToReqGetUserProfile(user *DBGetUser) *ReqGetUserProfile {
	return &ReqGetUserProfile{
		Username:    user.Username,
		Birthday:    transformSQLStringToString(user.Birthday),
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Icon:        transformSQLStringToString(user.Icon),
	}
}

func transformStringToSQLString(str string) *sql.NullString {
	if str != "" {
		return &sql.NullString{String: str, Valid: true}
	}
	return &sql.NullString{Valid: false}
}

func transformSQLStringToString(str sql.NullString) string {
	if str.Valid {
		return str.String
	}
	return ""
}
