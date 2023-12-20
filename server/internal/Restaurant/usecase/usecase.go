package usecase

import (
	"fmt"
	"math/rand"

	orderRep "server/internal/Order/repository"
	productRep "server/internal/Product/repository"
	restRep "server/internal/Restaurant/repository"
	sessionRep "server/internal/Session/repository"
	dto "server/internal/domain/dto"
	"server/internal/domain/entity"
	"sort"
)

//RestaurantUsecaseI interface
type RestaurantUsecaseI interface {
	GetRestaurants() (*dto.RestaurantWithCategoriesSlice, error)
	GetRestaurantByID(id uint) (*dto.RestaurantWithCategories, error)
	GetRestaurantProducts(id uint) (*dto.MenuTypeWithProductsSlice, error)
	GetRestaurantsByCategory(name string) (*dto.RestaurantWithCategoriesSlice, error)
	GetCategories() (*dto.StringSlice, error)
	Search(word string) (*dto.RestaurantWithCategoriesAndProductsSlice, error)
	GetRestaurantTips(SessionToken string) (*dto.RestaurantWithCategoriesSlice, error)
	GetRandomRestaurantTips() (*dto.RestaurantWithCategoriesSlice, error)
}

//RestaurantUsecase struct
type RestaurantUsecase struct {
	restaurantRepo restRep.RestaurantRepositoryI
	productRepo    productRep.ProductRepositoryI
	sessionRepo    sessionRep.SessionRepositoryI
	orderRepo      orderRep.OrderRepositoryI
}

//NewRestaurantUsecase creates new restaurant usecase
func NewRestaurantUsecase(resRep restRep.RestaurantRepositoryI, productRep productRep.ProductRepositoryI, sessionRep sessionRep.SessionRepositoryI, orderRep orderRep.OrderRepositoryI) *RestaurantUsecase {
	return &RestaurantUsecase{
		restaurantRepo: resRep,
		productRepo:    productRep,
		sessionRepo:    sessionRep,
		orderRepo:      orderRep,
	}

}

//GetRestaurants gets restaurants
func (res RestaurantUsecase) GetRestaurants() (*dto.RestaurantWithCategoriesSlice, error) {
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
		cats, err := res.restaurantRepo.GetCategoriesByRestaurantID(rest.ID)
		if err != nil {
			if err != entity.ErrNotFound {
				return nil, entity.ErrInternalServerError
			}
		}
		restwithcat := dto.ToRestaurantWithCategories(rest, cats)
		restswithcategories = append(restswithcategories, restwithcat)
	}

	var respRestWithCategories dto.RestaurantWithCategoriesSlice

	for _, rest := range restswithcategories {
		respRestWithCategories = append(respRestWithCategories, rest)
	}
	return &respRestWithCategories, nil
}

//GetRestaurantByID gets restaurant by id
func (res RestaurantUsecase) GetRestaurantByID(id uint) (*dto.RestaurantWithCategories, error) {
	rest, err := res.restaurantRepo.GetRestaurantByID(id)
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
	cats, err := res.restaurantRepo.GetCategoriesByRestaurantID(rest.ID)
	if err != nil {
		if err != entity.ErrNotFound {
			return nil, entity.ErrInternalServerError
		}
	}
	restwithcat := dto.ToRestaurantWithCategories(rest, cats)
	return restwithcat, nil
}

//GetRestaurantProducts gets products from restaurant
func (res RestaurantUsecase) GetRestaurantProducts(id uint) (*dto.MenuTypeWithProductsSlice, error) {
	menuTypes, err := res.restaurantRepo.GetMenuTypesByRestaurantID(id)
	if err != nil {
		return nil, err
	}
	var menuTypesWithProducts []*dto.MenuTypeWithProducts
	for _, menu := range menuTypes {
		products, err := res.productRepo.GetProductsByMenuTypeID(menu.ID)
		if err != nil {
			return nil, entity.ErrInternalServerError
		}
		menuTypeWithProducts := dto.MenuTypeWithProducts{
			MenuType: menu,
			Products: products,
		}
		menuTypesWithProducts = append(menuTypesWithProducts, &menuTypeWithProducts)
	}

	var respMenu dto.MenuTypeWithProductsSlice

	for _, rest := range menuTypesWithProducts {
		respMenu = append(respMenu, rest)
	}

	return &respMenu, nil
}

//GetRestaurantsByCategory gets restaurants by category
func (res RestaurantUsecase) GetRestaurantsByCategory(name string) (*dto.RestaurantWithCategoriesSlice, error) {
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
		cats, err := res.restaurantRepo.GetCategoriesByRestaurantID(rest.ID)
		if err != nil {
			if err != entity.ErrNotFound {
				return nil, entity.ErrInternalServerError
			}
		}
		restwithcat := dto.ToRestaurantWithCategories(rest, cats)
		restswithcategories = append(restswithcategories, restwithcat)
	}

	var respRestWithCategories dto.RestaurantWithCategoriesSlice

	for _, rest := range restswithcategories {
		respRestWithCategories = append(respRestWithCategories, rest)
	}
	return &respRestWithCategories, nil
}

//GetCategories gets categories
func (res RestaurantUsecase) GetCategories() (*dto.StringSlice, error) {
	cats, err := res.restaurantRepo.GetCategories()
	categories := dto.TransformCategoriesToStringSlice(cats)
	if err != nil {
		return nil, entity.ErrInternalServerError
	}

	var respRestWithCategories dto.StringSlice

	for _, rest := range *categories {
		fmt.Println("usecase ", rest)
		resp := &rest
		fmt.Println("usecase ", resp)
		respRestWithCategories = append(respRestWithCategories, rest)
		fmt.Println("usecase ", &respRestWithCategories)
	}
	//fmt.Println(&respRestWithCategories[0], &respRestWithCategories[1], &respRestWithCategories[2])
	return &respRestWithCategories, nil
}

//Search searches restaurant
func (res RestaurantUsecase) Search(word string) (*dto.RestaurantWithCategoriesAndProductsSlice, error) {
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
		restID, err := res.productRepo.GetRestaurantIDByProduct(prod.ID)
		if err != nil {
			return nil, entity.ErrInternalServerError
		}
		if !restset[restID] {
			restByID, err := res.restaurantRepo.GetRestaurantByID(restID)
			if err != nil {
				return nil, entity.ErrInternalServerError
			}
			restset[restID] = true
			rests = append(rests, restByID)
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
		cats, err := res.restaurantRepo.GetCategoriesByRestaurantID(rest.ID)
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
		restID, err := res.productRepo.GetRestaurantIDByProduct(prod.ID)
		if err != nil {
			return nil, entity.ErrInternalServerError
		}
		for _, rest := range restsWithCategoriesAndProducts {
			if rest.ID == restID {
				rest.Products = append(rest.Products, prod)
			}
		}
	}

	var respRestWithCategories dto.RestaurantWithCategoriesAndProductsSlice

	for _, rest := range restsWithCategoriesAndProducts {
		respRestWithCategories = append(respRestWithCategories, rest)
	}
	return &respRestWithCategories, nil
}

//GetRestaurantTips gets restaurant by tips
func (res RestaurantUsecase) GetRestaurantTips(SessionToken string) (*dto.RestaurantWithCategoriesSlice, error) {
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

	cookie, err := res.sessionRepo.Check(SessionToken)
	if err != nil {
		return nil, err
	}

	fmt.Println("usecase", cookie)

	var orders []*dto.RespGetOrder

	if cookie != nil {
		userID := cookie.UserID

		orders, err = res.orderRepo.GetOrders(userID)

		if err != nil {
			return nil, err
		}

		for _, ord := range orders {
			order, err := res.orderRepo.GetOrder(&dto.ReqGetOneOrder{OrderID: ord.ID, UserID: userID})

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
	}

	if len(orders) == 0 {
		for i, _ := range restvalslice {
			restvalslice[i].Value = rand.Intn(5)
		}
		restvalslice = restvalslice[:3]
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
			cats, err := res.restaurantRepo.GetCategoriesByRestaurantID(rest.ID)
			if err != nil {
				if err != entity.ErrNotFound {
					return nil, entity.ErrInternalServerError
				}
			}
			restwithcat := dto.ToRestaurantWithCategories(rest, cats)
			tiprests = append(tiprests, restwithcat)
		}
	}

	var respRestWithCategories dto.RestaurantWithCategoriesSlice

	for _, rest := range tiprests {
		respRestWithCategories = append(respRestWithCategories, rest)
	}

	return &respRestWithCategories, nil
}

//GetRandomRestaurantTips gets random restaurants as a tip
func (res RestaurantUsecase) GetRandomRestaurantTips() (*dto.RestaurantWithCategoriesSlice, error) {
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
			cats, err := res.restaurantRepo.GetCategoriesByRestaurantID(rest.ID)
			if err != nil {
				if err != entity.ErrNotFound {
					return nil, entity.ErrInternalServerError
				}
			}
			restwithcat := dto.ToRestaurantWithCategories(rest, cats)
			tiprests = append(tiprests, restwithcat)
		}
	}

	var respRestWithCategories dto.RestaurantWithCategoriesSlice

	for _, rest := range tiprests {
		respRestWithCategories = append(respRestWithCategories, rest)
	}

	return &respRestWithCategories, nil

}
