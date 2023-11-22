package repository

import (
	"database/sql"
	"fmt"
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
	insertOrder := `INSERT INTO orders (user_id, order_date, status, price, delivery_time)
				    VALUES ($1, $2, $3, $4, $5)
					RETURNING ID`
	var orderId uint
	err := repo.DB.QueryRow(insertOrder, order.UserId, order.Date, order.Status, order.Price, order.DeliveryTime).Scan(&orderId)
	if err != nil {
		return nil, entity.ErrInternalServerError
	}

	for _, product := range order.Products {
		insertProduct := `INSERT INTO orders_product (product_id, order_id, item_count) VALUES ($1, $2, $3)`
		_, err := repo.DB.Exec(insertProduct, product.ID, orderId, product.ItemCount)
		if err != nil {
			return nil, entity.ErrInternalServerError
		}
	}
	
	//Надо будет потом сделать так, чтобы не всгда инсертился, а селектил сначала

	insertAddress := `INSERT INTO address (city, street, house_number, flat_number)
				      VALUES ($1, $2, $3, $4)
					  RETURNING ID`
	var addressId uint
	err = repo.DB.QueryRow(insertAddress, order.Address.City, order.Address.Street, order.Address.House, order.Address.Flat).Scan(&addressId)
	fmt.Println(err)
	if err != nil {
		return nil, entity.ErrInternalServerError
	}

	insertOrderAddress := `INSERT INTO orders_address (orders_id, address_id)
				      	   VALUES ($1, $2)`
	_, err = repo.DB.Exec(insertOrderAddress, orderId, addressId)
	if err != nil {
		return nil, entity.ErrInternalServerError
	}

	return &dto.RespCreateOrder{
		Id:     orderId,
		Status: order.Status,
		Date:   order.Date,
		Price: order.Price,
		DeliveryTime: order.DeliveryTime,
	}, nil
}

func (repo *orderRepo) UpdateOrder(order *dto.ReqUpdateOrder) error {
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
	getOrders := `SELECT o.id, o.status, o.created_at, o.price, o.delivery_time, a.city, a.street, a.house_number, a.flat_number
			 	 FROM orders o
				 JOIN orders_address oa on o.id = oa.orders_id
				 JOIN address a on a.id = oa.address_id
				 WHERE o.user_id = $1
				 ORDER BY o.created_at DESC`

	rows, err := repo.DB.Query(getOrders, userId)
	if err != nil {
		return nil, entity.ErrInternalServerError
	}
	defer rows.Close()
	var orders = []*dto.RespGetOrder{}
	for rows.Next() {
		order := &dto.RespGetOrder{}
		address := &dto.RespOrderAddress{}
		err = rows.Scan(
			&order.Id,
			&order.Status,
			&order.Date,
			&order.Price,
			&order.DeliveryTime,
			&address.City,
			&address.Street,
			&address.House,
			&address.Flat,
		)
		order.Address = address
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

func (repo *orderRepo) GetOrder(reqOrder *dto.ReqGetOneOrder) (*dto.RespGetOneOrder, error) {
	getOrder :=`SELECT o.id, o.status, o.order_date, o.price, o.delivery_time, a.city, a.street, a.house_number, a.flat_number
				FROM orders o
   				JOIN orders_address oa on o.id = oa.orders_id
   				JOIN address a on a.id = oa.address_id
   				WHERE o.user_id = $1 AND o.id = $2` 
	order := dto.RespGetOneOrder{}
	address := dto.RespOrderAddress{}
	err := repo.DB.QueryRow(getOrder, reqOrder.OrderId, reqOrder.UserId).Scan(&order.Id, &order.Status, &order.Date,
							&order.Price, &order.DeliveryTime, &address.City, &address.Street, &address.House, &address.Flat)
	
	if err != nil {
		if err == sql.ErrNoRows {

			return nil, entity.ErrNotFound
		}
		return nil, entity.ErrInternalServerError
	}
	order.Address = &address
	getProducts := `SELECT p.id, p.name, p.price, p.icon, op.item_count
					FROM product p
					JOIN orders_product op ON p.id = op.product_id
					JOIN orders o ON o.id = op.order_id
					WHERE o.id = $1`
	rows, err := repo.DB.Query(getProducts, reqOrder.OrderId)
	if err != nil {
		return nil, entity.ErrInternalServerError
	}
	defer rows.Close()
	var products = []*dto.RespGetOrderProduct{}
	for rows.Next() {		
		product := &dto.RespGetOrderProduct{}
		err = rows.Scan(
			&product.Id,
			&product.Name,
			&product.Price,
			&product.Icon,
			&product.Count,
		)
		if err != nil {
			return nil, entity.ErrInternalServerError
		}
		products = append(products, product)
	}
	err = rows.Err()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, entity.ErrNotFound
		}
		return nil, entity.ErrInternalServerError
	}

	getRestaurant := `SELECT r.name
					  FROM restaurant r
					  JOIN menu_type mt ON mt.restaurant_id = r.id
					  JOIN product_menu_type pmt ON pmt.menu_type_id = mt.id 
					  JOIN product p ON p.id = pmt.id
					  WHERE p.id = $1`
	orderItems := &dto.OrderItems{}

	err = repo.DB.QueryRow(getRestaurant, products[0].Id).Scan(&orderItems.RestaurantName)
	if err != nil {
		return nil, entity.ErrInternalServerError
	}
	orderItems.Products = products
	order.OrderItems = append(order.OrderItems, orderItems)
	return &order, nil
}
