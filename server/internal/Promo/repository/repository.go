package repository

import "server/internal/domain/entity"

//PromoRepositoryI interface
type PromoRepositoryI interface {
	GetPromo(code string) (*entity.Promo, error)
	GetPromoByID(PromoID uint) (*entity.Promo, error)
	UsePromo(UserID uint, PromoID uint) error
	SetPromoToCart(cartID uint, PromoID uint) error
	CheckPromo(UserID uint, PromoID uint) (bool, error)
	DeletePromo(UserID uint, PromoID uint) error
	DeletePromoFromCart(cartID uint, PromoID uint) error
}
