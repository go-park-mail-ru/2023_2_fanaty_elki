package dto

import "server/internal/domain/entity"

//RestaurantWithCategories dto
type RestaurantWithCategories struct {
	ID              uint
	Name            string
	Rating          float32
	CommentsCount   int
	Icon            string
	Categories      []string
	MinDeliveryTime int
	MaxDeliveryTime int
	DeliveryPrice   int
}

//ToRestaurantWithCategories creates RestaurantWithCategories from Restaurant and Categories 
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

//TransformCategoriesToStringSlice transforms categories to string slice
func TransformCategoriesToStringSlice(categories []*entity.Category) *[]string {
	categorySlice := []string{}
	for _, cat := range categories {
		category := cat.Name
		categorySlice = append(categorySlice, category)
	}
	return &categorySlice
}

//MenuTypeWithProducts dto
type MenuTypeWithProducts struct {
	MenuType *entity.MenuType
	Products []*entity.Product
}

//RestaurantWithCategoriesAndProducts dto
type RestaurantWithCategoriesAndProducts struct {
	ID              uint
	Name            string
	Rating          float32
	CommentsCount   int
	Icon            string
	Categories      []string
	MinDeliveryTime int
	MaxDeliveryTime int
	DeliveryPrice   int
	Products        []*entity.Product
}

//ToRestaurantWithCategoriesAndProducts creates RestaurantWithCategoriesAndProducts from RestaurantWithCategories and Products 
func ToRestaurantWithCategoriesAndProducts(restaurant *RestaurantWithCategories, products []*entity.Product) *RestaurantWithCategoriesAndProducts {
	return &RestaurantWithCategoriesAndProducts{
		ID:              restaurant.ID,
		Name:            restaurant.Name,
		Rating:          restaurant.Rating,
		CommentsCount:   restaurant.CommentsCount,
		Icon:            restaurant.Icon,
		Categories:      restaurant.Categories,
		MinDeliveryTime: restaurant.MinDeliveryTime,
		MaxDeliveryTime: restaurant.MaxDeliveryTime,
		DeliveryPrice:   restaurant.DeliveryPrice,
		Products:        products,
	}
}

//easyjson:json
type RestaurantWithCategoriesSlice []*RestaurantWithCategories

//easyjson:json
type MenuTypeWithProductsSlice []*MenuTypeWithProducts

//easyjson:json
type StringSlice []string

//easyjson:json
type RestaurantWithCategoriesAndProductsSlice []*RestaurantWithCategoriesAndProducts
