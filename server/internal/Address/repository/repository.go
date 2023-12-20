package repository

import (
	"server/internal/domain/dto"
)

//AddressRepositoryI interface
type AddressRepositoryI interface {
	CreateAddress(address *dto.DBReqCreateUserAddress) error
	DeleteAddress(address *dto.DBReqDeleteUserAddress) error
	GetAddresses(UserID uint) (*dto.RespGetAddresses, error)
	SetAddress(address *dto.DBReqUpdateUserAddress) error
}
