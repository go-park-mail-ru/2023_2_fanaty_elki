package postgres

import (
	"database/sql"
	"errors"
	"reflect"
	"server/internal/domain/dto"
	"server/internal/domain/entity"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func TestFindUserByUsernameSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	var username string = "john_doe"

	rows := sqlmock.
		NewRows([]string{"id", "username", "password", "birthday", "phone_number", "email", "icon"})

	expect := &dto.DBGetUser{
		ID:          1,
		Username:    "john_doe",
		Password:    "secure_password",
		Birthday:    sql.NullString{Valid: true, String: "1995-04-04"},
		PhoneNumber: "89165342399",
		Email:       "john@example.com",
		Icon:        sql.NullString{Valid: true, String: "dificon"},
	}

	rows = rows.AddRow(expect.ID, expect.Username, expect.Password, expect.Birthday, expect.PhoneNumber, expect.Email, expect.Icon)

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
	if !reflect.DeepEqual(user, expect) {
		t.Errorf("results not match, want %v, have %v", expect, user)
		return
	}
}

func TestFindUserByUsernameFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	var username string = "john_doe"

	rows := sqlmock.
		NewRows([]string{"id", "username", "password", "birthday", "phone_number", "email", "icon"})

	expect := &dto.DBGetUser{
		ID:          1,
		Username:    "john_doe",
		Password:    "secure_password",
		Birthday:    sql.NullString{Valid: true, String: "1995-04-04"},
		PhoneNumber: "89165342399",
		Email:       "john@example.com",
		Icon:        sql.NullString{Valid: true, String: "dificon"},
	}

	rows = rows.AddRow(expect.ID, expect.Username, expect.Password, expect.Birthday, expect.PhoneNumber, expect.Email, expect.Icon)

	testErr := errors.New("testErr")

	mock.
		ExpectQuery("SELECT id, username, password, birthday, phone_number, email, icon FROM users WHERE").
		WithArgs(username).
		WillReturnError(testErr)

	repo := &UserRepo{
		DB: db,
	}

	_, err = repo.FindUserByUsername(username)
	if err != entity.ErrInternalServerError {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
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

	expect := &dto.DBGetUser{
		ID:          1,
		Username:    "john_doe",
		Password:    "secure_password",
		Birthday:    sql.NullString{Valid: true, String: "1995-04-04"},
		PhoneNumber: "89165342399",
		Email:       "john@example.com",
		Icon:        sql.NullString{Valid: true, String: "dificon"},
	}

	rows = rows.AddRow(expect.ID, expect.Username, expect.Password, expect.Birthday, expect.PhoneNumber, expect.Email, expect.Icon)

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
	if !reflect.DeepEqual(user, expect) {
		t.Errorf("results not match, want %v, have %v", expect, user)
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

	expect := &dto.DBGetUser{
		ID:          1,
		Username:    "john_doe",
		Password:    "secure_password",
		Birthday:    sql.NullString{Valid: true, String: "1995-04-04"},
		PhoneNumber: "89165342399",
		Email:       "john@example.com",
		Icon:        sql.NullString{Valid: true, String: "dificon"},
	}

	rows = rows.AddRow(expect.ID, expect.Username, expect.Password, expect.Birthday, expect.PhoneNumber, expect.Email, expect.Icon)

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
	if !reflect.DeepEqual(user, expect) {
		t.Errorf("results not match, want %v, have %v", expect, user)
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

	expect := &dto.DBGetUser{
		ID:          1,
		Username:    "john_doe",
		Password:    "secure_password",
		Birthday:    sql.NullString{Valid: true, String: "1995-04-04"},
		PhoneNumber: "89165342399",
		Email:       "john@example.com",
		Icon:        sql.NullString{Valid: true, String: "dificon"},
	}

	rows = rows.AddRow(expect.ID, expect.Username, expect.Password, expect.Birthday, expect.PhoneNumber, expect.Email, expect.Icon)

	mock.
		ExpectQuery("SELECT id, username, password, birthday, phone_number, email, icon FROM users WHERE").
		WithArgs(id).
		WillReturnRows(rows)

	repo := &UserRepo{
		DB: db,
	}

	user, err := repo.FindUserByID(id)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(user, expect) {
		t.Errorf("results not match, want %v, have %v", expect, user)
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

func TestUpdateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	user := &dto.DBUpdateUser{
		ID:          1,
		Username:    "john_doe",
		Password:    "secure_password",
		Birthday:    sql.NullString{String: "1985-08-22", Valid: true},
		PhoneNumber: "8916534239",
		Email:       "john@example.com",
		Icon:        sql.NullString{String: "deficon", Valid: true},
	}

	mock.
		ExpectExec(`UPDATE users SET `).
		WithArgs(user.Username, user.Password, user.Birthday, user.PhoneNumber, user.Email, user.Icon, user.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := &UserRepo{
		DB: db,
	}

	err = repo.UpdateUser(user)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
