package usecase

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github.com/microcosm-cc/bluemonday"
	"math/rand"
	sessionRep "server/internal/Session/repository"
	userRep "server/internal/User/repository"
	"server/internal/domain/dto"
	"server/internal/domain/entity"
	"time"
)

const sessKeyLen = 10

var (
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

type UsecaseI interface {
	Login(user *entity.User) (*entity.Cookie, error)
	Check(SessionToken string) (uint, error)
	Logout(cookie *entity.Cookie) error
	GetUserProfile(sessionToken string) (*dto.ReqGetUserProfile, error)
	GetIdByCookie(SessionToken string) (uint, error)
	CreateCookieAuth(cookie *entity.Cookie) (*dto.ReqGetUserProfile, error)
	CheckCsrf(sessionToken string, csrfToken string) error
	CreateCsrf(sessionToken string) (string, error)
}

type sessionUsecase struct {
	sessionRepo sessionRep.SessionRepositoryI
	userRepo    userRep.UserRepositoryI
	sanitizer   *bluemonday.Policy
}

func NewSessionUsecase(sessionRep sessionRep.SessionRepositoryI, userRep userRep.UserRepositoryI) *sessionUsecase {
	sanitizer := bluemonday.UGCPolicy()
	return &sessionUsecase{
		sessionRepo: sessionRep,
		userRepo:    userRep,
		sanitizer:   sanitizer,
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

	us, err := ss.userRepo.FindUserByUsername(user.Username)

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
		UserID:       us.ID,
		SessionToken: randStringRunes(sessKeyLen),
		MaxAge:       150 * time.Hour,
	}

	err = ss.sessionRepo.Create(cookie)
	if err != nil {
		return nil, err
	}

	return cookie, nil

}

func (ss sessionUsecase) Check(SessionToken string) (uint, error) {

	cookie, err := ss.sessionRepo.Check(SessionToken)
	if err != nil {
		return 0, err
	}
	if cookie == nil {
		return 0, nil
	}
	user, err := ss.userRepo.FindUserById(cookie.UserID)
	if err != nil {
		return 0, err
	}

	if user == nil {
		return 0, nil
	}
	return user.ID, nil
}

func (ss sessionUsecase) Logout(cookie *entity.Cookie) error {
	return ss.sessionRepo.Delete(dto.ToDBDeleteCookie(cookie))
}

func (ss sessionUsecase) GetUserProfile(sessionToken string) (*dto.ReqGetUserProfile, error) {
	cookie, err := ss.sessionRepo.Check(sessionToken)
	if err != nil {
		return nil, err
	}

	user, err := ss.userRepo.FindUserById(cookie.UserID)
	if err != nil {
		return nil, err
	}

	reqUser := dto.ToReqGetUserProfile(user)
	reqUser.Email = ss.sanitizer.Sanitize(reqUser.Email)
	reqUser.Birthday = ss.sanitizer.Sanitize(reqUser.Birthday)
	reqUser.Icon = ss.sanitizer.Sanitize(reqUser.Icon)
	reqUser.Username = ss.sanitizer.Sanitize(reqUser.Username)
	reqUser.PhoneNumber = ss.sanitizer.Sanitize(reqUser.PhoneNumber)

	return reqUser, nil
}

func (ss sessionUsecase) GetIdByCookie(SessionToken string) (uint, error) {

	cookie, err := ss.sessionRepo.Check(SessionToken)
	if err != nil || cookie == nil {
		return 0, err
	}

	return cookie.UserID, nil

}

func (ss sessionUsecase) CreateCookieAuth(cookie *entity.Cookie) (*dto.ReqGetUserProfile, error) {
	err := ss.sessionRepo.Expire(cookie)
	if err != nil {
		return nil, err
	}
	return ss.GetUserProfile(cookie.SessionToken)
}

func (ss sessionUsecase) CreateCsrf(sessionToken string) (string, error) {
	csrfToken := randStringRunes(10)
	redisCSRFToken := ss.getCSRFHash(csrfToken)
	err := ss.sessionRepo.CreateCsrf(sessionToken, redisCSRFToken)
	if err != nil {
		return "", err
	}
	return csrfToken, nil
}

func (ss sessionUsecase) CheckCsrf(sessionToken string, csrfToken string) error {
	redisCsrfToken, err := ss.sessionRepo.GetCsrf(sessionToken)
	hash := ss.getCSRFHash(csrfToken)

	if err != nil {
		return err
	}

	if redisCsrfToken == "" || hash != redisCsrfToken {
		return entity.ErrFailCSRF
	}

	return nil
}

func (ss sessionUsecase) getCSRFHash(csrfToken string) string {
	salt := "KOCTbILbSalt"
	hash := hmac.New(sha256.New, []byte(salt))
	hash.Write([]byte(csrfToken))
	hashInBytes := hash.Sum(nil)

	return hex.EncodeToString(hashInBytes)
}
