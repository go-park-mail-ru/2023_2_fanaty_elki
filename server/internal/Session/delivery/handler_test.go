package delivery

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"time"

	"server/config"
	mockS "server/internal/Session/usecase/mock_usecase"
	mockU "server/internal/User/usecase/mock_usecase"
	mw "server/internal/middleware"
	"testing"

	"server/internal/domain/dto"
	"server/internal/domain/entity"

	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/require"
)

func TestSignUpSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var logger *mw.ACLog
	apiPath := "/api/users"
	mockSes := mockS.NewMockSessionUsecaseI(ctrl)
	mockUs := mockU.NewMockUsecaseI(ctrl)
	handler := NewSessionHandler(mockSes, mockUs, logger)

	user := &entity.User{
		Username:    "ania",
		Password:    "anis1234",
		PhoneNumber: "89165342397",
		Email:       "ani@mail.ru",
	}

	var jsonuser = map[string]interface{}{
		"Username":    "ania",
		"Password":    "anis1234",
		"PhoneNumber": "89165342397",
		"Email":       "ani@mail.ru",
	}

	var UserID uint
	UserID = 1

	mockUs.EXPECT().CreateUser(user).Return(UserID, nil)

	body, err := json.Marshal(jsonuser)
	if err != nil {
		return
	}

	req := httptest.NewRequest("POST", apiPath, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.SignUp(w, req)

	resp := w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 201, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	//require.Contains(t, string(body), "Body")

}

func TestSignUpFail(t *testing.T) {
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
	apiPath := "/api/users"
	mockSes := mockS.NewMockSessionUsecaseI(ctrl)
	mockUs := mockU.NewMockUsecaseI(ctrl)
	handler := NewSessionHandler(mockSes, mockUs, logger)

	user := &entity.User{
		Username:    "ania",
		Password:    "anis1234",
		PhoneNumber: "89165342397",
		Email:       "ani@mail.ru",
	}

	var jsonuser = map[string]interface{}{
		"Username":    "ania",
		"Password":    "anis1234",
		"PhoneNumber": "89165342397",
		"Email":       "ani@mail.ru",
	}

	body, err := json.Marshal(jsonuser)
	if err != nil {
		return
	}

	req := httptest.NewRequest("POST", apiPath, bytes.NewReader(body))

	w := httptest.NewRecorder()

	handler.SignUp(w, req)

	resp := w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 400, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))

	req = httptest.NewRequest("POST", apiPath, nil)
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	handler.SignUp(w, req)

	resp = w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 400, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))

	mockUs.EXPECT().CreateUser(user).Return(uint(0), entity.ErrInternalServerError)

	body, err = json.Marshal(jsonuser)
	if err != nil {
		return
	}

	req = httptest.NewRequest("POST", apiPath, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	handler.SignUp(w, req)

	resp = w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 500, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))

	mockUs.EXPECT().CreateUser(user).Return(uint(0), entity.ErrInvalidBirthday)

	body, err = json.Marshal(jsonuser)
	if err != nil {
		return
	}

	req = httptest.NewRequest("POST", apiPath, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	handler.SignUp(w, req)

	resp = w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 400, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))

	mockUs.EXPECT().CreateUser(user).Return(uint(0), entity.ErrConflictEmail)

	body, err = json.Marshal(jsonuser)
	if err != nil {
		return
	}

	req = httptest.NewRequest("POST", apiPath, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	handler.SignUp(w, req)

	resp = w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 492, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))

	mockUs.EXPECT().CreateUser(user).Return(uint(0), entity.ErrConflictPhoneNumber)

	body, err = json.Marshal(jsonuser)
	if err != nil {
		return
	}

	req = httptest.NewRequest("POST", apiPath, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	handler.SignUp(w, req)

	resp = w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 493, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))

	mockUs.EXPECT().CreateUser(user).Return(uint(0), entity.ErrConflictUsername)

	body, err = json.Marshal(jsonuser)
	if err != nil {
		return
	}

	req = httptest.NewRequest("POST", apiPath, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	handler.SignUp(w, req)

	resp = w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 491, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))
}

func TestLoginSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var logger *mw.ACLog
	apiPath := "/api/login"
	mockSes := mockS.NewMockSessionUsecaseI(ctrl)
	mockUs := mockU.NewMockUsecaseI(ctrl)
	handler := NewSessionHandler(mockSes, mockUs, logger)

	user := &entity.User{
		Username: "ania",
		Password: "anis1234",
	}

	cookie := &entity.Cookie{
		UserID:       1,
		SessionToken: "TYebbYudb",
		MaxAge:       50 * time.Hour,
	}

	var jsonuser = map[string]interface{}{
		"Username": "ania",
		"Password": "anis1234",
	}

	profile := &dto.ReqGetUserProfile{
		Username: "ania",
	}

	mockSes.EXPECT().Login(user).Return(cookie, nil)
	mockSes.EXPECT().GetUserProfile(cookie.SessionToken).Return(profile, nil)

	body, err := json.Marshal(jsonuser)
	if err != nil {
		return
	}

	req := httptest.NewRequest("POST", apiPath, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.Login(w, req)

	resp := w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	//require.Contains(t, string(body), "Body")

}

func TestLoginFail(t *testing.T) {
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
	apiPath := "/api/login"
	mockSes := mockS.NewMockSessionUsecaseI(ctrl)
	mockUs := mockU.NewMockUsecaseI(ctrl)
	handler := NewSessionHandler(mockSes, mockUs, logger)

	user := &entity.User{
		Username: "ania",
		Password: "anis1234",
	}

	cookie := &entity.Cookie{
		UserID:       1,
		SessionToken: "TYebbYudb",
		MaxAge:       50 * time.Hour,
	}

	var jsonuser = map[string]interface{}{
		"Username": "ania",
		"Password": "anis1234",
	}

	body, err := json.Marshal(jsonuser)
	if err != nil {
		return
	}

	req := httptest.NewRequest("POST", apiPath, bytes.NewReader(body))

	w := httptest.NewRecorder()

	handler.Login(w, req)

	resp := w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 400, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))

	req = httptest.NewRequest("POST", apiPath, nil)
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	handler.Login(w, req)

	resp = w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 400, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))

	req = httptest.NewRequest("POST", apiPath, nil)
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	handler.Login(w, req)

	resp = w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 400, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))

	mockSes.EXPECT().Login(user).Return(nil, entity.ErrInternalServerError)
	body, err = json.Marshal(jsonuser)
	if err != nil {
		return
	}
	req = httptest.NewRequest("POST", apiPath, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	handler.Login(w, req)

	resp = w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 500, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))

	testErr := errors.New("testErr")

	mockSes.EXPECT().Login(user).Return(nil, testErr)
	body, err = json.Marshal(jsonuser)
	if err != nil {
		return
	}
	req = httptest.NewRequest("POST", apiPath, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	handler.Login(w, req)

	resp = w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 401, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))

	mockSes.EXPECT().Login(user).Return(cookie, nil)
	mockSes.EXPECT().GetUserProfile(cookie.SessionToken).Return(nil, entity.ErrInternalServerError)

	body, err = json.Marshal(jsonuser)
	if err != nil {
		return
	}
	req = httptest.NewRequest("POST", apiPath, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	handler.Login(w, req)

	resp = w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 500, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))

}

func TestLogoutSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var logger *mw.ACLog
	apiPath := "/api/logout"
	mockSes := mockS.NewMockSessionUsecaseI(ctrl)
	mockUs := mockU.NewMockUsecaseI(ctrl)
	handler := NewSessionHandler(mockSes, mockUs, logger)

	cookie := &entity.Cookie{
		SessionToken: "TYebbYudb",
	}

	mockSes.EXPECT().Logout(cookie).Return(nil)

	req := httptest.NewRequest("DELETE", apiPath, nil)
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(&http.Cookie{Name: "session_id", Value: cookie.SessionToken})
	w := httptest.NewRecorder()

	handler.Logout(w, req)

	resp := w.Result()

	require.Equal(t, 200, resp.StatusCode)
}

func TestLogoutFail(t *testing.T) {
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
	apiPath := "/api/logout"
	mockSes := mockS.NewMockSessionUsecaseI(ctrl)
	mockUs := mockU.NewMockUsecaseI(ctrl)
	handler := NewSessionHandler(mockSes, mockUs, logger)

	cookie := &entity.Cookie{
		SessionToken: "TYebbYudb",
	}

	mockSes.EXPECT().Logout(cookie).Return(entity.ErrInternalServerError)

	req := httptest.NewRequest("DELETE", apiPath, nil)
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(&http.Cookie{Name: "session_id", Value: cookie.SessionToken})
	w := httptest.NewRecorder()

	handler.Logout(w, req)

	resp := w.Result()

	require.Equal(t, 500, resp.StatusCode)
}

func TestAuthSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var logger *mw.ACLog
	apiPath := "/api/auth"
	mockSes := mockS.NewMockSessionUsecaseI(ctrl)
	mockUs := mockU.NewMockUsecaseI(ctrl)
	handler := NewSessionHandler(mockSes, mockUs, logger)

	cookie := &entity.Cookie{
		UserID:       1,
		SessionToken: "TYebbYudb",
	}

	profile := &dto.ReqGetUserProfile{
		Username: "ania",
	}

	mockSes.EXPECT().Check(cookie.SessionToken).Return(uint(1), nil)
	mockSes.EXPECT().CreateCookieAuth(cookie).Return(profile, nil)

	req := httptest.NewRequest("GET", apiPath, nil)
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(&http.Cookie{Name: "session_id", Value: cookie.SessionToken})
	w := httptest.NewRecorder()

	handler.Auth(w, req)

	resp := w.Result()

	require.Equal(t, 200, resp.StatusCode)
}

func TestProfileSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var logger *mw.ACLog
	apiPath := "/api/me"
	mockSes := mockS.NewMockSessionUsecaseI(ctrl)
	mockUs := mockU.NewMockUsecaseI(ctrl)
	handler := NewSessionHandler(mockSes, mockUs, logger)

	cookie := &entity.Cookie{
		UserID:       1,
		SessionToken: "TYebbYudb",
	}

	profile := &dto.ReqGetUserProfile{
		Username: "ania",
	}

	mockSes.EXPECT().GetUserProfile(cookie.SessionToken).Return(profile, nil)

	req := httptest.NewRequest("GET", apiPath, nil)
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(&http.Cookie{Name: "session_id", Value: cookie.SessionToken})
	w := httptest.NewRecorder()

	handler.Profile(w, req)

	resp := w.Result()

	require.Equal(t, 200, resp.StatusCode)
}

func TestUpdateProfileSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var logger *mw.ACLog
	apiPath := "/api/me"
	mockSes := mockS.NewMockSessionUsecaseI(ctrl)
	mockUs := mockU.NewMockUsecaseI(ctrl)
	handler := NewSessionHandler(mockSes, mockUs, logger)

	cookie := &entity.Cookie{
		UserID:       1,
		SessionToken: "TYebbYudb",
	}

	user := &entity.User{
		ID:    1,
		Email: "ani@mail.ru",
	}

	var jsonuser = map[string]interface{}{
		"Email": "ani@mail.ru",
	}

	mockSes.EXPECT().GetIDByCookie(cookie.SessionToken).Return(uint(1), nil)
	mockUs.EXPECT().UpdateUser(user).Return(nil)

	body, err := json.Marshal(jsonuser)
	if err != nil {
		return
	}

	req := httptest.NewRequest("GET", apiPath, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(&http.Cookie{Name: "session_id", Value: cookie.SessionToken})
	w := httptest.NewRecorder()

	handler.UpdateProfile(w, req)

	resp := w.Result()

	require.Equal(t, 200, resp.StatusCode)
}
