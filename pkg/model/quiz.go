package model

import "github.com/jinzhu/gorm"

type Quiz struct {
	gorm.Model
	Name        string
	Description string
	Platform    Platform
	Language    Language   `gorm:"foreignKey:LanguageId"`
	Questions   []Question `gorm:"many2many:quiz_questions;"`
	Categories  []Category `gorm:"many2many:quiz_categories;"`
}

func (q *Quiz) QuizGet() QuizGet {
	return QuizGet{
		ID:          q.ID,
		Name:        q.Name,
		Description: q.Description,
		LanguageID:  q.Language.ID,
		PlatformID:  q.Platform.ID,
	}
}

type QuizCreate struct {
	Name        string
	Description string
	LanguageID  uint
	PlatformID  uint64
	QuestionIds []uint64
	CategoryIds []uint64
}

type QuizGet struct {
	ID          uint
	Name        string
	Description string
	LanguageID  uint
	PlatformID  uint
}

type QuizGetExtended struct {
	QuizGet
	Questions []QuestionGet
}

type QuizCategoryUpdate struct {
	Categories []uint64
}

type QuizLanguageUpdate struct {
	LanguageID uint64
}

type QuizPlatformUpdate struct {
	PlatformID uint64
}

type QuizQuestionUpdate struct {
	QuestionIds []uint64
}
