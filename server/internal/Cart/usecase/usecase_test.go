package usecase

import (
	"testing"

	mockC "server/internal/Cart/repository/mock_repository"
	mockP "server/internal/Product/repository/mock_repository"
	mockS "server/internal/Session/repository/mock_repository"
	dto "server/internal/domain/dto"
	"server/internal/domain/entity"

	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
)

func TestGetUserCartSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCart := mockC.NewMockCartRepositoryI(ctrl)
	mockProd := mockP.NewMockProductRepositoryI(ctrl)
	mockSes := mockS.NewMockSessionRepositoryI(ctrl)
	usecase := NewCartUsecase(mockCart, mockProd, mockSes)

	cookie := &entity.Cookie{
		UserID:       1,
		SessionToken: "HJJvgsvd",
	}

	cart := &entity.Cart{
		ID:     1,
		UserID: 1,
	}

	prod := &entity.Product{
		ID:          1,
		Name:        "Burger",
		Price:       120.0,
		CookingTime: 23,
		Portion:     "160 г",
		Description: "Nice burger",
		Icon:        "deficon",
	}

	cartProds := []*entity.CartProduct{
		{
			ID:        1,
			ProductID: 1,
			CartID:    1,
			ItemCount: 1,
		},
	}

	res := []*dto.CartProduct{
		{Product: &entity.Product{
			ID:          1,
			Name:        "Burger",
			Price:       120.0,
			CookingTime: 23,
			Portion:     "160 г",
			Description: "Nice burger",
			Icon:        "deficon",
		},
			ItemCount: 1,
		},
	}

	var UserID uint
	UserID = 1

	mockSes.EXPECT().Check(cookie.SessionToken).Return(cookie, nil)
	mockCart.EXPECT().GetCartByUserID(UserID).Return(cart, nil)
	mockCart.EXPECT().GetCartProductsByCartID(cart.ID).Return(cartProds, nil)
	mockProd.EXPECT().GetProductByID(res[0].Product.ID).Return(prod, nil)
	actual, err := usecase.GetUserCart(cookie.SessionToken)
	assert.Equal(t, res, actual)
	assert.Nil(t, err)

}

func TestAddProductToCartSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCart := mockC.NewMockCartRepositoryI(ctrl)
	mockProd := mockP.NewMockProductRepositoryI(ctrl)
	mockSes := mockS.NewMockSessionRepositoryI(ctrl)
	usecase := NewCartUsecase(mockCart, mockProd, mockSes)

	cookie := &entity.Cookie{
		UserID:       1,
		SessionToken: "HJJvgsvd",
	}

	cart := &entity.Cart{
		ID:     1,
		UserID: 1,
	}

	prod := &entity.Product{
		ID:          1,
		Name:        "Burger",
		Price:       120.0,
		CookingTime: 23,
		Portion:     "160 г",
		Description: "Nice burger",
		Icon:        "deficon",
	}

	var UserID uint
	UserID = 1

	mockSes.EXPECT().Check(cookie.SessionToken).Return(cookie, nil)
	mockCart.EXPECT().GetCartByUserID(UserID).Return(cart, nil)
	mockCart.EXPECT().AddProductToCart(cart.ID, prod.ID).Return(nil)
	err := usecase.AddProductToCart(cookie.SessionToken, prod.ID)
	assert.Nil(t, err)
}

func TestDeleteProductFromCartSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCart := mockC.NewMockCartRepositoryI(ctrl)
	mockProd := mockP.NewMockProductRepositoryI(ctrl)
	mockSes := mockS.NewMockSessionRepositoryI(ctrl)
	usecase := NewCartUsecase(mockCart, mockProd, mockSes)

	cookie := &entity.Cookie{
		UserID:       1,
		SessionToken: "HJJvgsvd",
	}

	cart := &entity.Cart{
		ID:     1,
		UserID: 1,
	}

	prod := &entity.Product{
		ID:          1,
		Name:        "Burger",
		Price:       120.0,
		CookingTime: 23,
		Portion:     "160 г",
		Description: "Nice burger",
		Icon:        "deficon",
	}

	var UserID uint
	UserID = 1

	mockSes.EXPECT().Check(cookie.SessionToken).Return(cookie, nil)
	mockCart.EXPECT().GetCartByUserID(UserID).Return(cart, nil)
	mockCart.EXPECT().DeleteProductFromCart(cart.ID, prod.ID).Return(nil)
	err := usecase.DeleteProductFromCart(cookie.SessionToken, prod.ID)
	assert.Nil(t, err)
}

func TestUpdateItemCountUpSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCart := mockC.NewMockCartRepositoryI(ctrl)
	mockProd := mockP.NewMockProductRepositoryI(ctrl)
	mockSes := mockS.NewMockSessionRepositoryI(ctrl)
	usecase := NewCartUsecase(mockCart, mockProd, mockSes)

	cookie := &entity.Cookie{
		UserID:       1,
		SessionToken: "HJJvgsvd",
	}

	cart := &entity.Cart{
		ID:     1,
		UserID: 1,
	}

	prod := &entity.Product{
		ID:          1,
		Name:        "Burger",
		Price:       120.0,
		CookingTime: 23,
		Portion:     "160 г",
		Description: "Nice burger",
		Icon:        "deficon",
	}

	var UserID uint
	UserID = 1

	mockSes.EXPECT().Check(cookie.SessionToken).Return(cookie, nil)
	mockCart.EXPECT().GetCartByUserID(UserID).Return(cart, nil)
	mockCart.EXPECT().UpdateItemCountUp(cart.ID, prod.ID).Return(nil)
	err := usecase.UpdateItemCountUp(cookie.SessionToken, prod.ID)
	assert.Nil(t, err)
}

func TestUpdateItemCountDownSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCart := mockC.NewMockCartRepositoryI(ctrl)
	mockProd := mockP.NewMockProductRepositoryI(ctrl)
	mockSes := mockS.NewMockSessionRepositoryI(ctrl)
	usecase := NewCartUsecase(mockCart, mockProd, mockSes)

	cookie := &entity.Cookie{
		UserID:       1,
		SessionToken: "HJJvgsvd",
	}

	cart := &entity.Cart{
		ID:     1,
		UserID: 1,
	}

	prod := &entity.Product{
		ID:          1,
		Name:        "Burger",
		Price:       120.0,
		CookingTime: 23,
		Portion:     "160 г",
		Description: "Nice burger",
		Icon:        "deficon",
	}

	var UserID uint
	UserID = 1

	mockSes.EXPECT().Check(cookie.SessionToken).Return(cookie, nil)
	mockCart.EXPECT().GetCartByUserID(UserID).Return(cart, nil)
	mockCart.EXPECT().UpdateItemCountDown(cart.ID, prod.ID).Return(nil)
	err := usecase.UpdateItemCountDown(cookie.SessionToken, prod.ID)
	assert.Nil(t, err)
}
