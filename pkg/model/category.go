package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Category entity
type Category struct {
	gorm.Model
	Name        string
	Description string
	Platforms   []Platform `gorm:"many2many:platform_categories;"`
	Quizzes     []Quiz     `gorm:"many2many:category_quizzes"`
}

// CategoryGet getter method for Category type
func (c *Category) CategoryGet() CategoryGet {
	return CategoryGet{
		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
	}
}

// CategoryCreateUpdate to update categories
type CategoryCreateUpdate struct {
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	PlatformNames []string `json:"platforms"`
}

// CategoryCreateLite incase we add handlers to only create categories with name
type CategoryCreateLite struct {
	Name string `json:"name"`
}

// CategoryGet to read categories
type CategoryGet struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CategoryGetLite to get simple category specs
type CategoryGetLite struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
