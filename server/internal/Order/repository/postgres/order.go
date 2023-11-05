package order

import (
	"database/sql"
	"server/internal/domain/dto"
	"server/internal/domain/entity"
)

type orderRepo struct {
	DB *sql.DB
}

func NewOrderRepo(db *sql.DB) *orderRepo {
	return &orderRepo{
		DB: db,
	}
}

func (repo *orderRepo) CreateOrder(order *dto.DBReqCreateOrder) (*dto.RespCreateOrder, error) {
	insertOrder := `INSERT INTO orders (user_id, order_date, status) VALUES ($1, $2, $3) RETURNING ID`
	var orderId uint
	err := repo.DB.QueryRow(insertOrder, order.UserId, order.Date, order.Status).Scan(&orderId)
	if err != nil {
		return nil, entity.ErrInternalServerError
	}
	
	for product, count := range *order.Products {
		insertProduct := `INSERT INTO orders_product (product_id, order_id, item_count) VALUES ($1, $2, $3)`
		_, err := repo.DB.Exec(insertProduct, product, orderId, count)
		if err != nil {
			return nil, entity.ErrInternalServerError
		}
	}

	return &dto.RespCreateOrder{
		Id:orderId,
		Status: order.Status,
		Date: order.Date,
	}, nil
}