package model

import "github.com/jinzhu/gorm"

// Question to be added to quizzes
type Question struct {
	gorm.Model
	Content    string
	Hint       string
	LanguageID uint
	QuizID     uint
	Options    []Option
}

// QuestionGet is the getter for question type
func (q *Question) QuestionGet() QuestionGet {
	options := make([]OptionCreate, len(q.Options))
	for index, option := range q.Options {
		options[index] = OptionCreate{
			Content:   option.Content,
			IsCorrect: option.IsCorrect,
		}
	}

	return QuestionGet{
		ID:         q.ID,
		Content:    q.Content,
		Hint:       q.Hint,
		LanguageID: q.LanguageID,
		QuizID:     q.QuizID,
		Options:    options,
	}
}

// QuestionCreate type for creating questions
type QuestionCreate struct {
	Content    string
	Hint       string
	LanguageID uint
	QuizID     uint
	Options    []OptionCreate
}

// QuestionGet type is used to read questions
type QuestionGet struct {
	ID         uint
	Content    string
	Hint       string
	LanguageID uint
	QuizID     uint
	Options    []OptionCreate
}

// QuestionAnswerUpdate may be removed
type QuestionAnswerUpdate struct {
	Answers []string
}

// QuestionQuizzesUpdate may be removed
type QuestionQuizzesUpdate struct {
	QuizIds []uint64
}
