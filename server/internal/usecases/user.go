package usecases

import (
	"database/sql"
	"server/internal/domain/entity"
	"server/internal/repository"
)


type UserRepo interface{
	GetUserById(id uint) *entity.User
	Create(user *entity.User) error
	FindUserBy(field string, value string) (*entity.User, error) 
}

type UserUsecase struct {
	userRepo *repository.UserRepo
}

func NewUserUsecase(db *sql.DB) *UserUsecase {
	return &UserUsecase{
		userRepo: repository.NewUserRepo(db),
	}
}


func (us UserUsecase) GetUserById(id uint) *entity.User {
	return us.userRepo.GetUserById(id)	
}

func (us UserUsecase) Create(user *entity.User) error {
	return us.userRepo.Create(user) 
}

func (us UserUsecase) FindUserBy(field string, value string) (*entity.User, error) {
	return us.userRepo.FindUserBy(field, value)
}

