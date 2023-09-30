package store

import (
	"errors"
	"log"
	"sync"
)

type Restaurant struct {
	ID            uint
	Name          string
	Rating        float32
	CommentsCount int
	Icon          string
	DeliveryTime  int
	DeliveryPrice int
	Category      string
}

type User struct {
	ID          uint
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
	users  []*User
	mu     sync.RWMutex
	nextID uint
}

var Restaurants = []*Restaurant{{ID: 1, Name: "Burger King", Rating: 3.7, CommentsCount: 60, Icon: "defpath", DeliveryTime: 35, DeliveryPrice: 600, Category: "Fastfood"},
	{ID: 2, Name: "MacBurger", Rating: 3.8, CommentsCount: 69, Icon: "defpath", DeliveryTime: 35, DeliveryPrice: 600, Category: "Fastfood"},
	{ID: 3, Name: "Vcusno i tochka", Rating: 0.0, CommentsCount: 90, Icon: "defpath", DeliveryTime: 35, DeliveryPrice: 600, Category: "Fastfood"}}

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

func (us *UserStore) FindUser(username string) (*User, error) {
	for _, u := range us.users {
		if u.Username == username {
			return u, nil
		}
	}
	return nil, errors.New("No such user")

}

func (us *UserStore) SignUpUser(in *User) (uint, error) {
	log.Println("Signup called")

	us.mu.Lock()
	us.nextID++
	in.ID = us.nextID
	log.Println("nextID", us.nextID)
	us.users = append(us.users, in)
	us.mu.Unlock()

	return in.ID, nil
}
