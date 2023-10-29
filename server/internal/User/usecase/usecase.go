package usecase

import (
	"server/internal/domain/entity"
	"server/internal/domain/dto"
	userRep "server/internal/User/repository"
)


type UsecaseI interface{
	GetUserById(id uint) (*entity.User, error)
	CreateUser(new_user *entity.User) (uint, error)
	FindUserBy(field string, value string) (*entity.User, error) 
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

func (us userUsecase) CreateUser(new_user *entity.User) (uint, error) {
	
	user, err := us.userRepo.FindUserBy("Username", new_user.Username)
	if err != nil {
		return 0, entity.ErrInternalServerError
	}

	if user != nil {
		return 0, entity.ErrConflictUsername
	}

	user, err = us.userRepo.FindUserBy("Email", new_user.Email)
	if err != nil {
		return 0, entity.ErrInternalServerError
	}

	if user != nil {
		return 0, entity.ErrConflictEmail
	}

	user, err = us.userRepo.FindUserBy("PhoneNumber", new_user.PhoneNumber)
	if err != nil {
		return 0, entity.ErrInternalServerError
	}

	if user != nil {
		return 0, entity.ErrConflictPhoneNumber
	}

	return us.userRepo.CreateUser(dto.ToRepoUser(new_user)) 
}

func (us userUsecase) FindUserBy(field string, value string) (*entity.User, error) {
	return us.userRepo.FindUserBy(field, value)
}

