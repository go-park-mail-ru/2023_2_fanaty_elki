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

type UsecaseI interface {
	UsePromo(SessionToken string, promocode string) (*dto.RespPromo, error)
	DeletePromo(SessionToken string, promocode string) error
}

type promoUsecase struct {
	cartRepo       cartRep.CartRepositoryI
	promoRepo      promoRep.PromoRepositoryI
	sessionRepo    sessionRep.SessionRepositoryI
	restaurantRepo restaurantRep.RestaurantRepositoryI
}

func NewPromoUsecase(cartRep cartRep.CartRepositoryI, promoRep promoRep.PromoRepositoryI, sessionRep sessionRep.SessionRepositoryI, restaurantRep restaurantRep.RestaurantRepositoryI) *promoUsecase {
	return &promoUsecase{
		cartRepo:       cartRep,
		promoRepo:      promoRep,
		sessionRepo:    sessionRep,
		restaurantRepo: restaurantRep,
	}
}

func (pu promoUsecase) UsePromo(SessionToken string, promocode string) (*dto.RespPromo, error) {
	cookie, err := pu.sessionRepo.Check(SessionToken)
	if err != nil {
		return nil, err
	}

	userID := cookie.UserID
	cart, err := pu.cartRepo.GetCartByUserID(userID)
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

	if promo.RestaurantId != 0 {
		if cartWithRestaurant.RestaurantId != promo.RestaurantId {
			return nil, entity.ErrActionConditionsNotMet
		}
	}

	haspromo, err := pu.promoRepo.CheckPromo(userID, promo.ID)

	if err != nil {
		return nil, err
	}

	if haspromo {
		return nil, entity.ErrPromoIsAlreadyApplied
	}

	err = pu.promoRepo.UsePromo(userID, promo.ID)
	if err != nil {
		return nil, err
	}

	err = pu.promoRepo.SetPromoToCart(cart.ID, promo.ID)
	if err != nil {
		return nil, err
	}

	return dto.ToRespPromo(promo), nil
}

func (pu promoUsecase) DeletePromo(SessionToken string, promocode string) error {
	cookie, err := pu.sessionRepo.Check(SessionToken)
	if err != nil {
		return err
	}

	userID := cookie.UserID

	cart, err := pu.cartRepo.GetCartByUserID(userID)
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

	haspromo, err := pu.promoRepo.CheckPromo(userID, promo.ID)

	if err != nil {
		return err
	}

	if !haspromo {
		return entity.ErrNotFound
	} else {
		err := pu.promoRepo.DeletePromo(userID, promo.ID)
		if err != nil {
			return err
		}

		err = pu.promoRepo.DeletePromoFromCart(cart.ID, promo.ID)
		if err != nil {
			return err
		}
	}

	return nil
}
