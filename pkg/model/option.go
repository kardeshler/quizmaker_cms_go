package model

import "github.com/jinzhu/gorm"

/*
Option is the model for the choices that are added to each question.
IsCorrect indicates if the corresponding option is a correct answer or not.
*/
type Option struct {
	gorm.Model
	Content    string
	IsCorrect  bool
	QuestionID uint
}
