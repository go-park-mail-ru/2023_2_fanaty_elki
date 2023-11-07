package usecase

import (
	"math"
	"math/rand"
	productRep "server/internal/Product/repository"
	restRep "server/internal/Restaurant/repository"
	dto "server/internal/domain/dto"
	"server/internal/domain/entity"
)

type UsecaseI interface {
	GetRestaurants() ([]*entity.Restaurant, error)
	GetRestaurantById(id uint) (*dto.RestaurantWithProducts, error)
}

type restaurantUsecase struct {
	restaurantRepo restRep.RestaurantRepositoryI
	productRepo    productRep.ProductRepositoryI
}

func NewRestaurantUsecase(resRep restRep.RestaurantRepositoryI, productRep productRep.ProductRepositoryI) *restaurantUsecase {
	return &restaurantUsecase{
		restaurantRepo: resRep,
		productRepo:    productRep,
	}
}

func (res restaurantUsecase) GetRestaurants() ([]*entity.Restaurant, error) {
	rests, err := res.restaurantRepo.GetRestaurants()
	if err != nil {
		return nil, entity.ErrInternalServerError
	}
	for _, rest := range rests {
		deltime := rand.Intn(60)
		delprice := rand.Float64() * 1000
		delprice = math.Round(delprice*100) / 100
		rest.DeliveryTime = deltime
		rest.DeliveryPrice = float32(delprice)
	}
	return rests, nil
}

func (res restaurantUsecase) GetRestaurantById(id uint) (*dto.RestaurantWithProducts, error) {
	rest, err := res.restaurantRepo.GetRestaurantById(id)
	if err != nil {
		return nil, entity.ErrInternalServerError
	}
	deltime := rand.Intn(60)
	delprice := rand.Float64() * 1000
	delprice = math.Round(delprice*100) / 100
	rest.DeliveryTime = deltime
	rest.DeliveryPrice = float32(delprice)
	menuTypes, err := res.restaurantRepo.GetMenuTypesByRestaurantId(id)
	if err != nil {
		return nil, entity.ErrInternalServerError
	}
	var menuTypesWithProducts []*dto.MenuTypeWithProducts
	for _, menu := range menuTypes {
		products, err := res.productRepo.GetProductsByMenuTypeId(menu.ID)
		if err != nil {
			return nil, entity.ErrInternalServerError
		}
		menuTypeWithProducts := dto.MenuTypeWithProducts{
			MenuType: menu,
			Products: products,
		}
		menuTypesWithProducts = append(menuTypesWithProducts, &menuTypeWithProducts)
	}

	restaurantWithProducts := dto.RestaurantWithProducts{
		Restaurant:            rest,
		MenuTypesWithProducts: menuTypesWithProducts,
	}

	return &restaurantWithProducts, nil
}
