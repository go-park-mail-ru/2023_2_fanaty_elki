package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"io/ioutil"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/stretchr/testify/require"
// )

// var api = &Handler{
// 	restaurantstore: store.NewRestaurantRepo(),
// 	userstore:       store.NewUserRepo(),
// 	sessions:        make(map[string]uint, 10),
// }

// func TestGetRestaurantsList(t *testing.T) {
// 	t.Run("returns slice of restaurants", func(t *testing.T) {

// 		req := httptest.NewRequest("GET", "/restaurants", nil)
// 		w := httptest.NewRecorder()

// 		api.GetRestaurantList(w, req)

// 		resp := w.Result()
// 		body, err := ioutil.ReadAll(resp.Body)
// 		if err != nil {
// 			return
// 		}

// 		require.Equal(t, 200, resp.StatusCode)
// 		require.Equal(t, "application/json", resp.Header.Get("Content-Type"))
// 		require.Contains(t, string(body), "restaurants")
// 	})
// }

// func TestSignUp(t *testing.T) {

// 	var user1 = map[string]interface{}{
// 		"username":     "ania",
// 		"password":     "abc",
// 		"birthday":     "12-12-2007",
// 		"phone_number": "89165342399",
// 		"email":        "ana@mail.ru",
// 	}

// 	var user2 = map[string]interface{}{
// 		"username":     "ani",
// 		"password":     "abc",
// 		"phone_number": "89165342397",
// 		"email":        "ani@mail.ru",
// 	}

// 	var user3 = map[string]interface{}{
// 		"username":     "ani",
// 		"password":     "abc",
// 		"phone_number": "89165342397",
// 		"email":        "ani@mail.ru",
// 	}

// 	var user5 = map[string]interface{}{
// 		"username":     "anita",
// 		"password":     "abc",
// 		"phone_number": "89165342390",
// 		"email":        "ani@mail.ru",
// 	}

// 	var user6 = map[string]interface{}{
// 		"username":     "an",
// 		"password":     "abc",
// 		"phone_number": "89165342390",
// 		"email":        "ani@mail.ru",
// 	}

// 	var user7 = map[string]interface{}{
// 		"username":     "annnnnnnnnnnnnnnnnnnnnnnnnnnnnnnn",
// 		"password":     "abc",
// 		"phone_number": "89165342390",
// 		"email":        "ani@mail.ru",
// 	}

// 	var user8 = map[string]interface{}{
// 		"username":     "maria",
// 		"password":     "ab",
// 		"phone_number": "89165342390",
// 		"email":        "ani@mail.ru",
// 	}

// 	var user9 = map[string]interface{}{
// 		"username":     "lena",
// 		"password":     "abmmmmmmmmmmmmmmmmmmmmmmmmmm",
// 		"phone_number": "89165342390",
// 		"email":        "ani@mail.ru",
// 	}

// 	var user11 = map[string]interface{}{
// 		"username":     "lena",
// 		"password":     "abm",
// 		"phone_number": "89165342390",
// 		"email":        "lenamail.ru",
// 	}

// 	t.Run("returns ok and id on correct data", func(t *testing.T) {
// 		body, err := json.Marshal(user1)
// 		if err != nil {
// 			return
// 		}
// 		req := httptest.NewRequest("POST", "/users", bytes.NewReader(body))

// 		w := httptest.NewRecorder()

// 		api.SignUp(w, req)

// 		resp := w.Result()
// 		body, err = ioutil.ReadAll(resp.Body)
// 		if err != nil {
// 			return
// 		}

// 		require.Equal(t, 200, resp.StatusCode)
// 		require.Equal(t, "application/json", resp.Header.Get("Content-Type"))
// 		require.Contains(t, string(body), "id")
// 	})

// 	t.Run("returns ok and id on correct data without birthday", func(t *testing.T) {
// 		body, err := json.Marshal(user2)
// 		if err != nil {
// 			return
// 		}
// 		req := httptest.NewRequest("POST", "/users", bytes.NewReader(body))

// 		w := httptest.NewRecorder()

// 		api.SignUp(w, req)

// 		resp := w.Result()
// 		body, err = ioutil.ReadAll(resp.Body)
// 		if err != nil {
// 			return
// 		}

// 		require.Equal(t, 200, resp.StatusCode)
// 		require.Equal(t, "application/json", resp.Header.Get("Content-Type"))
// 		require.Contains(t, string(body), "id")
// 	})

// 	t.Run("returns 400 error when user already exists", func(t *testing.T) {
// 		body, err := json.Marshal(user3)
// 		if err != nil {
// 			return
// 		}
// 		req := httptest.NewRequest("POST", "/users", bytes.NewReader(body))

// 		w := httptest.NewRecorder()

// 		api.SignUp(w, req)

// 		resp := w.Result()
// 		body, err = ioutil.ReadAll(resp.Body)
// 		if err != nil {
// 			return
// 		}

// 		require.Equal(t, 400, resp.StatusCode)
// 		require.Equal(t, "application/json", resp.Header.Get("Content-Type"))
// 		require.Contains(t, string(body), "username already exists")
// 	})

// 	t.Run("returns 400 error when email already exists", func(t *testing.T) {
// 		body, err := json.Marshal(user5)
// 		if err != nil {
// 			return
// 		}
// 		req := httptest.NewRequest("POST", "/users", bytes.NewReader(body))

// 		w := httptest.NewRecorder()

// 		api.SignUp(w, req)

// 		resp := w.Result()
// 		body, err = ioutil.ReadAll(resp.Body)
// 		if err != nil {
// 			return
// 		}

// 		require.Equal(t, 400, resp.StatusCode)
// 		require.Equal(t, "application/json", resp.Header.Get("Content-Type"))
// 		require.Contains(t, string(body), "email already exists")
// 	})

// 	t.Run("returns 400 error when username is too short", func(t *testing.T) {
// 		body, err := json.Marshal(user6)
// 		if err != nil {
// 			return
// 		}
// 		req := httptest.NewRequest("POST", "/users", bytes.NewReader(body))

// 		w := httptest.NewRecorder()

// 		api.SignUp(w, req)

// 		resp := w.Result()
// 		body, err = ioutil.ReadAll(resp.Body)
// 		if err != nil {
// 			return
// 		}

// 		require.Equal(t, 400, resp.StatusCode)
// 		require.Equal(t, "application/json", resp.Header.Get("Content-Type"))
// 		require.Contains(t, string(body), "username is too short")
// 	})

// 	t.Run("returns 400 error when username is too long", func(t *testing.T) {
// 		body, err := json.Marshal(user7)
// 		if err != nil {
// 			return
// 		}
// 		req := httptest.NewRequest("POST", "/users", bytes.NewReader(body))

// 		w := httptest.NewRecorder()

// 		api.SignUp(w, req)

// 		resp := w.Result()
// 		body, err = ioutil.ReadAll(resp.Body)
// 		if err != nil {
// 			return
// 		}

// 		require.Equal(t, 400, resp.StatusCode)
// 		require.Equal(t, "application/json", resp.Header.Get("Content-Type"))
// 		require.Contains(t, string(body), "username is too long")
// 	})

// 	t.Run("returns 400 error when password is too short", func(t *testing.T) {
// 		body, err := json.Marshal(user8)
// 		if err != nil {
// 			return
// 		}
// 		req := httptest.NewRequest("POST", "/users", bytes.NewReader(body))

// 		w := httptest.NewRecorder()

// 		api.SignUp(w, req)

// 		resp := w.Result()
// 		body, err = ioutil.ReadAll(resp.Body)
// 		if err != nil {
// 			return
// 		}

// 		require.Equal(t, 400, resp.StatusCode)
// 		require.Equal(t, "application/json", resp.Header.Get("Content-Type"))
// 		require.Contains(t, string(body), "password is too short")
// 	})

// 	t.Run("returns 400 error when password is too long", func(t *testing.T) {
// 		body, err := json.Marshal(user9)
// 		if err != nil {
// 			return
// 		}
// 		req := httptest.NewRequest("POST", "/users", bytes.NewReader(body))

// 		w := httptest.NewRecorder()

// 		api.SignUp(w, req)

// 		resp := w.Result()
// 		body, err = ioutil.ReadAll(resp.Body)
// 		if err != nil {
// 			return
// 		}

// 		require.Equal(t, 400, resp.StatusCode)
// 		require.Equal(t, "application/json", resp.Header.Get("Content-Type"))
// 		require.Contains(t, string(body), "password is too long")
// 	})

// 	t.Run("returns 400 error when email is incorrect", func(t *testing.T) {
// 		body, err := json.Marshal(user11)
// 		if err != nil {
// 			return
// 		}
// 		req := httptest.NewRequest("POST", "/users", bytes.NewReader(body))

// 		w := httptest.NewRecorder()

// 		api.SignUp(w, req)

// 		resp := w.Result()
// 		body, err = ioutil.ReadAll(resp.Body)
// 		if err != nil {
// 			return
// 		}

// 		require.Equal(t, 400, resp.StatusCode)
// 		require.Equal(t, "application/json", resp.Header.Get("Content-Type"))
// 		require.Contains(t, string(body), "incorrect email")
// 	})
// }

// var Cookie []*http.Cookie
// var Cookie1 []*http.Cookie

// func TestLogin(t *testing.T) {
// 	var user12 = map[string]interface{}{
// 		"username": "ania",
// 		"password": "abc",
// 	}

// 	var user13 = map[string]interface{}{
// 		"username": "alina",
// 		"password": "abc",
// 	}

// 	var user14 = map[string]interface{}{
// 		"username": "ania",
// 		"password": "abbc",
// 	}

// 	var user15 = map[string]interface{}{
// 		"username": "ani",
// 		"password": "abc",
// 	}

// 	t.Run("returns ok when user is correct", func(t *testing.T) {
// 		body, err := json.Marshal(user12)
// 		if err != nil {
// 			return
// 		}
// 		req := httptest.NewRequest("POST", "/login", bytes.NewReader(body))

// 		w := httptest.NewRecorder()

// 		api.Login(w, req)

// 		resp := w.Result()
// 		body, err = ioutil.ReadAll(resp.Body)
// 		if err != nil {
// 			return
// 		}

// 		Cookie = resp.Cookies()

// 		require.Equal(t, 200, resp.StatusCode)
// 		require.Equal(t, "application/json", resp.Header.Get("Content-Type"))
// 		require.Contains(t, string(body), "cookie")
// 	})

// 	t.Run("returns ok when user is correct for cookie", func(t *testing.T) {
// 		body, err := json.Marshal(user15)
// 		if err != nil {
// 			return
// 		}
// 		req := httptest.NewRequest("POST", "/login", bytes.NewReader(body))

// 		w := httptest.NewRecorder()

// 		api.Login(w, req)

// 		resp := w.Result()
// 		body, err = ioutil.ReadAll(resp.Body)
// 		if err != nil {
// 			return
// 		}

// 		Cookie1 = resp.Cookies()

// 		require.Equal(t, 200, resp.StatusCode)
// 		require.Equal(t, "application/json", resp.Header.Get("Content-Type"))
// 		require.Contains(t, string(body), "cookie")
// 	})

// 	t.Run("returns 404 when user not found", func(t *testing.T) {
// 		body, err := json.Marshal(user13)
// 		if err != nil {
// 			return
// 		}
// 		req := httptest.NewRequest("POST", "/login", bytes.NewReader(body))

// 		w := httptest.NewRecorder()

// 		api.Login(w, req)

// 		resp := w.Result()
// 		body, err = ioutil.ReadAll(resp.Body)
// 		if err != nil {
// 			return
// 		}

// 		require.Equal(t, 404, resp.StatusCode)
// 		require.Equal(t, "application/json", resp.Header.Get("Content-Type"))
// 		require.Contains(t, string(body), "user not found")
// 	})

// 	t.Run("returns 400 when incorrect password ", func(t *testing.T) {
// 		body, err := json.Marshal(user14)
// 		if err != nil {
// 			return
// 		}
// 		req := httptest.NewRequest("POST", "/login", bytes.NewReader(body))

// 		w := httptest.NewRecorder()

// 		api.Login(w, req)

// 		resp := w.Result()
// 		body, err = ioutil.ReadAll(resp.Body)
// 		if err != nil {
// 			return
// 		}

// 		require.Equal(t, 400, resp.StatusCode)
// 		require.Equal(t, "application/json", resp.Header.Get("Content-Type"))
// 		require.Contains(t, string(body), "incorrect password")
// 	})
// }

// func TestLogout(t *testing.T) {
// 	t.Run("returns 400 when incorrect password ", func(t *testing.T) {

// 		req := httptest.NewRequest("GET", "/logout", nil)

// 		w := httptest.NewRecorder()

// 		api.Logout(w, req)

// 		resp := w.Result()

// 		require.Equal(t, 401, resp.StatusCode)
// 		require.Equal(t, "application/json", resp.Header.Get("Content-Type"))

// 	})

// 	t.Run("returns ok when user is correct and logout", func(t *testing.T) {
// 		req := httptest.NewRequest("GET", "/logout", nil)
// 		req.AddCookie(Cookie[0])

// 		w := httptest.NewRecorder()

// 		api.Logout(w, req)

// 		resp := w.Result()

// 		require.Equal(t, 200, resp.StatusCode)
// 		require.Equal(t, "application/json", resp.Header.Get("Content-Type"))
// 	})
// }

// func TestAuth(t *testing.T) {

// 	t.Run("returns 401 when not autorized", func(t *testing.T) {

// 		req := httptest.NewRequest("GET", "/auth", nil)

// 		w := httptest.NewRecorder()

// 		api.Auth(w, req)

// 		resp := w.Result()
// 		body, err := ioutil.ReadAll(resp.Body)
// 		if err != nil {
// 			return
// 		}

// 		require.Equal(t, 401, resp.StatusCode)
// 		require.Equal(t, "application/json", resp.Header.Get("Content-Type"))
// 		require.Contains(t, string(body), "unauthorized")

// 	})
// 	t.Run("returns id when autorized ", func(t *testing.T) {

// 		req := httptest.NewRequest("GET", "/auth", nil)
// 		req.AddCookie(Cookie1[0])

// 		w := httptest.NewRecorder()

// 		api.Auth(w, req)

// 		resp := w.Result()
// 		body, err := ioutil.ReadAll(resp.Body)
// 		if err != nil {
// 			return
// 		}

// 		require.Equal(t, 200, resp.StatusCode)
// 		require.Equal(t, "application/json", resp.Header.Get("Content-Type"))
// 		require.Contains(t, string(body), "username")
// 	})
// }
