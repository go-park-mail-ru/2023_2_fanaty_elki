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
	GetRestaurants() ([]*dto.RestaurantWithCategories, error)
	GetRestaurantById(id uint) (*dto.RestaurantWithCategories, error)
	GetRestaurantProducts(id uint) ([]*dto.MenuTypeWithProducts, error)
	GetRestaurantsByCategory(name string) ([]*dto.RestaurantWithCategories, error)
	GetCategories() (*[]string, error)
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

func (res restaurantUsecase) GetRestaurants() ([]*dto.RestaurantWithCategories, error) {
	rests, err := res.restaurantRepo.GetRestaurants()
	if err != nil {
		return nil, entity.ErrInternalServerError
	}
	restswithcategories := []*dto.RestaurantWithCategories{}
	for _, rest := range rests {
		mindeltime := rand.Intn(60)
		maxdeltime := mindeltime + rand.Intn(20)
		delprice := rand.Float64() * 1000
		delprice = math.Round(delprice*100) / 100
		rest.MinDeliveryTime = mindeltime
		rest.MaxDeliveryTime = maxdeltime
		rest.DeliveryPrice = float32(delprice)
		cats, err := res.restaurantRepo.GetCategoriesByRestaurantId(rest.ID)
		if err != nil {
			if err != entity.ErrNotFound {
				return nil, entity.ErrInternalServerError
			}
		}
		restwithcat := dto.ToRestaurantWithCategories(rest, cats)
		restswithcategories = append(restswithcategories, restwithcat)
	}
	return restswithcategories, nil
}

func (res restaurantUsecase) GetRestaurantById(id uint) (*dto.RestaurantWithCategories, error) {
	rest, err := res.restaurantRepo.GetRestaurantById(id)
	if err != nil {
		return nil, err
	}
	mindeltime := rand.Intn(60)
	maxdeltime := mindeltime + rand.Intn(20)
	delprice := rand.Float64() * 1000
	delprice = math.Round(delprice*100) / 100
	rest.MinDeliveryTime = mindeltime
	rest.MaxDeliveryTime = maxdeltime
	rest.DeliveryPrice = float32(delprice)
	cats, err := res.restaurantRepo.GetCategoriesByRestaurantId(rest.ID)
	if err != nil {
		if err != entity.ErrNotFound {
			return nil, entity.ErrInternalServerError
		}
	}
	restwithcat := dto.ToRestaurantWithCategories(rest, cats)
	return restwithcat, nil
}

func (res restaurantUsecase) GetRestaurantProducts(id uint) ([]*dto.MenuTypeWithProducts, error) {
	menuTypes, err := res.restaurantRepo.GetMenuTypesByRestaurantId(id)
	if err != nil {
		return nil, err
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

	return menuTypesWithProducts, nil
}

func (res restaurantUsecase) GetRestaurantsByCategory(name string) ([]*dto.RestaurantWithCategories, error) {
	rests, err := res.restaurantRepo.GetRestaurantsByCategory(name)
	if err != nil {
		if err == entity.ErrNotFound {
			return nil, entity.ErrNotFound
		}
		return nil, entity.ErrInternalServerError
	}
	restswithcategories := []*dto.RestaurantWithCategories{}
	for _, rest := range rests {
		mindeltime := rand.Intn(60)
		maxdeltime := mindeltime + rand.Intn(20)
		delprice := rand.Float64() * 1000
		delprice = math.Round(delprice*100) / 100
		rest.MinDeliveryTime = mindeltime
		rest.MaxDeliveryTime = maxdeltime
		rest.DeliveryPrice = float32(delprice)
		cats, err := res.restaurantRepo.GetCategoriesByRestaurantId(rest.ID)
		if err != nil {
			if err != entity.ErrNotFound {
				return nil, entity.ErrInternalServerError
			}
		}
		restwithcat := dto.ToRestaurantWithCategories(rest, cats)
		restswithcategories = append(restswithcategories, restwithcat)
	}
	return restswithcategories, nil
}

func (res restaurantUsecase) GetCategories() (*[]string, error) {
	cats, err := res.restaurantRepo.GetCategories()
	categories := dto.TransformCategoriesToStringSlice(cats)
	if err != nil {
		return nil, entity.ErrInternalServerError
	}
	return categories, nil
}
