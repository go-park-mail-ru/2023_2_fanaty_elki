package repository

import (
	"database/sql"
	"server/internal/domain/entity"
)

//ProductRepo struct
type ProductRepo struct {
	DB *sql.DB
}

//NewProductRepo craetes new product repo 
func NewProductRepo(db *sql.DB) *ProductRepo {
	return &ProductRepo{
		DB: db,
	}
}

//GetProductsByMenuTypeID gets products by menu type from db
func (repo *ProductRepo) GetProductsByMenuTypeID(id uint) ([]*entity.Product, error) {
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
		} 
		return nil, err
	}
	return Products, nil
}

//GetProductByID gets products by id from db
func (repo *ProductRepo) GetProductByID(id uint) (*entity.Product, error) {
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

//SearchProducts searches products in db
func (repo *ProductRepo) SearchProducts(word string) ([]*entity.Product, error) {
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
		}
		return nil, err
	}
	return Products, nil
}

//GetRestaurantIDByProduct get restaurant id by product from db
func (repo *ProductRepo) GetRestaurantIDByProduct(id uint) (uint, error) {
	var restID uint
	row := repo.DB.QueryRow(`SELECT restaurant_id FROM product_menu_type JOIN menu_type on menu_type.id = menu_type_id where product_id = $1`, id)
	err := row.Scan(
		&restID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, entity.ErrInternalServerError
		}
		return 0, entity.ErrInternalServerError
	}
	return restID, nil
}
