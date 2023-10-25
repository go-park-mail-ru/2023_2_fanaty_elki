package postgres

import (
	"context"
	"fmt"
	"server/internal/User/repository"
	"server/internal/domain/entity"

	sq "github.com/Masterminds/squirrel"
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
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	user := &entity.User{}
	switch field {
	case "Username":
		sql, _, err := psql.Select("id, username, password, birthday, phone_number, email, icon").From("users").Where("username = ?").ToSql()
		if err != nil {
			if err.Error() == pgx.ErrNoRows.Error() {
				return nil, nil
			}
			return nil, entity.ErrInternalServerError
		}
		row := repo.DB.QueryRow(context.Background(), sql, value)
		err = row.Scan(&user.ID, &user.Username, &user.Password, &user.Birthday, &user.PhoneNumber, &user.Email, &user.Icon)
		if err != nil {
			if err.Error() == pgx.ErrNoRows.Error() {
				return nil, nil
			}
			return nil, entity.ErrInternalServerError
		}
		return user, nil
	case "Email":
		sql, _, err := psql.Select("id, username, password, birthday, phone_number, email, icon").From("users").Where("email = ?").ToSql()
		if err != nil {
			if err.Error() == pgx.ErrNoRows.Error() {
				return nil, nil
			}
			return nil, entity.ErrInternalServerError
		}
		row := repo.DB.QueryRow(context.Background(), sql, value)
		err = row.Scan(&user.ID, &user.Username, &user.Password, &user.Birthday, &user.PhoneNumber, &user.Email, &user.Icon)
		if err != nil {
			if err.Error() == pgx.ErrNoRows.Error() {
				return nil, nil
			}
			return nil, entity.ErrInternalServerError
		}
		return user, nil
	case "PhoneNumber":
		sql, _, err := psql.Select("id, username, password, birthday, phone_number, email, icon").From("users").Where("phone_number = ?").ToSql()
		if err != nil {
			if err.Error() == pgx.ErrNoRows.Error() {
				return nil, nil
			}
			return nil, entity.ErrInternalServerError
		}
		row := repo.DB.QueryRow(context.Background(), sql, value)
		err = row.Scan(&user.ID, &user.Username, &user.Password, &user.Birthday, &user.PhoneNumber, &user.Email, &user.Icon)
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
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	user := &entity.User{}
	sql, _, err := psql.Select("id, username, password, birthday, email").From("users").Where("id = ?").ToSql()
	fmt.Println(sql)
	if err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			return nil, nil
		}
		return nil, entity.ErrInternalServerError
	}
	row := repo.DB.QueryRow(context.Background(), sql, id)
	err = row.Scan(&user.ID, &user.Username, &user.Password, &user.Birthday, &user.Email)
	if err != nil {
		fmt.Println(err.Error())
		if err.Error() == pgx.ErrNoRows.Error() {
			return nil, nil
		}
		return nil, entity.ErrInternalServerError
	}
	return user, nil
}

func (repo *UserRepo) CreateUser(in *entity.User) (uint, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	sql, _, err := psql.Insert("users").Columns("username", "password", "birthday", "phone_number", "email", "icon").Values("?", "?", "?", "?", "?", "?").ToSql()
	if err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			return 0, nil
		}
		return 0, entity.ErrInternalServerError
	}

	_, err = repo.DB.Exec(context.Background(), sql, in.Username, in.Password, in.Birthday, in.PhoneNumber, in.Email, in.Icon)
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
