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

	repo := &RestaurantRepo{
		DB: db,
	}

	rows := sqlmock.
		NewRows([]string{"id", "name", "rating", "comments_count", "icon"})
	expect := []*entity.Restaurant{
		{ID: 1,
			Name:          "Burger King",
			Rating:        3.7,
			CommentsCount: 60,
			Icon:          "img/burger_king.webp",
		},
		{ID: 2,
			Name:          "MacBurger",
			Rating:        3.8,
			CommentsCount: 69,
			Icon:          "img/mac_burger.webp",
		},
	}
	for _, restaurant := range expect {
		rows = rows.AddRow(restaurant.ID, restaurant.Name, restaurant.Rating, restaurant.CommentsCount, restaurant.Icon)
	}

	mock.
		ExpectQuery("SELECT id, name, rating, comments_count, icon FROM restaurant").
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

	emptyrows := sqlmock.NewRows([]string{"id", "name", "rating", "comments_count", "icon"})

	mock.ExpectQuery("SELECT id, name, rating, comments_count, icon FROM restaurant").
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

	repo := &RestaurantRepo{
		DB: db,
	}

	testErr := errors.New("test")
	mock.ExpectQuery("SELECT id, name, rating, comments_count, icon FROM restaurant").
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

	repo := &RestaurantRepo{
		DB: db,
	}

	row := sqlmock.
		NewRows([]string{"id", "name", "rating", "comments_count", "icon"})
	expect := &entity.Restaurant{
		ID:            1,
		Name:          "Burger King",
		Rating:        3.7,
		CommentsCount: 60,
		Icon:          "img/burger_king.webp",
	}

	row = row.AddRow(expect.ID, expect.Name, expect.Rating, expect.CommentsCount, expect.Icon)

	var elemID = 1

	mock.
		ExpectQuery("SELECT id, name, rating, comments_count, icon FROM restaurant WHERE").WithArgs(elemID).
		WillReturnRows(row)

	restaurant, err := repo.GetRestaurantByID(uint(elemID))
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if !reflect.DeepEqual(restaurant, expect) {
		t.Errorf("results not match, want %v, have %v", expect, restaurant)
		return
	}

	emptyrows := sqlmock.NewRows([]string{"id", "name", "rating", "comments_count", "icon"})

	mock.ExpectQuery("SELECT id, name, rating, comments_count, icon FROM restaurant").WithArgs(elemID).
		WillReturnRows(emptyrows)

	restaurant, err = repo.GetRestaurantByID(uint(elemID))
	if err != entity.ErrNotFound {
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

	repo := &RestaurantRepo{
		DB: db,
	}

	var elemID = 1

	testErr := errors.New("test")
	mock.ExpectQuery("SELECT id, name, rating, comments_count,  icon FROM restaurant").WithArgs(elemID).
		WillReturnError(testErr)

	restaurants, err := repo.GetRestaurantByID(uint(elemID))
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

	repo := &RestaurantRepo{
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

	for _, menuType := range expect {
		rows = rows.AddRow(menuType.ID, menuType.Name, menuType.RestaurantID)
	}

	mock.
		ExpectQuery("SELECT id, name, restaurant_id FROM menu_type WHERE").WithArgs(elemID).
		WillReturnRows(rows)

	menuTypes, err := repo.GetMenuTypesByRestaurantID(uint(elemID))
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if !reflect.DeepEqual(menuTypes[0], expect[0]) {
		t.Errorf("results not match, want %v, have %v", expect[0], menuTypes[0])
		return
	}

	emptyrows := sqlmock.NewRows([]string{"id", "name", "restaurant_id"})

	mock.ExpectQuery("SELECT id, name, restaurant_id FROM menu_type WHERE").WithArgs(elemID).
		WillReturnRows(emptyrows)

	menuTypes, err = repo.GetMenuTypesByRestaurantID(uint(elemID))
	if err != entity.ErrNotFound {
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

	repo := &RestaurantRepo{
		DB: db,
	}

	var elemID = 1

	testErr := errors.New("test")
	mock.ExpectQuery("SELECT id, name, restaurant_id FROM menu_type WHERE").WithArgs(elemID).
		WillReturnError(testErr)

	restaurants, err := repo.GetMenuTypesByRestaurantID(uint(elemID))
	if err != testErr {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if restaurants != nil {
		t.Errorf("restaurants not nil while error")
	}
}

func TestGetCategoriesByRestaurantIdSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &RestaurantRepo{
		DB: db,
	}

	var elemID = 1

	rows := sqlmock.
		NewRows([]string{"id", "name"})
	expect := []*entity.Category{
		{ID: 1,
			Name: "Бургеры",
		},
		{ID: 2,
			Name: "Завтраки",
		},
	}

	for _, category := range expect {
		rows = rows.AddRow(category.ID, category.Name)
	}

	mock.
		ExpectQuery(`SELECT category.id, category.name 
		FROM restaurant_category rc 
		INNER JOIN category ON rc.category_id=category.id 
		WHERE`).WithArgs(elemID).
		WillReturnRows(rows)

	categories, err := repo.GetCategoriesByRestaurantID(uint(elemID))
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if !reflect.DeepEqual(categories[0], expect[0]) {
		t.Errorf("results not match, want %v, have %v", expect[0], categories[0])
		return
	}

	emptyrows := sqlmock.NewRows([]string{"id", "name"})

	mock.ExpectQuery(`SELECT category.id, category.name 
	FROM restaurant_category rc 
	INNER JOIN category ON rc.category_id=category.id 
	WHERE`).WithArgs(elemID).
		WillReturnRows(emptyrows)

	categories, err = repo.GetCategoriesByRestaurantID(uint(elemID))
	if err != entity.ErrNotFound {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestGetCategoriesByRestaurantIdFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &RestaurantRepo{
		DB: db,
	}

	var elemID = 1

	rows := sqlmock.
		NewRows([]string{"id", "name"})
	expect := []*entity.Category{
		{ID: 1,
			Name: "Бургеры",
		},
		{ID: 2,
			Name: "Завтраки",
		},
	}

	for _, category := range expect {
		rows = rows.AddRow(category.ID, category.Name)
	}

	testErr := errors.New("test")

	mock.
		ExpectQuery(`SELECT category.id, category.name 
		FROM restaurant_category rc 
		INNER JOIN category ON rc.category_id=category.id 
		WHERE`).WithArgs(elemID).
		WillReturnError(testErr)

	_, err = repo.GetCategoriesByRestaurantID(uint(elemID))
	if err != testErr {
		t.Errorf("unexpected err: %s", err)
		return
	}

}

func TestGetRestaurantsByCategorySuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &RestaurantRepo{
		DB: db,
	}

	rows := sqlmock.
		NewRows([]string{"id", "name", "rating", "comments_count", "icon"})
	expect := []*entity.Restaurant{
		{ID: 1,
			Name:          "Burger King",
			Rating:        3.7,
			CommentsCount: 60,
			Icon:          "img/burger_king.webp",
		},
		{ID: 2,
			Name:          "MacBurger",
			Rating:        3.8,
			CommentsCount: 69,
			Icon:          "img/mac_burger.webp",
		},
	}
	for _, restaurant := range expect {
		rows = rows.AddRow(restaurant.ID, restaurant.Name, restaurant.Rating, restaurant.CommentsCount, restaurant.Icon)
	}

	var name = "Бургеры"

	mock.
		ExpectQuery(`SELECT restaurant.id, restaurant.name, rating, comments_count, icon 
		FROM restaurant_category rc 
		INNER JOIN restaurant ON rc.restaurant_id=restaurant.id
		INNER JOIN category ON rc.category_id=category.id 
		WHERE `).
		WillReturnRows(rows)

	restaurants, err := repo.GetRestaurantsByCategory(name)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if !reflect.DeepEqual(restaurants[0], expect[0]) {
		t.Errorf("results not match, want %v, have %v", expect[0], restaurants[0])
		return
	}

	emptyrows := sqlmock.NewRows([]string{"id", "name", "rating", "comments_count", "icon"})

	mock.ExpectQuery("SELECT id, name, rating, comments_count, icon FROM restaurant").
		WillReturnRows(emptyrows)

	restaurants, err = repo.GetRestaurants()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

}

func TestGetRestaurantsByCategoryFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &RestaurantRepo{
		DB: db,
	}

	rows := sqlmock.
		NewRows([]string{"id", "name", "rating", "comments_count", "icon"})
	expect := []*entity.Restaurant{
		{ID: 1,
			Name:          "Burger King",
			Rating:        3.7,
			CommentsCount: 60,
			Icon:          "img/burger_king.webp",
		},
		{ID: 2,
			Name:          "MacBurger",
			Rating:        3.8,
			CommentsCount: 69,
			Icon:          "img/mac_burger.webp",
		},
	}
	for _, restaurant := range expect {
		rows = rows.AddRow(restaurant.ID, restaurant.Name, restaurant.Rating, restaurant.CommentsCount, restaurant.Icon)
	}

	var name = "Бургеры"

	testErr := errors.New("test")

	mock.
		ExpectQuery(`SELECT restaurant.id, restaurant.name, rating, comments_count, icon 
		FROM restaurant_category rc 
		INNER JOIN restaurant ON rc.restaurant_id=restaurant.id
		INNER JOIN category ON rc.category_id=category.id 
		WHERE `).
		WillReturnError(testErr)

	_, err = repo.GetRestaurantsByCategory(name)
	if err != testErr {
		t.Errorf("unexpected err: %s", err)
		return
	}

}

func TestGetCategoriesSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &RestaurantRepo{
		DB: db,
	}

	rows := sqlmock.
		NewRows([]string{"id", "name"})
	expect := []*entity.Category{
		{ID: 1,
			Name: "Бургеры",
		},
		{ID: 2,
			Name: "Завтраки",
		},
	}

	for _, category := range expect {
		rows = rows.AddRow(category.ID, category.Name)
	}

	mock.
		ExpectQuery(`SELECT id, name FROM category`).
		WillReturnRows(rows)

	categories, err := repo.GetCategories()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if !reflect.DeepEqual(categories[0], expect[0]) {
		t.Errorf("results not match, want %v, have %v", expect[0], categories[0])
		return
	}
}

func TestGetCategoriesFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &RestaurantRepo{
		DB: db,
	}

	rows := sqlmock.
		NewRows([]string{"id", "name"})
	expect := []*entity.Category{
		{ID: 1,
			Name: "Бургеры",
		},
		{ID: 2,
			Name: "Завтраки",
		},
	}

	for _, category := range expect {
		rows = rows.AddRow(category.ID, category.Name)
	}

	testErr := errors.New("test")

	mock.
		ExpectQuery(`SELECT id, name FROM category`).
		WillReturnError(testErr)

	_, err = repo.GetCategories()
	if err != testErr {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestSearchRestaurantsSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &RestaurantRepo{
		DB: db,
	}

	rows := sqlmock.
		NewRows([]string{"id", "name", "rating", "comments_count", "icon"})
	expect := []*entity.Restaurant{
		{ID: 1,
			Name:          "Burger King",
			Rating:        3.7,
			CommentsCount: 60,
			Icon:          "img/burger_king.webp",
		},
		{ID: 2,
			Name:          "MacBurger",
			Rating:        3.8,
			CommentsCount: 69,
			Icon:          "img/mac_burger.webp",
		},
	}
	for _, restaurant := range expect {
		rows = rows.AddRow(restaurant.ID, restaurant.Name, restaurant.Rating, restaurant.CommentsCount, restaurant.Icon)
	}

	var word = "Бургер"

	mock.
		ExpectQuery(`SELECT id, name, rating, comments_count, icon
		FROM restaurant 
		WHERE `).WithArgs(word).
		WillReturnRows(rows)

	restaurants, err := repo.SearchRestaurants(word)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if !reflect.DeepEqual(restaurants[0], expect[0]) {
		t.Errorf("results not match, want %v, have %v", expect[0], restaurants[0])
		return
	}
}

func TestSearchCategoriesSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &RestaurantRepo{
		DB: db,
	}

	rows := sqlmock.
		NewRows([]string{"id", "name", "rating", "comments_count", "icon"})
	expect := []*entity.Restaurant{
		{ID: 1,
			Name:          "Burger King",
			Rating:        3.7,
			CommentsCount: 60,
			Icon:          "img/burger_king.webp",
		},
		{ID: 2,
			Name:          "MacBurger",
			Rating:        3.8,
			CommentsCount: 69,
			Icon:          "img/mac_burger.webp",
		},
	}
	for _, restaurant := range expect {
		rows = rows.AddRow(restaurant.ID, restaurant.Name, restaurant.Rating, restaurant.CommentsCount, restaurant.Icon)
	}

	var word = "Бургер"

	mock.
		ExpectQuery(`SELECT restaurant.id, restaurant.name, rating, comments_count, icon
		FROM restaurant_category rc 
		INNER JOIN restaurant on rc.restaurant_id=restaurant.id
		INNER JOIN category on rc.category_id=category.id 
		WHERE`).WithArgs(word).
		WillReturnRows(rows)

	restaurants, err := repo.SearchCategories(word)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if !reflect.DeepEqual(restaurants[0], expect[0]) {
		t.Errorf("results not match, want %v, have %v", expect[0], restaurants[0])
		return
	}
}
