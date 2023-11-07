package repository

import (
	"errors"
	"reflect"
	"server/internal/domain/entity"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func TestGetRestaurantsSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &restaurantRepo{
		DB: db,
	}

	rows := sqlmock.
		NewRows([]string{"id", "name", "rating", "comments_count", "category", "icon"})
	expect := []*entity.Restaurant{
		{ID: 1,
			Name:          "Burger King",
			Rating:        3.7,
			CommentsCount: 60,
			Category:      "Fastfood",
			Icon:          "img/burger_king.jpg",
		},
		{ID: 2,
			Name:          "MacBurger",
			Rating:        3.8,
			CommentsCount: 69,
			Category:      "Fastfood",
			Icon:          "img/mac_burger.jpg",
		},
	}
	for _, restaurant := range expect {
		rows = rows.AddRow(restaurant.ID, restaurant.Name, restaurant.Rating, restaurant.CommentsCount, restaurant.Category, restaurant.Icon)
	}

	mock.
		ExpectQuery("SELECT id, name, rating, comments_count, category, icon FROM restaurant").
		WillReturnRows(rows)

	restaurants, err := repo.GetRestaurants()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if !reflect.DeepEqual(restaurants[0], expect[0]) {
		t.Errorf("results not match, want %v, have %v", expect[0], restaurants[0])
		return
	}

	emptyrows := sqlmock.NewRows([]string{"id", "name", "rating", "comments_count", "category", "icon"})

	mock.ExpectQuery("SELECT id, name, rating, comments_count, category, icon FROM restaurant").
		WillReturnRows(emptyrows)

	restaurants, err = repo.GetRestaurants()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

}

func TestGetRestaurantsFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &restaurantRepo{
		DB: db,
	}

	testErr := errors.New("test")
	mock.ExpectQuery("SELECT id, name, rating, comments_count, category, icon FROM restaurant").
		WillReturnError(testErr)

	restaurants, err := repo.GetRestaurants()
	if err != testErr {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if restaurants != nil {
		t.Errorf("restaurants not nil while error")
	}

}

func TestGetRestaurantByIdSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &restaurantRepo{
		DB: db,
	}

	row := sqlmock.
		NewRows([]string{"id", "name", "rating", "comments_count", "category", "icon"})
	expect := &entity.Restaurant{
		ID:            1,
		Name:          "Burger King",
		Rating:        3.7,
		CommentsCount: 60,
		Category:      "Fastfood",
		Icon:          "img/burger_king.jpg",
	}

	row = row.AddRow(expect.ID, expect.Name, expect.Rating, expect.CommentsCount, expect.Category, expect.Icon)

	var elemID = 1

	mock.
		ExpectQuery("SELECT id, name, rating, comments_count, category, icon FROM restaurant WHERE").WithArgs(elemID).
		WillReturnRows(row)

	restaurant, err := repo.GetRestaurantById(uint(elemID))
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if !reflect.DeepEqual(restaurant, expect) {
		t.Errorf("results not match, want %v, have %v", expect, restaurant)
		return
	}

	emptyrows := sqlmock.NewRows([]string{"id", "name", "rating", "comments_count", "category", "icon"})

	mock.ExpectQuery("SELECT id, name, rating, comments_count, category, icon FROM restaurant").WithArgs(elemID).
		WillReturnRows(emptyrows)

	restaurant, err = repo.GetRestaurantById(uint(elemID))
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

}

func TestGetRestaurantByIdFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &restaurantRepo{
		DB: db,
	}

	var elemID = 1

	testErr := errors.New("test")
	mock.ExpectQuery("SELECT id, name, rating, comments_count, category, icon FROM restaurant").WithArgs(elemID).
		WillReturnError(testErr)

	restaurants, err := repo.GetRestaurantById(uint(elemID))
	if err != entity.ErrInternalServerError {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if restaurants != nil {
		t.Errorf("restaurants not nil while error")
	}
}

func TestGetMenuTypesByRestaurantIdSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &restaurantRepo{
		DB: db,
	}

	var elemID = 1

	rows := sqlmock.
		NewRows([]string{"id", "name", "restaurant_id"})
	expect := []*entity.MenuType{
		{ID: 1,
			Name:         "Популярное",
			RestaurantID: 1,
		},
		{ID: 2,
			Name:         "Новинки",
			RestaurantID: 1,
		},
	}

	for _, menu_type := range expect {
		rows = rows.AddRow(menu_type.ID, menu_type.Name, menu_type.RestaurantID)
	}

	mock.
		ExpectQuery("SELECT id, name, restaurant_id FROM menu_type WHERE").WithArgs(elemID).
		WillReturnRows(rows)

	menu_types, err := repo.GetMenuTypesByRestaurantId(uint(elemID))
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if !reflect.DeepEqual(menu_types[0], expect[0]) {
		t.Errorf("results not match, want %v, have %v", expect[0], menu_types[0])
		return
	}

	emptyrows := sqlmock.NewRows([]string{"id", "name", "restaurant_id"})

	mock.ExpectQuery("SELECT id, name, restaurant_id FROM menu_type WHERE").WithArgs(elemID).
		WillReturnRows(emptyrows)

	menu_types, err = repo.GetMenuTypesByRestaurantId(uint(elemID))
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestGetMenuTypesByRestaurantIdFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &restaurantRepo{
		DB: db,
	}

	var elemID = 1

	testErr := errors.New("test")
	mock.ExpectQuery("SELECT id, name, restaurant_id FROM menu_type WHERE").WithArgs(elemID).
		WillReturnError(testErr)

	restaurants, err := repo.GetMenuTypesByRestaurantId(uint(elemID))
	if err != testErr {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if restaurants != nil {
		t.Errorf("restaurants not nil while error")
	}
}
