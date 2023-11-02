package entity

type Cart struct {
	ID     uint
	UserID string
}

type CartProduct struct {
	ID        uint
	ProductID uint
	CartID    uint
	ItemCount int
}
