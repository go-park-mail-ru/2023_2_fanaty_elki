package usecase

import (
	mockA "server/internal/Address/repository/mock_repository"
	mockS "server/internal/Session/repository/mock_repository"
	mockU "server/internal/User/repository/mock_repository"
	"server/internal/domain/dto"
	"server/internal/domain/entity"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	//"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestLoginSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSes := mockS.NewMockSessionRepositoryI(ctrl)
	mockUs := mockU.NewMockUserRepositoryI(ctrl)
	mockAdd := mockA.NewMockAddressRepositoryI(ctrl)
	usecase := NewSessionUsecase(mockSes, mockUs, mockAdd)

	user := &entity.User{
		Username: "ania",
		Password: "anis1234",
	}

	dbuser := dto.DBGetUser{
		Username: "ania",
		Password: "anis1234",
	}

	mockUs.EXPECT().FindUserByUsername(user.Username).Return(&dbuser, nil)
	mockSes.EXPECT().Create(gomock.Any()).Return(nil)
	_, err := usecase.Login(user)
	assert.Nil(t, err)

}

func TestCheckSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSes := mockS.NewMockSessionRepositoryI(ctrl)
	mockUs := mockU.NewMockUserRepositoryI(ctrl)
	mockAdd := mockA.NewMockAddressRepositoryI(ctrl)
	usecase := NewSessionUsecase(mockSes, mockUs, mockAdd)

	sestok := "Uuehdbye"

	var UserID uint
	UserID = 1

	cookie := entity.Cookie{
		UserID:       1,
		SessionToken: "TYebbYudb",
		MaxAge:       50 * time.Hour,
	}

	dbuser := &dto.DBGetUser{
		ID:       1,
		Username: "ania",
		Password: "anis1234",
	}

	mockSes.EXPECT().Check(sestok).Return(&cookie, nil)
	mockUs.EXPECT().FindUserByID(UserID).Return(dbuser, nil)
	actual, err := usecase.Check(sestok)
	assert.Equal(t, cookie.UserID, actual)
	assert.Nil(t, err)
}

func TestLogoutSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSes := mockS.NewMockSessionRepositoryI(ctrl)
	mockUs := mockU.NewMockUserRepositoryI(ctrl)
	mockAdd := mockA.NewMockAddressRepositoryI(ctrl)
	usecase := NewSessionUsecase(mockSes, mockUs, mockAdd)

	cookie := entity.Cookie{
		UserID:       1,
		SessionToken: "TYebbYudb",
		MaxAge:       50 * time.Hour,
	}

	dbcookie := dto.DBDeleteCookie{
		SessionToken: "TYebbYudb",
	}

	mockSes.EXPECT().Delete(&dbcookie).Return(nil)
	err := usecase.Logout(&cookie)
	assert.Nil(t, err)
}

func TestGetUserProfileSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSes := mockS.NewMockSessionRepositoryI(ctrl)
	mockUs := mockU.NewMockUserRepositoryI(ctrl)
	mockAdd := mockA.NewMockAddressRepositoryI(ctrl)
	usecase := NewSessionUsecase(mockSes, mockUs, mockAdd)

	cookie := entity.Cookie{
		UserID:       1,
		SessionToken: "TYebbYudb",
		MaxAge:       50 * time.Hour,
	}

	dbuser := &dto.DBGetUser{
		ID:       1,
		Username: "ania",
		Password: "anis1234",
	}

	addresses := &dto.RespGetAddresses{
		Addresses:          []*dto.RespGetAddress{},
		CurrentAddressesID: 0,
	}

	profile := &dto.ReqGetUserProfile{
		Username:  "ania",
		Addresses: addresses.Addresses,
		Current:   0,
	}

	mockSes.EXPECT().Check(cookie.SessionToken).Return(&cookie, nil)
	mockUs.EXPECT().FindUserByID(cookie.UserID).Return(dbuser, nil)
	mockAdd.EXPECT().GetAddresses(cookie.UserID).Return(addresses, nil)
	actual, err := usecase.GetUserProfile(cookie.SessionToken)
	assert.Equal(t, profile, actual)
	assert.Nil(t, err)
}

func TestGetIdByCookieSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSes := mockS.NewMockSessionRepositoryI(ctrl)
	mockUs := mockU.NewMockUserRepositoryI(ctrl)
	mockAdd := mockA.NewMockAddressRepositoryI(ctrl)
	usecase := NewSessionUsecase(mockSes, mockUs, mockAdd)

	cookie := entity.Cookie{
		UserID:       1,
		SessionToken: "TYebbYudb",
		MaxAge:       50 * time.Hour,
	}

	mockSes.EXPECT().Check(cookie.SessionToken).Return(&cookie, nil)
	actual, err := usecase.GetIDByCookie(cookie.SessionToken)
	assert.Equal(t, cookie.UserID, actual)
	assert.Nil(t, err)
}

func TestCreateCookieAuthSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSes := mockS.NewMockSessionRepositoryI(ctrl)
	mockUs := mockU.NewMockUserRepositoryI(ctrl)
	mockAdd := mockA.NewMockAddressRepositoryI(ctrl)
	usecase := NewSessionUsecase(mockSes, mockUs, mockAdd)

	cookie := entity.Cookie{
		UserID:       1,
		SessionToken: "TYebbYudb",
		MaxAge:       50 * time.Hour,
	}

	dbuser := &dto.DBGetUser{
		ID:       1,
		Username: "ania",
		Password: "anis1234",
	}

	addresses := &dto.RespGetAddresses{
		Addresses:          []*dto.RespGetAddress{},
		CurrentAddressesID: 0,
	}

	profile := &dto.ReqGetUserProfile{
		Username:  "ania",
		Addresses: addresses.Addresses,
		Current:   0,
	}

	mockSes.EXPECT().Expire(&cookie).Return(nil)
	mockSes.EXPECT().Check(cookie.SessionToken).Return(&cookie, nil)
	mockUs.EXPECT().FindUserByID(cookie.UserID).Return(dbuser, nil)
	mockAdd.EXPECT().GetAddresses(cookie.UserID).Return(addresses, nil)
	actual, err := usecase.CreateCookieAuth(&cookie)
	assert.Equal(t, profile, actual)
	assert.Nil(t, err)
}

func TestCreateCsrfSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSes := mockS.NewMockSessionRepositoryI(ctrl)
	mockUs := mockU.NewMockUserRepositoryI(ctrl)
	mockAdd := mockA.NewMockAddressRepositoryI(ctrl)
	usecase := NewSessionUsecase(mockSes, mockUs, mockAdd)

	cookie := entity.Cookie{
		UserID:       1,
		SessionToken: "TYebbYudb",
		MaxAge:       50 * time.Hour,
	}

	mockSes.EXPECT().CreateCsrf(cookie.SessionToken, gomock.Any()).Return(nil)
	_, err := usecase.CreateCsrf(cookie.SessionToken)
	assert.Nil(t, err)
}

func TestCheckCsrfSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSes := mockS.NewMockSessionRepositoryI(ctrl)
	mockUs := mockU.NewMockUserRepositoryI(ctrl)
	mockAdd := mockA.NewMockAddressRepositoryI(ctrl)
	usecase := NewSessionUsecase(mockSes, mockUs, mockAdd)

	cookie := entity.Cookie{
		UserID:       1,
		SessionToken: "TYebbYudb",
		MaxAge:       50 * time.Hour,
	}

	rediscsfr := usecase.getCSRFHash("HBBGFCCDFG")

	mockSes.EXPECT().GetCsrf(cookie.SessionToken).Return(rediscsfr, nil)
	err := usecase.CheckCsrf(cookie.SessionToken, "HBBGFCCDFG")
	assert.Nil(t, err)
}
