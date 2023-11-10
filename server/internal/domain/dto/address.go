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
	Flat   *uint  `json:"Flat"`
}


type DBCreateOrderAddress struct {
	City   string 
	Street string 
	House  string   
	Flat   *uint  
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
	result := &DBCreateOrderAddress{}
	if address.Flat == 0 {
		result.Flat = nil
	} else {
		result.Flat = &address.Flat
	}
	result.City = address.City
	result.Street = address.Street
	result.House = address.House
	return result
}

