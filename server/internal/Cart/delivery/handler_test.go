package delivery

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"server/config"
	mockC "server/internal/Cart/usecase/mock_usecase"
	mw "server/internal/middleware"
	"testing"

	"server/internal/domain/dto"
	"server/internal/domain/entity"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

func TestGetCartSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiPath := "/api/cart"
	mock := mockC.NewMockCartUsecaseI(ctrl)
	var logger *mw.ACLog
	handler := NewCartHandler(mock, logger)

	cookie := &entity.Cookie{
		UserID:       1,
		SessionToken: "HJJvgsvd",
	}

	rest := &entity.Restaurant{
		ID:            1,
		Name:          "Burger King",
		Rating:        3.7,
		CommentsCount: 60,
		Icon:          "img/burger_king.webp",
	}

	res := &dto.CartWithRestaurant{
		Restaurant: rest,
		Products: []*dto.CartProduct{
			{
				Product: &entity.Product{
					ID:          1,
					Name:        "Burger",
					Price:       120.0,
					CookingTime: 23,
					Portion:     "160 г",
					Description: "Nice burger",
					Icon:        "deficon",
				},
				ItemCount: 6,
			},
			{
				Product: &entity.Product{
					ID:          2,
					Name:        "Burger",
					Price:       120.0,
					CookingTime: 23,
					Portion:     "160 г",
					Description: "Nice burger",
					Icon:        "deficon",
				},
				ItemCount: 6,
			},
		},
	}

	mock.EXPECT().GetUserCart(cookie.SessionToken).Return(res, nil)

	req := httptest.NewRequest("GET", apiPath, nil)
	req.AddCookie(&http.Cookie{Name: "session_id", Value: cookie.SessionToken})
	w := httptest.NewRecorder()

	handler.GetCart(w, req)

	resp := w.Result()
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return
	// }

	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	//require.Contains(t, string(body), "Body")

}

func TestGetCartFail(t *testing.T) {
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
	hitstats := &entity.HitStats{}
	logger := mw.NewACLog(baseLogger.Sugar(), errorLogger.Sugar(), *hitstats)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiPath := "/api/cart"
	mock := mockC.NewMockCartUsecaseI(ctrl)
	handler := NewCartHandler(mock, logger)

	testErr := errors.New("test")

	cookie := &entity.Cookie{
		UserID:       1,
		SessionToken: "HJJvgsvd",
	}

	mock.EXPECT().GetUserCart(cookie.SessionToken).Return(nil, testErr)

	req := httptest.NewRequest("GET", apiPath, nil)
	req.AddCookie(&http.Cookie{Name: "session_id", Value: cookie.SessionToken})
	w := httptest.NewRecorder()

	handler.GetCart(w, req)

	resp := w.Result()

	require.Equal(t, 500, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))

}

func TestAddProductToCartSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiPath := "/api/cart"
	mock := mockC.NewMockCartUsecaseI(ctrl)
	var logger *mw.ACLog
	handler := NewCartHandler(mock, logger)

	var id uint
	id = 1

	cookie := &entity.Cookie{
		UserID:       1,
		SessionToken: "HJJvgsvd",
	}

	body, err := json.Marshal(id)
	if err != nil {
		return
	}

	req := httptest.NewRequest("POST", apiPath, bytes.NewReader(body))
	req.AddCookie(&http.Cookie{Name: "session_id", Value: cookie.SessionToken})
	w := httptest.NewRecorder()

	mock.EXPECT().AddProductToCart(cookie.SessionToken, id).Return(nil)

	handler.AddProductToCart(w, req)

	resp := w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 201, resp.StatusCode)
}

func TestAddProductToCartFail(t *testing.T) {
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
	hitstats := &entity.HitStats{}
	logger := mw.NewACLog(baseLogger.Sugar(), errorLogger.Sugar(), *hitstats)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiPath := "/api/cart"
	mock := mockC.NewMockCartUsecaseI(ctrl)
	handler := NewCartHandler(mock, logger)

	var id uint
	id = 1

	body, err := json.Marshal(id)
	if err != nil {
		return
	}

	testErr := errors.New("test")

	cookie := &entity.Cookie{
		UserID:       1,
		SessionToken: "HJJvgsvd",
	}

	mock.EXPECT().AddProductToCart(cookie.SessionToken, id).Return(testErr)

	req := httptest.NewRequest("POST", apiPath, bytes.NewReader(body))
	req.AddCookie(&http.Cookie{Name: "session_id", Value: cookie.SessionToken})
	w := httptest.NewRecorder()

	handler.AddProductToCart(w, req)

	resp := w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 500, resp.StatusCode)

	mock.EXPECT().AddProductToCart(cookie.SessionToken, id).Return(entity.ErrNotFound)

	body, err = json.Marshal(id)
	if err != nil {
		return
	}

	req = httptest.NewRequest("POST", apiPath, bytes.NewReader(body))
	req.AddCookie(&http.Cookie{Name: "session_id", Value: cookie.SessionToken})
	w = httptest.NewRecorder()

	handler.AddProductToCart(w, req)

	resp = w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 404, resp.StatusCode)
}

func TestDeleteProductFromCartSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiPath := "/api/cart"
	mock := mockC.NewMockCartUsecaseI(ctrl)
	var logger *mw.ACLog
	handler := NewCartHandler(mock, logger)

	cookie := &entity.Cookie{
		UserID:       1,
		SessionToken: "HJJvgsvd",
	}

	req := httptest.NewRequest("DELETE", apiPath, nil)
	req.AddCookie(&http.Cookie{Name: "session_id", Value: cookie.SessionToken})
	w := httptest.NewRecorder()

	vars := map[string]string{
		"id": "1",
	}

	req = mux.SetURLVars(req, vars)

	var elemID uint
	elemID = 1

	mock.EXPECT().DeleteProductFromCart(cookie.SessionToken, elemID).Return(nil)

	handler.DeleteProductFromCart(w, req)

	resp := w.Result()

	require.Equal(t, 200, resp.StatusCode)
}

func TestDeleteProductFromCartFail(t *testing.T) {
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
	hitstats := &entity.HitStats{}
	logger := mw.NewACLog(baseLogger.Sugar(), errorLogger.Sugar(), *hitstats)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiPath := "/api/cart/1"
	mock := mockC.NewMockCartUsecaseI(ctrl)
	handler := NewCartHandler(mock, logger)

	testErr := errors.New("test")

	cookie := &entity.Cookie{
		UserID:       1,
		SessionToken: "HJJvgsvd",
	}

	vars := map[string]string{
		"id": "1",
	}

	mock.EXPECT().DeleteProductFromCart(cookie.SessionToken, uint(1)).Return(testErr)

	req := httptest.NewRequest("DELETE", apiPath, nil)
	req.AddCookie(&http.Cookie{Name: "session_id", Value: cookie.SessionToken})
	w := httptest.NewRecorder()
	req = mux.SetURLVars(req, vars)

	handler.DeleteProductFromCart(w, req)

	resp := w.Result()

	require.Equal(t, 500, resp.StatusCode)

	mock.EXPECT().DeleteProductFromCart(cookie.SessionToken, uint(1)).Return(entity.ErrNotFound)

	req = httptest.NewRequest("POST", apiPath, nil)
	req.AddCookie(&http.Cookie{Name: "session_id", Value: cookie.SessionToken})
	w = httptest.NewRecorder()
	req = mux.SetURLVars(req, vars)

	handler.DeleteProductFromCart(w, req)

	resp = w.Result()

	require.Equal(t, 404, resp.StatusCode)
}

func TestCleanCartSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiPath := "/api/cart/clear"
	mock := mockC.NewMockCartUsecaseI(ctrl)
	var logger *mw.ACLog
	handler := NewCartHandler(mock, logger)

	cookie := &entity.Cookie{
		UserID:       1,
		SessionToken: "HJJvgsvd",
	}

	req := httptest.NewRequest("POST", apiPath, nil)
	req.AddCookie(&http.Cookie{Name: "session_id", Value: cookie.SessionToken})
	w := httptest.NewRecorder()

	vars := map[string]string{
		"id": "1",
	}

	req = mux.SetURLVars(req, vars)

	mock.EXPECT().CleanCart(cookie.SessionToken).Return(nil)

	handler.CleanCart(w, req)

	resp := w.Result()

	require.Equal(t, 200, resp.StatusCode)
}

func TestCleanCartFail(t *testing.T) {
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
	hitstats := &entity.HitStats{}
	logger := mw.NewACLog(baseLogger.Sugar(), errorLogger.Sugar(), *hitstats)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiPath := "/api/cart/clear"
	mock := mockC.NewMockCartUsecaseI(ctrl)
	handler := NewCartHandler(mock, logger)

	testErr := errors.New("test")

	cookie := &entity.Cookie{
		UserID:       1,
		SessionToken: "HJJvgsvd",
	}

	mock.EXPECT().CleanCart(cookie.SessionToken).Return(testErr)

	req := httptest.NewRequest("POST", apiPath, nil)
	req.AddCookie(&http.Cookie{Name: "session_id", Value: cookie.SessionToken})
	w := httptest.NewRecorder()

	handler.CleanCart(w, req)

	resp := w.Result()

	require.Equal(t, 500, resp.StatusCode)
}
