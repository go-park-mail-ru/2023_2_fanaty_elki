package repository

import "server/internal/domain/entity"

type PromoRepositoryI interface {
	GetPromo(code string) (*entity.Promo, error)
	GetPromoById(PromoID uint) (*entity.Promo, error)
	UsePromo(UserID uint, PromoID uint) error
	SetPromoToCart(cartId uint, PromoID uint) error
	CheckPromo(UserID uint, PromoID uint) (bool, error)
	DeletePromo(UserID uint, PromoID uint) error
	DeletePromoFromCart(cartId uint, PromoID uint) error
}
