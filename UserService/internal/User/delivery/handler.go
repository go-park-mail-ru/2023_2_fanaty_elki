package delivery

import (
	userUsecase "UserService/internal/User/usecase"
	proto "UserService/proto"
	"context"
	"fmt"
)

type UserManager struct {
	proto.UnimplementedUserRPCServer
	userUC userUsecase.UserUsecaseI
}

func NewUserManager(uc userUsecase.UserUsecaseI) proto.UserRPCServer {
	return UserManager{
		userUC: uc,
	}
}

func (us UserManager) FindUserById(ctx context.Context, grpcid *proto.ID) (*proto.DBGetUser, error) {
	resp, err := us.userUC.FindUserById(grpcid)
	if resp == nil {
		return &proto.DBGetUser{Username: "@"}, nil
	}
	fmt.Println("resp", resp)
	return resp, err
}

func (us UserManager) CreateUser(ctx context.Context, grpcUser *proto.DBCreateUser) (*proto.ID, error) {
	resp, err := us.userUC.CreateUser(grpcUser)
	fmt.Println("resp cr us", resp)
	fmt.Println("err cr us", resp)
	return resp, err
}

func (us UserManager) UpdateUser(ctx context.Context, grpcuser *proto.DBUpdateUser) (*proto.Nothing, error) {
	resp, err := us.userUC.UpdateUser(grpcuser) 
	return resp, err
}

func (us UserManager) FindUserByUsername(ctx context.Context, value *proto.Value) (*proto.DBGetUser, error){
	resp, err := us.userUC.FindUserByUsername(value) 
	if resp == nil {
		return &proto.DBGetUser{Username: "@"}, nil
	}
	fmt.Println("resp", resp)
	return resp, err
}

func (us UserManager) FindUserByEmail(ctx context.Context, value *proto.Value) (*proto.DBGetUser, error){
	resp, err := us.userUC.FindUserByEmail(value) 
	if resp == nil {
		return &proto.DBGetUser{Username: "@"}, nil
	}
	return resp, err
}

func (us UserManager) FindUserByPhone(ctx context.Context, value *proto.Value) (*proto.DBGetUser, error){
	resp, err := us.userUC.FindUserByPhone(value) 
	fmt.Println("phone resp", resp)
	if resp == nil {
		return &proto.DBGetUser{Username: "@"}, nil
	}
	return resp, err
}
