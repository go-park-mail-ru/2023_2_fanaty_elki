package usecase

import (
	cartRep "server/internal/Cart/repository"
	promoRep "server/internal/Promo/repository"
	restaurantRep "server/internal/Restaurant/repository"
	sessionRep "server/internal/Session/repository"
	dto "server/internal/domain/dto"
	"server/internal/domain/entity"
	"time"
)

//PromoUsecaseI interface
type PromoUsecaseI interface {
	UsePromo(SessionToken string, promocode string) (*dto.RespPromo, error)
	DeletePromo(SessionToken string, promocode string) error
}

//PromoUsecase struct
type PromoUsecase struct {
	cartRepo       cartRep.CartRepositoryI
	promoRepo      promoRep.PromoRepositoryI
	sessionRepo    sessionRep.SessionRepositoryI
	restaurantRepo restaurantRep.RestaurantRepositoryI
}

//NewPromoUsecase creates promo usecase 
func NewPromoUsecase(cartRep cartRep.CartRepositoryI, promoRep promoRep.PromoRepositoryI, sessionRep sessionRep.SessionRepositoryI, restaurantRep restaurantRep.RestaurantRepositoryI) *PromoUsecase {
	return &PromoUsecase{
		cartRepo:       cartRep,
		promoRepo:      promoRep,
		sessionRepo:    sessionRep,
		restaurantRepo: restaurantRep,
	}
}

//UsePromo uses promocode
func (pu PromoUsecase) UsePromo(SessionToken string, promocode string) (*dto.RespPromo, error) {
	cookie, err := pu.sessionRepo.Check(SessionToken)
	if err != nil {
		return nil, err
	}

	UserID := cookie.UserID
	cart, err := pu.cartRepo.GetCartByUserID(UserID)
	if err != nil {
		return nil, err
	}

	cartWithRestaurant, err := pu.cartRepo.GetCartProductsByCartID(cart.ID)
	if err != nil {
		return nil, err
	}

	promo, err := pu.promoRepo.GetPromo(promocode)
	if err != nil {
		return nil, err
	}

	if promo == nil {
		return nil, entity.ErrNotFound
	}

	if time.Now().Before(promo.ActiveFrom) || time.Now().After(promo.ActiveTo) {

		return nil, entity.ErrActionConditionsNotMet
	}

	if promo.RestaurantID != 0 {
		if cartWithRestaurant.RestaurantID != promo.RestaurantID {
			return nil, entity.ErrActionConditionsNotMet
		}
	}

	haspromo, err := pu.promoRepo.CheckPromo(UserID, promo.ID)

	if err != nil {
		return nil, err
	}

	if haspromo {
		return nil, entity.ErrPromoIsAlreadyApplied
	}

	err = pu.promoRepo.UsePromo(UserID, promo.ID)
	if err != nil {
		return nil, err
	}

	err = pu.promoRepo.SetPromoToCart(cart.ID, promo.ID)
	if err != nil {
		return nil, err
	}

	return dto.ToRespPromo(promo), nil
}

//DeletePromo deletes promocode
func (pu PromoUsecase) DeletePromo(SessionToken string, promocode string) error {
	cookie, err := pu.sessionRepo.Check(SessionToken)
	if err != nil {
		return err
	}

	UserID := cookie.UserID

	cart, err := pu.cartRepo.GetCartByUserID(UserID)
	if err != nil {
		return err
	}

	promo, err := pu.promoRepo.GetPromo(promocode)
	if err != nil {
		return err
	}

	if promo == nil {
		return entity.ErrNotFound
	}

	haspromo, err := pu.promoRepo.CheckPromo(UserID, promo.ID)

	if err != nil {
		return err
	}

	if !haspromo {
		return entity.ErrNotFound
	}
	err = pu.promoRepo.DeletePromo(UserID, promo.ID)
	if err != nil {
		return err
	}

	err = pu.promoRepo.DeletePromoFromCart(cart.ID, promo.ID)
	if err != nil {
		return err
	}
	

	return nil
}
