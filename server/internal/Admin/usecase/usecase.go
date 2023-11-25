package usecase

import (
	//"fmt"
	"math/rand"
	adminRep "server/internal/Admin/repository"
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

type UsecaseI interface {
	Login(user *entity.Admin) (*entity.Cookie, error)
	Check(SessionToken string) (uint, error)
	Logout(cookie *entity.Cookie) error
	CreateCookieAuth(cookie *entity.Cookie) (*dto.ReqAuthAdmin, error)
}

type adminUsecase struct {
	adminRepo 	adminRep.AdminRepositoryI
	userRepo    userRep.UserRepositoryI
	sanitizer   *bluemonday.Policy
}

func NewadminUsecase(adminRep adminRep.AdminRepositoryI, userRep userRep.UserRepositoryI) *adminUsecase {
	sanitizer := bluemonday.UGCPolicy()
	return &adminUsecase{
		adminRepo: adminRep,
		userRepo:    userRep,
		sanitizer: sanitizer,
	}
}

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func (ss adminUsecase) Login(user *entity.Admin) (*entity.Cookie, error) {

	us, err := ss.userRepo.GetAdminByUsername(user.Username)
	
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
		UserID:       us.Id,
		SessionToken: randStringRunes(sessKeyLen),
		MaxAge:       150 * time.Hour,
	}

	err = ss.adminRepo.Create(cookie)

	if err != nil {
		return nil, err
	}

	return cookie, nil

}

func (ss adminUsecase) Check(SessionToken string) (uint, error) {
	
	cookie, err := ss.adminRepo.Check(SessionToken)
	
	if err != nil {
		return 0, err
	}
	
	if cookie == nil {
		return 0, nil
	}
	
	user, err := ss.userRepo.GetAdminById(cookie.UserID)
	if err != nil {
		return 0, err
	}
	
	if user == nil {
		return 0, nil
	}
	return user.Id, nil
}

func (ss adminUsecase) Logout(cookie *entity.Cookie) error {
	return ss.adminRepo.Delete(dto.ToDBDeleteCookie(cookie))	
}

func (ss adminUsecase) CreateCookieAuth(cookie *entity.Cookie) (*dto.ReqAuthAdmin, error) {
	err := ss.adminRepo.Expire(cookie)
	if err != nil {
		return nil, err
	}
	cook, err := ss.adminRepo.Check(cookie.SessionToken)
	if err != nil {
		return nil, err
	}
	admin, err := ss.userRepo.GetAdminById(cook.UserID)
	if err != nil {
		return nil, err
	}
	reqAd := &dto.ReqAuthAdmin{Username: admin.Username}
	return reqAd, nil
}