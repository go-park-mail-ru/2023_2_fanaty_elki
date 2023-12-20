package delivery

import (
	sessionUsecase "AuthService/internal/usecase"
	auth "AuthService/proto"
	"context"
)

//AuthManager struct
type AuthManager struct {
	auth.UnimplementedSessionRPCServer
	SessionUC sessionUsecase.SessionUsecaseI
}

//NewAuthManager creates new session rpc server
func NewAuthManager(uc sessionUsecase.SessionUsecaseI) auth.SessionRPCServer {
	return AuthManager{SessionUC: uc}
}

//Create handles create session request
func (am AuthManager) Create(ctx context.Context, grpccookie *auth.Cookie) (*auth.Nothing, error) {
	resp, err := am.SessionUC.Create(grpccookie)
	return resp, err
}

//Check handles check session request
func (am AuthManager) Check(ctx context.Context, grpcsessionToken *auth.SessionToken) (*auth.Cookie, error) {
	resp, err := am.SessionUC.Check(grpcsessionToken)
	return resp, err
}

//Delete handles Delete session request
func (am AuthManager) Delete(ctx context.Context, grpccookie *auth.DBDeleteCookie) (*auth.Nothing, error) {
	resp, err := am.SessionUC.Delete(grpccookie)
	return resp, err
}

//Expire handles update session request
func (am AuthManager) Expire(ctx context.Context, grpccookie *auth.Cookie) (*auth.Nothing, error) {
	resp, err := am.SessionUC.Expire(grpccookie)
	return resp, err
}

//CreateCsrf handles create csrf request
func (am AuthManager) CreateCsrf(ctx context.Context, grpcSessionAndCsrf *auth.SesionAndCsrf) (*auth.Nothing, error) {
	resp, err := am.SessionUC.CreateCsrf(grpcSessionAndCsrf)
	return resp, err
}

//GetCsrf handles get csrf request
func (am AuthManager) GetCsrf(ctx context.Context, grpcsessionToken *auth.SessionToken) (*auth.CsrfToken, error) {
	resp, err := am.SessionUC.GetCsrf(grpcsessionToken)
	return resp, err
}
