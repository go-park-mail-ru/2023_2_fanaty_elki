package usecase

import (
	"fmt"
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

// Iusecase is interface of user usecase
type Iusecase interface {
	CreateUser(newUser *entity.User) (uint, error)
	UpdateUser(newUser *entity.User) error
	UpdateAvatar(file multipart.File, filehandler *multipart.FileHeader, id uint) error
}

// UserUsecase provides usecase layer of User entity
type UserUsecase struct {
	userRepo userRep.UserRepositoryI
	cartRepo cartRep.CartRepositoryI
}

// NewUserUsecase creates UserUsecase object
func NewUserUsecase(userRepI userRep.UserRepositoryI, cartRepI cartRep.CartRepositoryI) *UserUsecase {
	return &UserUsecase{
		userRepo: userRepI,
		cartRepo: cartRepI,
	}
}

// GetUserByID gets user by id
func (us UserUsecase) GetUserByID(id uint) (*entity.User, error) {
	user, err := us.userRepo.FindUserByID(id)
	if err != nil {
		return nil, err
	}
	return dto.ToEntityGetUser(user), nil
}

// CreateUser creates user
func (us UserUsecase) CreateUser(newUser *entity.User) (uint, error) {

	err := us.checkUserFieldsCreate(newUser)

	if err != nil {
		return 0, err
	}

	_, err = us.checkUser(newUser)
	if err != nil {
		return 0, err
	}

	if newUser.Icon == "" {
		newUser.Icon = "img/defaultIcon.webp"
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

// UpdateUser updates user's data
func (us UserUsecase) UpdateUser(newUser *entity.User) error {
	err := us.checkUserFieldsUpdate(newUser)
	if err != nil {
		return err
	}

	_, err = us.checkUser(newUser)
	if err != nil {
		return err
	}

	user, err := us.GetUserByID(newUser.ID)
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

func (us UserUsecase) checkUser(checkUser *entity.User) (*entity.User, error) {
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

func (us UserUsecase) checkUserFieldsCreate(user *entity.User) error {
	re := regexp.MustCompile(`^[a-zA-Z0-9_-]{4,19}$`)
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

	re = regexp.MustCompile(`^\+7\s9[0-9]{2}\s[0-9]{3}-[0-9]{2}-[0-9]{2}$`)
	if user.PhoneNumber == "" || !re.MatchString(user.PhoneNumber) {
		return entity.ErrInvalidPhoneNumber
	}
	return nil
}

func (us UserUsecase) checkUserFieldsUpdate(user *entity.User) error {
	re := regexp.MustCompile(`^[a-zA-Z0-9_-]{4,19}$`)
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

	re = regexp.MustCompile(`^\+7\s9[0-9]{2}\s[0-9]{3}-[0-9]{2}-[0-9]{2}$`)
	if !re.MatchString(user.PhoneNumber) && len(user.PhoneNumber) != 0 {
		return entity.ErrInvalidPhoneNumber
	}
	return nil
}

// UpdateAvatar updates user's avatar
func (us UserUsecase) UpdateAvatar(file multipart.File, filehandler *multipart.FileHeader, id uint) error {
	endpoint := "bring-give.hb.ru-msk.vkcs.cloud"
	location := "bring-give"
	accessKeyID := "k1EeoX4ejNogUZS2TcirVq"
	secretAccessKey := "6Bo85qeL9A1bmrjdWH7a577wKwzbipc6ajVZGFoXTyaT"
	useSSL := true

	bucketName := "bring-give"
	objectName := filehandler.Filename
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		fmt.Println("New", err)
		return entity.ErrInternalServerError
	}

	err = us.uploadFile(minioClient, bucketName, location, objectName, file, filehandler.Size)
	if err != nil {
		fmt.Println("upload", err)
		return entity.ErrInternalServerError
	}

	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", "attachment; filename=\""+objectName+"\"")

	presignedURL, err := minioClient.PresignedGetObject(bucketName, objectName, time.Duration(168)*time.Hour, reqParams)
	if err != nil {
		fmt.Println("presigned", err)
		return entity.ErrInternalServerError
	}

	user, err := us.GetUserByID(id)
	if err != nil {
		fmt.Println("getUserbyId", err)
		return err
	}

	user.Icon = presignedURL.String()

	return us.userRepo.UpdateUser(dto.ToRepoUpdateUser(user))
}

func (us UserUsecase) uploadFile(minioClient *minio.Client, bucketName string, location string, objectName string, file multipart.File, filesize int64) error {
	err := minioClient.MakeBucket(bucketName, location)
	if err != nil {
		fmt.Println("make", err)
		exists, errBucketExists := minioClient.BucketExists(bucketName)
		if !(errBucketExists == nil && exists) {
			fmt.Println("buck", err)
			return err
		}
	}

	_, err = minioClient.PutObject(bucketName, objectName, file, filesize, minio.PutObjectOptions{})
	if err != nil {
		fmt.Println("put", err)
		return entity.ErrInternalServerError
	}

	return nil
}
