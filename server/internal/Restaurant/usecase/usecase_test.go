package usecase

import (
	"errors"
	mockP "server/internal/Product/repository/mock_repository"
	mockR "server/internal/Restaurant/repository/mock_repository"
	"server/internal/domain/dto"
	"server/internal/domain/entity"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetRestaurantsSucces(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRest := mockR.NewMockRestaurantRepositoryI(ctrl)
	mockProd := mockP.NewMockProductRepositoryI(ctrl)
	usecase := NewRestaurantUsecase(mockRest, mockProd)

	res := []*dto.RestaurantWithCategories{
		{ID: 1,
			Name:          "Burger King",
			Rating:        3.7,
			CommentsCount: 60,
			Categories:    []string{"Burger", "Breakfast"},
			Icon:          "img/burger_king.jpg",
		},
		{ID: 2,
			Name:          "MacBurger",
			Rating:        3.8,
			CommentsCount: 69,
			Categories:    []string{"Burger", "Breakfast"},
			Icon:          "img/mac_burger.jpg",
		},
	}

	rest := []*entity.Restaurant{
		{ID: 1,
			Name:          "Burger King",
			Rating:        3.7,
			CommentsCount: 60,
			Icon:          "img/burger_king.jpg",
		},
		{ID: 2,
			Name:          "MacBurger",
			Rating:        3.8,
			CommentsCount: 69,
			Icon:          "img/mac_burger.jpg",
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
	mockRest.EXPECT().GetCategoriesByRestaurantId(rest[0].ID).Return(categories, nil)
	mockRest.EXPECT().GetCategoriesByRestaurantId(rest[1].ID).Return(categories, nil)
	actual, err := usecase.GetRestaurants()
	assert.Equal(t, res[0].Name, actual[0].Name)
	assert.Nil(t, err)

}

func TestGetRestaurantsFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRest := mockR.NewMockRestaurantRepositoryI(ctrl)
	mockProd := mockP.NewMockProductRepositoryI(ctrl)
	usecase := NewRestaurantUsecase(mockRest, mockProd)

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
	usecase := NewRestaurantUsecase(mockRest, mockProd)

	rest := &entity.Restaurant{
		ID:            1,
		Name:          "Burger King",
		Rating:        3.7,
		CommentsCount: 60,
		Icon:          "img/burger_king.jpg",
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

	mockRest.EXPECT().GetRestaurantById(uint(elemID)).Return(rest, nil)
	mockRest.EXPECT().GetCategoriesByRestaurantId(uint(elemID)).Return(categories, nil)
	actual, err := usecase.GetRestaurantById(uint(elemID))
	assert.Equal(t, rest.Name, actual.Name)
	assert.Nil(t, err)
}

func TestGetRestaurantByIdFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRest := mockR.NewMockRestaurantRepositoryI(ctrl)
	mockProd := mockP.NewMockProductRepositoryI(ctrl)
	usecase := NewRestaurantUsecase(mockRest, mockProd)

	testErr := errors.New("test")
	var elemID = 1

	mockRest.EXPECT().GetRestaurantById(uint(elemID)).Return(&entity.Restaurant{}, testErr)
	actual, err := usecase.GetRestaurantById(uint(elemID))
	assert.Empty(t, actual)
	assert.Equal(t, testErr, err)

}

func TestGetRestaurantProductsSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRest := mockR.NewMockRestaurantRepositoryI(ctrl)
	mockProd := mockP.NewMockProductRepositoryI(ctrl)
	usecase := NewRestaurantUsecase(mockRest, mockProd)

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

	res := []*dto.MenuTypeWithProducts{
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

	mockRest.EXPECT().GetMenuTypesByRestaurantId(uint(elemID)).Return(menutypes, nil)
	mockProd.EXPECT().GetProductsByMenuTypeId(uint(elemID)).Return(products, nil)
	actual, err := usecase.GetRestaurantProducts(uint(elemID))
	assert.Equal(t, res, actual)
	assert.Nil(t, err)
}

func TestGetRestaurantProductsFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRest := mockR.NewMockRestaurantRepositoryI(ctrl)
	mockProd := mockP.NewMockProductRepositoryI(ctrl)
	usecase := NewRestaurantUsecase(mockRest, mockProd)

	menutypes := []*entity.MenuType{
		{
			ID:           1,
			Name:         "Популярное",
			RestaurantID: 1,
		},
	}

	testErr := errors.New("test")
	var elemID = 1

	mockRest.EXPECT().GetMenuTypesByRestaurantId(uint(elemID)).Return([]*entity.MenuType{}, testErr)
	actual, err := usecase.GetRestaurantProducts(uint(elemID))
	assert.Empty(t, actual)
	assert.Equal(t, testErr, err)

	mockRest.EXPECT().GetMenuTypesByRestaurantId(uint(elemID)).Return(menutypes, nil)
	mockProd.EXPECT().GetProductsByMenuTypeId(uint(elemID)).Return([]*entity.Product{}, testErr)
	actual, err = usecase.GetRestaurantProducts(uint(elemID))
	assert.Empty(t, actual)
	assert.Equal(t, entity.ErrInternalServerError, err)
}
