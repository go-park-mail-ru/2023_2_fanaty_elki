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
	insertOrder := `INSERT INTO orders (user_id, order_date, status)
				    VALUES ($1, $2, $3)
					RETURNING ID`
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

func (repo *orderRepo) UpdateOrder(order *dto.ReqUpdateOrder) (error) {
	updateOrder := `UPDATE orders
					SET status = $1
					WHERE id = $2`
	_, err := repo.DB.Exec(updateOrder, order.Status, order.Id)
	if err != nil {
		return entity.ErrInternalServerError
	}
	return nil
}

func (repo *orderRepo) GetOrders(userId uint) ([]*dto.RespGetOrder, error) {
	getOrder := `SELECT id, status, created_at, updated_at
			 	 FROM orders
				 WHERE user_id = $1
				 ORDER BY created_at DESC`

	rows, err := repo.DB.Query(getOrder, userId)
	if err != nil {
		return nil, entity.ErrInternalServerError
	}
	defer rows.Close()
	var orders = []*dto.RespGetOrder{}
	for rows.Next() {
		order := &dto.RespGetOrder{}
		err = rows.Scan(
			&order.Id,
			&order.Status,
			&order.Date,
			&order.UpdatedDate,
		)
		if err != nil {
			return nil, entity.ErrInternalServerError
		}
		orders = append(orders, order)
	}	
	err = rows.Err()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} 
		return nil, entity.ErrInternalServerError
	}
	return orders, nil
}