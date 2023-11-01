package usecase

import (
	userRep "server/internal/User/repository"
	"server/internal/domain/dto"
	"server/internal/domain/entity"
)

type UsecaseI interface {
  GetUserById(id uint) (*entity.User, error)
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
	return us.userRepo.GetUserById(id)	
}

func (us userUsecase) CreateUser(newUser *entity.User) (uint, error) {
	_, err := us.checkUser(newUser)
	if err != nil {
		return 0, err
	}
	return us.userRepo.CreateUser(dto.ToRepoCreateUser(newUser)) 
}


func (us userUsecase) UpdateUser(newUser *entity.User) (error) {
	_, err := us.checkUser(newUser)
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

	user, err = us.userRepo.FindUserByUsername(checkUser.Username)
	if err != nil {
		return nil, entity.ErrInternalServerError
	}

	if user != nil {
		return nil, entity.ErrConflictUsername
	}

	user, err = us.userRepo.FindUserByEmail(checkUser.Email)
	if err != nil {
		return nil, entity.ErrInternalServerError
	}

	if user != nil {
		return nil, entity.ErrConflictEmail
	}

	user, err = us.userRepo.FindUserByPhone(checkUser.PhoneNumber)
	if err != nil {
		return nil, entity.ErrInternalServerError
	}

	if user != nil {
		return nil, entity.ErrConflictPhoneNumber
	}
	
	return user, nil
}

	return us.userRepo.CreateUser(dto.ToRepoUser(new_user))
}

