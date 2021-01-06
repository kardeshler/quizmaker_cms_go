package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Name        string
	Description string
	Platforms   []Platform `gorm:"many2many:platform_categories"`
	Quizzes     []Quiz     `gorm:"many2many:category_quizzes"`
}

func (c *Category) CategoryGet() CategoryGet {
	return CategoryGet{
		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
	}
}

type CategoryCreate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CategoryGet struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
