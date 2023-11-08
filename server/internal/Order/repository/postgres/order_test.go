package repository

import (
	"reflect"
	"server/internal/domain/dto"
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

	order := &dto.DBReqCreateOrder{
		Products: &map[uint]int{1: 2},
		UserId:   1,
		Status:   "CREATED",
		Date:     time.Now(),
	}

	respOrder := &dto.RespCreateOrder{
		Id:     1,
		Status: "CREATED",
		Date:   time.Now(),
	}

	var userID uint
	userID = 1

	var orderID uint
	orderID = 1

	rows := sqlmock.
		NewRows([]string{"id"}).AddRow(1)

	mock.
		ExpectQuery("INSERT INTO orders").
		WithArgs(userID, order.Date, order.Status).
		WillReturnRows(rows)

	mock.
		ExpectExec("INSERT INTO orders_product ").
		WithArgs(1, orderID, 2).
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
		Status: "CREATED",
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
		NewRows([]string{"id", "status", "created_at", "updated_at"})
	expect := []*dto.RespGetOrder{
		{
			Id:          1,
			Status:      "CREATED",
			Date:        time.Now(),
			UpdatedDate: time.Now(),
		},
		{
			Id:          2,
			Status:      "CREATED",
			Date:        time.Now(),
			UpdatedDate: time.Now(),
		},
	}

	for _, order := range expect {
		rows = rows.AddRow(order.Id, order.Status, order.Date, order.UpdatedDate)
	}

	mock.
		ExpectQuery("SELECT id, status, created_at, updated_at FROM orders").
		WillReturnRows(rows)

	var userID uint
	userID = 1

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

	row := sqlmock.
		NewRows([]string{"status", "created_at", "updated_at"})

	reqorder := &dto.ReqGetOneOrder{
		OrderId: 1,
		UserId:  1,
	}

	order := dto.RespGetOneOrder{
		Status:      "CREATED",
		Date:        time.Now(),
		UpdatedDate: time.Now(),
		Products: []*dto.RespGetOrderProduct{
			{
				Name:        "Burger",
				Price:       134.0,
				CookingTime: 80,
				Portion:     "120 г",
				Icon:        "deih",
				Count:       3,
			},
			{
				Name:        "Burger1",
				Price:       134.0,
				CookingTime: 80,
				Portion:     "120 г",
				Icon:        "deih",
				Count:       3,
			},
		},
	}

	resporder := &dto.RespGetOrder{
		Id:          1,
		Status:      "CREATED",
		Date:        time.Now(),
		UpdatedDate: time.Now(),
	}

	row = row.AddRow(resporder.Status, resporder.Date, resporder.UpdatedDate)

	mock.
		ExpectQuery("SELECT status, order_date, updated_at FROM orders WHERE").WithArgs(reqorder.OrderId, reqorder.UserId).
		WillReturnRows(row)

	rows := sqlmock.
		NewRows([]string{"p.name", "p.price", "p.cooking_time", "p.portion", "p.icon", "op.item_count"})
	expectprod := []*dto.RespGetOrderProduct{
		{
			Name:        "Burger",
			Price:       134.0,
			CookingTime: 80,
			Portion:     "120 г",
			Icon:        "deih",
			Count:       3,
		},
		{
			Name:        "Burger1",
			Price:       134.0,
			CookingTime: 80,
			Portion:     "120 г",
			Icon:        "deih",
			Count:       3,
		},
	}

	for _, order := range expectprod {
		rows = rows.AddRow(order.Name, order.Price, order.CookingTime, order.Portion, order.Icon, order.Count)
	}

	mock.
		ExpectQuery("SELECT p.name, p.price, p.cooking_time, p.portion, p.icon, op.item_count FROM product p JOIN orders_product op ON p.id = op.product_id JOIN orders o ON o.id = op.order_id WHERE").
		WillReturnRows(rows)

	actual, err := repo.GetOrder(reqorder)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if !reflect.DeepEqual(actual.Products, order.Products) {
		t.Errorf("results not match, want %v, have %v", &actual.Products, &order.Products)
		return
	}

}
