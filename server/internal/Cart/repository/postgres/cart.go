package postgres

import (
	"database/sql"
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

func (repo *CartRepo) CreateCart(userID uint) (uint, error) {
	insertCart := `INSERT INTO cart (user_id) VALUES ($1)`
	_, err := repo.DB.Exec(insertCart, userID)
	if err != nil {
		return 0, entity.ErrInternalServerError
	}
	var ID uint
	row := repo.DB.QueryRow("SELECT ID FROM cart WHERE user_id = $1", userID)
	err = row.Scan(&ID)

	if err != nil {
		return 0, entity.ErrInternalServerError
	}

	return ID, nil
}

func (repo *CartRepo) GetCartByUserID(userID uint) (*entity.Cart, error) {
	cart := &entity.Cart{}
	row := repo.DB.QueryRow("SELECT id, user_id FROM cart WHERE user_id = $1", userID)
	err := row.Scan(
		&cart.ID,
		&cart.UserID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, entity.ErrInternalServerError
	}
	return cart, nil
}

func (repo *CartRepo) GetCartProductsByCartID(cartID uint) ([]*entity.CartProduct, error) {
	rows, err := repo.DB.Query("SELECT id, product_id, cart_id, item_count FROM cart_product")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var CartProducts = []*entity.CartProduct{}
	for rows.Next() {
		cartProduct := &entity.CartProduct{}
		err = rows.Scan(
			&cartProduct.ID,
			&cartProduct.ProductID,
			&cartProduct.CartID,
			&cartProduct.ItemCount,
		)
		if err != nil {
			return nil, err
		}
		CartProducts = append(CartProducts, cartProduct)
	}
	err = rows.Err()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return CartProducts, nil
}
