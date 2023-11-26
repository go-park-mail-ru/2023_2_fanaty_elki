package repository

import (
	"database/sql"
	"server/internal/domain/entity"
)

type productRepo struct {
	DB *sql.DB
}

func NewProductRepo(db *sql.DB) *productRepo {
	return &productRepo{
		DB: db,
	}
}

func (repo *productRepo) GetProductsByMenuTypeId(id uint) ([]*entity.Product, error) {
	rows, err := repo.DB.Query(`SELECT p.id, name, price, cooking_time, portion, description, icon  FROM product p 
	INNER JOIN product_menu_type pm ON pm.product_id = p.id AND menu_type_id = $1;`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var Products = []*entity.Product{}
	for rows.Next() {
		product := &entity.Product{}
		err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.CookingTime,
			&product.Portion,
			&product.Description,
			&product.Icon,
		)
		if err != nil {
			return nil, err
		}
		Products = append(Products, product)
	}
	err = rows.Err()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return Products, nil
}

func (repo *productRepo) GetProductByID(id uint) (*entity.Product, error) {
	product := &entity.Product{}
	row := repo.DB.QueryRow(`SELECT id, name, price, cooking_time, portion, description, icon FROM product WHERE id = $1`, id)
	err := row.Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.CookingTime,
		&product.Portion,
		&product.Description,
		&product.Icon,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, entity.ErrInternalServerError
	}
	return product, nil
}

func (repo *productRepo) SearchProducts(word string) ([]*entity.Product, error) {
	rows, err := repo.DB.Query(`SELECT id, name, price, cooking_time, portion, description, icon FROM product WHERE LOWER(name) LIKE LOWER('%' || $1 || '%')`, word)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var Products = []*entity.Product{}
	for rows.Next() {
		product := &entity.Product{}
		err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.CookingTime,
			&product.Portion,
			&product.Description,
			&product.Icon,
		)
		if err != nil {
			return nil, err
		}
		Products = append(Products, product)
	}
	err = rows.Err()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return Products, nil
}

func (repo *productRepo) GetRestaurantIdByProduct(id uint) (uint, error) {
	var restId uint
	row := repo.DB.QueryRow(`SELECT restaurant_id FROM product_menu_type JOIN menu_type on menu_type.id = menu_type_id where product_id = $1`, id)
	err := row.Scan(
		&restId,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, entity.ErrInternalServerError
		}
		return 0, entity.ErrInternalServerError
	}
	return restId, nil
}
