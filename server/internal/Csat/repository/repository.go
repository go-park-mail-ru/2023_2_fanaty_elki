package repository

import entity "server/internal/domain/entity"

type CsatRepositoryI interface {
	GetQuestionnaireByID(id uint) (*entity.Questionnaire, error)
	GetQuestionsByQuestionnaireID(id uint) ([]*entity.Question, error)
	AddAnswer(answer *entity.Answer) error
	GetAnswerTypeBYQuestionId(id uint) (uint, error)
	GetAnswerByQuestionId(id uint) ([]*entity.Answer, error)
}
