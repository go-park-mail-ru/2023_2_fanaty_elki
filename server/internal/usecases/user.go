package usecases

import (
	"database/sql"
	"server/internal/domain/entity"
	"server/repository"
)


type UserRepo interface{
	GetUserById(id uint) (*entity.User, error)
	CreateUser(user *entity.User) (uint, error)
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


func (us UserUsecase) GetUserById(id uint) (*entity.User, error) {
	return us.userRepo.GetUserById(id)	
}

func (us UserUsecase) CreateUser(user *entity.User) (uint, error) {
	return us.userRepo.CreateUser(user) 
}

func (us UserUsecase) FindUserBy(field string, value string) (*entity.User, error) {
	return us.userRepo.FindUserBy(field, value)
}

