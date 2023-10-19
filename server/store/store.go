package store

import (
	"database/sql"
	"fmt"
	"strings"
	"sync"
)

type Restaurant struct {
	ID            uint    `json:"ID"`
	Name          string  `json:"Name"`
	Rating        float32 `json:"Rating"`
	CommentsCount int     `json:"CommentsCount"`
	Icon          string  `json:"Icon"`
	DeliveryTime  int     `json:"DeliveryTime"`
	DeliveryPrice int     `json:"DeliveryPrice"`
	Category      string  `json:"Category"`
}

type User struct {
	ID          uint           `json:"ID"`
	Username    string         `json:"Username"`
	Password    string         `json:"Password"`
	Birthday    sql.NullString `json:"Birthday"`
	PhoneNumber string         `json:"PhoneNumber"`
	Email       string         `json:"Email"`
	Icon        sql.NullString `json:"Icon"`
}

type RestaurantRepo struct {
	DB *sql.DB
	mu sync.RWMutex
}

type UserRepo struct {
	DB *sql.DB
	mu sync.RWMutex
}

var Users = []*User{}

//var DB *sql.DB

func GetPostgres() (*sql.DB, error) {
	const (
		host     = "localhost"
		port     = 5432
		user     = "uliana"
		password = "uliana"
		dbname   = "prinesy-poday"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected!")
	return db, nil
}

func NewRestaurantRepo(db *sql.DB) *RestaurantRepo {
	return &RestaurantRepo{
		mu: sync.RWMutex{},
		DB: db,
	}
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		mu: sync.RWMutex{},
		DB: db,
	}
}

func (repo *RestaurantRepo) GetRestaurants() ([]*Restaurant, error) {

	repo.mu.RLock()
	defer repo.mu.RUnlock()

	rows, err := repo.DB.Query("SELECT id, name, rating, comments_count, category, delivery_price, delivery_time, icon FROM restaurant")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var Restaurants = []*Restaurant{}
	for rows.Next() {
		restaurant := &Restaurant{}
		err = rows.Scan(
			&restaurant.ID,
			&restaurant.Name,
			&restaurant.Rating,
			&restaurant.CommentsCount,
			&restaurant.Category,
			&restaurant.DeliveryPrice,
			&restaurant.DeliveryTime,
			&restaurant.Icon,
		)
		restaurant.Name = strings.TrimSpace(restaurant.Name)
		restaurant.Icon = strings.TrimSpace(restaurant.Icon)
		restaurant.Category = strings.TrimSpace(restaurant.Category)
		if err != nil {
			return nil, err
		}
		Restaurants = append(Restaurants, restaurant)
	}
	err = rows.Err()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return Restaurants, nil
}

func (repo *UserRepo) FindUserBy(field string, value string) (*User, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()
	user := &User{}
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

func (repo *UserRepo) GetUserById(id uint) (*User, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()
	user := &User{}
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

func (repo *UserRepo) SignUpUser(in *User) (uint, error) {

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
