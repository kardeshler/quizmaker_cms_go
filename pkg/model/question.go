package model

import "github.com/jinzhu/gorm"

type Question struct {
	gorm.Model
	Content string
	Hint    string
	Quizzes []Quiz `gorm:"many2many:question_quizzes;"`
}

type QuestionCreate struct {
	Content string
	Hint    string
	LangId  uint64
	QuizIds []uint64
	Answers []string
}

type QuestionGet struct {
	QuestionCreate
	ID uint64
}

type QuestionAnswerUpdate struct {
	Answers []string
}

type QuestionQuizzesUpdate struct {
	QuizIds []uint64
}
