package usecase

import (
	userRep "UserService/internal/User/repository"
	"UserService/internal/dto"
	user "UserService/proto"
	"fmt"
)

//UserUsecaseI interface
type UserUsecaseI interface {
	FindUserByID(grpcid *user.ID) (*user.DBGetUser, error)
	CreateUser(grpcUser *user.DBCreateUser) (*user.ID, error)
	UpdateUser(grpcuser *user.DBUpdateUser) (*user.Nothing, error)
	FindUserByUsername(value *user.Value) (*user.DBGetUser, error)
	FindUserByEmail(value *user.Value) (*user.DBGetUser, error)
	FindUserByPhone(value *user.Value) (*user.DBGetUser, error)
}

//UserUsecase struct
type UserUsecase struct {
	userRepo userRep.UserRepositoryI
}

//NewUserUsecase creates user usecase
func NewUserUsecase(userRepI userRep.UserRepositoryI) *UserUsecase {
	return &UserUsecase{
		userRepo: userRepI,
	}
}

//FindUserByID finds user by id
func (u UserUsecase) FindUserByID(grpcid *user.ID) (*user.DBGetUser, error) {
	id := uint(grpcid.ID)

	user, err := u.userRepo.FindUserByID(id)
	if err != nil {
		return nil, err
	}

	return dto.ToRespGetUser(user), nil	
}

//CreateUser creates user
func (u UserUsecase) CreateUser(grpcUser *user.DBCreateUser) (*user.ID, error) {
	newUser := dto.ToDBCreateUser(grpcUser)

	id, err := u.userRepo.CreateUser(newUser)
	fmt.Println("ms create", err)
	if err != nil {
		return nil, err
	}

	grpcid := &user.ID{ID: uint64(id)}

	return grpcid, nil	
}

//UpdateUser updates user 
func (u UserUsecase) UpdateUser(grpcuser *user.DBUpdateUser) (*user.Nothing, error) {
	upUser := dto.ToDBUpdateUser(grpcuser)

	err := u.userRepo.UpdateUser(upUser)
	return &user.Nothing{Dummy: true}, err
}

//FindUserByUsername finds user by username
func (u UserUsecase) FindUserByUsername(value *user.Value) (*user.DBGetUser, error) {
	username := value.Value

	fuser, err := u.userRepo.FindUserByUsername(username)
	fmt.Println("ms username err", err)
	
	if err != nil {
		return nil, err
	}
	return dto.ToRespGetUser(fuser), nil
}

//FindUserByEmail finds user by email
func (u UserUsecase) FindUserByEmail(value *user.Value) (*user.DBGetUser, error) {
	username := value.Value

	fuser, err := u.userRepo.FindUserByEmail(username)
	fmt.Println("ms email", err)
	if err != nil {
		return nil, err
	}

	return dto.ToRespGetUser(fuser), nil
}

//FindUserByPhone finds user by phone
func (u UserUsecase) FindUserByPhone(value *user.Value) (*user.DBGetUser, error) {
	username := value.Value

	fuser, err := u.userRepo.FindUserByPhone(username)
	fmt.Println("ms phone", err)
	if err != nil {
		return nil, err
	}

	return dto.ToRespGetUser(fuser), nil
}