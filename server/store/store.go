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
	ID       uint   `json:"ID"`
	Username string `json:"Username"`
	Password string `json:"Password"`
	Birthday string `json:"Birthday"`
	Email    string `json:"Email"`
}

type RestaurantStore struct {
	restourants []*Restaurant
	mu          sync.RWMutex
}

type UserStore struct {
	users  []*User
	mu     sync.RWMutex
	nextID uint
}

var Users = []*User{}

var DB *sql.DB

func NewRestaurantStore() *RestaurantStore {
	return &RestaurantStore{
		mu: sync.RWMutex{},
		//restourants: Restaurants{},
	}
}

func NewUserStore() *UserStore {
	return &UserStore{
		mu:    sync.RWMutex{},
		users: Users,
	}
}

func (rs *RestaurantStore) GetRestaurants() ([]*Restaurant, error) {

	rs.mu.RLock()
	defer rs.mu.RUnlock()

	rows, err := DB.Query("SELECT * FROM restaurant")
	if err != nil {
		fmt.Println("error while connecting")
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
			fmt.Println("error while scanning")
		}
		Restaurants = append(Restaurants, restaurant)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println("error while scanning")
	}
	return Restaurants, nil
}

func (us *UserStore) FindUserBy(field string, value string) *User {
	us.mu.RLock()
	defer us.mu.RUnlock()
	user := &User{}
	switch field {
	case "username":
		row := DB.QueryRow("SELECT id, username, password, birthday, email FROM users WHERE username = $1", value)
		err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Birthday, &user.Email)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			fmt.Println("error while scanning", err)
		}
		return user
	case "email":
		row := DB.QueryRow("SELECT id, username, password, birthday, email FROM users WHERE email = $1", value)
		err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Birthday, &user.Email)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			fmt.Println("error while scanning", err)
		}
		return user
	}

	return nil
}

func (us *UserStore) GetUserById(id uint) *User {
	us.mu.RLock()
	defer us.mu.RUnlock()
	user := &User{}
	row := DB.QueryRow("SELECT id, username, password, birthday, email FROM users WHERE id = $1", id)
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Birthday, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		fmt.Println("error while scanning", err)
	}
	return user
}

func (us *UserStore) SignUpUser(in *User) uint {

	us.mu.Lock()
	insertUser := `INSERT INTO users (username, password, birthday, phone_number, email, icon) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := DB.Exec(insertUser, in.Username, in.Password, in.Birthday, "+7916521", in.Email, "deficon")
	if err != nil {
		fmt.Println("error while inserting", err)
	}
	us.mu.Unlock()
	var ID uint
	row := DB.QueryRow("SELECT ID FROM users WHERE username = $1", in.Username)
	err = row.Scan(&ID)
	if err != nil {
		fmt.Println("error while scanning", err)
	}

	return ID
}
