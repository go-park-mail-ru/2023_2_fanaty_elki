package store

import (
	"sync"
)

type Restaurant struct {
	ID            int
	Name          string
	Rating        float32
	CommentsCount int
	Icon          string
	DeliveryTime  int
	DeliveryPrice int
	Category      string
}

type User struct {
	ID          int
	Username    string
	Password    string
	Birthday    string
	PhoneNumber string
	Email       string
	Icon        string
}

type RestaurantStore struct {
	restourants []*Restaurant
	mu          sync.RWMutex
}

type UserStore struct {
	user []*User
	mu   sync.RWMutex
}

var Restaurants = []*Restaurant{{ID: 1, Name: "Burger King", Rating: 3.7, CommentsCount: 60, Icon: "defpath", DeliveryTime: 35, DeliveryPrice: 600, Category: "Fastfood"},
	{ID: 2, Name: "MacBurger", Rating: 3.8, CommentsCount: 69, Icon: "defpath", DeliveryTime: 35, DeliveryPrice: 600, Category: "Fastfood"},
	{ID: 3, Name: "Vcusno i tochka", Rating: 0.0, CommentsCount: 90, Icon: "defpath", DeliveryTime: 35, DeliveryPrice: 600, Category: "Fastfood"}}

func NewRestaurantStore() *RestaurantStore {
	return &RestaurantStore{
		mu:          sync.RWMutex{},
		restourants: Restaurants,
	}
}

func (rs *RestaurantStore) GetRestaurants() ([]*Restaurant, error) {
	//log.Println("GetBooks called")

	rs.mu.RLock()
	defer rs.mu.RUnlock()

	return rs.restourants, nil
}
