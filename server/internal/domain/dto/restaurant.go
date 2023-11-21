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

type RestaurantWithCategories struct {
	ID              uint
	Name            string
	Rating          float32
	CommentsCount   int
	Icon            string
	Categories      []string
	MinDeliveryTime int
	MaxDeliveryTime int
	DeliveryPrice   float32
}

func ToRestaurantWithCategories(restaurant *entity.Restaurant, categories []*entity.Category) *RestaurantWithCategories {
	return &RestaurantWithCategories{
		ID:              restaurant.ID,
		Name:            restaurant.Name,
		Rating:          restaurant.Rating,
		CommentsCount:   restaurant.CommentsCount,
		Icon:            restaurant.Icon,
		Categories:      *TransformCategoriesToStringSlice(categories),
		MinDeliveryTime: restaurant.MinDeliveryTime,
		MaxDeliveryTime: restaurant.MaxDeliveryTime,
		DeliveryPrice:   restaurant.DeliveryPrice,
	}
}

func TransformCategoriesToStringSlice(categories []*entity.Category) *[]string {
	categorySlice := []string{}
	for _, cat := range categories {
		category := cat.Name
		categorySlice = append(categorySlice, category)
	}
	return &categorySlice
}

type MenuTypeWithProducts struct {
	MenuType *entity.MenuType
	Products []*entity.Product
}
