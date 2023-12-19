package repository

import (
	"errors"
	"reflect"
	"server/internal/domain/entity"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func TestCreateCartSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &CartRepo{
		DB: db,
	}

	var UserID uint
	UserID = 1

	rows := sqlmock.
		NewRows([]string{"id"}).AddRow(1)

	mock.
		ExpectExec(`INSERT INTO cart`).
		WithArgs(UserID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.
		ExpectQuery("SELECT ID FROM cart WHERE").
		WithArgs(UserID).
		WillReturnRows(rows)

	id, err := repo.CreateCart(UserID)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if id != 1 {
		t.Errorf("bad id: want %v, have %v", id, 1)
		return
	}

}

func TestCreateCartFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &CartRepo{
		DB: db,
	}

	var UserID uint
	UserID = 1

	testErr := errors.New("test")

	mock.
		ExpectExec(`INSERT INTO cart`).
		WithArgs(UserID).
		WillReturnError(testErr)

	id, err := repo.CreateCart(UserID)
	if err != entity.ErrInternalServerError {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if id != 0 {
		t.Errorf("bad id: want %v, have %v", id, 1)
		return
	}

	mock.
		ExpectExec(`INSERT INTO cart`).
		WithArgs(UserID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.
		ExpectQuery("SELECT ID FROM cart WHERE").
		WithArgs(UserID).
		WillReturnError(testErr)

	id, err = repo.CreateCart(UserID)
	if err != entity.ErrInternalServerError {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if id != 0 {
		t.Errorf("bad id: want %v, have %v", id, 1)
		return
	}

}

func TestGetCartByUserIDSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &CartRepo{
		DB: db,
	}

	row := sqlmock.
		NewRows([]string{"id", "user_id"})
	expect := &entity.Cart{
		ID:     1,
		UserID: 1,
	}

	row = row.AddRow(expect.ID, expect.UserID)

	var elemID = 1

	mock.
		ExpectQuery("SELECT id, user_id FROM cart WHERE").WithArgs(elemID).
		WillReturnRows(row)

	cart, err := repo.GetCartByUserID(uint(elemID))
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if !reflect.DeepEqual(cart, expect) {
		t.Errorf("results not match, want %v, have %v", expect, cart)
		return
	}

}

func TestGetCartByUserIDFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &CartRepo{
		DB: db,
	}

	testErr := errors.New("test")

	var elemID = 1

	mock.
		ExpectQuery("SELECT id, user_id FROM cart WHERE").WithArgs(elemID).
		WillReturnError(testErr)

	cart, err := repo.GetCartByUserID(uint(elemID))

	if err != entity.ErrInternalServerError {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if cart != nil {
		t.Errorf("not nil cart while error")
		return
	}
}

func TestGetCartProductsByCartIDSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &CartRepo{
		DB: db,
	}

	rows := sqlmock.
		NewRows([]string{"id", "product_id", "cart_id", "item_count"})
	expect := &entity.CartWithRestaurant{
		RestaurantID: 1,
		Products: []*entity.CartProduct{
			{
				ID:        1,
				ProductID: 1,
				CartID:    1,
				ItemCount: 6,
			},
			{
				ID:        2,
				ProductID: 3,
				CartID:    1,
				ItemCount: 6,
			},
		},
	}

	for _, product := range expect.Products {
		rows = rows.AddRow(product.ID, product.ProductID, product.CartID, product.ItemCount)
	}

	var elemID = 1

	mock.
		ExpectQuery("SELECT id, product_id, cart_id, item_count FROM cart_product WHERE").WithArgs(elemID).
		WillReturnRows(rows)

	restrows := sqlmock.
		NewRows([]string{"id"}).AddRow(1)

	promorows := sqlmock.
		NewRows([]string{"id"}).AddRow(1)

	mock.
		ExpectQuery(`SELECT mt.Restaurant_id
		FROM Product p
		JOIN Product_Menu_Type pmt ON p.id = pmt.Product_Id
		JOIN Menu_Type mt ON pmt.Menu_Type_id = mt.Id
		WHERE`).WithArgs(1).
		WillReturnRows(restrows)

	mock.
		ExpectQuery("SELECT promo_id FROM cart_promo WHERE").WithArgs(elemID).
		WillReturnRows(promorows)

	cartWithproducts, err := repo.GetCartProductsByCartID(uint(elemID))
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if !reflect.DeepEqual(cartWithproducts.RestaurantID, expect.RestaurantID) {
		t.Errorf("results not match, want %v, have %v", expect.RestaurantID, cartWithproducts.RestaurantID)
		return
	}

}

func TestGetCartProductsByCartIDFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &CartRepo{
		DB: db,
	}

	testErr := errors.New("test")

	var elemID = 1

	mock.
		ExpectQuery("SELECT id, product_id, cart_id, item_count FROM cart_product WHERE").WithArgs(elemID).
		WillReturnError(testErr)

	products, err := repo.GetCartProductsByCartID(uint(elemID))

	if err != testErr {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if products != nil {
		t.Errorf("carts not nil while error")
	}
}

func TestAddProductToCartSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &CartRepo{
		DB: db,
	}

	var cartID, productID uint
	cartID = 1
	productID = 1
	mock.
		ExpectExec(`INSERT INTO cart_product`).
		WithArgs(cartID, productID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.AddProductToCart(cartID, productID)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

}

func TestAddProductToCartFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &CartRepo{
		DB: db,
	}

	var cartID, productID uint
	cartID = 1
	productID = 1

	testErr := errors.New("test")

	mock.
		ExpectExec(`INSERT INTO cart_product`).
		WithArgs(cartID, productID).
		WillReturnError(testErr)

	err = repo.AddProductToCart(cartID, productID)
	if err != entity.ErrInternalServerError {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestDeleteProductFromCartSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &CartRepo{
		DB: db,
	}

	var cartID, productID uint
	cartID = 1
	productID = 1
	mock.
		ExpectExec(`DELETE FROM cart_product WHERE`).
		WithArgs(cartID, productID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.DeleteProductFromCart(cartID, productID)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestDeleteProductFromCartFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &CartRepo{
		DB: db,
	}

	var cartID, productID uint
	cartID = 1
	productID = 1

	testErr := errors.New("test")

	mock.
		ExpectExec(`DELETE FROM cart_product WHERE`).
		WithArgs(cartID, productID).
		WillReturnError(testErr)

	err = repo.DeleteProductFromCart(cartID, productID)
	if err != entity.ErrInternalServerError {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestUpdateItemCountUpSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &CartRepo{
		DB: db,
	}

	var cartID, productID uint
	cartID = 1
	productID = 1

	mock.
		ExpectExec(`UPDATE cart_product SET`).
		WithArgs(cartID, productID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.UpdateItemCountUp(cartID, productID)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

}

func TestUpdateItemCountUpFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &CartRepo{
		DB: db,
	}

	var cartID, productID uint
	cartID = 1
	productID = 1

	testErr := errors.New("test")

	mock.
		ExpectExec(`UPDATE cart_product SET`).
		WithArgs(cartID, productID).
		WillReturnError(testErr)

	err = repo.UpdateItemCountUp(cartID, productID)
	if err != entity.ErrInternalServerError {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestUpdateItemCountDownSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &CartRepo{
		DB: db,
	}

	var cartID, productID uint
	cartID = 1
	productID = 1

	mock.
		ExpectExec(`UPDATE cart_product SET`).
		WithArgs(cartID, productID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.UpdateItemCountDown(cartID, productID)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

}

func TestUpdateItemCountDownFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &CartRepo{
		DB: db,
	}

	var cartID, productID uint
	cartID = 1
	productID = 1

	testErr := errors.New("test")

	mock.
		ExpectExec(`UPDATE cart_product SET`).
		WithArgs(cartID, productID).
		WillReturnError(testErr)

	err = repo.UpdateItemCountDown(cartID, productID)
	if err != entity.ErrInternalServerError {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestCheckProductInCartSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &CartRepo{
		DB: db,
	}

	rows := sqlmock.
		NewRows([]string{"id", "product_id", "cart_id", "item_count"})
	expect := []*entity.CartProduct{
		{
			ID:        1,
			ProductID: 1,
			CartID:    1,
			ItemCount: 6,
		},
	}

	for _, product := range expect {
		rows = rows.AddRow(product.ID, product.ProductID, product.CartID, product.ItemCount)
	}

	var cartID, productID uint
	cartID = 1
	productID = 1

	mock.
		ExpectQuery("SELECT id, product_id, cart_id, item_count FROM cart_product WHERE").WithArgs(cartID, productID).
		WillReturnRows(rows)

	hasproducts, err := repo.CheckProductInCart(cartID, productID)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if !reflect.DeepEqual(true, hasproducts) {
		t.Errorf("results not match, want %v, have %v", true, hasproducts)
		return
	}

}

func TestCheckProductCountSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &CartRepo{
		DB: db,
	}

	rows := sqlmock.
		NewRows([]string{"id", "product_id", "cart_id", "item_count"})
	expect := []*entity.CartProduct{
		{
			ID:        1,
			ProductID: 1,
			CartID:    1,
			ItemCount: 6,
		},
	}

	for _, product := range expect {
		rows = rows.AddRow(product.ID, product.ProductID, product.CartID, product.ItemCount)
	}

	var cartID, productID uint
	cartID = 1
	productID = 1

	mock.
		ExpectQuery("SELECT id, product_id, cart_id, item_count FROM cart_product WHERE").WithArgs(cartID, productID).
		WillReturnRows(rows)

	productscount, err := repo.CheckProductCount(cartID, productID)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if !reflect.DeepEqual(uint(expect[0].ItemCount), productscount) {
		t.Errorf("results not match, want %v, have %v", uint(expect[0].ItemCount), productscount)
		return
	}

}

func TestCleanCartSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &CartRepo{
		DB: db,
	}

	var cartID uint
	cartID = 1
	mock.
		ExpectExec(`DELETE FROM cart_product WHERE`).
		WithArgs(cartID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.CleanCart(cartID)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}
