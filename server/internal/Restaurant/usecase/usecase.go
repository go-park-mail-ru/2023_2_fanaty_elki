package usecase

import (
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
	Search(word string) ([]*dto.RestaurantWithCategoriesAndProducts, error)
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
		mindeltime := rand.Intn(30) + 30
		maxdeltime := mindeltime + 10 + rand.Intn(10)
		delprice := rand.Intn(400) + 200
		rest.MinDeliveryTime = mindeltime
		rest.MaxDeliveryTime = maxdeltime
		rest.DeliveryPrice = delprice
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
	mindeltime := rand.Intn(30) + 30
	maxdeltime := mindeltime + 10 + rand.Intn(10)
	delprice := rand.Intn(400) + 200
	rest.MinDeliveryTime = mindeltime
	rest.MaxDeliveryTime = maxdeltime
	rest.DeliveryPrice = delprice
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
		mindeltime := rand.Intn(30) + 30
		maxdeltime := mindeltime + 10 + rand.Intn(10)
		delprice := rand.Intn(400) + 200
		rest.MinDeliveryTime = mindeltime
		rest.MaxDeliveryTime = maxdeltime
		rest.DeliveryPrice = delprice
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

func (res restaurantUsecase) Search(word string) ([]*dto.RestaurantWithCategoriesAndProducts, error) {
	rests, err := res.restaurantRepo.SearchRestaurants(word)
	if err != nil {
		return nil, entity.ErrInternalServerError
	}
	restsbycategory, err := res.restaurantRepo.SearchCategories(word)
	if err != nil {
		return nil, entity.ErrInternalServerError
	}
	restset := make(map[uint]bool)
	for _, rest := range rests {
		restset[rest.ID] = true
	}
	for i, rest := range restsbycategory {
		if restset[rest.ID] {
			restsbycategory = append(restsbycategory[:i], restsbycategory[i+1:]...)
		} else {
			restset[rest.ID] = true
		}
	}
	rests = append(rests, restsbycategory...)
	products, err := res.productRepo.SearchProducts(word)
	for _, prod := range products {
		restId, err := res.productRepo.GetRestaurantIdByProduct(prod.ID)
		if err != nil {
			return nil, entity.ErrInternalServerError
		}
		if !restset[restId] {
			restById, err := res.restaurantRepo.GetRestaurantById(restId)
			if err != nil {
				return nil, entity.ErrInternalServerError
			}
			restset[restId] = true
			rests = append(rests, restById)
		}
	}
	if err != nil {
		return nil, entity.ErrInternalServerError
	}
	restsWithCategoriesAndProducts := []*dto.RestaurantWithCategoriesAndProducts{}
	for _, rest := range rests {
		mindeltime := rand.Intn(30) + 30
		maxdeltime := mindeltime + 10 + rand.Intn(10)
		delprice := rand.Intn(400) + 200
		rest.MinDeliveryTime = mindeltime
		rest.MaxDeliveryTime = maxdeltime
		rest.DeliveryPrice = delprice
		cats, err := res.restaurantRepo.GetCategoriesByRestaurantId(rest.ID)
		if err != nil {
			if err != entity.ErrNotFound {
				return nil, entity.ErrInternalServerError
			}
		}
		restWithCat := dto.ToRestaurantWithCategories(rest, cats)
		restWithCatsAndProducts := dto.ToRestaurantWithCategoriesAndProducts(restWithCat, []*entity.Product{})
		restsWithCategoriesAndProducts = append(restsWithCategoriesAndProducts, restWithCatsAndProducts)
	}
	for _, prod := range products {
		restId, err := res.productRepo.GetRestaurantIdByProduct(prod.ID)
		if err != nil {
			return nil, entity.ErrInternalServerError
		}
		for _, rest := range restsWithCategoriesAndProducts {
			if rest.ID == restId {
				rest.Products = append(rest.Products, prod)
			}
		}
	}
	return restsWithCategoriesAndProducts, nil
}
