package usecase

import (
	//"errors"
	"database/sql"
	mockC "server/internal/Cart/repository/mock_repository"
	mockU "server/internal/User/repository/mock_repository"
	"server/internal/domain/dto"
	"server/internal/domain/entity"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetUserByIdSucces(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUs := mockU.NewMockUserRepositoryI(ctrl)
	mockCart := mockC.NewMockCartRepositoryI(ctrl)
	usecase := NewUserUsecase(mockUs, mockCart)

	dbuser := &dto.DBGetUser{
		ID:          1,
		Username:    "Иван Иванович",
		Password:    "secure_password",
		Birthday:    sql.NullString{Valid: true, String: "1995-04-04"},
		PhoneNumber: "+7 916 534-23-99",
		Email:       "john@example.com",
		Icon:        sql.NullString{Valid: true, String: "dificon"},
	}

	user := &entity.User{
		ID:          1,
		Username:    "Иван Иванович",
		Password:    "secure_password",
		PhoneNumber: "+7 916 534-23-99",
		Birthday:    "1995-04-04",
		Email:       "john@example.com",
		Icon:        "dificon",
	}

	mockUs.EXPECT().FindUserByID(uint(1)).Return(dbuser, nil)
	actual, err := usecase.GetUserByID(uint(1))
	assert.Equal(t, user, actual)
	assert.Nil(t, err)

}

func TestCreateUserSucces(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUs := mockU.NewMockUserRepositoryI(ctrl)
	mockCart := mockC.NewMockCartRepositoryI(ctrl)
	usecase := NewUserUsecase(mockUs, mockCart)

	dbuser := &dto.DBCreateUser{
		ID:          1,
		Username:    "Иван Иванович",
		Password:    "secure_password",
		Birthday:    sql.NullString{Valid: true, String: "1995-04-04"},
		PhoneNumber: "+7 916 534-23-99",
		Email:       "john@example.com",
		Icon:        sql.NullString{Valid: true, String: "dificon"},
	}

	user := &entity.User{
		ID:          1,
		Username:    "Иван Иванович",
		Password:    "secure_password",
		PhoneNumber: "+7 916 534-23-99",
		Birthday:    "1995-04-04",
		Email:       "john@example.com",
		Icon:        "dificon",
	}

	mockUs.EXPECT().FindUserByUsername(dbuser.Username).Return(nil, nil)
	mockUs.EXPECT().FindUserByEmail(dbuser.Email).Return(nil, nil)
	mockUs.EXPECT().FindUserByPhone(dbuser.PhoneNumber).Return(nil, nil)
	mockUs.EXPECT().CreateUser(dbuser).Return(uint(1), nil)
	mockCart.EXPECT().CreateCart(uint(1)).Return(uint(1), nil)
	actual, err := usecase.CreateUser(user)
	assert.Equal(t, uint(1), actual)
	assert.Nil(t, err)
}

func TestUpdateUserSucces(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUs := mockU.NewMockUserRepositoryI(ctrl)
	mockCart := mockC.NewMockCartRepositoryI(ctrl)
	usecase := NewUserUsecase(mockUs, mockCart)

	dbuser := &dto.DBUpdateUser{
		ID:          1,
		Username:    "Иван Иванович",
		Password:    "secure_password",
		Birthday:    sql.NullString{Valid: true, String: "1995-04-04"},
		PhoneNumber: "+7 916 534-23-99",
		Email:       "john@example.com",
		Icon:        sql.NullString{Valid: true, String: "dificon"},
	}

	dbgetuser := &dto.DBGetUser{
		ID:          1,
		Username:    "Иван Иванович",
		Password:    "secure_password",
		Birthday:    sql.NullString{Valid: true, String: "1995-04-04"},
		PhoneNumber: "+7 916 534-23-99",
		Email:       "john@example.com",
		Icon:        sql.NullString{Valid: true, String: "dificon"},
	}

	user := &entity.User{
		ID:          1,
		Username:    "Иван Иванович",
		Password:    "secure_password",
		PhoneNumber: "+7 916 534-23-99",
		Birthday:    "1995-04-04",
		Email:       "john@example.com",
		Icon:        "dificon",
	}

	mockUs.EXPECT().FindUserByUsername(dbuser.Username).Return(nil, nil)
	mockUs.EXPECT().FindUserByEmail(dbuser.Email).Return(nil, nil)
	mockUs.EXPECT().FindUserByPhone(dbuser.PhoneNumber).Return(nil, nil)
	mockUs.EXPECT().FindUserByID(uint(1)).Return(dbgetuser, nil)
	mockUs.EXPECT().UpdateUser(dbuser).Return(nil)
	err := usecase.UpdateUser(user)
	assert.Nil(t, err)

}
