package main

import (
	"backend/server/store"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func executeRequest(req *http.Request) *httptest.ResponseRecorder {
// 	rr := httptest.NewRecorder()
// 	mux := http.NewServeMux()
// 	mux.HandleFunc(req)
// 	return rr
// }

// func checkResponseCode(t *testing.T, expected, actual int) {
// 	if expected != actual {
// 		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
// 	}
// }

var api = &Handler{
	restaurantstore: store.NewRestaurantStore(),
	userstore:       store.NewUserStore(),
	sessions:        make(map[string]uint, 10),
}

func TestGetRestaurantsList(t *testing.T) {
	t.Run("returns slice of restaurants", func(t *testing.T) {

		req := httptest.NewRequest("GET", "/restaurants", nil)
		w := httptest.NewRecorder()

		api.getRestaurantList(w, req)

		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)

		assert.Equal(t, 200, resp.StatusCode)
		assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
		assert.Contains(t, string(body), "restaurants")
	})
}

func TestSignUp(t *testing.T) {

	req := httptest.NewRequest("POST", "/restaurants", nil)

	w := httptest.NewRecorder()

	api.getRestaurantList(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	assert.Contains(t, string(body), "restaurants")

}
