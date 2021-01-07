package handler

import (
	"fmt"
	"quizcms/pkg/database"
	"quizcms/pkg/model"

	"github.com/gofiber/fiber"
)

// GetAllQuestions returns all questions
func GetAllQuestions(c *fiber.Ctx) {
	db := database.DB
	var questions []model.Question

	db.Preload("Options").Find(&questions)
	questionList := make([]model.QuestionGet, len(questions))
	for index, question := range questions {
		questionList[index] = question.QuestionGet()
	}
	c.JSON(fiber.Map{"status": "success", "message": "All questions", "data": questionList})
}

// GetQuestion returns single question
func GetQuestion(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DB
	var question model.Question
	db.Preload("Options").First(&question, id)

	// todo: also add handler to return questions of a platform
	if question.ID == 0 {
		c.Status(404).JSON(fiber.Map{"status": "error", "message": "Question not found!", "data": nil})
		return
	}
	questionGet := question.QuestionGet()
	c.JSON(fiber.Map{"status": "success", "message": "Question found", "data": questionGet})
}

// CreateQuestion creates a question
func CreateQuestion(c *fiber.Ctx) {
	db := database.DB
	form := new(model.QuestionCreate)
	if err := c.BodyParser(form); err != nil {
		c.Status(400).JSON(fiber.Map{"status": "error", "message": "Could not parse request!", "data": nil})
		return
	}

	options := make([]model.Option, 0)
	if len(form.Options) > 0 {
		for _, v := range form.Options {
			options = append(options, model.Option{
				Content:   v.Content,
				IsCorrect: v.IsCorrect,
			})
		}
	}

	// todo: attach QuizID to a real quiz
	question := model.Question{
		Content:    form.Content,
		Hint:       form.Hint,
		LanguageID: form.LanguageID,
		QuizID:     form.QuizID,
		Options:    options,
	}

	db.Create(&question)
	if db.Error != nil {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not save question!", "data": nil})
		return
	}
	c.JSON(fiber.Map{"status": "success", "message": "Question created", "data": question.ID})
}

// UpdateQuestion updates a question
func UpdateQuestion(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DB

	newQuestion := new(model.Question)
	if err := c.BodyParser(newQuestion); err != nil {
		c.Status(400).JSON(fiber.Map{"status": "error", "message": "Could not parse request!", "data": nil})
		return
	}

	var oldQuestion model.Question
	db.Preload("Options").First(&oldQuestion, id)
	if oldQuestion.ID == 0 {
		c.Status(404).JSON(fiber.Map{"status": "error", "message": "Question not found!", "data": nil})
		return
	}

	oldQuestion.Content = newQuestion.Content
	oldQuestion.Hint = newQuestion.Hint
	oldQuestion.QuizID = newQuestion.QuizID

	oldOptions := oldQuestion.Options
	newOptions := newQuestion.Options
	filteredOptions := filterQptions(oldOptions, newOptions)
	fmt.Println("Adding new options to the question +", filteredOptions)
	deleteOptions := filterQptions(newOptions, oldOptions)
	fmt.Println("Removing unwanted options from the question +", deleteOptions)

	if len(filteredOptions) > 0 {
		fmt.Println("Adding new options to the question +", filteredOptions)
		db.Model(&oldQuestion).Association("Options").Append(filteredOptions)
	}
	if len(deleteOptions) > 0 {
		fmt.Println("Removing unwanted options from the question +", deleteOptions)
		db.Model(&oldQuestion).Association("Options").Delete(deleteOptions)
	}

	db.Save(&oldQuestion)

	if db.Error != nil {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not update question!", "data": nil})
		return
	}

	c.JSON(fiber.Map{"status": "success", "message": "Question updated", "data": oldQuestion.ID})
}

func filterQptions(base []model.Option, target []model.Option) (out []model.Option) {
	filterMap := make(map[uint]struct{}, len(base))
	for _, v := range base {
		filterMap[v.ID] = struct{}{}
	}

	for _, v := range target {
		if _, ok := filterMap[v.ID]; ok == false {
			out = append(out, v)
		}
	}
	return
}

// DeleteQuestion deletes a question
func DeleteQuestion(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DB

	var question model.Question
	db.First(&question, id)

	if question.ID == 0 {
		c.Status(404).JSON(fiber.Map{"status": "error", "message": "Question not found!", "data": nil})
		return
	}

	db.Delete(&question, id)
	if db.Error != nil {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not delete question!", "data": nil})
		return
	}

	c.JSON(fiber.Map{"status": "success", "message": "Deleted question", "data": nil})
}
