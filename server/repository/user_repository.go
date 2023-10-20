package repository

import (
	"database/sql"
	"sync"
	"server/internal/domain/entity"
)

// type User struct {
// 	ID          uint           `json:"ID"`
// 	Username    string         `json:"Username"`
// 	Password    string         `json:"Password"`
// 	Birthday    sql.NullString `json:"Birthday"`
// 	PhoneNumber string         `json:"PhoneNumber"`
// 	Email       string         `json:"Email"`
// 	Icon        sql.NullString `json:"Icon"`
// }

type UserRepo struct {
	DB *sql.DB
	mu sync.RWMutex
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		mu: sync.RWMutex{},
		DB: db,
	}
}

func (repo *UserRepo) FindUserBy(field string, value string) (*entity.User, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()
	user := &entity.User{}
	switch field {
	case "Username":
		row := repo.DB.QueryRow("SELECT id, username, password, birthday, phone_number, email, icon FROM users WHERE username = $1", value)
		err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Birthday, &user.PhoneNumber, &user.Email, &user.Icon)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}
		return user, nil
	case "Email":
		row := repo.DB.QueryRow("SELECT id, username, password, birthday, phone_number, email, icon FROM users WHERE email = $1", value)
		err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Birthday, &user.PhoneNumber, &user.Email, &user.Icon)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}
		return user, nil
	case "PhoneNumber":
		row := repo.DB.QueryRow("SELECT id, username, password, birthday, phone_number, email, icon FROM users WHERE phone_number = $1", value)
		err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Birthday, &user.PhoneNumber, &user.Email, &user.Icon)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}
		return user, nil
	}

	return nil, nil
}

func (repo *UserRepo) GetUserById(id uint) (*entity.User, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()
	user := &entity.User{}
	row := repo.DB.QueryRow("SELECT id, username, password, birthday, email FROM users WHERE id = $1", id)
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Birthday, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (repo *UserRepo) CreateUser(in *entity.User) (uint, error) {

	repo.mu.Lock()
	insertUser := `INSERT INTO users (username, password, birthday, phone_number, email, icon) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := repo.DB.Exec(insertUser, in.Username, in.Password, in.Birthday, in.PhoneNumber, in.Email, in.Icon)
	if err != nil {
		return 0, err
	}
	repo.mu.Unlock()
	var ID uint
	row := repo.DB.QueryRow("SELECT ID FROM users WHERE username = $1", in.Username)
	err = row.Scan(&ID)
	if err != nil {
		return 0, err
	}

	return ID, nil
}
