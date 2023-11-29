package delivery

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"server/config"
	mockO "server/internal/Order/usecase/mock_usecase"
	mockS "server/internal/Session/usecase/mock_usecase"
	"server/internal/domain/dto"
	"server/internal/domain/entity"
	mw "server/internal/middleware"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

func TestCreateOrderSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiPath := "/api/orders"
	mockO := mockO.NewMockUsecaseI(ctrl)
	mockS := mockS.NewMockUsecaseI(ctrl)
	var logger *mw.ACLog
	handler := NewOrderHandler(mockO, mockS, logger)

	cookie := &entity.Cookie{
		UserID:       1,
		SessionToken: "HJJvgsvd",
	}

	reqorder := &dto.ReqCreateOrder{
		Products: []uint{1, 2, 3},
		UserId:   cookie.UserID,
		Address: &dto.ReqCreateOrderAddress{
			City:   "Moscow",
			Street: "Tverskaya",
			House:  "2",
			Flat:   1,
		},
	}

	timenow := time.Now()

	resporder := &dto.RespCreateOrder{
		Id:     1,
		Status: "Wait",
		Date:   timenow,
	}

	mockS.EXPECT().GetIdByCookie(cookie.SessionToken).Return(cookie.UserID, nil)
	mockO.EXPECT().CreateOrder(reqorder).Return(resporder, nil)

	body, err := json.Marshal(reqorder)
	if err != nil {
		return
	}

	req := httptest.NewRequest("POST", apiPath, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(&http.Cookie{Name: "session_id", Value: cookie.SessionToken})
	w := httptest.NewRecorder()

	handler.CreateOrder(w, req)

	resp := w.Result()

	require.Equal(t, 201, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))
}

func TestCreateOrderFail(t *testing.T) {
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
	logger := mw.NewACLog(baseLogger.Sugar(), errorLogger.Sugar())
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiPath := "/api/orders"
	mockO := mockO.NewMockUsecaseI(ctrl)
	mockS := mockS.NewMockUsecaseI(ctrl)
	handler := NewOrderHandler(mockO, mockS, logger)

	cookie := &entity.Cookie{
		UserID:       1,
		SessionToken: "HJJvgsvd",
	}

	reqorder := &dto.ReqCreateOrder{
		Products: []uint{1, 2, 3},
		UserId:   cookie.UserID,
		Address: &dto.ReqCreateOrderAddress{
			City:   "Moscow",
			Street: "Tverskaya",
			House:  "2",
			Flat:   1,
		},
	}

	mockS.EXPECT().GetIdByCookie(cookie.SessionToken).Return(cookie.UserID, nil)
	mockO.EXPECT().CreateOrder(reqorder).Return(nil, entity.ErrInternalServerError)

	body, err := json.Marshal(reqorder)
	if err != nil {
		return
	}

	req := httptest.NewRequest("POST", apiPath, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(&http.Cookie{Name: "session_id", Value: cookie.SessionToken})
	w := httptest.NewRecorder()

	handler.CreateOrder(w, req)

	resp := w.Result()

	require.Equal(t, 500, resp.StatusCode)

	mockS.EXPECT().GetIdByCookie(cookie.SessionToken).Return(cookie.UserID, nil)
	mockO.EXPECT().CreateOrder(reqorder).Return(nil, entity.ErrBadRequest)

	body, err = json.Marshal(reqorder)
	if err != nil {
		return
	}

	req = httptest.NewRequest("POST", apiPath, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(&http.Cookie{Name: "session_id", Value: cookie.SessionToken})
	w = httptest.NewRecorder()

	handler.CreateOrder(w, req)

	resp = w.Result()

	require.Equal(t, 400, resp.StatusCode)

	body, err = json.Marshal(reqorder)
	if err != nil {
		return
	}

	req = httptest.NewRequest("POST", apiPath, bytes.NewReader(body))
	req.AddCookie(&http.Cookie{Name: "session_id", Value: cookie.SessionToken})
	w = httptest.NewRecorder()

	handler.CreateOrder(w, req)

	resp = w.Result()

	require.Equal(t, 400, resp.StatusCode)

}

func TestUpdateOrderSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiPath := "/api/orders"
	mockO := mockO.NewMockUsecaseI(ctrl)
	mockS := mockS.NewMockUsecaseI(ctrl)
	var logger *mw.ACLog
	handler := NewOrderHandler(mockO, mockS, logger)

	cookie := &entity.Cookie{
		UserID:       1,
		SessionToken: "HJJvgsvd",
	}

	reqorder := &dto.ReqUpdateOrder{
		Id:     1,
		Status: "Wait",
	}

	mockO.EXPECT().UpdateOrder(reqorder).Return(nil)

	body, err := json.Marshal(reqorder)
	if err != nil {
		return
	}

	req := httptest.NewRequest("PATCH", apiPath, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(&http.Cookie{Name: "session_id", Value: cookie.SessionToken})
	w := httptest.NewRecorder()

	handler.UpdateOrder(w, req)

	resp := w.Result()

	require.Equal(t, 200, resp.StatusCode)
}

func TestUpdateOrderFail(t *testing.T) {
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
	logger := mw.NewACLog(baseLogger.Sugar(), errorLogger.Sugar())
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiPath := "/api/orders"
	mockO := mockO.NewMockUsecaseI(ctrl)
	mockS := mockS.NewMockUsecaseI(ctrl)
	handler := NewOrderHandler(mockO, mockS, logger)

	cookie := &entity.Cookie{
		UserID:       1,
		SessionToken: "HJJvgsvd",
	}

	reqorder := &dto.ReqUpdateOrder{
		Id:     1,
		Status: "Wait",
	}

	mockO.EXPECT().UpdateOrder(reqorder).Return(entity.ErrInternalServerError)

	body, err := json.Marshal(reqorder)
	if err != nil {
		return
	}

	req := httptest.NewRequest("PATCH", apiPath, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(&http.Cookie{Name: "session_id", Value: cookie.SessionToken})
	w := httptest.NewRecorder()

	handler.UpdateOrder(w, req)

	resp := w.Result()

	require.Equal(t, 500, resp.StatusCode)

}

func TestGetOrdersSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiPath := "/api/orders"
	mockO := mockO.NewMockUsecaseI(ctrl)
	mockS := mockS.NewMockUsecaseI(ctrl)
	var logger *mw.ACLog
	handler := NewOrderHandler(mockO, mockS, logger)

	cookie := &entity.Cookie{
		UserID:       1,
		SessionToken: "HJJvgsvd",
	}

	timenow := time.Now()

	var flat uint
	flat = 1

	resporders := []*dto.RespGetOrder{
		{
			Id:     1,
			Status: "Wait",
			Date:   timenow,
			Address: &dto.RespOrderAddress{
				City:   "Moscow",
				Street: "Tverskaya",
				House:  "2",
				Flat:   &flat,
			},
		},
		{
			Id:     2,
			Status: "Wait",
			Date:   timenow,
			Address: &dto.RespOrderAddress{
				City:   "Moscow",
				Street: "Tverskaya",
				House:  "3",
				Flat:   &flat,
			},
		},
	}

	mockS.EXPECT().GetIdByCookie(cookie.SessionToken).Return(cookie.UserID, nil)
	mockO.EXPECT().GetOrders(cookie.UserID).Return(resporders, nil)

	req := httptest.NewRequest("GET", apiPath, nil)
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(&http.Cookie{Name: "session_id", Value: cookie.SessionToken})
	w := httptest.NewRecorder()

	handler.GetOrders(w, req)

	resp := w.Result()

	require.Equal(t, 200, resp.StatusCode)
}

func TestGetOrdersFail(t *testing.T) {
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
	logger := mw.NewACLog(baseLogger.Sugar(), errorLogger.Sugar())
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiPath := "/api/orders"
	mockO := mockO.NewMockUsecaseI(ctrl)
	mockS := mockS.NewMockUsecaseI(ctrl)
	handler := NewOrderHandler(mockO, mockS, logger)

	cookie := &entity.Cookie{
		UserID:       1,
		SessionToken: "HJJvgsvd",
	}

	mockS.EXPECT().GetIdByCookie(cookie.SessionToken).Return(cookie.UserID, nil)
	mockO.EXPECT().GetOrders(cookie.UserID).Return(nil, entity.ErrInternalServerError)

	req := httptest.NewRequest("GET", apiPath, nil)
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(&http.Cookie{Name: "session_id", Value: cookie.SessionToken})
	w := httptest.NewRecorder()

	handler.GetOrders(w, req)

	resp := w.Result()

	require.Equal(t, 500, resp.StatusCode)
}

func TestGetOrderSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiPath := "/api/orders/1"
	mockO := mockO.NewMockUsecaseI(ctrl)
	mockS := mockS.NewMockUsecaseI(ctrl)
	var logger *mw.ACLog
	handler := NewOrderHandler(mockO, mockS, logger)

	cookie := &entity.Cookie{
		UserID:       1,
		SessionToken: "HJJvgsvd",
	}

	timenow := time.Now()

	reqOrder := dto.ReqGetOneOrder{
		UserId:  1,
		OrderId: 1,
	}

	resporder := &dto.RespGetOneOrder{
		Status:      "Wait",
		Date:        timenow,
		UpdatedDate: timenow,
		Products:    []*dto.RespGetOrderProduct{},
	}

	mockS.EXPECT().GetIdByCookie(cookie.SessionToken).Return(cookie.UserID, nil)
	mockO.EXPECT().GetOrder(&reqOrder).Return(resporder, nil)

	req := httptest.NewRequest("GET", apiPath, nil)
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(&http.Cookie{Name: "session_id", Value: cookie.SessionToken})
	w := httptest.NewRecorder()

	vars := map[string]string{
		"id": "1",
	}

	req = mux.SetURLVars(req, vars)

	handler.GetOrder(w, req)

	resp := w.Result()

	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))
}

func TestGetOrderFail(t *testing.T) {
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
	logger := mw.NewACLog(baseLogger.Sugar(), errorLogger.Sugar())
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiPath := "/api/orders/1"
	mockO := mockO.NewMockUsecaseI(ctrl)
	mockS := mockS.NewMockUsecaseI(ctrl)
	handler := NewOrderHandler(mockO, mockS, logger)

	cookie := &entity.Cookie{
		UserID:       1,
		SessionToken: "HJJvgsvd",
	}

	reqOrder := dto.ReqGetOneOrder{
		UserId:  1,
		OrderId: 1,
	}

	mockS.EXPECT().GetIdByCookie(cookie.SessionToken).Return(cookie.UserID, nil)
	mockO.EXPECT().GetOrder(&reqOrder).Return(nil, entity.ErrInternalServerError)

	req := httptest.NewRequest("GET", apiPath, nil)
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(&http.Cookie{Name: "session_id", Value: cookie.SessionToken})
	w := httptest.NewRecorder()

	vars := map[string]string{
		"id": "1",
	}

	req = mux.SetURLVars(req, vars)

	handler.GetOrder(w, req)

	resp := w.Result()

	require.Equal(t, 500, resp.StatusCode)

	mockS.EXPECT().GetIdByCookie(cookie.SessionToken).Return(cookie.UserID, nil)
	mockO.EXPECT().GetOrder(&reqOrder).Return(nil, entity.ErrNotFound)

	req = httptest.NewRequest("GET", apiPath, nil)
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(&http.Cookie{Name: "session_id", Value: cookie.SessionToken})
	w = httptest.NewRecorder()

	vars = map[string]string{
		"id": "1",
	}

	req = mux.SetURLVars(req, vars)

	handler.GetOrder(w, req)

	resp = w.Result()

	require.Equal(t, 404, resp.StatusCode)
}
