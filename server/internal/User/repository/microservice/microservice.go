package microservice

import (
	"context"
	userRep "server/internal/User/repository"
	"server/internal/domain/dto"
	//"server/internal/domain/entity"
	userProto "server/proto/user"
	"fmt"
)

type UserMicroService struct {
	client userProto.UserRPCClient
}

func NewUserMicroService(client userProto.UserRPCClient) userRep.UserRepositoryI {
	return &UserMicroService {
		client: client,
	}
}

func (us *UserMicroService) FindUserById(id uint) (*dto.DBGetUser, error) {
	ctx := context.Background()

	grpcid := userProto.ID{ID: uint64(id)}

	grpcUser, err := us.client.FindUserById(ctx, &grpcid)
	if err != nil {
		return nil, err
	}
	if grpcUser.Username ==  "@" {
		fmt.Println("HFS")
		return nil, nil
	}
	return dto.ToDBGetUser(grpcUser), err
}

func(us *UserMicroService) CreateUser(user *dto.DBCreateUser) (uint, error) {
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

func (us *UserMicroService) UpdateUser(user *dto.DBUpdateUser) error {
	ctx := context.Background()

	grpcUser := dto.ToRespUpdateUser(user)

	_, err := us.client.UpdateUser(ctx, grpcUser)
	return err
}

func (us *UserMicroService) FindUserByUsername(value string) (*dto.DBGetUser, error) {
	ctx := context.Background()

	grpcValue := userProto.Value{Value: value}

	grpcUser, err := us.client.FindUserByUsername(ctx, &grpcValue)
	fmt.Println("grpcUser ", grpcUser)
	fmt.Println("grpcUser err", err)
	if err != nil {
		return nil, err
	}
	
	if grpcUser.Username ==  "@" {
		fmt.Println("HFS")
		return nil, nil
	}
	return dto.ToDBGetUser(grpcUser), nil
}

func (us *UserMicroService) FindUserByEmail(value string) (*dto.DBGetUser, error) {
	ctx := context.Background()

	grpcValue := userProto.Value{Value: value}

	grpcUser, err := us.client.FindUserByEmail(ctx, &grpcValue)
	if err != nil {
		return nil, err
	}
	if grpcUser.Username ==  "@" {
		fmt.Println("HFS")
		return nil, nil
	}
	return dto.ToDBGetUser(grpcUser), nil
}

func (us *UserMicroService) FindUserByPhone(value string) (*dto.DBGetUser, error) {
	ctx := context.Background()

	grpcValue := userProto.Value{Value: value}

	grpcUser, err := us.client.FindUserByPhone(ctx, &grpcValue)
	if err != nil {
		return nil, err
	}
	if grpcUser.Username ==  "@" {
		fmt.Println("HFS")
		return nil, nil
	}
	return dto.ToDBGetUser(grpcUser), nil
}