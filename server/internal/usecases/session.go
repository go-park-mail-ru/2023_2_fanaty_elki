package usecases

import (
	"server/internal/domain/entity"
	"server/internal/repository"
	"github.com/gomodule/redigo/redis"
)

type SessionRepo interface {
	Create(in *entity.Session) (*entity.SessionID, error)
	Check(in *entity.SessionID) (*entity.Session, error)
	Delete(in *entity.SessionID) error
}

type SessionUsecase struct {
	sessionRepo *repository.SessionManager
}

func NewSessionUsecase(conn redis.Conn) *SessionUsecase {
	return &SessionUsecase{
		sessionRepo: repository.NewSessionManager(conn),
	}
}

func (ss SessionUsecase) Create(in *entity.Session) (*entity.SessionID, error) {
	return ss.sessionRepo.Create(in)
}

func (ss SessionUsecase) Check(in *entity.SessionID) (*entity.Session, error) {
	return ss.sessionRepo.Check(in)
}

func (ss SessionUsecase) Delete(in *entity.SessionID)  error {
	return ss.sessionRepo.Delete(in)
}