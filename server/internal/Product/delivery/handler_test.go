package delivery

import (
	"errors"
	"fmt"
	//"io/ioutil"
	"net/http/httptest"
	"server/config"
	mockP "server/internal/Product/usecase/mock_usecase"
	mw "server/internal/middleware"
	"testing"

	"server/internal/domain/entity"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

func TestGetProductSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiPath := "/api/products/1"
	mock := mockP.NewMockUsecaseI(ctrl)
	var logger *mw.ACLog
	handler := NewProductHandler(mock, logger)

	product := &entity.Product{
		ID:          1,
		Name:        "Burger",
		Price:       120.0,
		CookingTime: 23,
		Portion:     "160 г",
		Description: "Nice burger",
		Icon:        "deficon",
	}

	var elemID = 1

	mock.EXPECT().GetProductByID(uint(elemID)).Return(product, nil)

	req := httptest.NewRequest("GET", apiPath, nil)
	w := httptest.NewRecorder()

	vars := map[string]string{
		"id": "1",
	}

	req = mux.SetURLVars(req, vars)

	handler.GetProduct(w, req)

	resp := w.Result()
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return
	// }

	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	//require.Contains(t, string(body), "Body")

}

func TestGetProductFail(t *testing.T) {
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
	apiPath := "/api/products/fdfd"
	mock := mockP.NewMockUsecaseI(ctrl)
	hitstats := entity.HitStats{}
	logger := mw.NewACLog(baseLogger.Sugar(), errorLogger.Sugar(), hitstats)
	handler := NewProductHandler(mock, logger)

	req := httptest.NewRequest("GET", apiPath, nil)
	w := httptest.NewRecorder()

	handler.GetProduct(w, req)

	resp := w.Result()

	require.Equal(t, 400, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))

	req = httptest.NewRequest("GET", apiPath, nil)
	w = httptest.NewRecorder()

	vars := map[string]string{
		"id": "fdfd",
	}

	req = mux.SetURLVars(req, vars)

	handler.GetProduct(w, req)

	resp = w.Result()

	require.Equal(t, 400, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))

	testErr := errors.New("test")
	var elemID = 1

	mock.EXPECT().GetProductByID(uint(elemID)).Return(nil, testErr)

	req = httptest.NewRequest("GET", apiPath, nil)
	w = httptest.NewRecorder()

	vars = map[string]string{
		"id": "1",
	}

	req = mux.SetURLVars(req, vars)

	handler.GetProduct(w, req)

	resp = w.Result()

	require.Equal(t, 500, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))

	mock.EXPECT().GetProductByID(uint(elemID)).Return(nil, entity.ErrNotFound)

	req = httptest.NewRequest("GET", apiPath, nil)
	w = httptest.NewRecorder()

	vars = map[string]string{
		"id": "1",
	}

	req = mux.SetURLVars(req, vars)

	handler.GetProduct(w, req)

	resp = w.Result()

	require.Equal(t, 404, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))

}
