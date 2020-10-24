package model

import "github.com/jinzhu/gorm"

// Platform type
type Platform struct {
	gorm.Model
	Name       string
	Languages  []Language `gorm:"many2many:platform_languages;"`
	Categories []Category `gorm:"many2many:platform_categories;"`
}

// PlatformGet getter for Platform
func (p *Platform) PlatformGet() PlatformGet {
	// first mapping languages and categories into PlatformGet type
	languageGets := make([]LanguageGet, len(p.Languages))
	for index, lang := range p.Languages {
		var langGet LanguageGet
		langGet.ID = lang.ID
		langGet.Name = lang.Name
		langGet.ShortName = lang.ShortName
		langGet.CreatedAt = lang.CreatedAt
		langGet.UpdatedAt = lang.UpdatedAt
		languageGets[index] = langGet
	}

	categoryGets := make([]CategoryGet, len(p.Categories))
	for index, ctg := range p.Categories {
		var ctgGet CategoryGet
		ctgGet.ID = ctg.ID
		ctgGet.Name = ctg.Name
		ctgGet.Description = ctg.Description
		ctgGet.CreatedAt = ctg.CreatedAt
		ctgGet.UpdatedAt = ctg.UpdatedAt
		categoryGets[index] = ctgGet
	}

	return PlatformGet{
		ID:         p.ID,
		Name:       p.Name,
		Languages:  languageGets,
		Categories: categoryGets,
	}
}

type PlatformCreateUpdate struct {
	Name       string
	Categories []uint64
	Quizzes    []uint64
	Languages  []uint64
}

type PlatformGet struct {
	ID         uint
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
