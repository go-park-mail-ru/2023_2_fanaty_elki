package usecase

import (
	"AuthService/dto"
	"AuthService/entity"
	sessionRep "AuthService/internal/repository"
	auth "AuthService/proto"
	"time"
)

type SessionUsecaseI interface {
	Create(grpccookie *auth.Cookie) (*auth.Nothing, error)
	Check(grpcsessionToken *auth.SessionToken) (*auth.Cookie, error)
	Delete(grpccookie *auth.DBDeleteCookie) (*auth.Nothing, error)
	Expire(grpccookie *auth.Cookie) (*auth.Nothing, error)
	CreateCsrf(grpcSessionAndCsrf *auth.SesionAndCsrf) (*auth.Nothing, error)
	GetCsrf(grpcsessionToken *auth.SessionToken) (*auth.CsrfToken, error)
}

type sessionUsecase struct {
	sessionRepo sessionRep.SessionRepositoryI
}

func NewSessionUsecase(sessionRep sessionRep.SessionRepositoryI) *sessionUsecase {
	return &sessionUsecase{
		sessionRepo: sessionRep,
	}
}

func (su sessionUsecase) Create(grpccookie *auth.Cookie) (*auth.Nothing, error) {
	cookie := &entity.Cookie{
		UserID:       uint(grpccookie.UserID),
		SessionToken: grpccookie.SessionToken,
		MaxAge:       time.Duration(grpccookie.MaxAge),
	}

	err := su.sessionRepo.Create(cookie)
	if err != nil {
		return nil, err
	}

	return &auth.Nothing{Dummy: true}, nil
}

func (su sessionUsecase) Check(grpcsessionToken *auth.SessionToken) (*auth.Cookie, error) {
	cookie, err := su.sessionRepo.Check(grpcsessionToken.Token)

	if err != nil {
		return nil, err
	}

	grpccookie := &auth.Cookie{
		UserID:       uint64(cookie.UserID),
		SessionToken: cookie.SessionToken,
		MaxAge:       int64(cookie.MaxAge),
	}

	return grpccookie, nil
}

func (su sessionUsecase) Delete(grpccookie *auth.DBDeleteCookie) (*auth.Nothing, error) {
	cookie := &dto.DBDeleteCookie{
		SessionToken: grpccookie.SessionToken,
	}

	err := su.sessionRepo.Delete(cookie)

	if err != nil {
		return nil, err
	}

	return &auth.Nothing{Dummy: true}, nil
}

func (su sessionUsecase) Expire(grpccookie *auth.Cookie) (*auth.Nothing, error) {
	cookie := &entity.Cookie{
		UserID:       uint(grpccookie.UserID),
		SessionToken: grpccookie.SessionToken,
		MaxAge:       time.Duration(grpccookie.MaxAge),
	}

	err := su.sessionRepo.Expire(cookie)

	if err != nil {
		return nil, err
	}

	return &auth.Nothing{Dummy: true}, nil
}

func (su sessionUsecase) CreateCsrf(grpcSessionAndCsrf *auth.SesionAndCsrf) (*auth.Nothing, error) {
	sessionToken := grpcSessionAndCsrf.SessionToken
	csrfToken := grpcSessionAndCsrf.CsrfToken

	err := su.sessionRepo.CreateCsrf(sessionToken, csrfToken)

	if err != nil {
		return nil, err
	}

	return &auth.Nothing{Dummy: true}, nil
}

func (su sessionUsecase) GetCsrf(grpcsessionToken *auth.SessionToken) (*auth.CsrfToken, error) {
	sessionToken := grpcsessionToken.Token

	csrfToken, err := su.sessionRepo.GetCsrf(sessionToken)

	if err != nil {
		return nil, err
	}

	return &auth.CsrfToken{Token: csrfToken}, nil
}
