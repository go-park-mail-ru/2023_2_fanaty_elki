package dto

import "server/internal/domain/entity"

// type ReqGetRestaurant struct {
// 	ID            uint    `json:"ID"`
// 	Name          string  `json:"Name"`
// 	Rating        float32 `json:"Rating"`
// 	CommentsCount int     `json:"CommentsCount"`
// 	Icon          string  `json:"Icon"`
// 	Category      string  `json:"Category"`
// }

// func ToEntityRestaurant(restaurant *ReqGetRestaurant) *entity.Restaurant {
// 	return &entity.Restaurant{
// 		ID: restaurant.ID,
// 		Name: restaurant.Name,
// 		Rating: restaurant.Rating,
// 		CommentsCount: restaurant.CommentsCount,
// 		Icon: restaurant.Icon,
// 		Category: restaurant.Category,
// 	}
// }

type MenuTypeWithProducts struct {
	MenuType *entity.MenuType
	Products []*entity.Product
}
type RestaurantWithProducts struct {
	Restaurant            *entity.Restaurant
	MenuTypesWithProducts []*MenuTypeWithProducts
}
