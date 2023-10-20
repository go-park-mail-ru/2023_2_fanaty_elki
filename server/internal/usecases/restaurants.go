package usecases

import (
	"database/sql"
	"server/internal/domain/entity"
	"server/repository"
)

type RestaurantRepo interface {
	GetRestaurants() ([]*entity.Restaurant, error)
}

type RestaurantUsecase struct {
	RestaurantRepo *repository.RestaurantRepo
}

func NewRestaurantUsecase(db *sql.DB) *RestaurantUsecase {
	return &RestaurantUsecase{
		RestaurantRepo: repository.NewRestaurantRepo(db),
	}
}

func (res RestaurantUsecase) GetRestaurants() ([]*entity.Restaurant, error) {
	return res.RestaurantRepo.GetRestaurants()
}
