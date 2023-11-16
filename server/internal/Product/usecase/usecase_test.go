package usecase

import (
	//"errors"
	"errors"
	mockP "server/internal/Product/repository/mock_repository"
	//"server/internal/domain/dto"
	"server/internal/domain/entity"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetProductByIDSucces(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProd := mockP.NewMockProductRepositoryI(ctrl)
	usecase := NewProductUsecase(mockProd)

	res := &entity.Product{
		ID:          1,
		Name:        "Burger",
		Price:       120.0,
		CookingTime: 23,
		Portion:     "160 г",
		Description: "Nice burger",
		Icon:        "deficon",
	}

	mockProd.EXPECT().GetProductByID(uint(1)).Return(res, nil)
	actual, err := usecase.GetProductByID(uint(1))
	assert.Equal(t, res, actual)
	assert.Nil(t, err)

}

func TestGetProductByIDFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProd := mockP.NewMockProductRepositoryI(ctrl)
	usecase := NewProductUsecase(mockProd)

	testErr := errors.New("testErr")

	mockProd.EXPECT().GetProductByID(uint(1)).Return(nil, testErr)
	actual, err := usecase.GetProductByID(uint(1))
	assert.Equal(t, err, entity.ErrInternalServerError)
	assert.Nil(t, actual)

	mockProd.EXPECT().GetProductByID(uint(1)).Return(nil, nil)
	actual, err = usecase.GetProductByID(uint(1))
	assert.Equal(t, err, entity.ErrNotFound)
	assert.Nil(t, actual)
}
