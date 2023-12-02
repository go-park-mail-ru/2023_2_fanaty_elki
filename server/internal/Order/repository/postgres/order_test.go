package repository

import (
	"reflect"
	"server/internal/domain/dto"
	"server/internal/domain/entity"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func TestCreateOrderSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &orderRepo{
		DB: db,
	}

	var flat uint
	flat = 1

	product := &entity.CartProduct{
		ID:        1,
		ProductID: 1,
		CartID:    1,
		ItemCount: 1,
	}
	products := []*entity.CartProduct{product}
	order := &dto.DBReqCreateOrder{
		Products: products,
		UserId:   1,
		Status:   0,
		Date:     time.Now(),
		Address: &dto.DBCreateOrderAddress{
			City:   "Moscow",
			Street: "Tverskaya",
			House:  "2",
			Flat:   flat,
		},
	}

	respOrder := &dto.RespCreateOrder{
		Id:           1,
		Status:       0,
		Date:         order.Date,
		Price:        109,
		DeliveryTime: 30,
		Address: &entity.Address{
			City:   "Moscow",
			Street: "Tverskaya",
			House:  "2",
			Flat:   flat,
		},
	}

	var userID uint
	userID = 1

	var orderID uint
	orderID = 1

	var addressID uint
	addressID = 1

	rows := sqlmock.
		NewRows([]string{"id"}).AddRow(1)

	mock.
		ExpectQuery("INSERT INTO orders").
		WithArgs(userID, order.Date, order.Status, order.Price, order.DeliveryTime).
		WillReturnRows(rows)

	mock.
		ExpectExec("INSERT INTO orders_product ").
		WithArgs(1, orderID, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	rows = sqlmock.
		NewRows([]string{"id"}).AddRow(1)

	mock.
		ExpectQuery("INSERT INTO address").
		WithArgs(order.Address.City, order.Address.Street, order.Address.House, order.Address.Flat).
		WillReturnRows(rows)

	mock.
		ExpectExec("INSERT INTO orders_address").
		WithArgs(orderID, addressID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	actual, err := repo.CreateOrder(order)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if respOrder.Id != actual.Id {
		t.Errorf("bad id: want %v, have %v", respOrder.Id, actual.Id)
		return
	}

}

func TestUpdateOrderSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &orderRepo{
		DB: db,
	}

	order := &dto.ReqUpdateOrder{
		Status: 1,
		Id:     1,
	}

	mock.
		ExpectExec("UPDATE orders SET").
		WithArgs(order.Status, order.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.UpdateOrder(order)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestGetOrdersSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &orderRepo{
		DB: db,
	}

	rows := sqlmock.
		NewRows([]string{"id", "status", "created_at", "price", "delivery_time", "city", "street", "house_numbe", "flat_number"})
	expect := []*dto.RespGetOrder{
		{
			Id:     1,
			Status: 0,
			Date:   time.Now(),
			Address: &dto.RespOrderAddress{
				City:   "Moscow",
				Street: "Tverskaya",
				House:  "2",
				Flat:   0,
			},
		},
		{
			Id:     2,
			Status: 0,
			Date:   time.Now(),
			Address: &dto.RespOrderAddress{
				City:   "Moscow",
				Street: "Tverskaya",
				House:  "3",
				Flat:   0,
			},
		},
	}

	for _, order := range expect {
		rows = rows.AddRow(order.Id, order.Status, order.Date, order.Price, order.DeliveryTime, order.Address.City, order.Address.Street, order.Address.House, order.Address.Flat)
	}
	var userID uint
	userID = 1
	mock.
		ExpectQuery(`SELECT o.id, o.status, o.created_at, o.price, o.delivery_time, a.city, a.street, a.house_number, a.flat_number
		FROM orders o
	   JOIN orders_address oa on o.id = oa.orders_id
	   JOIN address a on a.id = oa.address_id
	   WHERE`).
		WithArgs(userID).
		WillReturnRows(rows)

	orders, err := repo.GetOrders(userID)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if !reflect.DeepEqual(orders[0], expect[0]) {
		t.Errorf("results not match, want %v, have %v", expect[0], orders[0])
		return
	}
}

func TestGetOrderSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &orderRepo{
		DB: db,
	}

	reqorder := &dto.ReqGetOneOrder{
		OrderId: 1,
		UserId:  1,
	}

	products := []*dto.RespGetOrderProduct{
		{
			Id:    1,
			Name:  "Burger",
			Price: 134.0,
			Icon:  "deih",
			Count: 3,
		},
		{
			Id:    2,
			Name:  "Burger1",
			Price: 134.0,
			Icon:  "deih",
			Count: 3,
		},
	}

	address := &dto.RespOrderAddress{
		City:   "Moscow",
		Street: "Tverskaya",
		House:  "2",
		Flat:   0,
	}

	order := dto.RespGetOneOrder{
		Id:     1,
		Status: 0,
		Date:   time.Now(),
		OrderItems: []*dto.OrderItems{
			{
				RestaurantName: "Burger King",
				Products:       products,
			},
		},
		Address: address,
	}

	resporder := &dto.RespGetOneOrder{
		Id:     1,
		Status: 0,
		Date:   time.Now(),
		OrderItems: []*dto.OrderItems{
			{
				RestaurantName: "Burger King",
				Products:       products,
			},
		},
		Address: address,
	}

	row := sqlmock.
		NewRows([]string{"id", "status", "order_date", "price", "delivery_time", "city", "street", "house_number", "falt_number"})
	row = row.AddRow(resporder.Id, resporder.Status, resporder.Date, resporder.Price, resporder.DeliveryTime, resporder.Address.City, resporder.Address.Street, resporder.Address.House, resporder.Address.Flat)

	mock.
		ExpectQuery(`SELECT o.id, o.status, o.order_date, o.price, o.delivery_time, a.city, a.street, a.house_number, a.flat_number
		FROM orders o
		   JOIN orders_address oa on o.id = oa.orders_id
		   JOIN address a on a.id = oa.address_id
		   WHERE`).WithArgs(reqorder.UserId, reqorder.OrderId).
		WillReturnRows(row)

	rows := sqlmock.
		NewRows([]string{"p.id", "p.name", "p.price", "p.icon", "op.item_count"})
	expectprod := []*dto.RespGetOrderProduct{
		{
			Id:    1,
			Name:  "Burger",
			Price: 134.0,
			Icon:  "deih",
			Count: 3,
		},
		{
			Id:    2,
			Name:  "Burger1",
			Price: 134.0,
			Icon:  "deih",
			Count: 3,
		},
	}

	for _, product := range expectprod {
		rows = rows.AddRow(product.Id, product.Name, product.Price, product.Icon, product.Count)
	}

	mock.
		ExpectQuery(`SELECT p.id, p.name, p.price, p.icon, op.item_count
		FROM product p
		JOIN orders_product op ON p.id = op.product_id
		JOIN orders o ON o.id = op.order_id
		WHERE`).WithArgs(order.Id).
		WillReturnRows(rows)

	rows = sqlmock.
		NewRows([]string{"r.name"})
	expectres := "Burger King"

	rows.AddRow(expectres)

	mock.
		ExpectQuery(`SELECT r.name
		FROM restaurant r
		JOIN menu_type mt ON mt.restaurant_id = r.id
		JOIN product_menu_type pmt ON pmt.menu_type_id = mt.id 
		JOIN product p ON p.id = pmt.id
		WHERE`).WithArgs(order.Id).
		WillReturnRows(rows)

	actual, err := repo.GetOrder(reqorder)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if !reflect.DeepEqual(actual.OrderItems, order.OrderItems) {
		t.Errorf("results not match, want %v, have %v", &actual.OrderItems, &order.OrderItems)
		return
	}

}
