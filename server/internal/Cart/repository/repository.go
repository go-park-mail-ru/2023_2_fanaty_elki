package repository

import entity "server/internal/domain/entity"

type CartRepositoryI interface {
	CreateCart(userID uint) (uint, error)
	GetCartByUserID(userID uint) (*entity.Cart, error)
	GetCartProductsByCartID(cartID uint) ([]*entity.CartProduct, error)
}
