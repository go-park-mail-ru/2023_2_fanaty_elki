package dto

import "server/internal/domain/entity"

type Config struct {
	Title     string
	Questions []*entity.Question
}

type AnswerConfig struct {
	Type    uint
	Answers []*entity.Answer
}
