package postgres

import (
	"database/sql"
	"server/internal/User/repository"
	"server/internal/domain/dto"
	"server/internal/domain/entity"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) repository.UserRepositoryI {
	return &UserRepo{
		DB: db,
	}
}

func (repo *UserRepo) FindUserByUsername(value string) (*entity.User, error) {
	user := &entity.User{}
	row := repo.DB.QueryRow("SELECT id, username, password, birthday, phone_number, email, icon FROM users WHERE username = $1", value)
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Birthday, &user.PhoneNumber, &user.Email, &user.Icon)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, entity.ErrInternalServerError
	}
	return user, nil
}

func (repo *UserRepo) FindUserByEmail(value string) (*entity.User, error) {
	user := &entity.User{}
	row := repo.DB.QueryRow("SELECT id, username, password, birthday, phone_number, email, icon FROM users WHERE email = $1", value)
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Birthday, &user.PhoneNumber, &user.Email, &user.Icon)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, entity.ErrInternalServerError
	}
	return user, nil
}

func (repo *UserRepo) FindUserByPhone(value string) (*entity.User, error) {
	user := &entity.User{}
	row := repo.DB.QueryRow("SELECT id, username, password, birthday, phone_number, email, icon FROM users WHERE phone_number = $1", value)
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Birthday, &user.PhoneNumber, &user.Email, &user.Icon)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, entity.ErrInternalServerError
	}
	return user, nil
}

func (repo *UserRepo) FindUserById(id uint) (*entity.User, error) {
	user := &entity.User{}
	row := repo.DB.QueryRow("SELECT id, username, password, birthday, email FROM users WHERE id = $1", id)
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Birthday, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, entity.ErrInternalServerError
	}
	return user, nil
}

func (repo *UserRepo) CreateUser(in *dto.DBCreateUser) (uint, error) {

	insertUser := `INSERT INTO users (username, password, birthday, phone_number, email, icon) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := repo.DB.Exec(insertUser, in.Username, in.Password, in.Birthday, in.PhoneNumber, in.Email, in.Icon)
	if err != nil {
		return 0, entity.ErrInternalServerError
	}
	var ID uint
	row := repo.DB.QueryRow("SELECT ID FROM users WHERE username = $1", in.Username)
	err = row.Scan(&ID)
	if err != nil {
		return 0, entity.ErrInternalServerError
	}

	return ID, nil
}
