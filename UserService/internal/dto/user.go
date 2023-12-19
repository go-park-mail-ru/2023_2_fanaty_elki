package dto

import (
	proto "UserService/proto"
	"database/sql"
)

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

func ToRespGetUser(user *DBGetUser) *proto.DBGetUser {
	if user == nil {
		return nil
	}
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

func ToDBGetUser(user *proto.DBGetUser) *DBGetUser {
	if user == nil {
		return nil
	}
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

func ToRespCreateUser(user *DBCreateUser) *proto.DBCreateUser {
	if user == nil {
		return nil
	}
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

func ToDBCreateUser(user *proto.DBCreateUser) *DBCreateUser {
	if user == nil {
		return nil
	}
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

func ToRespUpdateUser(user *DBUpdateUser) *proto.DBUpdateUser {
	if user == nil {
		return nil
	}
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

func ToDBUpdateUser(user *proto.DBUpdateUser) *DBUpdateUser {
	if user == nil {
		return nil
	}
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
