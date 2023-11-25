package usecase

import (
	csatRep "server/internal/Csat/repository"
	"server/internal/domain/dto"
	"server/internal/domain/entity"
)

type UsecaseI interface {
	GetQuestionnaireByID(id uint) (*dto.Config, error)
	AddAnswer(answer *entity.Answer) error
	GetAnswersByQuestionId(id uint) (*dto.AnswerConfig, error)
}

type csatUsecase struct {
	csatRepo csatRep.CsatRepositoryI
}

func NewCsatUsecase(csatRep csatRep.CsatRepositoryI) *csatUsecase {
	return &csatUsecase{
		csatRepo: csatRep,
	}
}

func (cu csatUsecase) GetQuestionnaireByID(id uint) (*dto.Config, error) {
	questionnaire, err := cu.csatRepo.GetQuestionnaireByID(id)
	if err != nil {
		return nil, err
	}
	questions, err := cu.csatRepo.GetQuestionsByQuestionnaireID(id)
	if err != nil {
		return nil, err
	}

	config := dto.Config{
		Title:     questionnaire.Name,
		Questions: questions,
	}
	return &config, nil
}

func (cu csatUsecase) AddAnswer(answer *entity.Answer) error {
	err := cu.csatRepo.AddAnswer(answer)
	if err != nil {
		return err
	}
	return nil
}

func (cu csatUsecase) GetAnswersByQuestionId(id uint) (*dto.AnswerConfig, error) {
	answerType, err := cu.csatRepo.GetAnswerTypeBYQuestionId(id)
	if err != nil {
		return nil, err
	}
	answers, err := cu.csatRepo.GetAnswerByQuestionId(id)
	if err != nil {
		return nil, err
	}
	config := dto.AnswerConfig{
		Type:    answerType,
		Answers: answers,
	}
	return &config, nil
}
