package model

import "github.com/jinzhu/gorm"

type Question struct {
	gorm.Model
	Content  string
	Hint     string
	Language Language `gorm:"foreignKey:LanguageId"`
	Quizzes  []Quiz   `gorm:"many2many:question_quizzes;"`
	Answers  []string
}

// QuestionGet is the getter for question type
func (q *Question) QuestionGet() QuestionGet {
	quizIDs := make([]uint, len(q.Quizzes))
	for index, quiz := range q.Quizzes {
		quizIDs[index] = quiz.ID
	}

	questionCreate := QuestionCreate{
		Content: q.Content,
		Hint:    q.Hint,
		LangID:  q.Language.ID,
		QuizIDs: quizIDs,
		Answers: q.Answers,
	}

	return QuestionGet{
		ID:             q.ID,
		QuestionCreate: questionCreate,
	}
}

// QuestionCreate type for creating questions
type QuestionCreate struct {
	Content string
	Hint    string
	LangID  uint
	QuizIDs []uint
	Answers []string
}

// QuestionGet type is used to read questions
type QuestionGet struct {
	ID uint
	QuestionCreate
}

type QuestionAnswerUpdate struct {
	Answers []string
}

type QuestionQuizzesUpdate struct {
	QuizIds []uint64
}
