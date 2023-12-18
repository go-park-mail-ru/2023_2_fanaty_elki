package usecase

import (
	"math/rand"

	orderRep "server/internal/Order/repository"
	productRep "server/internal/Product/repository"
	restRep "server/internal/Restaurant/repository"
	sessionRep "server/internal/Session/repository"
	dto "server/internal/domain/dto"
	"server/internal/domain/entity"
	"sort"
)

type UsecaseI interface {
	GetRestaurants() ([]*dto.RestaurantWithCategories, error)
	GetRestaurantById(id uint) (*dto.RestaurantWithCategories, error)
	GetRestaurantProducts(id uint) ([]*dto.MenuTypeWithProducts, error)
	GetRestaurantsByCategory(name string) ([]*dto.RestaurantWithCategories, error)
	GetCategories() (*[]string, error)
	Search(word string) ([]*dto.RestaurantWithCategoriesAndProducts, error)
	GetRestaurantTips(SessionToken string) ([]*dto.RestaurantWithCategories, error)
	GetRandomRestaurantTips() ([]*dto.RestaurantWithCategories, error)
}

type restaurantUsecase struct {
	restaurantRepo restRep.RestaurantRepositoryI
	productRepo    productRep.ProductRepositoryI
	sessionRepo    sessionRep.SessionRepositoryI
	orderRepo      orderRep.OrderRepositoryI
}

func NewRestaurantUsecase(resRep restRep.RestaurantRepositoryI, productRep productRep.ProductRepositoryI, sessionRep sessionRep.SessionRepositoryI, orderRep orderRep.OrderRepositoryI) *restaurantUsecase {
	return &restaurantUsecase{
		restaurantRepo: resRep,
		productRepo:    productRep,
		sessionRepo:    sessionRep,
		orderRepo:      orderRep,
	}

}

func (res restaurantUsecase) GetRestaurants() ([]*dto.RestaurantWithCategories, error) {
	rests, err := res.restaurantRepo.GetRestaurants()
	if err != nil {
		return nil, entity.ErrInternalServerError
	}
	restswithcategories := []*dto.RestaurantWithCategories{}
	for _, rest := range rests {
		mindeltime := len(rest.Name) + 15
		mindeltime = mindeltime - (mindeltime % 5)
		maxdeltime := mindeltime + 10
		maxdeltime = maxdeltime - (maxdeltime % 5)
		delprice := len(rest.Name)*8 + 200
		delprice = delprice - (delprice % 10)
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
	mindeltime := len(rest.Name) + 15
	mindeltime = mindeltime - (mindeltime % 5)
	maxdeltime := mindeltime + 10
	maxdeltime = maxdeltime - (maxdeltime % 5)
	delprice := len(rest.Name)*8 + 200
	delprice = delprice - (delprice % 10)
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
		mindeltime := len(rest.Name) + 15
		mindeltime = mindeltime - (mindeltime % 5)
		maxdeltime := mindeltime + 10
		maxdeltime = maxdeltime - (maxdeltime % 5)
		delprice := len(rest.Name)*8 + 200
		delprice = delprice - (delprice % 10)
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
		mindeltime := len(rest.Name) + 15
		mindeltime = mindeltime - (mindeltime % 5)
		maxdeltime := mindeltime + 10
		maxdeltime = maxdeltime - (maxdeltime % 5)
		delprice := len(rest.Name)*8 + 200
		delprice = delprice - (delprice % 10)
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

func (res restaurantUsecase) GetRestaurantTips(SessionToken string) ([]*dto.RestaurantWithCategories, error) {
	cookie, err := res.sessionRepo.Check(SessionToken)
	if err != nil {
		return nil, err
	}

	userID := cookie.UserID

	restaurants, err := res.restaurantRepo.GetRestaurants()

	if err != nil {
		return nil, err
	}

	type RestVal struct {
		Restaurant string
		Value      int
	}

	var restvalslice []RestVal

	for _, restaurant := range restaurants {
		restval := RestVal{Restaurant: restaurant.Name, Value: 0}
		restvalslice = append(restvalslice, restval)
	}

	orders, err := res.orderRepo.GetOrders(userID)

	if err != nil {
		return nil, err
	}

	if len(orders) == 0 {
		for i, _ := range restvalslice {
			restvalslice[i].Value = rand.Intn(5)
		}
		restvalslice = restvalslice[:3]
	}

	for _, ord := range orders {
		order, err := res.orderRepo.GetOrder(&dto.ReqGetOneOrder{OrderId: ord.Id, UserId: userID})

		if err != nil {
			return nil, err
		}

		restname := order.OrderItems[0].RestaurantName

		for i, restval := range restvalslice {
			if restval.Restaurant == restname {
				restvalslice[i].Value++
			}
		}
	}

	sort.Slice(restvalslice, func(i, j int) bool {
		return restvalslice[i].Value > restvalslice[j].Value
	})

	var tiprests []*dto.RestaurantWithCategories

	for _, restval := range restvalslice {
		if restval.Value > 0 {
			rest, err := res.restaurantRepo.GetRestaurantByName(restval.Restaurant)
			if err != nil {
				return nil, err
			}
			mindeltime := len(rest.Name) + 15
			mindeltime = mindeltime - (mindeltime % 5)
			maxdeltime := mindeltime + 10
			maxdeltime = maxdeltime - (maxdeltime % 5)
			delprice := len(rest.Name)*8 + 200
			delprice = delprice - (delprice % 10)
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
			tiprests = append(tiprests, restwithcat)
		}
	}

	return tiprests, nil
}

func (res restaurantUsecase) GetRandomRestaurantTips() ([]*dto.RestaurantWithCategories, error) {
	restaurants, err := res.restaurantRepo.GetRestaurants()

	if err != nil {
		return nil, err
	}

	type RestVal struct {
		Restaurant string
		Value      int
	}

	var restvalslice []RestVal

	for _, restaurant := range restaurants {
		restval := RestVal{Restaurant: restaurant.Name, Value: rand.Intn(5)}
		restvalslice = append(restvalslice, restval)
	}

	sort.Slice(restvalslice, func(i, j int) bool {
		return restvalslice[i].Value > restvalslice[j].Value
	})

	restvalslice = restvalslice[:3]

	var tiprests []*dto.RestaurantWithCategories

	for _, restval := range restvalslice {
		if restval.Value > 0 {
			rest, err := res.restaurantRepo.GetRestaurantByName(restval.Restaurant)
			if err != nil {
				return nil, err
			}
			mindeltime := len(rest.Name) + 15
			mindeltime = mindeltime - (mindeltime % 5)
			maxdeltime := mindeltime + 10
			maxdeltime = maxdeltime - (maxdeltime % 5)
			delprice := len(rest.Name)*8 + 200
			delprice = delprice - (delprice % 10)
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
			tiprests = append(tiprests, restwithcat)
		}
	}

	return tiprests, nil

}
