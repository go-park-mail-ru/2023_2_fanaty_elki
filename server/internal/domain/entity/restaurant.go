package entity

type Restaurant struct {
	ID            uint    `json:"ID"`
	Name          string  `json:"Name"`
	Rating        float32 `json:"Rating"`
	CommentsCount int     `json:"CommentsCount"`
	Icon          string  `json:"Icon"`
	Category      string  `json:"Category"`
}