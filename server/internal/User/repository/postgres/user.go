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

func (repo *UserRepo) FindUserByUsername(value string) (*dto.DBGetUser, error) {
	user := &dto.DBGetUser{}
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

func (repo *UserRepo) FindUserByEmail(value string) (*dto.DBGetUser, error) {
	user := &dto.DBGetUser{}
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

func (repo *UserRepo) FindUserByPhone(value string) (*dto.DBGetUser, error) {
	user := &dto.DBGetUser{}
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

func (repo *UserRepo) FindUserById(id uint) (*dto.DBGetUser, error) {
	user := &dto.DBGetUser{}
	row := repo.DB.QueryRow("SELECT id, username, password, birthday, phone_number, email, icon FROM users WHERE id = $1", id)
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Birthday, &user.PhoneNumber, &user.Email, &user.Icon)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, entity.ErrInternalServerError
	}
	return user, nil
}

func (repo *UserRepo) CreateUser(user *dto.DBCreateUser) (uint, error) {

	insertUser := `INSERT INTO users (username, password, birthday, phone_number, email, icon) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := repo.DB.Exec(insertUser, user.Username, user.Password, user.Birthday, user.PhoneNumber, user.Email, user.Icon)
	if err != nil {
		return 0, entity.ErrInternalServerError
	}
	var ID uint
	row := repo.DB.QueryRow("SELECT ID FROM users WHERE username = $1", user.Username)
	err = row.Scan(&ID)
	if err != nil {
		return 0, entity.ErrInternalServerError
	}

	return ID, nil
}

func (repo *UserRepo) UpdateUser(user *dto.DBUpdateUser) (error) {
	updateUser := `UPDATE users 
				   SET username = $1, password = $2, birthday = $3, phone_number = $4, email = $5, icon = $6
				   WHERE id = $7`
	_, err := repo.DB.Exec(updateUser, user.Username, user.Password, user.Birthday, user.PhoneNumber, user.Email, user.Icon, user.ID)
	if err != nil {
		return entity.ErrInternalServerError
	}
	return nil
}
