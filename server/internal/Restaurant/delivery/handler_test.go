package delivery

import (
	"errors"
	"io/ioutil"
	"net/http/httptest"
	mockR "server/internal/Restaurant/usecase/mock_usecase"
	"testing"

	"server/internal/domain/dto"
	"server/internal/domain/entity"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

func TestGetRestaurantsListSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiPath := "/api/restaurants"
	mock := mockR.NewMockUsecaseI(ctrl)
	handler := NewRestaurantHandler(mock)

	rests := []*entity.Restaurant{
		{ID: 1,
			Name:          "Burger King",
			Rating:        3.7,
			CommentsCount: 60,
			Category:      "Fastfood",
			Icon:          "img/burger_king.jpg",
		},
		{ID: 2,
			Name:          "MacBurger",
			Rating:        3.8,
			CommentsCount: 69,
			Category:      "Fastfood",
			Icon:          "img/mac_burger.jpg",
		},
	}

	mock.EXPECT().GetRestaurants().Return(rests, nil)

	req := httptest.NewRequest("GET", apiPath, nil)
	w := httptest.NewRecorder()

	handler.GetRestaurantList(w, req)

	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	require.Contains(t, string(body), "restaurants")

}

func TestGetRestaurantsListFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiPath := "/api/restaurants"
	mock := mockR.NewMockUsecaseI(ctrl)
	handler := NewRestaurantHandler(mock)

	testErr := errors.New("test")

	mock.EXPECT().GetRestaurants().Return(nil, testErr)

	req := httptest.NewRequest("GET", apiPath, nil)
	w := httptest.NewRecorder()

	handler.GetRestaurantList(w, req)

	resp := w.Result()

	require.Equal(t, 500, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))
}

func TestGetRestaurantByIdSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiPath := "/api/restaurants"
	mock := mockR.NewMockUsecaseI(ctrl)
	handler := NewRestaurantHandler(mock)

	rest := &dto.RestaurantWithProducts{
		Restaurant: &entity.Restaurant{
			ID:            1,
			Name:          "Burger King",
			Rating:        3.7,
			CommentsCount: 60,
			Category:      "Fastfood",
			Icon:          "img/burger_king.jpg",
		},
		MenuTypesWithProducts: []*dto.MenuTypeWithProducts{
			{
				MenuType: &entity.MenuType{
					ID:           1,
					Name:         "Популярное",
					RestaurantID: 1,
				},
				Products: []*entity.Product{
					{
						ID:          1,
						Name:        "Burger",
						Price:       120.0,
						CookingTime: 23,
						Portion:     "160 г",
						Description: "Nice burger",
						Icon:        "deficon",
					},
				},
			},
		},
	}

	var elemID = 1

	mock.EXPECT().GetRestaurantById(uint(elemID)).Return(rest, nil)

	req := httptest.NewRequest("GET", apiPath, nil)
	w := httptest.NewRecorder()

	vars := map[string]string{
		"id": "1",
	}

	req = mux.SetURLVars(req, vars)

	handler.GetRestaurantById(w, req)

	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	require.Contains(t, string(body), "RestaurantWithProducts")

}

func TestGetRestaurantByIdFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiPath := "/api/restaurants/fdfd"
	mock := mockR.NewMockUsecaseI(ctrl)
	handler := NewRestaurantHandler(mock)

	req := httptest.NewRequest("GET", apiPath, nil)
	w := httptest.NewRecorder()

	handler.GetRestaurantById(w, req)

	resp := w.Result()

	require.Equal(t, 400, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))

	req = httptest.NewRequest("GET", apiPath, nil)
	w = httptest.NewRecorder()

	vars := map[string]string{
		"id": "uue",
	}

	req = mux.SetURLVars(req, vars)

	handler.GetRestaurantById(w, req)

	resp = w.Result()

	require.Equal(t, 400, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))

	testErr := errors.New("test")
	var elemID = 1

	mock.EXPECT().GetRestaurantById(uint(elemID)).Return(nil, testErr)

	req = httptest.NewRequest("GET", apiPath, nil)
	w = httptest.NewRecorder()

	vars = map[string]string{
		"id": "1",
	}

	req = mux.SetURLVars(req, vars)

	handler.GetRestaurantById(w, req)

	resp = w.Result()

	require.Equal(t, 500, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))

}
