package handler

import (
	"quizcms/pkg/database"
	"quizcms/pkg/model"

	"github.com/gofiber/fiber"
)

// GetAllQuestions returns all questions
func GetAllQuestions(c *fiber.Ctx) {
	db := database.DB
	var questions []model.Question

	db.Find(&questions)
	questionList := make([]model.Question, len(questions))
	for index, question := range questions {
		questionList[index] = question
	}
	c.JSON(fiber.Map{"status": "success", "message": "All questions", "data": questionList})
}

// GetQuestion returns single question
func GetQuestion(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DB
	var question model.Question
	db.First(&question, id)

	// todo: also add handler to return questions of a platform
	if question.ID == 0 {
		c.Status(404).JSON(fiber.Map{"status": "error", "message": "Question not found!", "data": nil})
		return
	}
	c.JSON(fiber.Map{"status": "success", "message": "Question found", "data": question})
}

// CreateQuestion creates a question
func CreateQuestion(c *fiber.Ctx) {
	db := database.DB
	question := new(model.Question)
	if err := c.BodyParser(question); err != nil {
		c.Status(400).JSON(fiber.Map{"status": "error", "message": "Could not parse request!", "data": nil})
		return
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
	db.First(&oldQuestion, id)
	if oldQuestion.ID == 0 {
		c.Status(404).JSON(fiber.Map{"status": "error", "message": "Question not found!", "data": nil})
		return
	}

	oldQuestion.Content = newQuestion.Content
	oldQuestion.Hint = newQuestion.Hint
	oldQuestion.Quizzes = newQuestion.Quizzes
	db.Save(&oldQuestion)

	if db.Error != nil {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not update question!", "data": nil})
		return
	}

	c.JSON(fiber.Map{"status": "success", "message": "Question updated", "data": oldQuestion.ID})
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
