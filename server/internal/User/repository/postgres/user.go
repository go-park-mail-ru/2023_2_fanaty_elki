package postgres

import (
	"database/sql"
	"sync"
	"server/internal/domain/entity"
	"server/internal/User/repository"
	"server/internal/domain/dto"
)

type UserRepo struct {
	DB *sql.DB
	mu sync.RWMutex
}

func NewUserRepo(db *sql.DB) repository.UserRepositoryI {
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
			return nil, entity.ErrInternalServerError
		}
		return user, nil
	case "Email":
		row := repo.DB.QueryRow("SELECT id, username, password, birthday, phone_number, email, icon FROM users WHERE email = $1", value)
		err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Birthday, &user.PhoneNumber, &user.Email, &user.Icon)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, entity.ErrInternalServerError
		}
		return user, nil
	case "PhoneNumber":
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
		return nil, entity.ErrInternalServerError
	}
	return user, nil
}

func (repo *UserRepo) CreateUser(in *dto.DBCreateUser) (uint, error) {

	repo.mu.Lock()
	insertUser := `INSERT INTO users (username, password, birthday, phone_number, email, icon) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := repo.DB.Exec(insertUser, in.Username, in.Password, in.Birthday, in.PhoneNumber, in.Email, in.Icon)
	if err != nil {
		return 0, entity.ErrInternalServerError
	}
	repo.mu.Unlock()
	var ID uint
	row := repo.DB.QueryRow("SELECT ID FROM users WHERE username = $1", in.Username)
	err = row.Scan(&ID)
	if err != nil {
		return 0, entity.ErrInternalServerError
	}

	return ID, nil
}


// Ульяне надо разобраться куда это засунуть. И не забудь разедлить FindUSerBy на несколько. И заменить ошибки

// if len(username) < 3 {
// 	w.WriteHeader(http.StatusBadRequest)
// 	err = json.NewEncoder(w).Encode(&Error{Err: "username is too short"})
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 	}
// 	return
// }

// if len(username) > 30 {
// 	w.WriteHeader(http.StatusBadRequest)
// 	err = json.NewEncoder(w).Encode(&Error{Err: "username is too long"})
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 	}
// 	return
// }

// if len(password) < 3 {
// 	w.WriteHeader(http.StatusBadRequest)
// 	err = json.NewEncoder(w).Encode(&Error{Err: "password is too short"})
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 	}
// 	return
// }

// if len(password) > 20 {
// 	w.WriteHeader(http.StatusBadRequest)
// 	err = json.NewEncoder(w).Encode(&Error{Err: "password is too long"})
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 	}
// 	return
// }

// re := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
// if birthday != "" && !re.MatchString(birthday) {
// 	w.WriteHeader(http.StatusBadRequest)
// 	err = json.NewEncoder(w).Encode(&Error{Err: "incorrect birthday"})
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 	}
// 	return
// }

// re = regexp.MustCompile(`^[+]?[0-9]{3,25}$`)
// if phoneNumber != "" && !re.MatchString(phoneNumber) {
// 	w.WriteHeader(http.StatusBadRequest)
// 	err = json.NewEncoder(w).Encode(&Error{Err: "incorrect phone number"})
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 	}
// 	return
// }

// re = regexp.MustCompile(`\S*@\S*`)
// if !re.MatchString(email) {
// 	w.WriteHeader(http.StatusBadRequest)
// 	err = json.NewEncoder(w).Encode(&Error{Err: "incorrect email"})
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 	}
// 	return
// }