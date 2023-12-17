package usecase

import (
	cartRep "server/internal/Cart/repository"
	productRep "server/internal/Product/repository"
	promoRep "server/internal/Promo/repository"
	restaurantRep "server/internal/Restaurant/repository"
	sessionRep "server/internal/Session/repository"
	dto "server/internal/domain/dto"
	"server/internal/domain/entity"
)

type UsecaseI interface {
	GetUserCart(SessionToken string) (*dto.CartWithRestaurant, error)
	AddProductToCart(SessionToken string, productID uint) error
	DeleteProductFromCart(SessionToken string, productID uint) error
	CleanCart(SessionToken string) error
	GetCartTips(SessionToken string) ([]*entity.Product, error)
}

type cartUsecase struct {
	cartRepo       cartRep.CartRepositoryI
	productRepo    productRep.ProductRepositoryI
	sessionRepo    sessionRep.SessionRepositoryI
	restaurantRepo restaurantRep.RestaurantRepositoryI
	promoRepo      promoRep.PromoRepositoryI
}

func NewCartUsecase(cartRep cartRep.CartRepositoryI, productRep productRep.ProductRepositoryI, sessionRep sessionRep.SessionRepositoryI, restaurantRep restaurantRep.RestaurantRepositoryI, promoRep promoRep.PromoRepositoryI) *cartUsecase {
	return &cartUsecase{
		cartRepo:       cartRep,
		productRepo:    productRep,
		sessionRepo:    sessionRep,
		restaurantRepo: restaurantRep,
		promoRepo:      promoRep,
	}
}

func (cu cartUsecase) GetUserCart(SessionToken string) (*dto.CartWithRestaurant, error) {
	cookie, err := cu.sessionRepo.Check(SessionToken)
	if err != nil {
		return nil, err
	}

	userID := cookie.UserID
	cart, err := cu.cartRepo.GetCartByUserID(userID)
	if err != nil {
		return nil, err
	}

	cartWithRestaurant, err := cu.cartRepo.GetCartProductsByCartID(cart.ID)
	if err != nil {
		return nil, err
	}

	CartWithRestaurant := &dto.CartWithRestaurant{}

	for _, cartProduct := range cartWithRestaurant.Products {
		product, err := cu.productRepo.GetProductByID(cartProduct.ProductID)
		if err != nil {
			return nil, entity.ErrInternalServerError
		}
		CartProduct := dto.CartProduct{
			Product:   product,
			ItemCount: cartProduct.ItemCount,
		}
		CartWithRestaurant.Products = append(CartWithRestaurant.Products, &CartProduct)
	}
	if cartWithRestaurant.RestaurantId == 0 {
		return CartWithRestaurant, nil
	}

	restaurant, err := cu.restaurantRepo.GetRestaurantById(cartWithRestaurant.RestaurantId)
	if err != nil {
		return nil, err
	}

	CartWithRestaurant.Restaurant = restaurant

	if cartWithRestaurant.PromoId != 0 {

		promo, err := cu.promoRepo.GetPromoById(cartWithRestaurant.PromoId)
		if err != nil {
			return nil, err
		}

		resppromo := dto.ToRespPromo(promo)
		CartWithRestaurant.Promo = resppromo

	}

	return CartWithRestaurant, nil
}

func (cu cartUsecase) AddProductToCart(SessionToken string, productID uint) error {
	cookie, err := cu.sessionRepo.Check(SessionToken)
	if err != nil {
		return err
	}

	product, err := cu.productRepo.GetProductByID(productID)

	if product == nil {
		return entity.ErrNotFound
	}

	userID := cookie.UserID
	cart, err := cu.cartRepo.GetCartByUserID(userID)
	if err != nil {
		return err
	}

	hasproduct, err := cu.cartRepo.CheckProductInCart(cart.ID, productID)

	if err != nil {
		return err
	}

	if hasproduct {
		err = cu.cartRepo.UpdateItemCountUp(cart.ID, productID)
		if err != nil {
			return err
		}
	} else {
		err = cu.cartRepo.AddProductToCart(cart.ID, productID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (cu cartUsecase) DeleteProductFromCart(SessionToken string, productID uint) error {
	cookie, err := cu.sessionRepo.Check(SessionToken)
	if err != nil {
		return err
	}

	product, err := cu.productRepo.GetProductByID(productID)

	if product == nil {
		return entity.ErrNotFound
	}

	userID := cookie.UserID
	cart, err := cu.cartRepo.GetCartByUserID(userID)
	if err != nil {
		return err
	}

	hasproduct, err := cu.cartRepo.CheckProductInCart(cart.ID, productID)

	if err != nil {
		return err
	}

	if hasproduct {
		itemcount, err := cu.cartRepo.CheckProductCount(cart.ID, productID)
		if err != nil {
			return err
		}

		if itemcount == 1 {
			err = cu.cartRepo.DeleteProductFromCart(cart.ID, productID)
			if err != nil {
				return err
			}
		}
		err = cu.cartRepo.UpdateItemCountDown(cart.ID, productID)
		if err != nil {
			return err
		}
	} else {
		return entity.ErrNotFound
	}

	return nil
}

func (cu cartUsecase) CleanCart(SessionToken string) error {
	cookie, err := cu.sessionRepo.Check(SessionToken)
	if err != nil {
		return err
	}

	userID := cookie.UserID
	cart, err := cu.cartRepo.GetCartByUserID(userID)
	if err != nil {
		return err
	}

	cartWithRestaurant, err := cu.cartRepo.GetCartProductsByCartID(cart.ID)
	if err != nil {
		return err
	}

	err = cu.promoRepo.DeletePromoFromCart(cart.ID, cartWithRestaurant.PromoId)
	if err != nil {
		return err
	}

	err = cu.cartRepo.CleanCart(cart.ID)
	if err != nil {
		return err
	}

	return nil
}

func (cu cartUsecase) GetCartTips(SessionToken string) ([]*entity.Product, error) {
	cookie, err := cu.sessionRepo.Check(SessionToken)
	if err != nil {
		return nil, err
	}

	userID := cookie.UserID
	cart, err := cu.cartRepo.GetCartByUserID(userID)
	if err != nil {
		return nil, err
	}

	cartWithRestaurant, err := cu.cartRepo.GetCartProductsByCartID(cart.ID)
	if err != nil {
		return nil, err
	}

	prods := make(map[uint]bool)

	for _, product := range cartWithRestaurant.Products {
		prods[product.ProductID] = true
	}

	if cartWithRestaurant.RestaurantId == 0 {
		return nil, entity.ErrNotFound
	}

	menuTypes, err := cu.restaurantRepo.GetMenuTypesByRestaurantId(cartWithRestaurant.RestaurantId)
	if err != nil {
		return nil, err
	}
	var restProducts []*entity.Product
	for _, menu := range menuTypes {
		products, err := cu.productRepo.GetProductsByMenuTypeId(menu.ID)
		if err != nil {
			return nil, entity.ErrInternalServerError
		}
		restProducts = append(restProducts, products...)
	}

	var tipProducts []*entity.Product
	for _, product := range restProducts {
		if !prods[product.ID] {
			tipProducts = append(tipProducts, product)
		}
	}

	return tipProducts, nil
}
