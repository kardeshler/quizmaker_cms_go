package model

import "github.com/jinzhu/gorm"

type Platform struct {
	gorm.Model
	Name       string
	Languages  []Language `gorm:"many2many:platform_languages;"`
	Categories []Category `gorm:"many2many:platform_categories;"`
}

type PlatformCreateUpdate struct {
	Name       string
	Categories []uint64
	Quizzes    []uint64
	Languages  []uint64
}

type PlatformGet struct {
	ID         uint64
	Name       string
	Categories []CategoryGet
	Quizzes    []QuizGet
	Languages  []LanguageGet
}

type PlatformCategoryUpdate struct {
	Categories []uint64
}

type PlatformLangUpdate struct {
	Languages []uint64
}

type PlatformLight struct {
	ID uint64
}

type PlatformLightGet struct {
	ID         uint64
	Categories []uint64
	Quizzes    []uint64
	Languages  []uint64
}
