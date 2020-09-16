package model

import "github.com/jinzhu/gorm"

type Language struct {
	gorm.Model
	Name      string
	ShortName string
	Platforms []Platform `gorm:"many2many:language_platforms;"`
}

type LanguageCreate struct {
	Name      string
	ShortName string
}

type LanguageGet struct {
	LanguageCreate
	ID uint64
}
