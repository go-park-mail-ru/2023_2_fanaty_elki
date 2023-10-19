package store

// import (
// 	"testing"

// 	"github.com/stretchr/testify/require"
// )

// func TestCheckGetRestaurantsSuccess(t *testing.T) {
// 	restaurant1 := Restaurant{
// 		ID:            1,
// 		Name:          "Burger King",
// 		Rating:        3.7,
// 		CommentsCount: 60,
// 		Icon:          "defpath",
// 		DeliveryTime:  35,
// 		DeliveryPrice: 600,
// 		Category:      "Fastfood",
// 	}

// 	restaurant2 := Restaurant{
// 		ID:            2,
// 		Name:          "MacBurger",
// 		Rating:        3.8,
// 		CommentsCount: 69,
// 		Icon:          "defpath",
// 		DeliveryTime:  35,
// 		DeliveryPrice: 600,
// 		Category:      "Fastfood",
// 	}

// 	res := RestaurantStore{
// 		restourants: []*Restaurant{&restaurant1, &restaurant2}}

// 	t.Run("Get Restaurnt Success", func(t *testing.T) {
// 		result, err := res.GetRestaurants()
// 		require.Equal(t, []*Restaurant{&restaurant1, &restaurant2}, result)
// 		require.NoError(t, err)
// 	})
// }

// func TestCheckGetUsersSuccess(t *testing.T) {
// 	user1 := User{
// 		ID:       1,
// 		Username: "john_doe",
// 		Password: "secure_password",
// 		Birthday: "1990-05-15",
// 		Email:    "john@example.com",
// 	}

// 	user2 := User{
// 		ID:       2,
// 		Username: "jane_smith",
// 		Password: "another_password",
// 		Birthday: "1985-08-22",
// 		Email:    "jane@example.com",
// 	}

// 	res := UserStore{
// 		users:  []*User{&user1, &user2},
// 		nextID: 3}

// 	t.Run("Get User Success", func(t *testing.T) {
// 		result, err := res.GetUsers()
// 		require.Equal(t, []*User{&user1, &user2}, result)
// 		require.NoError(t, err)
// 	})
// }

// func TestCheckFindUserBySuccess(t *testing.T) {
// 	user1 := User{
// 		ID:       1,
// 		Username: "john_doe",
// 		Password: "secure_password",
// 		Birthday: "1990-05-15",
// 		Email:    "john@example.com",
// 	}

// 	user2 := User{
// 		ID:       2,
// 		Username: "jane_smith",
// 		Password: "another_password",
// 		Birthday: "1985-08-22",
// 		Email:    "jane@example.com",
// 	}

// 	res := UserStore{
// 		users:  []*User{&user1, &user2},
// 		nextID: 3}

// 	t.Run("Get User Success", func(t *testing.T) {
// 		result := res.FindUserBy("username", "john_doe")
// 		require.Equal(t, &user1, result)
// 		result = res.FindUserBy("email", "jane@example.com")
// 		require.Equal(t, &user2, result)
// 	})
// }

// func TestCheckFindUserByFail(t *testing.T) {
// 	user1 := User{
// 		ID:       1,
// 		Username: "john_doe",
// 		Password: "secure_password",
// 		Birthday: "1990-05-15",
// 		Email:    "john@example.com",
// 	}

// 	user2 := User{
// 		ID:       2,
// 		Username: "jane_smith",
// 		Password: "another_password",
// 		Birthday: "1985-08-22",
// 		Email:    "jane@example.com",
// 	}

// 	res := UserStore{
// 		users:  []*User{&user1, &user2},
// 		nextID: 3}

// 	t.Run("Get User Fail", func(t *testing.T) {
// 		result := res.FindUserBy("username", "john")
// 		require.Equal(t, (*User)(nil), result)
// 		result = res.FindUserBy("email", "jane@exame.com")
// 		require.Equal(t, (*User)(nil), result)
// 		result = res.FindUserBy("phone_number", "+14567890")
// 		require.Equal(t, (*User)(nil), result)
// 	})
// }

// func TestCheckSignUpUserSuccess(t *testing.T) {
// 	user1 := User{
// 		ID:       1,
// 		Username: "john_doe",
// 		Password: "secure_password",
// 		Birthday: "1990-05-15",
// 		Email:    "john@example.com",
// 	}

// 	user2 := User{
// 		ID:       2,
// 		Username: "jane_smith",
// 		Password: "another_password",
// 		Birthday: "1985-08-22",
// 		Email:    "jane@example.com",
// 	}

// 	user3 := User{
// 		Username: "jane_smit",
// 		Password: "another_passwor",
// 		Birthday: "1985-08-23",
// 		Email:    "jan@example.com",
// 	}

// 	res := UserStore{
// 		users:  []*User{&user1, &user2},
// 		nextID: 2}

// 	t.Run("Get User Success", func(t *testing.T) {
// 		result := res.SignUpUser(&user3)
// 		require.Equal(t, uint(3), result)
// 	})
// }
