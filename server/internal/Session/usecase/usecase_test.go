package usecase

import (
	mockS "server/internal/Session/repository/mock_repository"
	mockU "server/internal/User/repository/mock_repository"
	"server/internal/domain/dto"
	"server/internal/domain/entity"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestLoginSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSes := mockS.NewMockSessionRepositoryI(ctrl)
	mockUs := mockU.NewMockUserRepositoryI(ctrl)
	usecase := NewSessionUsecase(mockSes, mockUs)

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
	usecase := NewSessionUsecase(mockSes, mockUs)

	sestok := "Uuehdbye"

	var userID uint
	userID = 1

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
	mockUs.EXPECT().FindUserById(userID).Return(dbuser, nil)
	actual, err := usecase.Check(sestok)
	assert.Equal(t, cookie.UserID, actual)
	assert.Nil(t, err)
}

func TestLogoutSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSes := mockS.NewMockSessionRepositoryI(ctrl)
	mockUs := mockU.NewMockUserRepositoryI(ctrl)
	usecase := NewSessionUsecase(mockSes, mockUs)

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
	usecase := NewSessionUsecase(mockSes, mockUs)

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

	profile := &dto.ReqGetUserProfile{
		Username: "ania",
	}

	mockSes.EXPECT().Check(cookie.SessionToken).Return(&cookie, nil)
	mockUs.EXPECT().FindUserById(cookie.UserID).Return(dbuser, nil)
	actual, err := usecase.GetUserProfile(cookie.SessionToken)
	assert.Equal(t, profile, actual)
	assert.Nil(t, err)
}

func TestGetIdByCookieSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSes := mockS.NewMockSessionRepositoryI(ctrl)
	mockUs := mockU.NewMockUserRepositoryI(ctrl)
	usecase := NewSessionUsecase(mockSes, mockUs)

	cookie := entity.Cookie{
		UserID:       1,
		SessionToken: "TYebbYudb",
		MaxAge:       50 * time.Hour,
	}

	mockSes.EXPECT().Check(cookie.SessionToken).Return(&cookie, nil)
	actual, err := usecase.GetIdByCookie(cookie.SessionToken)
	assert.Equal(t, cookie.UserID, actual)
	assert.Nil(t, err)
}

func TestCreateCookieAuthSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSes := mockS.NewMockSessionRepositoryI(ctrl)
	mockUs := mockU.NewMockUserRepositoryI(ctrl)
	usecase := NewSessionUsecase(mockSes, mockUs)

	cookie := entity.Cookie{
		UserID:       1,
		SessionToken: "TYebbYudb",
		MaxAge:       50 * time.Hour,
	}

	profile := &dto.ReqGetUserProfile{
		Username: "ania",
	}

	dbuser := &dto.DBGetUser{
		ID:       1,
		Username: "ania",
		Password: "anis1234",
	}

	mockSes.EXPECT().Expire(&cookie).Return(nil)
	mockSes.EXPECT().Check(cookie.SessionToken).Return(&cookie, nil)
	mockUs.EXPECT().FindUserById(cookie.UserID).Return(dbuser, nil)
	actual, err := usecase.CreateCookieAuth(&cookie)
	assert.Equal(t, profile, actual)
	assert.Nil(t, err)
}
