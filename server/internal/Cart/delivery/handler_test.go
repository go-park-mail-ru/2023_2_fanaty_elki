package delivery

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	mockC "server/internal/Cart/usecase/mock_usecase"
	"testing"

	"server/internal/domain/dto"
	"server/internal/domain/entity"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetCartSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiPath := "/api/cart"
	mock := mockC.NewMockUsecaseI(ctrl)
	handler := NewCartHandler(mock)

	cookie := &entity.Cookie{
		UserID:       1,
		SessionToken: "HJJvgsvd",
	}

	res := []*dto.CartProduct{
		{Product: &entity.Product{
			ID:          1,
			Name:        "Burger",
			Price:       120.0,
			CookingTime: 23,
			Portion:     "160 г",
			Description: "Nice burger",
			Icon:        "deficon",
		},
			ItemCount: 1,
		},
	}

	mock.EXPECT().GetUserCart(cookie.SessionToken).Return(res, nil)

	req := httptest.NewRequest("GET", apiPath, nil)
	req.AddCookie(&http.Cookie{Name: "session_id", Value: cookie.SessionToken})
	w := httptest.NewRecorder()

	handler.GetCart(w, req)

	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	require.Contains(t, string(body), "Cart")

}

func TestGetCartFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiPath := "/api/cart"
	mock := mockC.NewMockUsecaseI(ctrl)
	handler := NewCartHandler(mock)

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

	req = httptest.NewRequest("GET", apiPath, nil)
	w = httptest.NewRecorder()

	handler.GetCart(w, req)

	resp = w.Result()

	require.Equal(t, 401, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))
}

func TestAddProductToCartSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiPath := "/api/cart/add"
	mock := mockC.NewMockUsecaseI(ctrl)
	handler := NewCartHandler(mock)

	var product = map[string]interface{}{
		"ProductID": 1,
	}

	cookie := &entity.Cookie{
		UserID:       1,
		SessionToken: "HJJvgsvd",
	}

	body, err := json.Marshal(product)
	if err != nil {
		return
	}

	req := httptest.NewRequest("POST", apiPath, bytes.NewReader(body))
	req.AddCookie(&http.Cookie{Name: "session_id", Value: cookie.SessionToken})
	w := httptest.NewRecorder()

	reqProduct := dto.ReqProductID{ProductID: 1}

	mock.EXPECT().AddProductToCart(cookie.SessionToken, reqProduct.ProductID).Return(nil)

	handler.AddProductToCart(w, req)

	resp := w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))
}

func TestAddProductToCartFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiPath := "/api/cart/add"
	mock := mockC.NewMockUsecaseI(ctrl)
	handler := NewCartHandler(mock)

	var product = map[string]interface{}{
		"ProductID": 1,
	}

	body, err := json.Marshal(product)
	if err != nil {
		return
	}

	req := httptest.NewRequest("POST", apiPath, bytes.NewReader(body))
	w := httptest.NewRecorder()

	handler.AddProductToCart(w, req)

	resp := w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 401, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))
}

func TestDeleteProductFromCartSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiPath := "/api/cart/delete"
	mock := mockC.NewMockUsecaseI(ctrl)
	handler := NewCartHandler(mock)

	var product = map[string]interface{}{
		"ProductID": 1,
	}

	cookie := &entity.Cookie{
		UserID:       1,
		SessionToken: "HJJvgsvd",
	}

	body, err := json.Marshal(product)
	if err != nil {
		return
	}

	req := httptest.NewRequest("POST", apiPath, bytes.NewReader(body))
	req.AddCookie(&http.Cookie{Name: "session_id", Value: cookie.SessionToken})
	w := httptest.NewRecorder()

	reqProduct := dto.ReqProductID{ProductID: 1}

	mock.EXPECT().DeleteProductFromCart(cookie.SessionToken, reqProduct.ProductID).Return(nil)

	handler.DeleteProductFromCart(w, req)

	resp := w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))
}

func TestUpdateItemCountUpSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiPath := "/api/cart/update/up"
	mock := mockC.NewMockUsecaseI(ctrl)
	handler := NewCartHandler(mock)

	var product = map[string]interface{}{
		"ProductID": 1,
	}

	cookie := &entity.Cookie{
		UserID:       1,
		SessionToken: "HJJvgsvd",
	}

	body, err := json.Marshal(product)
	if err != nil {
		return
	}

	req := httptest.NewRequest("PATCH", apiPath, bytes.NewReader(body))
	req.AddCookie(&http.Cookie{Name: "session_id", Value: cookie.SessionToken})
	w := httptest.NewRecorder()

	reqProduct := dto.ReqProductID{ProductID: 1}

	mock.EXPECT().UpdateItemCountUp(cookie.SessionToken, reqProduct.ProductID).Return(nil)

	handler.UpdateItemCountUp(w, req)

	resp := w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))
}

func TestUpdateItemCountDownSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiPath := "/api/cart/update/down"
	mock := mockC.NewMockUsecaseI(ctrl)
	handler := NewCartHandler(mock)

	var product = map[string]interface{}{
		"ProductID": 1,
	}

	cookie := &entity.Cookie{
		UserID:       1,
		SessionToken: "HJJvgsvd",
	}

	body, err := json.Marshal(product)
	if err != nil {
		return
	}

	req := httptest.NewRequest("PATCH", apiPath, bytes.NewReader(body))
	req.AddCookie(&http.Cookie{Name: "session_id", Value: cookie.SessionToken})
	w := httptest.NewRecorder()

	reqProduct := dto.ReqProductID{ProductID: 1}

	mock.EXPECT().UpdateItemCountDown(cookie.SessionToken, reqProduct.ProductID).Return(nil)

	handler.UpdateItemCountDown(w, req)

	resp := w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))
}
