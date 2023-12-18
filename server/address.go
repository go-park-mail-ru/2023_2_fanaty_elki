package dto

import (
	"server/internal/domain/entity"
)

type ReqCreateOrderAddress struct {
	City   string `json:"City"`
	Street string `json:"Street"`
	House  string `json:"House"`
	Flat   uint  `json:"Flat"`
}

type RespOrderAddress struct {
	City   string `json:"City"`
	Street string `json:"Street"`
	House  string `json:"House"`
	Flat   uint  `json:"Flat"`
}


type DBCreateOrderAddress struct {
	City   string 
	Street string 
	House  string   
	Flat   uint  
}

type DBReqCreateUserAddress struct {
	UserId uint	 
	City   string 
	Street string 
	House  string   
	Flat   uint  
}

type ReqCreateAddress struct {
	City   string `json:"City"`
	Street string `json:"Street"`
	House  string `json:"House"`
	Flat   uint   `json:"Flat"`
	Cookie string 
}



type DBReqDeleteUserAddress struct {
	UserId 	  uint
	AddressId uint 
}

type DBReqUpdateUserAddress struct {
	UserId 	  uint
	AddressId uint 
}

type RespGetAddress struct {
	Id 	   uint   `json:"Id"`
	City   string `json:"City"`
	Street string `json:"Street"`
	House  string `json:"House"`
	Flat   uint   `json:"Flat"`
}

type RespGetAddresses struct {
	Addresses 		   []*RespGetAddress `json:"Addresses"`
	CurrentAddressesId uint 			 `json:"CurrentAddressesId"`
}

func ToEntityCreateOrderAddress(address *ReqCreateOrderAddress) *entity.Address{
	return &entity.Address{
		City: address.City,
		Street: address.Street,
		House: address.House,
		Flat: address.Flat,
	}
}


func ToDBCreateOrderAddress(address *entity.Address) *DBCreateOrderAddress{
	return &DBCreateOrderAddress{
		Flat: address.Flat,
		City: address.City,
		Street: address.Street,
		House: address.House,
	}
	
}

func ToEntityCreateAddress(address *ReqCreateAddress) *entity.Address{
	return &entity.Address{
		City: address.City,
		Street: address.Street,
		House: address.House,
		Flat: address.Flat,
	}
}


func ToDBCreateAddress(address *entity.Address, userId uint) *DBReqCreateUserAddress{
	return &DBReqCreateUserAddress{
		UserId: userId,
		Flat: address.Flat,
		City: address.City,
		Street: address.Street,
		House: address.House,
	}
}

