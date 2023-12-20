package microservice

import (
	"context"
	userRep "server/internal/User/repository"
	"server/internal/domain/dto"
	//"server/internal/domain/entity"
	"fmt"
	userProto "server/proto/user"
)

//UserMicroService provides management with user microservice
type UserMicroService struct {
	client userProto.UserRPCClient
}

//NewUserMicroService creates new UserRepository interface
func NewUserMicroService(client userProto.UserRPCClient) userRep.UserRepositoryI {
	return &UserMicroService{
		client: client,
	}
}

//FindUserByID finds user by id in db
func (us *UserMicroService) FindUserByID(id uint) (*dto.DBGetUser, error) {
	ctx := context.Background()

	grpcid := userProto.ID{ID: uint64(id)}

	grpcUser, err := us.client.FindUserByID(ctx, &grpcid)
	if err != nil {
		return nil, err
	}
	if grpcUser.Username == "@" {
		fmt.Println("HFS")
		return nil, nil
	}
	return dto.ToDBGetUser(grpcUser), err
}

//CreateUser creates user in db
func (us *UserMicroService) CreateUser(user *dto.DBCreateUser) (uint, error) {
	ctx := context.Background()

	grpcUser := dto.ToRespCreateUser(user)

	grpcid, err := us.client.CreateUser(ctx, grpcUser)
	fmt.Println("cre rep err", err)
	fmt.Println("cre rep id", grpcid)
	if err != nil {
		return 0, err
	}

	return uint(grpcid.ID), nil
}

//UpdateUser updates user in db
func (us *UserMicroService) UpdateUser(user *dto.DBUpdateUser) error {
	ctx := context.Background()

	grpcUser := dto.ToRespUpdateUser(user)

	_, err := us.client.UpdateUser(ctx, grpcUser)
	return err
}

//FindUserByUsername finds user by username in db
func (us *UserMicroService) FindUserByUsername(value string) (*dto.DBGetUser, error) {
	ctx := context.Background()

	grpcValue := userProto.Value{Value: value}

	grpcUser, err := us.client.FindUserByUsername(ctx, &grpcValue)
	fmt.Println("grpcUser ", grpcUser)
	fmt.Println("grpcUser err", err)
	if err != nil {
		return nil, err
	}

	if grpcUser.Username == "@" {
		fmt.Println("HFS")
		return nil, nil
	}
	return dto.ToDBGetUser(grpcUser), nil
}

//FindUserByEmail finds user by email in db
func (us *UserMicroService) FindUserByEmail(value string) (*dto.DBGetUser, error) {
	ctx := context.Background()

	grpcValue := userProto.Value{Value: value}

	grpcUser, err := us.client.FindUserByEmail(ctx, &grpcValue)
	if err != nil {
		return nil, err
	}
	if grpcUser.Username == "@" {
		fmt.Println("HFS")
		return nil, nil
	}
	return dto.ToDBGetUser(grpcUser), nil
}

//FindUserByPhone finds user by phone number in db
func (us *UserMicroService) FindUserByPhone(value string) (*dto.DBGetUser, error) {
	ctx := context.Background()

	grpcValue := userProto.Value{Value: value}

	grpcUser, err := us.client.FindUserByPhone(ctx, &grpcValue)
	if err != nil {
		return nil, err
	}
	if grpcUser.Username == "@" {
		fmt.Println("HFS")
		return nil, nil
	}
	return dto.ToDBGetUser(grpcUser), nil
}
