package usecase

import (
	"math/rand"
	"regexp"
	sessionRep "server/internal/Session/repository"
	userRep "server/internal/User/repository"
	"server/internal/domain/entity"
	"time"
)

const sessKeyLen = 10

var (
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

type UsecaseI interface {
	SignUp(user *entity.User) (uint, error)
	Login(user *entity.User) (*entity.Cookie, error)
	Check(SessionToken string) (*string, error)
	Logout(cookie *entity.Cookie) error
}

type sessionUsecase struct {
	sessionRepo sessionRep.SessionRepositoryI
	userRepo    userRep.UserRepositoryI
}

func NewSessionUsecase(sessionRep sessionRep.SessionRepositoryI, userRep userRep.UserRepositoryI) *sessionUsecase {
	return &sessionUsecase{
		sessionRepo: sessionRep,
		userRepo:    userRep,
	}
}

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func (ss sessionUsecase) SignUp(user *entity.User) (uint, error) {
	if len(user.Username) < 3 || len(user.Username) > 30 {
		return 0, entity.ErrInvalidUsername
	}

	if len(user.Username) < 3 || len(user.Username) > 50 {
		return 0, entity.ErrInvalidPassword
	}

	re := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
	if user.Birthday.String != "" && !re.MatchString(user.Birthday.String) {
		return 0, entity.ErrInvalidBirthday
	}

	re = regexp.MustCompile(`^[+]?[0-9]{3,25}$`)
	if user.PhoneNumber != "" && !re.MatchString(user.PhoneNumber) {
		return 0, entity.ErrConflictPhoneNumber
	}

	re = regexp.MustCompile(`\S*@\S*`)
	if !re.MatchString(user.Email) {
		return 0, entity.ErrInvalidEmail
	}

	us, err := ss.userRepo.FindUserBy("Username", user.Username)
	if err != nil {
		return 0, err
	}

	if us != nil {
		return 0, entity.ErrConflictUsername
	}

	us, err = ss.userRepo.FindUserBy("Email", user.Email)
	if err != nil {
		return 0, err
	}

	if us != nil {
		return 0, entity.ErrConflictEmail
	}

	us, err = ss.userRepo.FindUserBy("PhoneNumber", user.PhoneNumber)
	if err != nil {
		return 0, err
	}

	if us != nil {
		return 0, entity.ErrConflictPhoneNumber
	}

	return ss.userRepo.CreateUser(user)
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
		UserID:       us.ID,
		SessionToken: randStringRunes(sessKeyLen),
		MaxAge:       50 * time.Hour,
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
		//fmt.Println(err.Error())
		return nil, err
	}

	user, err := ss.userRepo.GetUserById(cookie.UserID)
	if err != nil {
		//fmt.Println(err.Error())
		return nil, err
	}

	return &user.Username, nil
}

func (ss sessionUsecase) Logout(cookie *entity.Cookie) error {
	return ss.sessionRepo.Delete(cookie)

}
