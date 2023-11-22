package dto

type RespGetOrderProduct struct {
	Id 			uint 	`json:"Id"`
	Name        string  `json:"Name"`
	Price       float32 `json:"Price"`
	Icon        string	`json:"Icon"`
	Count 		int 	`json:"Count"`
}