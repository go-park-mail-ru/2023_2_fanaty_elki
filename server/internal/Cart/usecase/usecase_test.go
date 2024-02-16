package usecase

import (
	"testing"

	mockC "server/internal/Cart/repository/mock_repository"
	mockP "server/internal/Product/repository/mock_repository"
	mockPr "server/internal/Promo/repository/mock_repository"
	mockR "server/internal/Restaurant/repository/mock_repository"
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
	mockRest := mockR.NewMockRestaurantRepositoryI(ctrl)
	mockPromo := mockPr.NewMockPromoRepositoryI(ctrl)
	usecase := NewCartUsecase(mockCart, mockProd, mockSes, mockRest, mockPromo)

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

	cartwithrest := &entity.CartWithRestaurant{
		RestaurantID: 1,
		Products: []*entity.CartProduct{
			{
				ID:        1,
				ProductID: 1,
				CartID:    1,
				ItemCount: 6,
			},
			{
				ID:        2,
				ProductID: 3,
				CartID:    1,
				ItemCount: 6,
			},
		},
	}

	rest := &entity.Restaurant{
		ID:            1,
		Name:          "Burger King",
		Rating:        3.7,
		CommentsCount: 60,
		Icon:          "img/burger_king.webp",
	}

	res := &dto.CartWithRestaurant{
		Restaurant: rest,
		Products: []*dto.CartProduct{
			{
				Product: &entity.Product{
					ID:          1,
					Name:        "Burger",
					Price:       120.0,
					CookingTime: 23,
					Portion:     "160 г",
					Description: "Nice burger",
					Icon:        "deficon",
				},
				ItemCount: 6,
			},
			{
				Product: &entity.Product{
					ID:          2,
					Name:        "Burger",
					Price:       120.0,
					CookingTime: 23,
					Portion:     "160 г",
					Description: "Nice burger",
					Icon:        "deficon",
				},
				ItemCount: 6,
			},
		},
	}

	var UserID uint
	UserID = 1

	mockSes.EXPECT().Check(cookie.SessionToken).Return(cookie, nil)
	mockCart.EXPECT().GetCartByUserID(UserID).Return(cart, nil)
	mockCart.EXPECT().GetCartProductsByCartID(cart.ID).Return(cartwithrest, nil)
	mockProd.EXPECT().GetProductByID(uint(1)).Return(prod, nil)
	mockProd.EXPECT().GetProductByID(uint(3)).Return(prod, nil)
	mockRest.EXPECT().GetRestaurantByID(uint(1)).Return(rest, nil)
	actual, err := usecase.GetUserCart(cookie.SessionToken)
	assert.Equal(t, res.Restaurant, actual.Restaurant)
	assert.Nil(t, err)

}

func TestAddProductToCartSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCart := mockC.NewMockCartRepositoryI(ctrl)
	mockProd := mockP.NewMockProductRepositoryI(ctrl)
	mockSes := mockS.NewMockSessionRepositoryI(ctrl)
	mockRest := mockR.NewMockRestaurantRepositoryI(ctrl)
	mockPromo := mockPr.NewMockPromoRepositoryI(ctrl)
	usecase := NewCartUsecase(mockCart, mockProd, mockSes, mockRest, mockPromo)

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
	mockProd.EXPECT().GetProductByID(prod.ID).Return(prod, nil)
	mockCart.EXPECT().GetCartByUserID(UserID).Return(cart, nil)
	mockCart.EXPECT().CheckProductInCart(cart.ID, prod.ID).Return(false, nil)
	mockCart.EXPECT().AddProductToCart(cart.ID, prod.ID).Return(nil)
	err := usecase.AddProductToCart(cookie.SessionToken, prod.ID)
	assert.Nil(t, err)

	mockSes.EXPECT().Check(cookie.SessionToken).Return(cookie, nil)
	mockProd.EXPECT().GetProductByID(prod.ID).Return(prod, nil)
	mockCart.EXPECT().GetCartByUserID(UserID).Return(cart, nil)
	mockCart.EXPECT().CheckProductInCart(cart.ID, prod.ID).Return(true, nil)
	mockCart.EXPECT().UpdateItemCountUp(cart.ID, prod.ID).Return(nil)
	err = usecase.AddProductToCart(cookie.SessionToken, prod.ID)
	assert.Nil(t, err)
}

func TestDeleteProductFromCartSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCart := mockC.NewMockCartRepositoryI(ctrl)
	mockProd := mockP.NewMockProductRepositoryI(ctrl)
	mockSes := mockS.NewMockSessionRepositoryI(ctrl)
	mockRest := mockR.NewMockRestaurantRepositoryI(ctrl)
	mockPromo := mockPr.NewMockPromoRepositoryI(ctrl)
	usecase := NewCartUsecase(mockCart, mockProd, mockSes, mockRest, mockPromo)

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
	mockProd.EXPECT().GetProductByID(prod.ID).Return(prod, nil)
	mockCart.EXPECT().GetCartByUserID(UserID).Return(cart, nil)
	mockCart.EXPECT().CheckProductInCart(cart.ID, prod.ID).Return(true, nil)
	mockCart.EXPECT().CheckProductCount(cart.ID, prod.ID).Return(uint(2), nil)
	mockCart.EXPECT().UpdateItemCountDown(cart.ID, prod.ID).Return(nil)
	err := usecase.DeleteProductFromCart(cookie.SessionToken, prod.ID)
	assert.Nil(t, err)
}

func TestCleanCartSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCart := mockC.NewMockCartRepositoryI(ctrl)
	mockProd := mockP.NewMockProductRepositoryI(ctrl)
	mockSes := mockS.NewMockSessionRepositoryI(ctrl)
	mockRest := mockR.NewMockRestaurantRepositoryI(ctrl)
	mockPromo := mockPr.NewMockPromoRepositoryI(ctrl)
	usecase := NewCartUsecase(mockCart, mockProd, mockSes, mockRest, mockPromo)

	cookie := &entity.Cookie{
		UserID:       1,
		SessionToken: "HJJvgsvd",
	}

	cart := &entity.Cart{
		ID:     1,
		UserID: 1,
	}

	cartwithrest := &entity.CartWithRestaurant{
		RestaurantID: 1,
		Products: []*entity.CartProduct{
			{
				ID:        1,
				ProductID: 1,
				CartID:    1,
				ItemCount: 6,
			},
			{
				ID:        2,
				ProductID: 3,
				CartID:    1,
				ItemCount: 6,
			},
		},
	}

	var UserID uint
	UserID = 1

	var CartID uint
	CartID = 1

	var PromoID uint
	PromoID = 0

	mockSes.EXPECT().Check(cookie.SessionToken).Return(cookie, nil)
	mockCart.EXPECT().GetCartByUserID(UserID).Return(cart, nil)
	mockCart.EXPECT().GetCartProductsByCartID(CartID).Return(cartwithrest, nil)
	mockPromo.EXPECT().DeletePromoFromCart(CartID, PromoID).Return(nil)
	mockCart.EXPECT().CleanCart(UserID).Return(nil)
	err := usecase.CleanCart(cookie.SessionToken)
	assert.Nil(t, err)

}
