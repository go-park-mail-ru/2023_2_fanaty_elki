package repository

import (
	"database/sql"
	"fmt"
	"server/internal/Cart/repository"
	"server/internal/domain/entity"
)

type CartRepo struct {
	DB *sql.DB
}

func NewCartRepo(db *sql.DB) repository.CartRepositoryI {
	return &CartRepo{
		DB: db,
	}
}

func (repo *CartRepo) CreateCart(UserID uint) (uint, error) {
	insertCart := `INSERT INTO cart (user_id) VALUES ($1)`
	_, err := repo.DB.Exec(insertCart, UserID)
	if err != nil {
		return 0, entity.ErrInternalServerError
	}
	var ID uint
	row := repo.DB.QueryRow("SELECT ID FROM cart WHERE user_id = $1", UserID)
	err = row.Scan(&ID)

	if err != nil {
		return 0, entity.ErrInternalServerError
	}

	return ID, nil
}

func (repo *CartRepo) GetCartByUserID(UserID uint) (*entity.Cart, error) {
	cart := &entity.Cart{}
	row := repo.DB.QueryRow("SELECT id, user_id FROM cart WHERE user_id = $1", UserID)
	err := row.Scan(
		&cart.ID,
		&cart.UserID,
	)
	if err != nil {
		return nil, entity.ErrInternalServerError
	}
	return cart, nil
}

func (repo *CartRepo) GetCartProductsByCartID(cartID uint) (*entity.CartWithRestaurant, error) {
	CartWithRestaurant := &entity.CartWithRestaurant{}

	cartRows, err := repo.DB.Query("SELECT id, product_id, cart_id, item_count FROM cart_product WHERE cart_id = $1", cartID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}
	defer cartRows.Close()

	var cartProducts = []*entity.CartProduct{}
	for cartRows.Next() {
		cartProduct := &entity.CartProduct{}
		err = cartRows.Scan(
			&cartProduct.ID,
			&cartProduct.ProductID,
			&cartProduct.CartID,
			&cartProduct.ItemCount,
		)
		if err != nil {
			return nil, err
		}
		cartProducts = append(cartProducts, cartProduct)
	}

	CartWithRestaurant.Products = cartProducts

	if len(cartProducts) == 0 {
		return CartWithRestaurant, nil
	}

	restaurantRows, err := repo.DB.Query(
		`SELECT mt.Restaurant_id
		FROM Product p
		JOIN Product_Menu_Type pmt ON p.id = pmt.Product_Id
		JOIN Menu_Type mt ON pmt.Menu_Type_id = mt.Id
		WHERE pmt.Product_Id = $1`,
		cartProducts[0].ProductID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}
	defer restaurantRows.Close()

	restaurantRows.Next()
	var restaurantId uint
	err = restaurantRows.Scan(&restaurantId)
	if err != nil {
		return nil, err
	}

	promoRow := repo.DB.QueryRow(`SELECT promo_id FROM cart_promo WHERE cart_id = $1`, cartID)
	// if err != nil {
	// 	fmt.Println("query", err)
	// 	if err == sql.ErrNoRows {
	// 		return nil, nil
	// 	} else {
	// 		return nil, err
	// 	}
	// }
	// defer promoRows.Close()

	// promoRows.Next()
	var PromoID uint
	err = promoRow.Scan(&PromoID)
	if err != nil {
		fmt.Println("sacn", err)
		if err == sql.ErrNoRows {
			PromoID = 0
		} else {
			return nil, err
		}
	}

	CartWithRestaurant.RestaurantID = restaurantId
	CartWithRestaurant.PromoID = PromoID

	return CartWithRestaurant, nil
}

func (repo *CartRepo) AddProductToCart(cartID uint, productID uint) error {
	insertProduct := `INSERT INTO cart_product (product_id, cart_id) VALUES ($1, $2)`
	_, err := repo.DB.Exec(insertProduct, productID, cartID)
	if err != nil {
		return entity.ErrInternalServerError
	}
	return nil
}

func (repo *CartRepo) DeleteProductFromCart(cartID uint, productID uint) error {
	deleteProduct := `DELETE FROM cart_product WHERE cart_id = $1 AND product_id = $2`
	_, err := repo.DB.Exec(deleteProduct, cartID, productID)
	if err != nil {
		return entity.ErrInternalServerError
	}
	return nil
}

func (repo *CartRepo) UpdateItemCountUp(cartID uint, productID uint) error {
	updateProduct := `UPDATE cart_product SET item_count = item_count + 1 WHERE cart_id = $1 AND product_id = $2`
	_, err := repo.DB.Exec(updateProduct, cartID, productID)
	if err != nil {
		return entity.ErrInternalServerError
	}
	return nil
}

func (repo *CartRepo) UpdateItemCountDown(cartID uint, productID uint) error {
	updateProduct := `UPDATE cart_product SET item_count = item_count - 1 WHERE cart_id = $1 AND product_id = $2`
	_, err := repo.DB.Exec(updateProduct, cartID, productID)
	if err != nil {
		return entity.ErrInternalServerError
	}
	return nil
}

func (repo *CartRepo) CheckProductInCart(cartID uint, productID uint) (bool, error) {
	productCart := &entity.CartProduct{}
	row := repo.DB.QueryRow("SELECT id, product_id, cart_id, item_count FROM cart_product WHERE cart_id = $1 and product_id = $2", cartID, productID)
	err := row.Scan(
		&productCart.ID,
		&productCart.ProductID,
		&productCart.CartID,
		&productCart.ItemCount,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func (repo *CartRepo) CheckProductCount(cartID uint, productID uint) (uint, error) {
	productCart := &entity.CartProduct{}
	row := repo.DB.QueryRow("SELECT id, product_id, cart_id, item_count FROM cart_product WHERE cart_id = $1 and product_id = $2", cartID, productID)
	err := row.Scan(
		&productCart.ID,
		&productCart.ProductID,
		&productCart.CartID,
		&productCart.ItemCount,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		} else {
			return 0, err
		}
	}
	return uint(productCart.ItemCount), nil
}

func (repo *CartRepo) CleanCart(cartID uint) error {
	deleteProducts := `DELETE FROM cart_product WHERE cart_id = $1`
	_, err := repo.DB.Exec(deleteProducts, cartID)
	if err != nil {
		return entity.ErrInternalServerError
	}
	return nil
}
