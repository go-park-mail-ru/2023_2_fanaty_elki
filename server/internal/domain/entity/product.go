package entity

//Product entity
type Product struct {
	ID          uint    `json:"ID"`
	Name        string  `json:"Name"`
	Price       float32 `json:"Price"`
	CookingTime int     `json:"CookingTime"`
	Portion     string  `json:"Portion"`
	Description string  `json:"Description"`
	Icon        string  `json:"Icon"`
}
