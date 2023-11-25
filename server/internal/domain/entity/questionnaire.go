package entity

type Questionnaire struct {
	Id   uint
	Name string
}

type Question struct {
	Id              uint
	QuestionnaireId uint
	Text            string
	AnswerType      uint
}

type Answer struct {
	Id         uint
	QuestionId uint
	Text       string
}

type Admin struct {
	Id       uint
	Username string
	Password string
}
