package repository

import entity "server/internal/domain/entity"

//CartRepositoryI interface
type CartRepositoryI interface {
	CreateCart(UserID uint) (uint, error)
	GetCartByUserID(UserID uint) (*entity.Cart, error)
	GetCartProductsByCartID(cartID uint) (*entity.CartWithRestaurant, error)
	AddProductToCart(cartID uint, productID uint) error
	DeleteProductFromCart(cartID uint, productID uint) error
	UpdateItemCountUp(cartID uint, productID uint) error
	UpdateItemCountDown(cartID uint, productID uint) error
	CheckProductInCart(cartID uint, productID uint) (bool, error)
	CheckProductCount(cartID uint, productID uint) (uint, error)
	CleanCart(cartID uint) error
}
