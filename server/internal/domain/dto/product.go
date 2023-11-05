package dto

type RespGetOrderProduct struct {
	Name        string  `json:"Name"`
	Price       float32 `json:"Price"`
	CookingTime int		`json:"CookingTime"`
	Portion     string	`json:"Portion"`
	Icon        string	`json:"Icon"`
}