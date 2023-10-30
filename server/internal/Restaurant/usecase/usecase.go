package usecase

import (
	restRep "server/internal/Restaurant/repository"
	"server/internal/domain/entity"
)

type UsecaseI interface {
	GetRestaurants() ([]*entity.Restaurant, error)
}

type restaurantUsecase struct {
	RestaurantRepo restRep.RestaurantRepositoryI
}

func NewRestaurantUsecase(resRep restRep.RestaurantRepositoryI) *restaurantUsecase {
	return &restaurantUsecase{
		RestaurantRepo: resRep,
	}
}

func (res restaurantUsecase) GetRestaurants() ([]*entity.Restaurant, error) {
	return res.RestaurantRepo.GetRestaurants()
}
