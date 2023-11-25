package postgres

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

func (repo *CsatRepo) AddAnswer(answer *entity.Answer) error {
	insertProduct := `INSERT INTO answer (question_id, text) VALUES ($1, $2)`
	_, err := repo.DB.Exec(insertProduct, answer.QuestionId, answer.Text)
	if err != nil {
		return err
	}
	return nil
}

func (repo *CsatRepo) GetAnswerTypeBYQuestionId(id uint) (uint, error) {
	var answertype uint
	row := repo.DB.QueryRow("SELECT answer_type FROM question WHERE id = $1", id)
	err := row.Scan(
		&answertype,
	)
	if err != nil {
		return 0, err
	}
	return answertype, nil
}

func (repo *CsatRepo) GetAnswerByQuestionId(id uint) ([]*entity.Answer, error) {
	rows, err := repo.DB.Query("SELECT id, text FROM answer WHERE question_id = $1", id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}
	defer rows.Close()

	var Answers = []*entity.Answer{}
	for rows.Next() {
		answer := &entity.Answer{}
		err = rows.Scan(
			&answer.Id,
			&answer.Text,
		)
		if err != nil {
			return nil, err
		}
		Answers = append(Answers, answer)
	}
	return Answers, nil
}
