package repository

import (
	"errors"
	"reflect"
	"server/internal/domain/entity"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func TestGetProductsByMenuTypeIdSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &productRepo{
		DB: db,
	}

	rows := sqlmock.
		NewRows([]string{"id", "name", "price", "cooking_time", "portion", "description", "icon"})
	expect := []*entity.Product{
		{
			ID:          1,
			Name:        "Burger",
			Price:       120.0,
			CookingTime: 23,
			Portion:     "160 г",
			Description: "Nice burger",
			Icon:        "deficon",
		},
		{
			ID:          2,
			Name:        "Burger",
			Price:       120.0,
			CookingTime: 23,
			Portion:     "160 г",
			Description: "Nice burger",
			Icon:        "deficon",
		},
	}
	for _, product := range expect {
		rows = rows.AddRow(product.ID, product.Name, product.Price, product.CookingTime, product.Portion, product.Description, product.Icon)
	}

	mock.
		ExpectQuery("SELECT p.id, name, price, cooking_time, portion, description, icon  FROM product p INNER JOIN").WithArgs(uint(1)).
		WillReturnRows(rows)

	products, err := repo.GetProductsByMenuTypeId(uint(1))
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if !reflect.DeepEqual(products[0], expect[0]) {
		t.Errorf("results not match, want %v, have %v", expect[0], products[0])
		return
	}

}

func TestGetProductsByMenuTypeIdFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &productRepo{
		DB: db,
	}

	rows := sqlmock.
		NewRows([]string{"id", "name", "price", "cooking_time", "portion", "description", "icon"})
	expect := []*entity.Product{
		{
			ID:          1,
			Name:        "Burger",
			Price:       120.0,
			CookingTime: 23,
			Portion:     "160 г",
			Description: "Nice burger",
			Icon:        "deficon",
		},
		{
			ID:          2,
			Name:        "Burger",
			Price:       120.0,
			CookingTime: 23,
			Portion:     "160 г",
			Description: "Nice burger",
			Icon:        "deficon",
		},
	}
	for _, product := range expect {
		rows = rows.AddRow(product.ID, product.Name, product.Price, product.CookingTime, product.Portion, product.Description, product.Icon)
	}

	testErr := errors.New("test")

	mock.
		ExpectQuery("SELECT p.id, name, price, cooking_time, portion, description, icon  FROM product p INNER JOIN").WithArgs(uint(1)).
		WillReturnError(testErr)

	_, err = repo.GetProductsByMenuTypeId(uint(1))
	if err != testErr {
		t.Errorf("unexpected err: %s", err)
		return
	}

}

func TestGetProductByIDSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &productRepo{
		DB: db,
	}

	rows := sqlmock.
		NewRows([]string{"id", "name", "price", "cooking_time", "portion", "description", "icon"})
	expect := &entity.Product{
		ID:          1,
		Name:        "Burger",
		Price:       120.0,
		CookingTime: 23,
		Portion:     "160 г",
		Description: "Nice burger",
		Icon:        "deficon",
	}

	rows = rows.AddRow(expect.ID, expect.Name, expect.Price, expect.CookingTime, expect.Portion, expect.Description, expect.Icon)

	mock.
		ExpectQuery("SELECT id, name, price, cooking_time, portion, description, icon FROM product WHERE").WithArgs(uint(1)).
		WillReturnRows(rows)

	product, err := repo.GetProductByID(uint(1))
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if !reflect.DeepEqual(product, expect) {
		t.Errorf("results not match, want %v, have %v", expect, product)
		return
	}
}

func TestGetProductByIDFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &productRepo{
		DB: db,
	}

	rows := sqlmock.
		NewRows([]string{"id", "name", "price", "cooking_time", "portion", "description", "icon"})
	expect := &entity.Product{
		ID:          1,
		Name:        "Burger",
		Price:       120.0,
		CookingTime: 23,
		Portion:     "160 г",
		Description: "Nice burger",
		Icon:        "deficon",
	}

	rows = rows.AddRow(expect.ID, expect.Name, expect.Price, expect.CookingTime, expect.Portion, expect.Description, expect.Icon)

	testErr := errors.New("test")

	mock.
		ExpectQuery("SELECT id, name, price, cooking_time, portion, description, icon FROM product WHERE").WithArgs(uint(1)).
		WillReturnError(testErr)

	_, err = repo.GetProductByID(uint(1))
	if err != entity.ErrInternalServerError {
		t.Errorf("unexpected err: %s", err)
		return
	}
}
