package dto

import (
	"server/internal/domain/entity"
)

//ReqCreateOrderAddress dto
type ReqCreateOrderAddress struct {
	City   string `json:"City"`
	Street string `json:"Street"`
	House  string `json:"House"`
	Flat   uint   `json:"Flat"`
}

//RespOrderAddress dto
type RespOrderAddress struct {
	City   string `json:"City"`
	Street string `json:"Street"`
	House  string `json:"House"`
	Flat   uint   `json:"Flat"`
}

//DBCreateOrderAddress dto
type DBCreateOrderAddress struct {
	City   string
	Street string
	House  string
	Flat   uint
}

//DBReqCreateUserAddress dto
type DBReqCreateUserAddress struct {
	UserID uint
	City   string
	Street string
	House  string
	Flat   uint
}

//ReqCreateAddress dto
type ReqCreateAddress struct {
	City   string `json:"City"`
	Street string `json:"Street"`
	House  string `json:"House"`
	Flat   uint   `json:"Flat"`
	Cookie string
}

//DBReqDeleteUserAddress dto
type DBReqDeleteUserAddress struct {
	UserID    uint
	AddressID uint
}

//DBReqUpdateUserAddress dto
type DBReqUpdateUserAddress struct {
	UserID    uint
	AddressID uint
}

//RespGetAddress dto
type RespGetAddress struct {
	ID     uint   `json:"Id"`
	City   string `json:"City"`
	Street string `json:"Street"`
	House  string `json:"House"`
	Flat   uint   `json:"Flat"`
}

//RespGetAddresses dto
type RespGetAddresses struct {
	Addresses          []*RespGetAddress `json:"Addresses"`
	CurrentAddressesID uint              `json:"CurrentAddressesId"`
}

//ToEntityCreateOrderAddress transforms ReqCreateOrderAddress to Address
func ToEntityCreateOrderAddress(address *ReqCreateOrderAddress) *entity.Address {
	return &entity.Address{
		City:   address.City,
		Street: address.Street,
		House:  address.House,
		Flat:   address.Flat,
	}
}

//ToDBCreateOrderAddress transforms Address to DBCreateOrderAddress
func ToDBCreateOrderAddress(address *entity.Address) *DBCreateOrderAddress {
	return &DBCreateOrderAddress{
		Flat:   address.Flat,
		City:   address.City,
		Street: address.Street,
		House:  address.House,
	}

}

//ToEntityCreateAddress transforms ReqCreateAddress to Address
func ToEntityCreateAddress(address *ReqCreateAddress) *entity.Address {
	return &entity.Address{
		City:   address.City,
		Street: address.Street,
		House:  address.House,
		Flat:   address.Flat,
	}
}

//ToDBCreateAddress transforms address and user id to ReqCreateUserAddress
func ToDBCreateAddress(address *entity.Address, UserID uint) *DBReqCreateUserAddress {
	return &DBReqCreateUserAddress{
		UserID: UserID,
		Flat:   address.Flat,
		City:   address.City,
		Street: address.Street,
		House:  address.House,
	}
}
