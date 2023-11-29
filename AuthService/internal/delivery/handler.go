package delivery

import (
	sessionUsecase "AuthService/internal/usecase"
	auth "AuthService/proto"
	"context"
)

type AuthManager struct {
	auth.UnimplementedSessionRPCServer
	SessionUC sessionUsecase.SessionUsecaseI
}

func NewAuthManager(uc sessionUsecase.SessionUsecaseI) auth.SessionRPCServer {
	return AuthManager{SessionUC: uc}
}

func (am AuthManager) Create(ctx context.Context, grpccookie *auth.Cookie) (*auth.Nothing, error) {
	resp, err := am.SessionUC.Create(grpccookie)
	return resp, err
}

func (am AuthManager) Check(ctx context.Context, grpcsessionToken *auth.SessionToken) (*auth.Cookie, error) {
	resp, err := am.SessionUC.Check(grpcsessionToken)
	return resp, err
}

func (am AuthManager) Delete(ctx context.Context, grpccookie *auth.DBDeleteCookie) (*auth.Nothing, error) {
	resp, err := am.SessionUC.Delete(grpccookie)
	return resp, err
}

func (am AuthManager) Expire(ctx context.Context, grpccookie *auth.Cookie) (*auth.Nothing, error) {
	resp, err := am.SessionUC.Expire(grpccookie)
	return resp, err
}

func (am AuthManager) CreateCsrf(ctx context.Context, grpcSessionAndCsrf *auth.SesionAndCsrf) (*auth.Nothing, error) {
	resp, err := am.SessionUC.CreateCsrf(grpcSessionAndCsrf)
	return resp, err
}

func (am AuthManager) GetCsrf(ctx context.Context, grpcsessionToken *auth.SessionToken) (*auth.CsrfToken, error) {
	resp, err := am.SessionUC.GetCsrf(grpcsessionToken)
	return resp, err
}
