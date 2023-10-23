package usecase

import (
	"server/internal/domain/entity"
	restRep "server/internal/Restaurant/repository"
)

type UsecaseI interface {
	GetRestaurants() ([]*entity.Restaurant, error)
}

type restaurantUsecase struct {
	RestaurantRepo restRep.RestaurantI
}

func NewRestaurantUsecase(resRep restRep.RestaurantI) *restaurantUsecase {
	return &restaurantUsecase{
		RestaurantRepo: resRep,
	}
}

func (res restaurantUsecase) GetRestaurants() ([]*entity.Restaurant, error) {
	return res.RestaurantRepo.GetRestaurants()
}
