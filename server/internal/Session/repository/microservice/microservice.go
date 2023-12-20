package microservice

import (
	"context"

	sessionRep "server/internal/Session/repository"
	"server/internal/domain/dto"
	"server/internal/domain/entity"
	auth "server/proto/auth"
	"time"
)

type microService struct {
	client auth.SessionRPCClient
}

//NewMicroService creates new session repository
func NewMicroService(client auth.SessionRPCClient) sessionRep.SessionRepositoryI {
	return &microService{
		client: client,
	}
}

func (ms *microService) Create(cookie *entity.Cookie) error {
	ctx := context.Background()

	grpccookie := &auth.Cookie{
		UserID:       uint64(cookie.UserID),
		SessionToken: cookie.SessionToken,
		MaxAge:       int64(cookie.MaxAge),
	}

	_, err := ms.client.Create(ctx, grpccookie)

	if err != nil {
		return err
	}

	return nil

}

func (ms *microService) Check(sessionToken string) (*entity.Cookie, error) {
	ctx := context.Background()

	grpcsessionToken := &auth.SessionToken{Token: sessionToken}

	grpccookie, err := ms.client.Check(ctx, grpcsessionToken)

	if err != nil {
		return nil, err
	}

	if grpccookie.UserID == 0 {
		return nil, nil
	}

	cookie := entity.Cookie{
		UserID:       uint(grpccookie.UserID),
		SessionToken: grpccookie.SessionToken,
		MaxAge:       time.Duration(grpccookie.MaxAge),
	}

	return &cookie, nil
}

func (ms *microService) Delete(cookie *dto.DBDeleteCookie) error {
	ctx := context.Background()

	grpcDBDeleteCookie := &auth.DBDeleteCookie{
		SessionToken: cookie.SessionToken,
	}

	_, err := ms.client.Delete(ctx, grpcDBDeleteCookie)

	if err != nil {
		return err
	}

	return nil
}

func (ms *microService) Expire(cookie *entity.Cookie) error {
	ctx := context.Background()

	grpccookie := &auth.Cookie{
		UserID:       uint64(cookie.UserID),
		SessionToken: cookie.SessionToken,
		MaxAge:       int64(cookie.MaxAge),
	}

	_, err := ms.client.Expire(ctx, grpccookie)

	if err != nil {
		return err
	}

	return nil
}

func (ms *microService) CreateCsrf(sessionToken string, csrfToken string) error {
	ctx := context.Background()

	grpcSessionAndCsrf := &auth.SesionAndCsrf{
		SessionToken: sessionToken,
		CsrfToken:    csrfToken,
	}

	_, err := ms.client.CreateCsrf(ctx, grpcSessionAndCsrf)

	if err != nil {
		return err
	}

	return nil
}

func (ms *microService) GetCsrf(sessionToken string) (string, error) {
	ctx := context.Background()

	grpcsessionToken := &auth.SessionToken{Token: sessionToken}

	grpcCsrfToken, err := ms.client.GetCsrf(ctx, grpcsessionToken)

	if err != nil {
		return "", err
	}

	csrfToken := grpcCsrfToken.Token

	return csrfToken, nil
}
