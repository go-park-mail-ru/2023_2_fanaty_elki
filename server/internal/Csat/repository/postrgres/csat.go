package repository

import (
	"database/sql"
	"server/internal/Csat/repository"
	"server/internal/domain/entity"
)

type CsatRepo struct {
	DB *sql.DB
}

func NewCsatRepo(db *sql.DB) repository.CsatRepositoryI {
	return &CsatRepo{
		DB: db,
	}
}

func (repo *CsatRepo) GetQuestionnaireByID(id uint) (*entity.Questionnaire, error) {
	questionnaire := &entity.Questionnaire{}
	row := repo.DB.QueryRow("SELECT id, name FROM questionnaire WHERE id = $1", id)
	err := row.Scan(
		&questionnaire.Id,
		&questionnaire.Name,
	)
	if err != nil {
		return nil, err
	}
	return questionnaire, nil
}

func (repo *CsatRepo) GetQuestionsByQuestionnaireID(id uint) ([]*entity.Question, error) {
	rows, err := repo.DB.Query("SELECT id, text, answer_type FROM question WHERE questionnaire_id = $1", id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}
	defer rows.Close()

	var Questions = []*entity.Question{}
	for rows.Next() {
		question := &entity.Question{}
		err = rows.Scan(
			&question.Id,
			&question.Text,
			&question.AnswerType,
		)
		if err != nil {
			return nil, err
		}
		Questions = append(Questions, question)
	}
	return Questions, nil
}
