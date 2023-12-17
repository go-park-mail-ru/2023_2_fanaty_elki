package repository

import "server/internal/domain/entity"

type PromoRepositoryI interface {
	GetPromo(code string) (*entity.Promo, error)
	GetPromoById(promoId uint) (*entity.Promo, error)
	UsePromo(userId uint, promoId uint) error
	SetPromoToCart(cartId uint, promoId uint) error
	CheckPromo(userId uint, promoId uint) (bool, error)
	DeletePromo(userId uint, promoId uint) error
	DeletePromoFromCart(cartId uint, promoId uint) error
}
