package usecase

import (
	userRep "server/internal/User/repository"
	"server/internal/domain/dto"
	"server/internal/domain/entity"
)

type UsecaseI interface {
	CreateUser(new_user *entity.User) (uint, error)
}

type userUsecase struct {
	userRepo userRep.UserRepositoryI
}

func NewUserUsecase(repI userRep.UserRepositoryI) *userUsecase {
	return &userUsecase{
		userRepo: repI,
	}
}

func (us userUsecase) CreateUser(new_user *entity.User) (uint, error) {

	user, err := us.userRepo.FindUserByUsername(new_user.Username)
	if err != nil {
		return 0, entity.ErrInternalServerError
	}

	if user != nil {
		return 0, entity.ErrConflictUsername
	}

	user, err = us.userRepo.FindUserByEmail(new_user.Email)
	if err != nil {
		return 0, entity.ErrInternalServerError
	}

	if user != nil {
		return 0, entity.ErrConflictEmail
	}

	user, err = us.userRepo.FindUserByPhone(new_user.PhoneNumber)
	if err != nil {
		return 0, entity.ErrInternalServerError
	}

	if user != nil {
		return 0, entity.ErrConflictPhoneNumber
	}

	return us.userRepo.CreateUser(dto.ToRepoUser(new_user))

}
