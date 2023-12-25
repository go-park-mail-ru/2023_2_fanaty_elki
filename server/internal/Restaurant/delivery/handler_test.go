package delivery

import (
	"errors"
	"fmt"

	//"io/ioutil"
	"net/http/httptest"
	"server/config"
	mockR "server/internal/Restaurant/usecase/mock_usecase"
	mw "server/internal/middleware"
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
	var logger *mw.ACLog
	apiPath := "/api/restaurants"
	mock := mockR.NewMockRestaurantUsecaseI(ctrl)
	handler := NewRestaurantHandler(mock, logger)

	rests := &dto.RestaurantWithCategoriesSlice{
		{ID: 1,
			Name:          "Burger King",
			Rating:        3.7,
			CommentsCount: 60,
			Categories:    []string{"Burger", "Breakfast"},
			Icon:          "img/burger_king.webp",
		},
		{ID: 2,
			Name:          "MacBurger",
			Rating:        3.8,
			CommentsCount: 69,
			Categories:    []string{"Burger", "Breakfast"},
			Icon:          "img/mac_burger.webp",
		},
	}

	mock.EXPECT().GetRestaurants().Return(rests, nil)

	req := httptest.NewRequest("GET", apiPath, nil)
	w := httptest.NewRecorder()

	handler.GetRestaurantList(w, req)

	resp := w.Result()
	//body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return
	// }

	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))

}

func TestGetRestaurantsListFail(t *testing.T) {
	baseLogger, err := config.Cfg.Build()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer baseLogger.Sync()

	errorLogger, err := config.ErrorCfg.Build()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer errorLogger.Sync()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiPath := "/api/restaurants"
	mock := mockR.NewMockRestaurantUsecaseI(ctrl)
	hitstats := &entity.HitStats{}
	logger := mw.NewACLog(baseLogger.Sugar(), errorLogger.Sugar(), *hitstats)
	handler := NewRestaurantHandler(mock, logger)

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
	apiPath := "/api/restaurants/1"
	mock := mockR.NewMockRestaurantUsecaseI(ctrl)
	var logger *mw.ACLog
	handler := NewRestaurantHandler(mock, logger)

	restaurant := &dto.RestaurantWithCategories{
		ID:            1,
		Name:          "Burger King",
		Rating:        3.7,
		CommentsCount: 60,
		Categories:    []string{"Burger", "Breakfast"},
		Icon:          "img/burger_king.webp",
	}

	var elemID = 1

	mock.EXPECT().GetRestaurantByID(uint(elemID)).Return(restaurant, nil)

	req := httptest.NewRequest("GET", apiPath, nil)
	w := httptest.NewRecorder()

	vars := map[string]string{
		"id": "1",
	}

	req = mux.SetURLVars(req, vars)

	handler.GetRestaurantByID(w, req)

	resp := w.Result()
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return
	// }

	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	//require.Contains(t, string(body), "Body")

}

func TestGetRestaurantByIdFail(t *testing.T) {
	baseLogger, err := config.Cfg.Build()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer baseLogger.Sync()

	errorLogger, err := config.ErrorCfg.Build()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer errorLogger.Sync()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiPath := "/api/restaurants/fdfd"
	mock := mockR.NewMockRestaurantUsecaseI(ctrl)
	hitstats := &entity.HitStats{}
	logger := mw.NewACLog(baseLogger.Sugar(), errorLogger.Sugar(), *hitstats)
	handler := NewRestaurantHandler(mock, logger)

	req := httptest.NewRequest("GET", apiPath, nil)
	w := httptest.NewRecorder()

	handler.GetRestaurantByID(w, req)

	resp := w.Result()

	require.Equal(t, 400, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))

	req = httptest.NewRequest("GET", apiPath, nil)
	w = httptest.NewRecorder()

	vars := map[string]string{
		"id": "uue",
	}

	req = mux.SetURLVars(req, vars)

	handler.GetRestaurantByID(w, req)

	resp = w.Result()

	require.Equal(t, 400, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))

	testErr := errors.New("test")
	var elemID = 1

	mock.EXPECT().GetRestaurantByID(uint(elemID)).Return(nil, testErr)

	req = httptest.NewRequest("GET", apiPath, nil)
	w = httptest.NewRecorder()

	vars = map[string]string{
		"id": "1",
	}

	req = mux.SetURLVars(req, vars)

	handler.GetRestaurantByID(w, req)

	resp = w.Result()

	require.Equal(t, 500, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))

}

func TestGetRestaurantProductsSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiPath := "/api/restaurants"
	mock := mockR.NewMockRestaurantUsecaseI(ctrl)
	var logger *mw.ACLog
	handler := NewRestaurantHandler(mock, logger)

	MenuTypesWithProducts := &dto.MenuTypeWithProductsSlice{
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
	}

	var elemID = 1

	mock.EXPECT().GetRestaurantProducts(uint(elemID)).Return(MenuTypesWithProducts, nil)

	req := httptest.NewRequest("GET", apiPath, nil)
	w := httptest.NewRecorder()

	vars := map[string]string{
		"id": "1",
	}

	req = mux.SetURLVars(req, vars)

	handler.GetRestaurantProducts(w, req)

	resp := w.Result()
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return
	// }

	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	//require.Contains(t, string(body), "Body")

}

func TestGetRestaurantProductsFail(t *testing.T) {
	baseLogger, err := config.Cfg.Build()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer baseLogger.Sync()

	errorLogger, err := config.ErrorCfg.Build()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer errorLogger.Sync()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiPath := "/api/restaurants/fdfd/products"
	mock := mockR.NewMockRestaurantUsecaseI(ctrl)
	hitstats := &entity.HitStats{}
	logger := mw.NewACLog(baseLogger.Sugar(), errorLogger.Sugar(), *hitstats)
	handler := NewRestaurantHandler(mock, logger)

	req := httptest.NewRequest("GET", apiPath, nil)
	w := httptest.NewRecorder()

	handler.GetRestaurantProducts(w, req)

	resp := w.Result()

	require.Equal(t, 400, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))

	req = httptest.NewRequest("GET", apiPath, nil)
	w = httptest.NewRecorder()

	vars := map[string]string{
		"id": "uue",
	}

	req = mux.SetURLVars(req, vars)

	handler.GetRestaurantProducts(w, req)

	resp = w.Result()

	require.Equal(t, 400, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))

	testErr := errors.New("test")
	var elemID = 1

	mock.EXPECT().GetRestaurantProducts(uint(elemID)).Return(nil, testErr)

	req = httptest.NewRequest("GET", apiPath, nil)
	w = httptest.NewRecorder()

	vars = map[string]string{
		"id": "1",
	}

	req = mux.SetURLVars(req, vars)

	handler.GetRestaurantProducts(w, req)

	resp = w.Result()

	require.Equal(t, 500, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))

	mock.EXPECT().GetRestaurantProducts(uint(elemID)).Return(nil, entity.ErrNotFound)

	req = httptest.NewRequest("GET", apiPath, nil)
	w = httptest.NewRecorder()

	vars = map[string]string{
		"id": "1",
	}

	req = mux.SetURLVars(req, vars)

	handler.GetRestaurantProducts(w, req)

	resp = w.Result()

	require.Equal(t, 404, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))

}

func TestGetRestaurantListByCategorySuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiPath := "/api/restaurants/Бургеры"
	mock := mockR.NewMockRestaurantUsecaseI(ctrl)
	var logger *mw.ACLog
	handler := NewRestaurantHandler(mock, logger)

	rests := &dto.RestaurantWithCategoriesSlice{
		{ID: 1,
			Name:          "Burger King",
			Rating:        3.7,
			CommentsCount: 60,
			Categories:    []string{"Burger", "Breakfast"},
			Icon:          "img/burger_king.webp",
		},
		{ID: 2,
			Name:          "MacBurger",
			Rating:        3.8,
			CommentsCount: 69,
			Categories:    []string{"Burger", "Breakfast"},
			Icon:          "img/mac_burger.webp",
		},
	}

	mock.EXPECT().GetRestaurantsByCategory("Бургеры").Return(rests, nil)

	vars := map[string]string{
		"category": "Бургеры",
	}

	req := httptest.NewRequest("GET", apiPath, nil)

	req = mux.SetURLVars(req, vars)

	w := httptest.NewRecorder()

	handler.GetRestaurantListByCategory(w, req)

	resp := w.Result()
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return
	// }

	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	//require.Contains(t, string(body), "Body")

}
