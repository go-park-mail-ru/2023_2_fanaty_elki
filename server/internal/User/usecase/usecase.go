package usecase

import (
	"regexp"
	userRep "server/internal/User/repository"
	"server/internal/domain/dto"
	"server/internal/domain/entity"
)

type UsecaseI interface {
	CreateUser(new_user *entity.User) (uint, error)
	UpdateUser(newUser *entity.User) (error) 
}

type userUsecase struct {
	userRepo userRep.UserRepositoryI
}

func NewUserUsecase(repI userRep.UserRepositoryI) *userUsecase {
	return &userUsecase{
		userRepo: repI,
	}
}

func (us userUsecase) GetUserById(id uint) (*entity.User, error) {
	return us.userRepo.FindUserById(id)	
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
	return us.userRepo.CreateUser(dto.ToRepoCreateUser(newUser)) 
}


func (us userUsecase) UpdateUser(newUser *entity.User) (error) {
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

		if newUser.Password != "" {
			user.Password = newUser.Password
		}

		if newUser.Birthday != "" {
			user.Birthday = newUser.Birthday
		}

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
	var user *entity.User

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
	return user, nil
}

func (us userUsecase) checkUserFieldsCreate(user *entity.User) error {
	if len(user.Username) < 3 || len(user.Username) > 30 {
		return entity.ErrInvalidUsername
	}

	if len(user.Password) < 3 || len(user.Password) > 30 {
		return entity.ErrInvalidPassword
	}

	re := regexp.MustCompile(`\d{2}-\d{2}-\d{4}`) 
	if user.Birthday != "" && !re.MatchString(user.Birthday){
		return entity.ErrInvalidBirthday
	}

	re = regexp.MustCompile(`@`)
	if user.Email == "" || !re.MatchString(user.Email) {
		return entity.ErrInvalidEmail
	}

	re = regexp.MustCompile(`^[+]?[0-9]{3,25}$`)
	if user.PhoneNumber == "" || !re.MatchString(user.PhoneNumber) {
		return entity.ErrInvalidPhoneNumber
	}
	return nil
}

func (us userUsecase) checkUserFieldsUpdate(user *entity.User) error {
	if ((len(user.Username) < 3 || len(user.Username) > 30)) && user.Username != "" {
		return entity.ErrInvalidUsername
	}

	if (len(user.Password) < 3 || len(user.Password) > 30) && user.Password != "" {
		return entity.ErrInvalidPassword
	}

	re := regexp.MustCompile(`\d{2}-\d{2}-\d{4}`) 
	if !re.MatchString(user.Birthday) && user.Birthday != ""{
		return entity.ErrInvalidBirthday
	}

	re = regexp.MustCompile(`@`)
	if !re.MatchString(user.Email) && user.Email != ""{
		return entity.ErrInvalidEmail
	}

	re = regexp.MustCompile(`^[+]?[0-9]{3,25}$`)
	if !re.MatchString(user.PhoneNumber) && user.PhoneNumber != ""{
		return entity.ErrInvalidPhoneNumber
	}
	return nil
}