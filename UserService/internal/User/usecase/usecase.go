package usecase

import (
	userRep "UserService/internal/User/repository"
	"UserService/internal/dto"
	user "UserService/proto"
	"fmt"
)

type UserUsecaseI interface {
	FindUserById(grpcid *user.ID) (*user.DBGetUser, error)
	CreateUser(grpcUser *user.DBCreateUser) (*user.ID, error)
	UpdateUser(grpcuser *user.DBUpdateUser) (*user.Nothing, error)
	FindUserByUsername(value *user.Value) (*user.DBGetUser, error)
	FindUserByEmail(value *user.Value) (*user.DBGetUser, error)
	FindUserByPhone(value *user.Value) (*user.DBGetUser, error)
}

type userUsecase struct {
	userRepo userRep.UserRepositoryI
}

func NewUserUsecase(userRepI userRep.UserRepositoryI) *userUsecase {
	return &userUsecase{
		userRepo: userRepI,
	}
}

func (u userUsecase) FindUserById(grpcid *user.ID) (*user.DBGetUser, error) {
	id := uint(grpcid.ID)

	user, err := u.userRepo.FindUserById(id)
	if err != nil {
		return nil, err
	}

	return dto.ToRespGetUser(user), nil	
}

func (u userUsecase) CreateUser(grpcUser *user.DBCreateUser) (*user.ID, error) {
	newUser := dto.ToDBCreateUser(grpcUser)

	id, err := u.userRepo.CreateUser(newUser)
	fmt.Println("ms create", err)
	if err != nil {
		return nil, err
	}

	grpcid := &user.ID{ID: uint64(id)}

	return grpcid, nil	
}

func (u userUsecase) UpdateUser(grpcuser *user.DBUpdateUser) (*user.Nothing, error) {
	upUser := dto.ToDBUpdateUser(grpcuser)

	err := u.userRepo.UpdateUser(upUser)
	return &user.Nothing{Dummy: true}, err
}

func (u userUsecase) FindUserByUsername(value *user.Value) (*user.DBGetUser, error) {
	username := value.Value

	fuser, err := u.userRepo.FindUserByUsername(username)
	fmt.Println("ms username err", err)
	
	if err != nil {
		return nil, err
	}
	return dto.ToRespGetUser(fuser), nil
}


func (u userUsecase) FindUserByEmail(value *user.Value) (*user.DBGetUser, error) {
	username := value.Value

	fuser, err := u.userRepo.FindUserByEmail(username)
	fmt.Println("ms email", err)
	if err != nil {
		return nil, err
	}

	return dto.ToRespGetUser(fuser), nil
}


func (u userUsecase) FindUserByPhone(value *user.Value) (*user.DBGetUser, error) {
	username := value.Value

	fuser, err := u.userRepo.FindUserByPhone(username)
	fmt.Println("ms phone", err)
	if err != nil {
		return nil, err
	}

	return dto.ToRespGetUser(fuser), nil
}