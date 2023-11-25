package usecase

import (
	csatRep "server/internal/Csat/repository"
	dto "server/internal/domain/dto"
	"server/internal/domain/entity"
)

type UsecaseI interface {
	GetQuestionnaireByID(id uint) ([]*entity.Question, error)
}

type csatUsecase struct {
	csatRepo csatRep.CsatRepositoryI
}

func NewCsatUsecase(csatRep csatRep.CsatRepositoryI) *cartUsecase {
	return &cartUsecase{
		cartRepo:    cartRep,
		productRepo: productRep,
		sessionRepo: sessionRep,
	}
}

func (cu cartUsecase) GetUserCart(SessionToken string) ([]*dto.CartProduct, error) {
	cookie, err := cu.sessionRepo.Check(SessionToken)
	if err != nil {
		return nil, err
	}

	userID := cookie.UserID
	cart, err := cu.cartRepo.GetCartByUserID(userID)
	if err != nil {
		return nil, err
	}

	cartProducts, err := cu.cartRepo.GetCartProductsByCartID(cart.ID)
	if err != nil {
		return nil, err
	}
	var CartProducts []*dto.CartProduct
	for _, cartProduct := range cartProducts {
		product, err := cu.productRepo.GetProductByID(cartProduct.ProductID)
		if err != nil {
			return nil, entity.ErrInternalServerError
		}
		CartProduct := dto.CartProduct{
			Product:   product,
			ItemCount: cartProduct.ItemCount,
		}
		CartProducts = append(CartProducts, &CartProduct)
	}
	return CartProducts, nil
}
