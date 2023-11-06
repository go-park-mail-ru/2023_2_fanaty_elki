package usecase

import (
	cartRep "server/internal/Cart/repository"
	productRep "server/internal/Product/repository"
	sessionRep "server/internal/Session/repository"
	dto "server/internal/domain/dto"
	"server/internal/domain/entity"
)

type UsecaseI interface {
	GetUserCart(SessionToken string) ([]*dto.CartProduct, error)
	AddProductToCart(SessionToken string, productID uint) error
	DeleteProductFromCart(SessionToken string, productID uint) error
	UpdateItemCountUp(SessionToken string, productID uint) error
	UpdateItemCountDown(SessionToken string, productID uint) error
}

type cartUsecase struct {
	CartRepo    cartRep.CartRepositoryI
	ProductRepo productRep.ProductRepositoryI
	SessionRepo sessionRep.SessionRepositoryI
}

func NewCartUsecase(cartRep cartRep.CartRepositoryI, productRep productRep.ProductRepositoryI, sessionRep sessionRep.SessionRepositoryI) *cartUsecase {
	return &cartUsecase{
		CartRepo:    cartRep,
		ProductRepo: productRep,
		SessionRepo: sessionRep,
	}
}

func (cu cartUsecase) GetUserCart(SessionToken string) ([]*dto.CartProduct, error) {
	cookie, err := cu.SessionRepo.Check(SessionToken)
	if err != nil {
		return nil, err
	}

	userID := cookie.UserID
	cart, err := cu.CartRepo.GetCartByUserID(userID)
	if err != nil {
		return nil, err
	}

	cartProducts, err := cu.CartRepo.GetCartProductsByCartID(cart.ID)
	if err != nil {
		return nil, err
	}
	var CartProducts []*dto.CartProduct
	for _, cartProduct := range cartProducts {
		product, err := cu.ProductRepo.GetProductByID(cartProduct.ProductID)
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

func (cu cartUsecase) AddProductToCart(SessionToken string, productID uint) error {
	cookie, err := cu.SessionRepo.Check(SessionToken)
	if err != nil {
		return err
	}

	userID := cookie.UserID
	cart, err := cu.CartRepo.GetCartByUserID(userID)
	if err != nil {
		return err
	}

	err = cu.CartRepo.AddProductToCart(cart.ID, productID)
	if err != nil {
		return err
	}

	return nil
}

func (cu cartUsecase) DeleteProductFromCart(SessionToken string, productID uint) error {
	cookie, err := cu.SessionRepo.Check(SessionToken)
	if err != nil {
		return err
	}

	userID := cookie.UserID
	cart, err := cu.CartRepo.GetCartByUserID(userID)
	if err != nil {
		return err
	}

	err = cu.CartRepo.DeleteProductFromCart(cart.ID, productID)
	if err != nil {
		return err
	}

	return nil
}

func (cu cartUsecase) UpdateItemCountUp(SessionToken string, productID uint) error {
	cookie, err := cu.SessionRepo.Check(SessionToken)
	if err != nil {
		return err
	}

	userID := cookie.UserID
	cart, err := cu.CartRepo.GetCartByUserID(userID)
	if err != nil {
		return err
	}
	err = cu.CartRepo.UpdateItemCountUp(cart.ID, productID)
	if err != nil {
		return err
	}

	return nil
}

func (cu cartUsecase) UpdateItemCountDown(SessionToken string, productID uint) error {
	cookie, err := cu.SessionRepo.Check(SessionToken)
	if err != nil {
		return err
	}

	userID := cookie.UserID
	cart, err := cu.CartRepo.GetCartByUserID(userID)
	if err != nil {
		return err
	}
	err = cu.CartRepo.UpdateItemCountDown(cart.ID, productID)
	if err != nil {
		return err
	}

	return nil
}
