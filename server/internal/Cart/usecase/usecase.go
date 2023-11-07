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
	cartRepo    cartRep.CartRepositoryI
	productRepo productRep.ProductRepositoryI
	sessionRepo sessionRep.SessionRepositoryI
}

func NewCartUsecase(cartRep cartRep.CartRepositoryI, productRep productRep.ProductRepositoryI, sessionRep sessionRep.SessionRepositoryI) *cartUsecase {
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

func (cu cartUsecase) AddProductToCart(SessionToken string, productID uint) error {
	cookie, err := cu.sessionRepo.Check(SessionToken)
	if err != nil {
		return err
	}

	userID := cookie.UserID
	cart, err := cu.cartRepo.GetCartByUserID(userID)
	if err != nil {
		return err
	}

	err = cu.cartRepo.AddProductToCart(cart.ID, productID)
	if err != nil {
		return err
	}

	return nil
}

func (cu cartUsecase) DeleteProductFromCart(SessionToken string, productID uint) error {
	cookie, err := cu.sessionRepo.Check(SessionToken)
	if err != nil {
		return err
	}

	userID := cookie.UserID
	cart, err := cu.cartRepo.GetCartByUserID(userID)
	if err != nil {
		return err
	}

	err = cu.cartRepo.DeleteProductFromCart(cart.ID, productID)
	if err != nil {
		return err
	}

	return nil
}

func (cu cartUsecase) UpdateItemCountUp(SessionToken string, productID uint) error {
	cookie, err := cu.sessionRepo.Check(SessionToken)
	if err != nil {
		return err
	}

	userID := cookie.UserID
	cart, err := cu.cartRepo.GetCartByUserID(userID)
	if err != nil {
		return err
	}
	err = cu.cartRepo.UpdateItemCountUp(cart.ID, productID)
	if err != nil {
		return err
	}

	return nil
}

func (cu cartUsecase) UpdateItemCountDown(SessionToken string, productID uint) error {
	cookie, err := cu.sessionRepo.Check(SessionToken)
	if err != nil {
		return err
	}

	userID := cookie.UserID
	cart, err := cu.cartRepo.GetCartByUserID(userID)
	if err != nil {
		return err
	}
	err = cu.cartRepo.UpdateItemCountDown(cart.ID, productID)
	if err != nil {
		return err
	}

	return nil
}
