package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Language entity
type Language struct {
	gorm.Model
	Name      string
	ShortName string
	Platforms []Platform `gorm:"many2many:platform_languages;"`
}

// LanguageCreate incase we add new handler function to create language with those
type LanguageCreate struct {
	Name      string
	ShortName string
}

// LanguageGet to read language
type LanguageGet struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	ShortName string    `json:"short_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// LanguageGetLite only to get simple language specs
type LanguageGetLite struct {
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
}

// LanguageGet getter method for Language type
func (l *Language) LanguageGet() LanguageGet {
	return LanguageGet{
		ID:        l.ID,
		Name:      l.Name,
		ShortName: l.ShortName,
		CreatedAt: l.CreatedAt,
		UpdatedAt: l.UpdatedAt,
	}
}
