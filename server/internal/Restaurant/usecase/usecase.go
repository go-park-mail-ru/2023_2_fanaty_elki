package usecase

import (
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
	GetRestaurants() ([]*dto.RestaurantWithCategories, error)
	GetRestaurantByID(id uint) (*dto.RestaurantWithCategories, error)
	GetRestaurantProducts(id uint) ([]*dto.MenuTypeWithProducts, error)
	GetRestaurantsByCategory(name string) ([]*dto.RestaurantWithCategories, error)
	GetCategories() (*[]string, error)
	Search(word string) ([]*dto.RestaurantWithCategoriesAndProducts, error)
	GetRestaurantTips(SessionToken string) ([]*dto.RestaurantWithCategories, error)
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
func (res RestaurantUsecase) GetRestaurants() ([]*dto.RestaurantWithCategories, error) {
	rests, err := res.restaurantRepo.GetRestaurants()
	if err != nil {
		return nil, entity.ErrInternalServerError
	}
	restswithcategories := []*dto.RestaurantWithCategories{}
	for _, rest := range rests {
		mindeltime := len(rest.Name) + 15
		maxdeltime := mindeltime + 10
		delprice := len(rest.Name)*8 + 200
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
	return restswithcategories, nil
}

//GetRestaurantByID gets restaurant by id
func (res RestaurantUsecase) GetRestaurantByID(id uint) (*dto.RestaurantWithCategories, error) {
	rest, err := res.restaurantRepo.GetRestaurantByID(id)
	if err != nil {
		return nil, err
	}
	mindeltime := len(rest.Name) + 15
	maxdeltime := mindeltime + 10
	delprice := len(rest.Name)*8 + 200
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
func (res RestaurantUsecase) GetRestaurantProducts(id uint) ([]*dto.MenuTypeWithProducts, error) {
	menuTypes, err := res.restaurantRepo.GetMenuTypesByRestaurantID(id)
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

//GetRestaurantsByCategory gets restaurants by category
func (res RestaurantUsecase) GetRestaurantsByCategory(name string) ([]*dto.RestaurantWithCategories, error) {
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
		maxdeltime := mindeltime + 10
		delprice := len(rest.Name)*8 + 200
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
	return restswithcategories, nil
}

//GetCategories gets categories
func (res RestaurantUsecase) GetCategories() (*[]string, error) {
	cats, err := res.restaurantRepo.GetCategories()
	categories := dto.TransformCategoriesToStringSlice(cats)
	if err != nil {
		return nil, entity.ErrInternalServerError
	}
	return categories, nil
}

//Search searches restaurant
func (res RestaurantUsecase) Search(word string) ([]*dto.RestaurantWithCategoriesAndProducts, error) {
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
		restID, err := res.productRepo.GetRestaurantIdByProduct(prod.ID)
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
		maxdeltime := mindeltime + 10
		delprice := len(rest.Name)*8 + 200
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
		restID, err := res.productRepo.GetRestaurantIdByProduct(prod.ID)
		if err != nil {
			return nil, entity.ErrInternalServerError
		}
		for _, rest := range restsWithCategoriesAndProducts {
			if rest.ID == restID {
				rest.Products = append(rest.Products, prod)
			}
		}
	}
	return restsWithCategoriesAndProducts, nil
}

//GetRestaurantTips gets restaurant by tips
func (res RestaurantUsecase) GetRestaurantTips(SessionToken string) ([]*dto.RestaurantWithCategories, error) {
	cookie, err := res.sessionRepo.Check(SessionToken)
	if err != nil {
		return nil, err
	}

	UserID := cookie.UserID

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

	orders, err := res.orderRepo.GetOrders(UserID)

	if err != nil {
		return nil, err
	}

	for _, ord := range orders {
		order, err := res.orderRepo.GetOrder(&dto.ReqGetOneOrder{OrderID: ord.ID, UserID: UserID})

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
			maxdeltime := mindeltime + 10
			delprice := len(rest.Name)*8 + 200
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

	return tiprests, nil
}
