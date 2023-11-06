package repository

import entity "server/internal/domain/entity"

type CartRepositoryI interface {
	CreateCart(userID uint) (uint, error)
	GetCartByUserID(userID uint) (*entity.Cart, error)
	GetCartProductsByCartID(cartID uint) ([]*entity.CartProduct, error)
	AddProductToCart(cartID uint, productID uint) error
	DeleteProductFromCart(cartID uint, productID uint) error
	UpdateItemCountUp(cartID uint, productID uint) error
	UpdateItemCountDown(cartID uint, productID uint) error
}
