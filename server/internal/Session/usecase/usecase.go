package usecase

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	addressRep "server/internal/Address/repository"
	sessionRep "server/internal/Session/repository"
	userRep "server/internal/User/repository"
	"server/internal/domain/dto"
	"server/internal/domain/entity"
	"time"

	"github.com/microcosm-cc/bluemonday"
)

const sessKeyLen = 10

var (
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

//SessionUsecaseI interface of session usecase
type SessionUsecaseI interface {
	Login(user *entity.User) (*entity.Cookie, error)
	Check(SessionToken string) (uint, error)
	Logout(cookie *entity.Cookie) error
	GetUserProfile(sessionToken string) (*dto.ReqGetUserProfile, error)
	GetIDByCookie(SessionToken string) (uint, error)
	CreateCookieAuth(cookie *entity.Cookie) (*dto.ReqGetUserProfile, error)
	CheckCsrf(sessionToken string, csrfToken string) error
	CreateCsrf(sessionToken string) (string, error)
}

//SessionUsecase manage sessions
type SessionUsecase struct {
	sessionRepo sessionRep.SessionRepositoryI
	userRepo    userRep.UserRepositoryI
	addressRepo addressRep.AddressRepositoryI
	sanitizer   *bluemonday.Policy
}

//NewSessionUsecase creates new session usecase object
func NewSessionUsecase(sessionRep sessionRep.SessionRepositoryI, userRep userRep.UserRepositoryI, addressRep addressRep.AddressRepositoryI) *SessionUsecase {
	sanitizer := bluemonday.UGCPolicy()
	return &SessionUsecase{
		sessionRepo: sessionRep,
		userRepo:    userRep,
		addressRepo: addressRep,
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

//Login creates session 
func (ss SessionUsecase) Login(user *entity.User) (*entity.Cookie, error) {

	us, err := ss.userRepo.FindUserByUsername(user.Username)
	fmt.Println("login err", err)
	fmt.Println("login us", us, " ", us)

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

//Check checks session of user
func (ss SessionUsecase) Check(SessionToken string) (uint, error) {

	cookie, err := ss.sessionRepo.Check(SessionToken)
	if err != nil {
		return 0, err
	}
	if cookie == nil {
		return 0, nil
	}
	user, err := ss.userRepo.FindUserByID(cookie.UserID)
	if err != nil {
		return 0, err
	}

	if user == nil {
		return 0, nil
	}
	return user.ID, nil
}

//Logout deletes cookie
func (ss SessionUsecase) Logout(cookie *entity.Cookie) error {
	return ss.sessionRepo.Delete(dto.ToDBDeleteCookie(cookie))
}

//GetUserProfile gets user's profile
func (ss SessionUsecase) GetUserProfile(sessionToken string) (*dto.ReqGetUserProfile, error) {
	cookie, err := ss.sessionRepo.Check(sessionToken)
	if err != nil {
		return nil, err
	}

	user, err := ss.userRepo.FindUserByID(cookie.UserID)
	if err != nil {
		return nil, err
	}

	reqUser := dto.ToReqGetUserProfile(user)
	addresses, err := ss.addressRepo.GetAddresses(cookie.UserID)
	if err != nil {
		return nil, err
	}
	reqUser.Addresses = addresses.Addresses
	reqUser.Current = addresses.CurrentAddressesID
	reqUser.Email = ss.sanitizer.Sanitize(reqUser.Email)
	reqUser.Birthday = ss.sanitizer.Sanitize(reqUser.Birthday)
	//reqUser.Icon = ss.sanitizer.Sanitize(reqUser.Icon)
	reqUser.Username = ss.sanitizer.Sanitize(reqUser.Username)
	reqUser.PhoneNumber = ss.sanitizer.Sanitize(reqUser.PhoneNumber)

	return reqUser, nil
}

//GetIDByCookie gets id of user by cookie
func (ss SessionUsecase) GetIDByCookie(SessionToken string) (uint, error) {

	cookie, err := ss.sessionRepo.Check(SessionToken)
	if err != nil || cookie == nil {
		return 0, err
	}

	return cookie.UserID, nil

}

//CreateCookieAuth updates cookie expire
func (ss SessionUsecase) CreateCookieAuth(cookie *entity.Cookie) (*dto.ReqGetUserProfile, error) {
	err := ss.sessionRepo.Expire(cookie)
	if err != nil {
		return nil, err
	}
	return ss.GetUserProfile(cookie.SessionToken)
}

//CreateCsrf creates csrf token
func (ss SessionUsecase) CreateCsrf(sessionToken string) (string, error) {
	csrfToken := randStringRunes(10)
	redisCSRFToken := ss.getCSRFHash(csrfToken)
	err := ss.sessionRepo.CreateCsrf(sessionToken, redisCSRFToken)
	if err != nil {
		return "", err
	}
	return csrfToken, nil
}

//CheckCsrf checks csrf token
func (ss SessionUsecase) CheckCsrf(sessionToken string, csrfToken string) error {
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

func (ss SessionUsecase) getCSRFHash(csrfToken string) string {
	salt := "KOCTbILbSalt"
	hash := hmac.New(sha256.New, []byte(salt))
	hash.Write([]byte(csrfToken))
	hashInBytes := hash.Sum(nil)

	return hex.EncodeToString(hashInBytes)
}
