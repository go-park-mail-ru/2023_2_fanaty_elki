package usecase

import (
	"errors"
	//"fmt"
	mockO "server/internal/Order/repository/mock_repository"
	mockP "server/internal/Product/repository/mock_repository"
	mockR "server/internal/Restaurant/repository/mock_repository"
	mockS "server/internal/Session/repository/mock_repository"
	"server/internal/domain/dto"
	"server/internal/domain/entity"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetRestaurantsSucces(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRest := mockR.NewMockRestaurantRepositoryI(ctrl)
	mockProd := mockP.NewMockProductRepositoryI(ctrl)
	mockSess := mockS.NewMockSessionRepositoryI(ctrl)
	mockOrd := mockO.NewMockOrderRepositoryI(ctrl)
	usecase := NewRestaurantUsecase(mockRest, mockProd, mockSess, mockOrd)

	// res := []*dto.RestaurantWithCategories{
	// 	{ID: 1,
	// 		Name:          "Burger King",
	// 		Rating:        3.7,
	// 		CommentsCount: 60,
	// 		Categories:    []string{"Burger", "Breakfast"},
	// 		Icon:          "img/burger_king.webp",
	// 	},
	// 	{ID: 2,
	// 		Name:          "MacBurger",
	// 		Rating:        3.8,
	// 		CommentsCount: 69,
	// 		Categories:    []string{"Burger", "Breakfast"},
	// 		Icon:          "img/mac_burger.webp",
	// 	},
	// }

	rest := []*entity.Restaurant{
		{ID: 1,
			Name:          "Burger King",
			Rating:        3.7,
			CommentsCount: 60,
			Icon:          "img/burger_king.webp",
		},
		{ID: 2,
			Name:          "MacBurger",
			Rating:        3.8,
			CommentsCount: 69,
			Icon:          "img/mac_burger.webp",
		},
	}

	categories := []*entity.Category{
		{
			ID:   1,
			Name: "Burger",
		},
		{
			ID:   2,
			Name: "Breakfast",
		},
	}

	mockRest.EXPECT().GetRestaurants().Return(rest, nil)
	mockRest.EXPECT().GetCategoriesByRestaurantID(rest[0].ID).Return(categories, nil)
	mockRest.EXPECT().GetCategoriesByRestaurantID(rest[1].ID).Return(categories, nil)
	_, err := usecase.GetRestaurants()
	//assert.Equal(t, res[0].Name, actual)
	assert.Nil(t, err)

}

func TestGetRestaurantsFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRest := mockR.NewMockRestaurantRepositoryI(ctrl)
	mockProd := mockP.NewMockProductRepositoryI(ctrl)
	mockSess := mockS.NewMockSessionRepositoryI(ctrl)
	mockOrd := mockO.NewMockOrderRepositoryI(ctrl)
	usecase := NewRestaurantUsecase(mockRest, mockProd, mockSess, mockOrd)

	testErr := errors.New("test")

	mockRest.EXPECT().GetRestaurants().Return([]*entity.Restaurant{}, testErr)
	actual, err := usecase.GetRestaurants()
	assert.Empty(t, actual)
	assert.Equal(t, entity.ErrInternalServerError, err)
}

func TestGetRestaurantByIdSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRest := mockR.NewMockRestaurantRepositoryI(ctrl)
	mockProd := mockP.NewMockProductRepositoryI(ctrl)
	mockSess := mockS.NewMockSessionRepositoryI(ctrl)
	mockOrd := mockO.NewMockOrderRepositoryI(ctrl)
	usecase := NewRestaurantUsecase(mockRest, mockProd, mockSess, mockOrd)

	rest := &entity.Restaurant{
		ID:            1,
		Name:          "Burger King",
		Rating:        3.7,
		CommentsCount: 60,
		Icon:          "img/burger_king.webp",
	}

	var elemID = 1

	categories := []*entity.Category{
		{
			ID:   1,
			Name: "Burger",
		},
		{
			ID:   2,
			Name: "Breakfast",
		},
	}

	mockRest.EXPECT().GetRestaurantByID(uint(elemID)).Return(rest, nil)
	mockRest.EXPECT().GetCategoriesByRestaurantID(uint(elemID)).Return(categories, nil)
	_, err := usecase.GetRestaurantByID(uint(elemID))
	//assert.Equal(t, rest.Name, actual.Name)
	assert.Nil(t, err)
}

func TestGetRestaurantByIdFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRest := mockR.NewMockRestaurantRepositoryI(ctrl)
	mockProd := mockP.NewMockProductRepositoryI(ctrl)
	mockSess := mockS.NewMockSessionRepositoryI(ctrl)
	mockOrd := mockO.NewMockOrderRepositoryI(ctrl)
	usecase := NewRestaurantUsecase(mockRest, mockProd, mockSess, mockOrd)

	testErr := errors.New("test")
	var elemID = 1

	mockRest.EXPECT().GetRestaurantByID(uint(elemID)).Return(&entity.Restaurant{}, testErr)
	actual, err := usecase.GetRestaurantByID(uint(elemID))
	assert.Empty(t, actual)
	assert.Equal(t, testErr, err)

}

func TestGetRestaurantProductsSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRest := mockR.NewMockRestaurantRepositoryI(ctrl)
	mockProd := mockP.NewMockProductRepositoryI(ctrl)
	mockSess := mockS.NewMockSessionRepositoryI(ctrl)
	mockOrd := mockO.NewMockOrderRepositoryI(ctrl)
	usecase := NewRestaurantUsecase(mockRest, mockProd, mockSess, mockOrd)

	menutypes := []*entity.MenuType{
		{
			ID:           1,
			Name:         "Популярное",
			RestaurantID: 1,
		},
	}

	products := []*entity.Product{
		{
			ID:          1,
			Name:        "Burger",
			Price:       120.0,
			CookingTime: 23,
			Portion:     "160 г",
			Description: "Nice burger",
			Icon:        "deficon",
		},
	}

	res := &dto.MenuTypeWithProductsSlice{
		{
			MenuType: &entity.MenuType{
				ID:           1,
				Name:         "Популярное",
				RestaurantID: 1,
			},
			Products: []*entity.Product{
				{
					ID:          1,
					Name:        "Burger",
					Price:       120.0,
					CookingTime: 23,
					Portion:     "160 г",
					Description: "Nice burger",
					Icon:        "deficon",
				},
			},
		},
	}

	var elemID = 1

	mockRest.EXPECT().GetMenuTypesByRestaurantID(uint(elemID)).Return(menutypes, nil)
	mockProd.EXPECT().GetProductsByMenuTypeID(uint(elemID)).Return(products, nil)
	actual, err := usecase.GetRestaurantProducts(uint(elemID))
	assert.Equal(t, res, actual)
	assert.Nil(t, err)
}

func TestGetRestaurantProductsFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRest := mockR.NewMockRestaurantRepositoryI(ctrl)
	mockProd := mockP.NewMockProductRepositoryI(ctrl)
	mockSess := mockS.NewMockSessionRepositoryI(ctrl)
	mockOrd := mockO.NewMockOrderRepositoryI(ctrl)
	usecase := NewRestaurantUsecase(mockRest, mockProd, mockSess, mockOrd)

	menutypes := []*entity.MenuType{
		{
			ID:           1,
			Name:         "Популярное",
			RestaurantID: 1,
		},
	}

	testErr := errors.New("test")
	var elemID = 1

	mockRest.EXPECT().GetMenuTypesByRestaurantID(uint(elemID)).Return([]*entity.MenuType{}, testErr)
	actual, err := usecase.GetRestaurantProducts(uint(elemID))
	assert.Empty(t, actual)
	assert.Equal(t, testErr, err)

	mockRest.EXPECT().GetMenuTypesByRestaurantID(uint(elemID)).Return(menutypes, nil)
	mockProd.EXPECT().GetProductsByMenuTypeID(uint(elemID)).Return([]*entity.Product{}, testErr)
	actual, err = usecase.GetRestaurantProducts(uint(elemID))
	assert.Empty(t, actual)
	assert.Equal(t, entity.ErrInternalServerError, err)
}

func TestGetRestaurantsByCategorySuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRest := mockR.NewMockRestaurantRepositoryI(ctrl)
	mockProd := mockP.NewMockProductRepositoryI(ctrl)
	mockSess := mockS.NewMockSessionRepositoryI(ctrl)
	mockOrd := mockO.NewMockOrderRepositoryI(ctrl)
	usecase := NewRestaurantUsecase(mockRest, mockProd, mockSess, mockOrd)

	// res := []*dto.RestaurantWithCategories{
	// 	{ID: 1,
	// 		Name:          "Burger King",
	// 		Rating:        3.7,
	// 		CommentsCount: 60,
	// 		Categories:    []string{"Burger", "Breakfast"},
	// 		Icon:          "img/burger_king.webp",
	// 	},
	// 	{ID: 2,
	// 		Name:          "MacBurger",
	// 		Rating:        3.8,
	// 		CommentsCount: 69,
	// 		Categories:    []string{"Burger", "Breakfast"},
	// 		Icon:          "img/mac_burger.webp",
	// 	},
	// }

	rest := []*entity.Restaurant{
		{ID: 1,
			Name:          "Burger King",
			Rating:        3.7,
			CommentsCount: 60,
			Icon:          "img/burger_king.webp",
		},
		{ID: 2,
			Name:          "MacBurger",
			Rating:        3.8,
			CommentsCount: 69,
			Icon:          "img/mac_burger.webp",
		},
	}

	categories := []*entity.Category{
		{
			ID:   1,
			Name: "Burger",
		},
		{
			ID:   2,
			Name: "Breakfast",
		},
	}

	mockRest.EXPECT().GetRestaurantsByCategory("Burger").Return(rest, nil)
	mockRest.EXPECT().GetCategoriesByRestaurantID(rest[0].ID).Return(categories, nil)
	mockRest.EXPECT().GetCategoriesByRestaurantID(rest[1].ID).Return(categories, nil)
	_, err := usecase.GetRestaurantsByCategory("Burger")
	//assert.Equal(t, res[0].Name, actual[0].Name)
	assert.Nil(t, err)

}

func TestGetRestaurantsByCategoryFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRest := mockR.NewMockRestaurantRepositoryI(ctrl)
	mockProd := mockP.NewMockProductRepositoryI(ctrl)
	mockSess := mockS.NewMockSessionRepositoryI(ctrl)
	mockOrd := mockO.NewMockOrderRepositoryI(ctrl)
	usecase := NewRestaurantUsecase(mockRest, mockProd, mockSess, mockOrd)

	testErr := errors.New("test")

	mockRest.EXPECT().GetRestaurantsByCategory("Burger").Return([]*entity.Restaurant{}, testErr)
	actual, err := usecase.GetRestaurantsByCategory("Burger")
	assert.Empty(t, actual)
	assert.Equal(t, entity.ErrInternalServerError, err)
}

func TestGetCategoriesSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRest := mockR.NewMockRestaurantRepositoryI(ctrl)
	mockProd := mockP.NewMockProductRepositoryI(ctrl)
	mockSess := mockS.NewMockSessionRepositoryI(ctrl)
	mockOrd := mockO.NewMockOrderRepositoryI(ctrl)
	usecase := NewRestaurantUsecase(mockRest, mockProd, mockSess, mockOrd)

	categories := []*entity.Category{
		{
			ID:   1,
			Name: "Burger",
		},
		{
			ID:   2,
			Name: "Breakfast",
		},
	}

	categoriesstring := dto.StringSlice([]string{"Burger", "Breakfast"})

	mockRest.EXPECT().GetCategories().Return(categories, nil)

	actual, err := usecase.GetCategories()
	assert.Equal(t, &categoriesstring, actual)
	assert.Nil(t, err)

}

func TestGetCategoriesFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRest := mockR.NewMockRestaurantRepositoryI(ctrl)
	mockProd := mockP.NewMockProductRepositoryI(ctrl)
	mockSess := mockS.NewMockSessionRepositoryI(ctrl)
	mockOrd := mockO.NewMockOrderRepositoryI(ctrl)
	usecase := NewRestaurantUsecase(mockRest, mockProd, mockSess, mockOrd)

	testErr := errors.New("test")

	mockRest.EXPECT().GetCategories().Return([]*entity.Category{}, testErr)
	actual, err := usecase.GetCategories()
	assert.Empty(t, actual)
	assert.Equal(t, entity.ErrInternalServerError, err)
}

func TestSearchSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRest := mockR.NewMockRestaurantRepositoryI(ctrl)
	mockProd := mockP.NewMockProductRepositoryI(ctrl)
	mockSess := mockS.NewMockSessionRepositoryI(ctrl)
	mockOrd := mockO.NewMockOrderRepositoryI(ctrl)
	usecase := NewRestaurantUsecase(mockRest, mockProd, mockSess, mockOrd)

	rests := []*entity.Restaurant{
		{ID: 1,
			Name:          "Burger King",
			Rating:        3.7,
			CommentsCount: 60,
			Icon:          "img/burger_king.webp",
		},
		{ID: 2,
			Name:          "MacBurger",
			Rating:        3.8,
			CommentsCount: 69,
			Icon:          "img/mac_burger.webp",
		},
	}

	restscat := []*entity.Restaurant{
		{ID: 2,
			Name:          "Burger King",
			Rating:        3.7,
			CommentsCount: 60,
			Icon:          "img/burger_king.webp",
		},
		{ID: 3,
			Name:          "MacBurger",
			Rating:        3.8,
			CommentsCount: 69,
			Icon:          "img/mac_burger.webp",
		},
	}

	products := []*entity.Product{
		{
			ID:          1,
			Name:        "Burger",
			Price:       120.0,
			CookingTime: 23,
			Portion:     "160 г",
			Description: "Nice burger",
			Icon:        "deficon",
		},
	}

	categories := []*entity.Category{
		{
			ID:   1,
			Name: "Burger",
		},
		{
			ID:   2,
			Name: "Breakfast",
		},
	}

	// expect := []*dto.RestaurantWithCategoriesAndProducts{
	// 	{ID: 1,
	// 		Name:          "Burger King",
	// 		Rating:        3.7,
	// 		CommentsCount: 60,
	// 		Icon:          "img/burger_king.webp",
	// 	},
	// 	{ID: 2,
	// 		Name:          "MacBurger",
	// 		Rating:        3.8,
	// 		CommentsCount: 69,
	// 		Icon:          "img/mac_burger.webp",
	// 	},
	// }

	var ID uint
	ID = 1

	mockRest.EXPECT().SearchRestaurants("Бургер").Return(rests, nil)
	mockRest.EXPECT().SearchCategories("Бургер").Return(restscat, nil)
	mockProd.EXPECT().SearchProducts("Бургер").Return(products, nil)
	mockProd.EXPECT().GetRestaurantIDByProduct(products[0].ID).Return(ID, nil)
	mockRest.EXPECT().GetCategoriesByRestaurantID(rests[0].ID).Return(categories, nil)
	mockRest.EXPECT().GetCategoriesByRestaurantID(rests[1].ID).Return(categories, nil)
	mockRest.EXPECT().GetCategoriesByRestaurantID(restscat[1].ID).Return(categories, nil)
	mockProd.EXPECT().GetRestaurantIDByProduct(products[0].ID).Return(ID, nil)

	_, err := usecase.Search("Бургер")
	//assert.Equal(t, expect[0].Name, actual[0].Name)
	assert.Nil(t, err)

}

func TestSearchFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRest := mockR.NewMockRestaurantRepositoryI(ctrl)
	mockProd := mockP.NewMockProductRepositoryI(ctrl)
	mockSess := mockS.NewMockSessionRepositoryI(ctrl)
	mockOrd := mockO.NewMockOrderRepositoryI(ctrl)
	usecase := NewRestaurantUsecase(mockRest, mockProd, mockSess, mockOrd)

	testErr := errors.New("test")

	mockRest.EXPECT().SearchRestaurants("Бургер").Return([]*entity.Restaurant{}, testErr)
	actual, err := usecase.Search("Бургер")
	assert.Empty(t, actual)
	assert.Equal(t, entity.ErrInternalServerError, err)
}

func TestGetRestaurantTipsSuccess(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRest := mockR.NewMockRestaurantRepositoryI(ctrl)
	mockProd := mockP.NewMockProductRepositoryI(ctrl)
	mockSess := mockS.NewMockSessionRepositoryI(ctrl)
	mockOrd := mockO.NewMockOrderRepositoryI(ctrl)
	usecase := NewRestaurantUsecase(mockRest, mockProd, mockSess, mockOrd)

	sestok := "Uuehdbye"

	var UserID uint
	UserID = 1

	cookie := entity.Cookie{
		UserID:       1,
		SessionToken: "TYebbYudb",
		MaxAge:       50 * time.Hour,
	}

	rests := []*entity.Restaurant{
		{ID: 1,
			Name:          "Burger King",
			Rating:        3.7,
			CommentsCount: 60,
			Icon:          "img/burger_king.webp",
		},
		{ID: 2,
			Name:          "MacBurger",
			Rating:        3.8,
			CommentsCount: 69,
			Icon:          "img/mac_burger.webp",
		},
	}

	resporders := []*dto.RespGetOrder{
		{
			ID:     1,
			Status: 0,
			Address: &dto.RespOrderAddress{
				City:   "Moscow",
				Street: "Tverskaya",
				House:  "2",
			},
			Price: 100,
		},
	}

	products := &dto.RespGetOrderProduct{
		ID:    1,
		Name:  "Burger",
		Price: 100,
		Icon:  "def",
		Count: 1,
	}

	orderItems := &dto.OrderItems{
		RestaurantName: "Burger King",
		Products:       []*dto.RespGetOrderProduct{products},
	}

	resporder := &dto.RespGetOneOrder{
		ID:     1,
		Status: 0,
		Address: &dto.RespOrderAddress{
			City:   "Moscow",
			Street: "Tverskaya",
			House:  "2",
			Flat:   1,
		},
		OrderItems: []*dto.OrderItems{orderItems},
	}

	categories := []*entity.Category{
		{
			ID:   1,
			Name: "Burger",
		},
		{
			ID:   2,
			Name: "Breakfast",
		},
	}

	// expect := &dto.RestaurantWithCategoriesSlice{
	// 	{ID: 1,
	// 		Name:          "Burger King",
	// 		Rating:        3.7,
	// 		CommentsCount: 60,
	// 		Categories:    []string{"Burger", "Breakfast"},
	// 		Icon:          "img/burger_king.webp",
	// 	},
	// 	{ID: 2,
	// 		Name:          "MacBurger",
	// 		Rating:        3.8,
	// 		CommentsCount: 69,
	// 		Categories:    []string{"Burger", "Breakfast"},
	// 		Icon:          "img/mac_burger.webp",
	// 	},
	// }

	reqorder := dto.ReqGetOneOrder{OrderID: 1, UserID: UserID}

	var ID uint
	ID = 1

	mockSess.EXPECT().Check(sestok).Return(&cookie, nil)
	mockRest.EXPECT().GetRestaurants().Return(rests, nil)
	mockOrd.EXPECT().GetOrders(ID).Return(resporders, nil)
	mockOrd.EXPECT().GetOrder(&reqorder).Return(resporder, nil)
	mockRest.EXPECT().GetRestaurantByName(orderItems.RestaurantName).Return(rests[0], nil)
	mockRest.EXPECT().GetCategoriesByRestaurantID(ID).Return(categories, nil)

	_, err := usecase.GetRestaurantTips(sestok)
	//assert.Equal(t, expect[0].Name, actual[0].Name)
	assert.Nil(t, err)
}

// func TestGetRandomRestaurantTipsSuccess(t *testing.T) {

// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockRest := mockR.NewMockRestaurantRepositoryI(ctrl)
// 	mockProd := mockP.NewMockProductRepositoryI(ctrl)
// 	mockSess := mockS.NewMockSessionRepositoryI(ctrl)
// 	mockOrd := mockO.NewMockOrderRepositoryI(ctrl)
// 	usecase := NewRestaurantUsecase(mockRest, mockProd, mockSess, mockOrd)

// 	rests := []*entity.Restaurant{
// 		{ID: 1,
// 			Name:          "Burger King",
// 			Rating:        3.7,
// 			CommentsCount: 60,
// 			Icon:          "img/burger_king.webp",
// 		},
// 		{ID: 2,
// 			Name:          "MacBurger",
// 			Rating:        3.8,
// 			CommentsCount: 69,
// 			Icon:          "img/mac_burger.webp",
// 		},
// 		{ID: 3,
// 			Name:          "MacBurger",
// 			Rating:        3.8,
// 			CommentsCount: 69,
// 			Icon:          "img/mac_burger.webp",
// 		},
// 	}

// 	categories := []*entity.Category{
// 		{
// 			ID:   1,
// 			Name: "Burger",
// 		},
// 		{
// 			ID:   2,
// 			Name: "Breakfast",
// 		},
// 	}

// 	// expect := []*dto.RestaurantWithCategories{
// 	// 	{ID: 1,
// 	// 		Name:          "MacBurger",
// 	// 		Rating:        3.7,
// 	// 		CommentsCount: 60,
// 	// 		Categories:    []string{"Burger", "Breakfast"},
// 	// 		Icon:          "img/burger_king.webp",
// 	// 	},
// 	// 	{ID: 2,
// 	// 		Name:          "MacBurger",
// 	// 		Rating:        3.8,
// 	// 		CommentsCount: 69,
// 	// 		Categories:    []string{"Burger", "Breakfast"},
// 	// 		Icon:          "img/mac_burger.webp",
// 	// 	},
// 	// }

// 	mockRest.EXPECT().GetRestaurants().Return(rests, nil)
// 	mockRest.EXPECT().GetRestaurantByName(rests[0].Name).Return(rests[0], nil)
// 	mockRest.EXPECT().GetRestaurantByName(rests[1].Name).Return(rests[1], nil)
// 	//mockRest.EXPECT().GetRestaurantByName(rests[2].Name).Return(rests[2], nil)
// 	mockRest.EXPECT().GetCategoriesByRestaurantID(rests[0].ID).Return(categories, nil)
// 	mockRest.EXPECT().GetCategoriesByRestaurantID(rests[1].ID).Return(categories, nil)
// 	//mockRest.EXPECT().GetCategoriesByRestaurantID(rests[2].ID).Return(categories, nil)

// 	_, err := usecase.GetRandomRestaurantTips()
// 	//fmt.Println(actual)
// 	//assert.Equal(t, expect[0].Name, actual[0].Name)
// 	assert.Nil(t, err)
// }
