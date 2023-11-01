package usecase

import (
	"errors"
	mockR "server/internal/Restaurant/repository/mock_repository"
	"server/internal/domain/entity"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetRestaurants(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRest := mockR.NewMockRestaurantRepositoryI(ctrl)
	usecase := NewRestaurantUsecase(mockRest)

	res := []*entity.Restaurant{
		{1,
			"Burger King",
			3.7,
			60,
			"Fastfood",
			"img/burger_king.jpg",
		},
		{2,
			"MacBurger",
			3.8,
			69,
			"Fastfood",
			"img/mac_burger.jpg",
		},
	}

	mockRest.EXPECT().GetRestaurants().Return(res, nil)
	actual, err := usecase.GetRestaurants()
	assert.Equal(t, res, actual)
	assert.Nil(t, err)

	testErr := errors.New("test")

	mockRest.EXPECT().GetRestaurants().Return([]*entity.Restaurant{}, testErr)
	actual, err = usecase.GetRestaurants()
	assert.Equal(t, []*entity.Restaurant{}, actual)
	assert.Equal(t, testErr, err)

}
