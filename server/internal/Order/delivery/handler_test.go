package delivery

import (
	"bytes"
	"encoding/json"
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
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/require"
)

func TestCreateOrderSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiPath := "/api/orders"
	mockO := mockO.NewMockOrderUsecaseI(ctrl)
	mockS := mockS.NewMockSessionUsecaseI(ctrl)
	var logger *mw.ACLog
	handler := NewOrderHandler(mockO, mockS, logger)

	cookie := &entity.Cookie{
		UserID:       1,
		SessionToken: "HJJvgsvd",
	}

	reqorder := &dto.ReqCreateOrder{
		UserID: cookie.UserID,
		Address: &dto.ReqCreateOrderAddress{
			City:   "Moscow",
			Street: "Tverskaya",
			House:  "2",
			Flat:   1,
		},
	}

	timenow := time.Now()

	resporder := &dto.RespCreateOrder{
		ID:     1,
		Status: 0,
		Price:  100,
		Date:   timenow,
		Address: &entity.Address{
			City:   "Moscow",
			Street: "Tverskaya",
			House:  "2",
			Flat:   1,
		},

		DeliveryTime: 30,
	}

	mockS.EXPECT().GetIDByCookie(cookie.SessionToken).Return(cookie.UserID, nil)
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
		return
	}
	defer baseLogger.Sync()

	errorLogger, err := config.ErrorCfg.Build()
	if err != nil {
		return
	}
	defer errorLogger.Sync()
	var OKHitCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "ok_request_count",
			Help: "200 status counter",
		},
	)

	var InternalServerErrorCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "internal_server_error_request_count",
			Help: "500 status counter",
		},
	)

	var NotFoundErrorCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "no_found_server_error_request_count",
			Help: "400 status counter",
		},
	)

	var hits = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "hits",
	}, []string{"status", "path"})

	var timerhits = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "timerhits",
	}, []string{"status", "path"})

	hitstats := &entity.HitStats{
		Ok:                  OKHitCounter,
		InternalServerError: InternalServerErrorCounter,
		NotFoundError:       NotFoundErrorCounter,
		URLMetric:           *hits,
		Timing:              *timerhits,
	}

	logger := mw.NewACLog(baseLogger.Sugar(), errorLogger.Sugar(), *hitstats)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiPath := "/api/orders"
	mockO := mockO.NewMockOrderUsecaseI(ctrl)
	mockS := mockS.NewMockSessionUsecaseI(ctrl)
	handler := NewOrderHandler(mockO, mockS, logger)

	cookie := &entity.Cookie{
		UserID:       1,
		SessionToken: "HJJvgsvd",
	}

	reqorder := &dto.ReqCreateOrder{
		UserID: cookie.UserID,
		Address: &dto.ReqCreateOrderAddress{
			City:   "Moscow",
			Street: "Tverskaya",
			House:  "2",
			Flat:   1,
		},
	}

	mockS.EXPECT().GetIDByCookie(cookie.SessionToken).Return(cookie.UserID, nil)
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

	mockS.EXPECT().GetIDByCookie(cookie.SessionToken).Return(cookie.UserID, nil)
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
	mockO := mockO.NewMockOrderUsecaseI(ctrl)
	mockS := mockS.NewMockSessionUsecaseI(ctrl)
	var logger *mw.ACLog
	handler := NewOrderHandler(mockO, mockS, logger)

	cookie := &entity.Cookie{
		UserID:       1,
		SessionToken: "HJJvgsvd",
	}

	reqorder := &dto.ReqUpdateOrder{
		ID:     1,
		Status: 1,
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
		return
	}
	defer baseLogger.Sync()

	errorLogger, err := config.ErrorCfg.Build()
	if err != nil {
		return
	}
	defer errorLogger.Sync()

	var OKHitCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "ok_request_count",
			Help: "200 status counter",
		},
	)

	var InternalServerErrorCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "internal_server_error_request_count",
			Help: "500 status counter",
		},
	)

	var NotFoundErrorCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "no_found_server_error_request_count",
			Help: "400 status counter",
		},
	)

	var hits = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "hits",
	}, []string{"status", "path"})

	var timerhits = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "timerhits",
	}, []string{"status", "path"})

	hitstats := &entity.HitStats{
		Ok:                  OKHitCounter,
		InternalServerError: InternalServerErrorCounter,
		NotFoundError:       NotFoundErrorCounter,
		URLMetric:           *hits,
		Timing:              *timerhits,
	}

	logger := mw.NewACLog(baseLogger.Sugar(), errorLogger.Sugar(), *hitstats)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiPath := "/api/orders"
	mockO := mockO.NewMockOrderUsecaseI(ctrl)
	mockS := mockS.NewMockSessionUsecaseI(ctrl)
	handler := NewOrderHandler(mockO, mockS, logger)

	cookie := &entity.Cookie{
		UserID:       1,
		SessionToken: "HJJvgsvd",
	}

	reqorder := &dto.ReqUpdateOrder{
		ID:     1,
		Status: 1,
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
	mockO := mockO.NewMockOrderUsecaseI(ctrl)
	mockS := mockS.NewMockSessionUsecaseI(ctrl)
	var logger *mw.ACLog
	handler := NewOrderHandler(mockO, mockS, logger)

	cookie := &entity.Cookie{
		UserID:       1,
		SessionToken: "HJJvgsvd",
	}

	timenow := time.Now()

	var flat uint
	flat = 1

	resporders := &dto.RespOrders{
		{
			ID:     1,
			Status: 0,
			Date:   timenow,
			Address: &dto.RespOrderAddress{
				City:   "Moscow",
				Street: "Tverskaya",
				House:  "2",
				Flat:   flat,
			},
		},
		{
			ID:     2,
			Status: 0,
			Date:   timenow,
			Address: &dto.RespOrderAddress{
				City:   "Moscow",
				Street: "Tverskaya",
				House:  "3",
				Flat:   flat,
			},
		},
	}

	mockS.EXPECT().GetIDByCookie(cookie.SessionToken).Return(cookie.UserID, nil)
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
		return
	}
	defer baseLogger.Sync()

	errorLogger, err := config.ErrorCfg.Build()
	if err != nil {
		return
	}
	defer errorLogger.Sync()

	var OKHitCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "ok_request_count",
			Help: "200 status counter",
		},
	)

	var InternalServerErrorCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "internal_server_error_request_count",
			Help: "500 status counter",
		},
	)

	var NotFoundErrorCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "no_found_server_error_request_count",
			Help: "400 status counter",
		},
	)

	var hits = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "hits",
	}, []string{"status", "path"})

	var timerhits = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "timerhits",
	}, []string{"status", "path"})

	hitstats := &entity.HitStats{
		Ok:                  OKHitCounter,
		InternalServerError: InternalServerErrorCounter,
		NotFoundError:       NotFoundErrorCounter,
		URLMetric:           *hits,
		Timing:              *timerhits,
	}

	logger := mw.NewACLog(baseLogger.Sugar(), errorLogger.Sugar(), *hitstats)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiPath := "/api/orders"
	mockO := mockO.NewMockOrderUsecaseI(ctrl)
	mockS := mockS.NewMockSessionUsecaseI(ctrl)
	handler := NewOrderHandler(mockO, mockS, logger)

	cookie := &entity.Cookie{
		UserID:       1,
		SessionToken: "HJJvgsvd",
	}

	mockS.EXPECT().GetIDByCookie(cookie.SessionToken).Return(cookie.UserID, nil)
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
	mockO := mockO.NewMockOrderUsecaseI(ctrl)
	mockS := mockS.NewMockSessionUsecaseI(ctrl)
	var logger *mw.ACLog
	handler := NewOrderHandler(mockO, mockS, logger)

	cookie := &entity.Cookie{
		UserID:       1,
		SessionToken: "HJJvgsvd",
	}

	timenow := time.Now()

	reqOrder := dto.ReqGetOneOrder{
		UserID:  1,
		OrderID: 1,
	}

	products := &dto.RespGetOrderProduct{
		ID:    1,
		Name:  "Burger",
		Price: 100,
		Icon:  "def",
		Count: 1,
	}

	orderItems := &dto.OrderItems{
		RestaurantName: "BK",
		Products:       []*dto.RespGetOrderProduct{products},
	}

	resporder := &dto.RespGetOneOrder{
		ID:     1,
		Status: 0,
		Date:   timenow,
		Address: &dto.RespOrderAddress{
			City:   "Moscow",
			Street: "Tverskaya",
			House:  "3",
			Flat:   1,
		},
		OrderItems:   []*dto.OrderItems{orderItems},
		Price:        100,
		DeliveryTime: 30,
	}

	mockS.EXPECT().GetIDByCookie(cookie.SessionToken).Return(cookie.UserID, nil)
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
		return
	}
	defer baseLogger.Sync()

	errorLogger, err := config.ErrorCfg.Build()
	if err != nil {
		return
	}
	defer errorLogger.Sync()
	var OKHitCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "ok_request_count",
			Help: "200 status counter",
		},
	)

	var InternalServerErrorCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "internal_server_error_request_count",
			Help: "500 status counter",
		},
	)

	var NotFoundErrorCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "no_found_server_error_request_count",
			Help: "400 status counter",
		},
	)

	var hits = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "hits",
	}, []string{"status", "path"})

	var timerhits = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "timerhits",
	}, []string{"status", "path"})

	hitstats := &entity.HitStats{
		Ok:                  OKHitCounter,
		InternalServerError: InternalServerErrorCounter,
		NotFoundError:       NotFoundErrorCounter,
		URLMetric:           *hits,
		Timing:              *timerhits,
	}
	logger := mw.NewACLog(baseLogger.Sugar(), errorLogger.Sugar(), *hitstats)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiPath := "/api/orders/1"
	mockO := mockO.NewMockOrderUsecaseI(ctrl)
	mockS := mockS.NewMockSessionUsecaseI(ctrl)
	handler := NewOrderHandler(mockO, mockS, logger)

	cookie := &entity.Cookie{
		UserID:       1,
		SessionToken: "HJJvgsvd",
	}

	reqOrder := dto.ReqGetOneOrder{
		UserID:  1,
		OrderID: 1,
	}

	mockS.EXPECT().GetIDByCookie(cookie.SessionToken).Return(cookie.UserID, nil)
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

	mockS.EXPECT().GetIDByCookie(cookie.SessionToken).Return(cookie.UserID, nil)
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
