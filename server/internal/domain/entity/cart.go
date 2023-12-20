package entity

//Cart entity
type Cart struct {
	ID     uint
	UserID uint
}

//CartProduct entity
type CartProduct struct {
	ID        uint
	ProductID uint
	CartID    uint
	ItemCount int
}

//CartWithRestaurant entity
type CartWithRestaurant struct {
	RestaurantID uint
	Products     []*CartProduct
	PromoID      uint
}
