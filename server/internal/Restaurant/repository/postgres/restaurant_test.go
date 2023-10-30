package repository

import (
	"reflect"
	"server/internal/domain/entity"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func TestGetRestaurants(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	rows := sqlmock.
		NewRows([]string{"id", "name", "rating", "comments_count", "category", "icon"})
	expect := []*entity.Restaurant{
		{1, "Burger King", 3.7, 60, "Fastfood", "img/burger_king.jpg"},
		{2, "MacBurger", 3.8, 69, "Fastfood", "img/mac_burger.jpg"},
	}
	for _, restaurant := range expect {
		rows = rows.AddRow(restaurant.ID, restaurant.Name, restaurant.Rating, restaurant.CommentsCount, restaurant.Category, restaurant.Icon)
	}

	mock.
		ExpectQuery("SELECT id, name, rating, comments_count, category, icon FROM restaurant").
		WillReturnRows(rows)

	repo := &restaurantRepo{
		DB: db,
	}
	restaurants, err := repo.GetRestaurants()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(restaurants[0], expect[0]) {
		t.Errorf("results not match, want %v, have %v", expect[0], restaurants[1])
		return
	}
}
