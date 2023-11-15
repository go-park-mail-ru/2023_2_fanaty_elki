package usecase

import (
	mockO "server/internal/Order/repository/mock_repository"
	dto "server/internal/domain/dto"
	"server/internal/domain/entity"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrderSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOrder := mockO.NewMockOrderRepositoryI(ctrl)
	usecase := NewOrderUsecase(mockOrder)

	var flat uint
	flat = 1

	timenow := time.Now()

	reqorder := &dto.ReqCreateOrder{
		Products: []uint{1, 2, 3},
		UserId:   1,
		Address: &dto.ReqCreateOrderAddress{
			City:   "Moscow",
			Street: "Tverskaya",
			House:  "2",
			Flat:   flat,
		},
	}

	resporder := &dto.RespCreateOrder{
		Id:     1,
		Status: "CREATED",
		Date:   timenow,
	}

	products := make(map[uint]int)
	for _, product := range reqorder.Products {
		products[product]++
	}

	mockOrder.EXPECT().CreateOrder(gomock.Any()).Return(resporder, nil)
	actual, err := usecase.CreateOrder(reqorder)
	assert.Equal(t, resporder, actual)
	assert.Nil(t, err)

}

func TestCreateOrderFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOrder := mockO.NewMockOrderRepositoryI(ctrl)
	usecase := NewOrderUsecase(mockOrder)

	var flat uint
	flat = 1

	reqorder := &dto.ReqCreateOrder{
		Products: []uint{1, 2, 3},
		UserId:   1,
		Address: &dto.ReqCreateOrderAddress{
			City:   "",
			Street: "Tverskaya",
			House:  "2",
			Flat:   flat,
		},
	}

	actual, err := usecase.CreateOrder(reqorder)
	assert.Equal(t, entity.ErrBadRequest, err)
	assert.Nil(t, actual)

}

func TestUpdateOrderSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOrder := mockO.NewMockOrderRepositoryI(ctrl)
	usecase := NewOrderUsecase(mockOrder)

	dbreqorder := dto.ReqUpdateOrder{
		Id:     1,
		Status: "Wait",
	}

	reqorder := dto.ReqUpdateOrder{
		Id:     1,
		Status: "Wait",
	}

	mockOrder.EXPECT().UpdateOrder(&dbreqorder).Return(nil)
	err := usecase.UpdateOrder(&reqorder)

	assert.Nil(t, err)

}

func TestGetOrdersSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOrder := mockO.NewMockOrderRepositoryI(ctrl)
	usecase := NewOrderUsecase(mockOrder)

	var UserID uint
	UserID = 1

	var flat uint
	flat = 1

	timenow := time.Now()

	resporders := []*dto.RespGetOrder{
		{
			Id:     1,
			Status: "Wait",
			Date:   timenow,
			Address: &dto.RespOrderAddress{
				City:   "Moscow",
				Street: "Tverskaya",
				House:  "2",
				Flat:   &flat,
			},
		},
		{
			Id:     2,
			Status: "Wait",
			Date:   timenow,
			Address: &dto.RespOrderAddress{
				City:   "Moscow",
				Street: "Tverskaya",
				House:  "3",
				Flat:   &flat,
			},
		},
	}

	mockOrder.EXPECT().GetOrders(UserID).Return(resporders, nil)
	actual, err := usecase.GetOrders(UserID)
	assert.Equal(t, resporders, actual)
	assert.Nil(t, err)

}

func TestGetOrderSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOrder := mockO.NewMockOrderRepositoryI(ctrl)
	usecase := NewOrderUsecase(mockOrder)

	timenow := time.Now()

	reqorder := &dto.ReqGetOneOrder{
		OrderId: 1,
		UserId:  1,
	}

	resporder := &dto.RespGetOneOrder{
		Status:      "Wait",
		Date:        timenow,
		UpdatedDate: timenow,
		Products:    []*dto.RespGetOrderProduct{},
	}

	mockOrder.EXPECT().GetOrder(reqorder).Return(resporder, nil)
	actual, err := usecase.GetOrder(reqorder)
	assert.Equal(t, resporder, actual)
	assert.Nil(t, err)

}
