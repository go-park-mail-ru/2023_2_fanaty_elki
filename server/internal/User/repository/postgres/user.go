package postgres

import (
	"context"
	"fmt"
	"server/internal/User/repository"
	"server/internal/domain/entity"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo struct {
	DB *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) repository.UserRepositoryI {
	return &UserRepo{
		DB: db,
	}
}

func (repo *UserRepo) FindUserBy(field string, value string) (*entity.User, error) {
	user := &entity.User{}
	switch field {
	case "Username":
		row := repo.DB.QueryRow(context.Background(), "SELECT id, username, password, birthday, phone_number, email, icon FROM users WHERE username = $1", value)
		err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Birthday, &user.PhoneNumber, &user.Email, &user.Icon)
		if err != nil {
			fmt.Println()
			if err.Error() == pgx.ErrNoRows.Error() {

				return nil, nil
			}
			return nil, entity.ErrInternalServerError
		}
		return user, nil
	case "Email":
		row := repo.DB.QueryRow(context.Background(), "SELECT id, username, password, birthday, phone_number, email, icon FROM users WHERE email = $1", value)
		err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Birthday, &user.PhoneNumber, &user.Email, &user.Icon)
		if err != nil {
			if err.Error() == pgx.ErrNoRows.Error() {
				return nil, nil
			}
			return nil, entity.ErrInternalServerError
		}
		return user, nil
	case "PhoneNumber":
		row := repo.DB.QueryRow(context.Background(), "SELECT id, username, password, birthday, phone_number, email, icon FROM users WHERE phone_number = $1", value)
		err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Birthday, &user.PhoneNumber, &user.Email, &user.Icon)
		if err != nil {
			if err.Error() == pgx.ErrNoRows.Error() {
				return nil, nil
			}
			return nil, entity.ErrInternalServerError
		}
		return user, nil
	}

	return nil, nil
}

func (repo *UserRepo) GetUserById(id uint) (*entity.User, error) {
	user := &entity.User{}
	row := repo.DB.QueryRow(context.Background(), "SELECT id, username, password, birthday, email FROM users WHERE id = $1", id)
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Birthday, &user.Email)
	if err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			return nil, nil
		}
		return nil, entity.ErrInternalServerError
	}
	return user, nil
}

func (repo *UserRepo) CreateUser(in *entity.User) (uint, error) {
	insertUser := `INSERT INTO users (username, password, birthday, phone_number, email, icon) VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := repo.DB.Exec(context.Background(), insertUser, in.Username, in.Password, in.Birthday, in.PhoneNumber, in.Email, in.Icon)
	if err != nil {
		return 0, entity.ErrInternalServerError

	}
	var ID uint
	row := repo.DB.QueryRow(context.Background(), "SELECT ID FROM users WHERE username = $1", in.Username)
	err = row.Scan(&ID)
	if err != nil {
		return 0, entity.ErrInternalServerError
	}

	return ID, nil
}
