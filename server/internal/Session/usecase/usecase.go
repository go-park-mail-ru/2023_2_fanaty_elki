package usecase

import (
	"server/internal/domain/entity"
	"server/internal/domain/dto"
	sessionRep "server/internal/Session/repository"
	userRep "server/internal/User/repository"
	"time"
	"math/rand"
)

const sessKeyLen = 10

var (
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

type UsecaseI interface {
	Login(user *entity.User) (*entity.Cookie, error)
	Check(SessionToken string) (*string, error)
	Logout(cookie *entity.Cookie) error
	GetUserProfile(sessionToken string) (*dto.ReqGetUserProfile, error)
	GetIdByCookie(SessionToken string) (uint, error)
}

type sessionUsecase struct {
	sessionRepo sessionRep.SessionRepositoryI
	userRepo userRep.UserRepositoryI
}

func NewSessionUsecase(sessionRep sessionRep.SessionRepositoryI, userRep userRep.UserRepositoryI) *sessionUsecase {
	return &sessionUsecase{
		sessionRepo: sessionRep,
		userRepo: userRep,
	}
}

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func (ss sessionUsecase) Login(user *entity.User) (*entity.Cookie, error) {
	
	us, err := ss.userRepo.FindUserBy("Username", user.Username)
	
	if err != nil {
		return nil, err
	}
	
	if us == nil {
		return nil, entity.ErrBadRequest
	}

	if user.Password != us.Password {
		return nil, entity.ErrBadRequest
	}

	cookie := &entity.Cookie{
		UserID: us.ID,
		SessionToken: randStringRunes(sessKeyLen),
		MaxAge: 50 * time.Hour,	
	}

	err = ss.sessionRepo.Create(cookie)

	if err != nil {
		return nil, err
	}

	return cookie, nil

}

func (ss sessionUsecase) Check(SessionToken string) (*string, error) {
	
	cookie, err := ss.sessionRepo.Check(SessionToken)
	if err != nil {
		return nil, err
	}

	user, err := ss.userRepo.GetUserById(cookie.UserID)
	if err != nil {
		return nil, err
	}

	return &user.Username, nil
}

func (ss sessionUsecase) Logout(cookie *entity.Cookie) error {
	return ss.sessionRepo.Delete(dto.ToDBDeleteCookie(cookie))	
}

func (ss sessionUsecase) GetUserProfile(sessionToken string) (*dto.ReqGetUserProfile, error) {
	cookie, err := ss.sessionRepo.Check(sessionToken)
	if err != nil {
		return nil, err
	}

	user, err := ss.userRepo.GetUserById(cookie.UserID)
	if err != nil{
		return nil, err
	}
	
	return dto.ToReqGetUserProfile(user), nil
}

func (ss sessionUsecase) GetIdByCookie(SessionToken string) (uint, error) {
	
	cookie, err := ss.sessionRepo.Check(SessionToken)
	if err != nil {
		return 0, err
	}

	return cookie.UserID, nil
}
