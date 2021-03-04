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
	languageGets := make([]LanguageGetLite, len(p.Languages))
	for index, lang := range p.Languages {
		var langGet LanguageGetLite
		langGet.Name = lang.Name
		langGet.ShortName = lang.ShortName
		languageGets[index] = langGet
	}

	categoryGets := make([]CategoryGetLite, len(p.Categories))
	for index, ctg := range p.Categories {
		var ctgGet CategoryGetLite
		ctgGet.Name = ctg.Name
		ctgGet.Description = ctg.Description
		categoryGets[index] = ctgGet
	}

	return PlatformGet{
		ID:         p.ID,
		Name:       p.Name,
		Languages:  languageGets,
		Categories: categoryGets,
	}
}

// PlatformCreateUpdate to create & update platforms
type PlatformCreateUpdate struct {
	Name       string
	Categories []string
	Quizzes    []uint
	Languages  []string
}

// PlatformGet to read platform
type PlatformGet struct {
	ID         uint              `json:"id"`
	Name       string            `json:"name"`
	Categories []CategoryGetLite `json:"categories"`
	Quizzes    []QuizGet         `json:"quizzes"`
	Languages  []LanguageGetLite `json:"languages"`
}

// PlatformCategoryUpdate incase new handlers are added
type PlatformCategoryUpdate struct {
	Categories []uint64
}

// PlatformLangUpdate incase new handlers are added
type PlatformLangUpdate struct {
	Languages []uint64
}

// PlatformLight incase we may init with id
type PlatformLight struct {
	ID uint64
}

// PlatformLightGet to read platform in a new potential handler
type PlatformLightGet struct {
	ID         uint64
	Categories []uint64
	Quizzes    []uint64
	Languages  []uint64
}
