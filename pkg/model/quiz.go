package model

import "github.com/jinzhu/gorm"

type Quiz struct {
	gorm.Model
	Name        string
	Description string
	Platform    Platform
	Questions   []Question `gorm:"many2many:quiz_questions;"`
	Categories  []Category `gorm:"many2many:quiz_categories;"`
}

type QuizCreate struct {
	Name        string
	Description string
	LanguageID  uint64
	PlatformID  uint64
	QuestionIds []uint64
	CategoryIds []uint64
}

type QuizGet struct {
	ID          uint64
	Name        string
	Description string
	LanguageID  uint64
	PlatformID  uint64
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
