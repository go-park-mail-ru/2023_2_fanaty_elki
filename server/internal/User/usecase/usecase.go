package usecase

import (
	"mime/multipart"
	"net/url"
	"regexp"
	cartRep "server/internal/Cart/repository"
	userRep "server/internal/User/repository"
	"server/internal/domain/dto"
	"server/internal/domain/entity"
	"time"

	"github.com/minio/minio-go/v6"
)

type UsecaseI interface {
	CreateUser(new_user *entity.User) (uint, error)
	UpdateUser(newUser *entity.User) error
	UpdateAvatar(file multipart.File, filehandler *multipart.FileHeader, id uint) error
}

type userUsecase struct {
	userRepo userRep.UserRepositoryI
	cartRepo cartRep.CartRepositoryI
}

func NewUserUsecase(userRepI userRep.UserRepositoryI, cartRepI cartRep.CartRepositoryI) *userUsecase {
	return &userUsecase{
		userRepo: userRepI,
		cartRepo: cartRepI,
	}
}

func (us userUsecase) GetUserById(id uint) (*entity.User, error) {
	user, err := us.userRepo.FindUserById(id)
	if err != nil {
		return nil, err
	}
	return dto.ToEntityGetUser(user), nil
}

func (us userUsecase) CreateUser(newUser *entity.User) (uint, error) {

	err := us.checkUserFieldsCreate(newUser)

	if err != nil {
		return 0, err
	}

	_, err = us.checkUser(newUser)
	if err != nil {
		return 0, err
	}

	if newUser.Icon == "" {
		newUser.Icon = "img/defaultIcon.png"
	}
	user, err := us.userRepo.CreateUser(dto.ToRepoCreateUser(newUser))
	if err != nil {
		return 0, entity.ErrInternalServerError
	}

	_, err = us.cartRepo.CreateCart(user)
	if err != nil {
		return 0, entity.ErrInternalServerError
	}

	return user, nil
}

func (us userUsecase) UpdateUser(newUser *entity.User) error {
	err := us.checkUserFieldsUpdate(newUser)
	if err != nil {
		return err
	}

	_, err = us.checkUser(newUser)
	if err != nil {
		return err
	}

	user, err := us.GetUserById(newUser.ID)
	if err != nil {
		return err
	}
	if user != nil {
		if newUser.Username != "" {
			user.Username = newUser.Username
		}

		// if newUser.Password != "" {
		// 	user.Password = newUser.Password
		// }

		if newUser.PhoneNumber != "" {
			user.PhoneNumber = newUser.PhoneNumber
		}

		if newUser.Email != "" {
			user.Email = newUser.Email
		}

		if newUser.Icon != "" {
			user.Icon = newUser.Icon
		}
		return us.userRepo.UpdateUser(dto.ToRepoUpdateUser(user))
	}

	return entity.ErrNotFound

}

func (us userUsecase) checkUser(checkUser *entity.User) (*entity.User, error) {
	var user *dto.DBGetUser

	if checkUser.Username != "" {
		user, err := us.userRepo.FindUserByUsername(checkUser.Username)
		if err != nil {
			return nil, entity.ErrInternalServerError
		}

		if user != nil {
			return nil, entity.ErrConflictUsername
		}
	}

	if checkUser.Email != "" {

		user, err := us.userRepo.FindUserByEmail(checkUser.Email)
		if err != nil {

			return nil, entity.ErrInternalServerError
		}

		if user != nil {
			return nil, entity.ErrConflictEmail
		}
	}

	if checkUser.PhoneNumber != "" {
		user, err := us.userRepo.FindUserByPhone(checkUser.PhoneNumber)
		if err != nil {
			return nil, entity.ErrInternalServerError
		}

		if user != nil {
			return nil, entity.ErrConflictPhoneNumber
		}
	}
	return dto.ToEntityGetUser(user), nil
}

func (us userUsecase) checkUserFieldsCreate(user *entity.User) error {
	re := regexp.MustCompile(`^[a-zA-Z0-9_]{4,29}$`)
	if !re.MatchString(user.Username) {
		return entity.ErrInvalidUsername
	}

	if len(user.Password) < 8 {
		return entity.ErrInvalidPassword
	}

	re = regexp.MustCompile(`\d{4}-\d{1,2}-\d{1,2}`)
	if !re.MatchString(user.Birthday) {
		return entity.ErrInvalidBirthday
	}

	re = regexp.MustCompile(`@`)
	if user.Email == "" || !re.MatchString(user.Email) {
		return entity.ErrInvalidEmail
	}

	re = regexp.MustCompile(`^((8|\+7)[\- ]?)?(\(?\d{3}\)?[\- ]?)?[\d\- ]{6,10}$`)
	if user.PhoneNumber == "" || !re.MatchString(user.PhoneNumber) {
		return entity.ErrInvalidPhoneNumber
	}
	return nil
}

func (us userUsecase) checkUserFieldsUpdate(user *entity.User) error {
	re := regexp.MustCompile(`^[a-zA-Z0-9_]{4,29}$`)
	if len(user.Username) != 0 && !re.MatchString(user.Username) {
		return entity.ErrInvalidUsername
	}

	// if (len(user.Password) < 3 || len(user.Password) > 30) && user.Password != "" {
	// 	return entity.ErrInvalidPassword
	// }

	re = regexp.MustCompile(`@`)
	if !re.MatchString(user.Email) && user.Email != "" {
		return entity.ErrInvalidEmail
	}

	re = regexp.MustCompile(`^((8|\+7)[\- ]?)?(\(?\d{3}\)?[\- ]?)?[\d\- ]{6,10}$`)
	if !re.MatchString(user.PhoneNumber) && len(user.PhoneNumber) != 0 {
		return entity.ErrInvalidPhoneNumber
	}
	return nil
}

func (us userUsecase) UpdateAvatar(file multipart.File, filehandler *multipart.FileHeader, id uint) error {
	endpoint := "bring-give.hb.ru-msk.vkcs.cloud"
	location := "bring-give"
	accessKeyID := "s7X63TovV3DHPNCcsuhM5H"
	secretAccessKey := "fUxGkBzPBRdvfxuQWJ1urt1BFdfq85V9gfoE65drMLe"
	useSSL := true

	bucketName := "bring-give"
	objectName := filehandler.Filename
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		return entity.ErrInternalServerError
	}

	err = us.uploadFile(minioClient, bucketName, location, objectName, file, filehandler.Size)
	if err != nil {
		return entity.ErrInternalServerError
	}

	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", "attachment; filename=\""+objectName+"\"")

	presignedURL, err := minioClient.PresignedGetObject(bucketName, objectName, time.Duration(168)*time.Hour, reqParams)
	if err != nil {
		return entity.ErrInternalServerError
	}

	user, err := us.GetUserById(id)
	if err != nil {
		return err
	}

	user.Icon = presignedURL.String()

	return us.userRepo.UpdateUser(dto.ToRepoUpdateUser(user))
}

func (us userUsecase) uploadFile(minioClient *minio.Client, bucketName string, location string, objectName string, file multipart.File, filesize int64) error {
	err := minioClient.MakeBucket(bucketName, location)
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(bucketName)
		if !(errBucketExists == nil && exists) {
			return err
		}
	}

	_, err = minioClient.PutObject(bucketName, objectName, file, filesize, minio.PutObjectOptions{})
	if err != nil {
		return entity.ErrInternalServerError
	}

	return nil
}

