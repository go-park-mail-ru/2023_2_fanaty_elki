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

	res := []*entity.Restaurant{
		{ID: 1,
			Name:          "Burger King",
			Rating:        3.7,
			CommentsCount: 60,
			Category:      "Fastfood",
			Icon:          "img/burger_king.jpg",
		},
		{ID: 2,
			Name:          "MacBurger",
			Rating:        3.8,
			CommentsCount: 69,
			Category:      "Fastfood",
			Icon:          "img/mac_burger.jpg",
		},
	}

	mockRest.EXPECT().GetRestaurants().Return(res, nil)
	actual, err := usecase.GetRestaurants()
	assert.Equal(t, res, actual)
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
		Category:      "Fastfood",
		Icon:          "img/burger_king.jpg",
	}

	var elemID = 1

	mockRest.EXPECT().GetRestaurantById(uint(elemID)).Return(rest, nil)
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

	// rest := &entity.Restaurant{
	// 	ID:            1,
	// 	Name:          "Burger King",
	// 	Rating:        3.7,
	// 	CommentsCount: 60,
	// 	Category:      "Fastfood",
	// 	Icon:          "img/burger_king.jpg",
	// }

	// menutypes := []*entity.MenuType{
	// 	{
	// 		ID:           1,
	// 		Name:         "Популярное",
	// 		RestaurantID: 1,
	// 	},
	// }

	testErr := errors.New("test")
	var elemID = 1

	mockRest.EXPECT().GetRestaurantById(uint(elemID)).Return(&entity.Restaurant{}, testErr)
	actual, err := usecase.GetRestaurantById(uint(elemID))
	assert.Empty(t, actual)
	assert.Equal(t, testErr, err)

	// mockRest.EXPECT().GetRestaurantById(uint(elemID)).Return(rest, nil)
	// mockRest.EXPECT().GetMenuTypesByRestaurantId(uint(elemID)).Return([]*entity.MenuType{}, testErr)
	// actual, err = usecase.GetRestaurantById(uint(elemID))
	// assert.Empty(t, actual)
	// assert.Equal(t, entity.ErrInternalServerError, err)

	// mockRest.EXPECT().GetRestaurantById(uint(elemID)).Return(rest, nil)
	// mockRest.EXPECT().GetMenuTypesByRestaurantId(uint(elemID)).Return(menutypes, nil)
	// mockProd.EXPECT().GetProductsByMenuTypeId(uint(elemID)).Return([]*entity.Product{}, testErr)
	// actual, err = usecase.GetRestaurantById(uint(elemID))
	// assert.Empty(t, actual)
	// assert.Equal(t, entity.ErrInternalServerError, err)

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
