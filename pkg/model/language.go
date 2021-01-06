package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Language struct {
	gorm.Model
	Name      string
	ShortName string
	Platforms []Platform `gorm:"many2many:platform_languages;"`
}

type LanguageCreate struct {
	Name      string
	ShortName string
}

type LanguageGet struct {
	ID        uint
	Name      string
	ShortName string
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (l *Language) LanguageGet() LanguageGet {
	return LanguageGet{
		ID:        l.ID,
		Name:      l.Name,
		ShortName: l.ShortName,
		CreatedAt: l.CreatedAt,
		UpdatedAt: l.UpdatedAt,
	}
}
