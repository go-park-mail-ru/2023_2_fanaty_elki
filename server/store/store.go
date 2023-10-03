package store

import (
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

var Restaurants = []*Restaurant{
	{
		ID:            1,
		Name:          "Burger King",
		Rating:        3.7,
		CommentsCount: 60,
		Icon:          "img/burger_king.jpg",
		DeliveryTime:  15,
		DeliveryPrice: 600,
		Category:      "Fastfood",
	},
	{
		ID:     2,
		Name:   "MacBurger",
		Rating: 3.8, CommentsCount: 69,
		Icon:          "img/mac_burger.jpg",
		DeliveryTime:  35,
		DeliveryPrice: 500,
		Category:      "Fastfood",
	},
	{
		ID: 3, Name: "Вкусно и точка",
		Rating: 3.2, CommentsCount: 90,
		Icon:          "img/tasty_and..jpg",
		DeliveryTime:  20,
		DeliveryPrice: 100,
		Category:      "Fastfood",
	},
	{
		ID: 3, Name: "KFC",
		Rating: 4.0, CommentsCount: 90,
		Icon:          "img/kfc.jpg",
		DeliveryTime:  40,
		DeliveryPrice: 600,
		Category:      "Fastfood",
	},
	{
		ID: 3, Name: "Шоколадница",
		Rating: 4.5, CommentsCount: 90,
		Icon:          "img/chocolate.jpeg",
		DeliveryTime:  30,
		DeliveryPrice: 400,
		Category:      "Fastfood",
	},
	{
		ID: 3, Name: "Корчма Тарас Бульба",
		Rating: 5.0, CommentsCount: 90,
		Icon:          "img/bulba.jpg",
		DeliveryTime:  30,
		DeliveryPrice: 800,
		Category:      "Fastfood",
	},
	{
		ID: 3, Name: "Subway",
		Rating: 3.0, CommentsCount: 90,
		Icon:          "img/subway.jpeg",
		DeliveryTime:  40,
		DeliveryPrice: 600,
		Category:      "Fastfood",
	},
	{
		ID: 3, Name: "Sushiwok",
		Rating: 4.5, CommentsCount: 90,
		Icon:          "img/sushi_wok.png",
		DeliveryTime:  10,
		DeliveryPrice: 300,
		Category:      "Fastfood",
	},
}

var Users = []*User{}

func NewRestaurantStore() *RestaurantStore {
	return &RestaurantStore{
		mu:          sync.RWMutex{},
		restourants: Restaurants,
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

	return rs.restourants, nil
}

func (us *UserStore) GetUsers() ([]*User, error) {
	us.mu.RLock()
	defer us.mu.RUnlock()

	return us.users, nil
}

func (us *UserStore) FindUserBy(field string, value string) *User {
	for _, u := range us.users {
		switch field {
		case "username":
			if u.Username == value {
				return u
			}
		case "email":
			if u.Email == value {
				return u
			}
		}
	}
	return nil
}

func (us *UserStore) GetUserById(id uint) *User {
	return us.users[id]
}

func (us *UserStore) SignUpUser(in *User) (uint, error) {

	us.mu.Lock()
	us.nextID++
	in.ID = us.nextID
	us.users = append(us.users, in)
	us.mu.Unlock()

	return in.ID, nil
}
