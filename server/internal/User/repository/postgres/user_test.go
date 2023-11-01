package postgres

import (
	"database/sql"
	"reflect"
	"server/internal/domain/dto"
	"server/internal/domain/entity"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func TestFindUserByUsername(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	var username string = "john_doe"

	rows := sqlmock.
		NewRows([]string{"id", "username", "password", "birthday", "phone_number", "email", "icon"})

	expect := []*entity.User{
		{1, "john_doe", "secure_password", "1990-05-15", "89165342399", "john@example.com", "deficon"},
	}
	for _, user := range expect {
		rows = rows.AddRow(user.ID, user.Username, user.Password, user.Birthday, user.PhoneNumber, user.Email, user.Icon)
	}

	mock.
		ExpectQuery("SELECT id, username, password, birthday, phone_number, email, icon FROM users WHERE").
		WithArgs(username).
		WillReturnRows(rows)

	repo := &UserRepo{
		DB: db,
	}

	user, err := repo.FindUserByUsername(username)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(user, expect[0]) {
		t.Errorf("results not match, want %v, have %v", expect[0], user)
		return
	}
}

func TestFindUserByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	var email string = "jane@example.com"

	rows := sqlmock.
		NewRows([]string{"id", "username", "password", "birthday", "phone_number", "email", "icon"})

	expect := []*entity.User{
		{2, "jane_smith", "another_password", "1985-08-22", "8916534239", "jane@example.com", "deficon"},
	}

	for _, user := range expect {
		rows = rows.AddRow(user.ID, user.Username, user.Password, user.Birthday, user.PhoneNumber, user.Email, user.Icon)
	}

	mock.
		ExpectQuery("SELECT id, username, password, birthday, phone_number, email, icon FROM users WHERE").
		WithArgs(email).
		WillReturnRows(rows)

	repo := &UserRepo{
		DB: db,
	}

	user, err := repo.FindUserByEmail(email)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(user, expect[0]) {
		t.Errorf("results not match, want %v, have %v", expect[0], user)
		return
	}
}

func TestFindUserByPhone(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	var phone string = "8916534239"

	rows := sqlmock.
		NewRows([]string{"id", "username", "password", "birthday", "phone_number", "email", "icon"})

	expect := []*entity.User{
		{2, "jane_smith", "another_password", "1985-08-22", "8916534239", "jane@example.com", "deficon"},
	}

	for _, user := range expect {
		rows = rows.AddRow(user.ID, user.Username, user.Password, user.Birthday, user.PhoneNumber, user.Email, user.Icon)
	}

	mock.
		ExpectQuery("SELECT id, username, password, birthday, phone_number, email, icon FROM users WHERE").
		WithArgs(phone).
		WillReturnRows(rows)

	repo := &UserRepo{
		DB: db,
	}

	user, err := repo.FindUserByPhone(phone)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(user, expect[0]) {
		t.Errorf("results not match, want %v, have %v", expect[0], user)
		return
	}
}

func TestFindUserById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	var id uint = 2

	rows := sqlmock.
		NewRows([]string{"id", "username", "password", "birthday", "phone_number", "email", "icon"})

	expect := []*entity.User{
		{2, "jane_smith", "another_password", "1985-08-22", "8916534239", "jane@example.com", "deficon"},
	}

	for _, user := range expect {
		rows = rows.AddRow(user.ID, user.Username, user.Password, user.Birthday, user.PhoneNumber, user.Email, user.Icon)
	}

	mock.
		ExpectQuery("SELECT id, username, password, birthday, phone_number, email, icon FROM users WHERE").
		WithArgs(id).
		WillReturnRows(rows)

	repo := &UserRepo{
		DB: db,
	}

	user, err := repo.FindUserById(id)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(user, expect[0]) {
		t.Errorf("results not match, want %v, have %v", expect[0], user)
		return
	}
}

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	user := &dto.DBCreateUser{
		Username:    "john_doe",
		Password:    "secure_password",
		Birthday:    sql.NullString{String: "1985-08-22", Valid: true},
		PhoneNumber: "8916534239",
		Email:       "john@example.com",
		Icon:        sql.NullString{String: "deficon", Valid: true},
	}

	rows := sqlmock.
		NewRows([]string{"id"}).AddRow(1)

	mock.
		ExpectExec(`INSERT INTO users`).
		WithArgs(user.Username, user.Password, user.Birthday, user.PhoneNumber, user.Email, user.Icon).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.
		ExpectQuery("SELECT ID FROM users WHERE").
		WithArgs(user.Username).
		WillReturnRows(rows)

	repo := &UserRepo{
		DB: db,
	}

	id, err := repo.CreateUser(user)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if id != 1 {
		t.Errorf("bad id: want %v, have %v", id, 1)
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
