package usecase

import (
	mockC "server/internal/Cart/repository/mock_repository"
	mockO "server/internal/Order/repository/mock_repository"
	mockP "server/internal/Product/repository/mock_repository"
	dto "server/internal/domain/dto"
	"server/internal/domain/entity"
	"testing"

	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// func TestCreateOrderSuccess(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockOrder := mockO.NewMockOrderRepositoryI(ctrl)
// 	mockCart := mockC.NewMockCartRepositoryI(ctrl)
// 	mockProduct := mockP.NewMockProductRepositoryI(ctrl)
// 	usecase := NewOrderUsecase(mockOrder, mockCart, mockProduct)

// 	var flat uint
// 	flat = 1

// 	timenow := time.Now()
// 	reqorder := &dto.ReqCreateOrder{
// 		UserID: 1,
// 		Address: &dto.ReqCreateOrderAddress{
// 			City:   "Moscow",
// 			Street: "Tverskaya",
// 			House:  "2",
// 			Flat:   flat,
// 		},
// 	}

// 	order := dto.ToEntityCreateOrder(reqorder)
// 	order.Date = timenow

// 	cart := &entity.Cart{
// 		ID:     1,
// 		UserID: 1,
// 	}

// 	mockCart.EXPECT().GetCartByUserID(order.UserID).Return(cart, nil)

// 	cartwithrest := &entity.CartWithRestaurant{
// 		RestaurantID: 1,
// 		Products: []*entity.CartProduct{
// 			{
// 				ID:        1,
// 				ProductID: 1,
// 				CartID:    1,
// 				ItemCount: 6,
// 			},
// 			{
// 				ID:        2,
// 				ProductID: 3,
// 				CartID:    1,
// 				ItemCount: 6,
// 			},
// 		},
// 	}

// 	mockCart.EXPECT().GetCartProductsByCartID(cart.ID).Return(cartwithrest, nil)

// 	product := &entity.Product{
// 		ID:          1,
// 		Name:        "Burger",
// 		Price:       100,
// 		CookingTime: 10,
// 		Portion:     "120g",
// 		Icon:        "def",
// 	}
// 	mockProduct.EXPECT().GetProductByID(cartwithrest.Products[0].ProductID).Return(product, nil)
// 	mockProduct.EXPECT().GetProductByID(cartwithrest.Products[1].ProductID).Return(product, nil)

// 	order.Price += uint(product.Price) * uint(cartwithrest.Products[0].ItemCount)

// 	resporder := &dto.RespCreateOrder{
// 		ID:     1,
// 		Status: 0,
// 		Price:  uint(product.Price) * uint(cartwithrest.Products[0].ItemCount),
// 		Date:   timenow,
// 		Address: &entity.Address{
// 			City:   "Moscow",
// 			Street: "Tverskaya",
// 			House:  "2",
// 			Flat:   flat,
// 		},

// 		DeliveryTime: 30,
// 	}

// 	resorder := &dto.RespCreateOrder{
// 		ID:           1,
// 		Status:       0,
// 		Price:        uint(product.Price) * uint(cartwithrest.Products[0].ItemCount),
// 		Date:         timenow,
// 		DeliveryTime: 30,
// 		Address:      nil,
// 	}

// 	mockOrder.EXPECT().CreateOrder(gomock.Any()).Return(resorder, nil)
// 	actual, err := usecase.CreateOrder(reqorder)
// 	assert.Equal(t, resporder, actual)
// 	assert.Nil(t, err)

// }

func TestCreateOrderFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOrder := mockO.NewMockOrderRepositoryI(ctrl)
	mockCart := mockC.NewMockCartRepositoryI(ctrl)
	mockProduct := mockP.NewMockProductRepositoryI(ctrl)

	usecase := NewOrderUsecase(mockOrder, mockCart, mockProduct)

	var flat uint
	flat = 1

	reqorder := &dto.ReqCreateOrder{
		UserID: 1,
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
	mockCart := mockC.NewMockCartRepositoryI(ctrl)
	mockProduct := mockP.NewMockProductRepositoryI(ctrl)
	usecase := NewOrderUsecase(mockOrder, mockCart, mockProduct)

	dbreqorder := dto.ReqUpdateOrder{
		ID:     1,
		Status: 1,
	}

	reqorder := dto.ReqUpdateOrder{
		ID:     1,
		Status: 1,
	}

	mockOrder.EXPECT().UpdateOrder(&dbreqorder).Return(nil)
	err := usecase.UpdateOrder(&reqorder)

	assert.Nil(t, err)

}

func TestGetOrdersSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOrder := mockO.NewMockOrderRepositoryI(ctrl)
	mockCart := mockC.NewMockCartRepositoryI(ctrl)
	mockProduct := mockP.NewMockProductRepositoryI(ctrl)
	usecase := NewOrderUsecase(mockOrder, mockCart, mockProduct)

	var UserID uint
	UserID = 1

	var flat uint
	flat = 1

	timenow := time.Now()

	resporders := []*dto.RespGetOrder{
		{
			ID:     1,
			Status: 0,
			Date:   timenow,
			Address: &dto.RespOrderAddress{
				City:   "Moscow",
				Street: "Tverskaya",
				House:  "2",
				Flat:   flat,
			},
			Price: 100,
		},
		{
			ID:     2,
			Status: 0,
			Date:   timenow,
			Address: &dto.RespOrderAddress{
				City:   "Moscow",
				Street: "Tverskaya",
				House:  "3",
				Flat:   flat,
			},
			Price: 200,
		},
	}

	exp := &dto.RespOrders{
		{
			ID:     1,
			Status: 0,
			Date:   timenow,
			Address: &dto.RespOrderAddress{
				City:   "Moscow",
				Street: "Tverskaya",
				House:  "2",
				Flat:   flat,
			},
			Price: 100,
		},
		{
			ID:     2,
			Status: 0,
			Date:   timenow,
			Address: &dto.RespOrderAddress{
				City:   "Moscow",
				Street: "Tverskaya",
				House:  "3",
				Flat:   flat,
			},
			Price: 200,
		},
	}

	mockOrder.EXPECT().GetOrders(UserID).Return(resporders, nil)
	actual, err := usecase.GetOrders(UserID)
	assert.Equal(t, exp, actual)
	assert.Nil(t, err)

}

func TestGetOrderSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOrder := mockO.NewMockOrderRepositoryI(ctrl)
	mockCart := mockC.NewMockCartRepositoryI(ctrl)
	mockProduct := mockP.NewMockProductRepositoryI(ctrl)
	usecase := NewOrderUsecase(mockOrder, mockCart, mockProduct)

	timenow := time.Now()

	reqorder := &dto.ReqGetOneOrder{
		OrderID: 1,
		UserID:  1,
	}

	products := &dto.RespGetOrderProduct{
		ID:    1,
		Name:  "Burger",
		Price: 100,
		Icon:  "def",
		Count: 1,
	}

	orderItems := &dto.OrderItems{
		RestaurantName: "BK",
		Products:       []*dto.RespGetOrderProduct{products},
	}

	resporder := &dto.RespGetOneOrder{
		ID:     1,
		Status: 0,
		Date:   timenow,
		Address: &dto.RespOrderAddress{
			City:   "Moscow",
			Street: "Tverskaya",
			House:  "2",
			Flat:   1,
		},
		OrderItems: []*dto.OrderItems{orderItems},
	}

	mockOrder.EXPECT().GetOrder(reqorder).Return(resporder, nil)
	actual, err := usecase.GetOrder(reqorder)
	assert.Equal(t, resporder, actual)
	assert.Nil(t, err)

}
