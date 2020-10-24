package handler

import (
	"quizcms/pkg/database"
	"quizcms/pkg/model"

	"github.com/gofiber/fiber"
)

// GetAllQuizzes returns all quizzes
func GetAllQuizzes(c *fiber.Ctx) {
	db := database.DB
	var quizzes []model.Quiz
	db.Find(&quizzes)
	quizList := make([]model.Quiz, len(quizzes))
	for index, quiz := range quizzes {
		quizList[index] = quiz
	}
	c.JSON(fiber.Map{"status": "success", "message": "All quizzes", "data": quizList})
}

// GetQuiz returns single quiz
func GetQuiz(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DB
	var quiz model.Quiz
	db.First(&quiz, id)

	if quiz.ID == 0 {
		c.Status(404).JSON(fiber.Map{"status": "error", "message": "Quiz not found!", "data": nil})
		return
	}
	// todo: QuizGet(), QuestionGet()
	c.JSON(fiber.Map{"status": "success", "message": "Quiz found", "data": quiz})
}

// CreateQuiz creates a quiz
func CreateQuiz(c *fiber.Ctx) {
	db := database.DB
	quiz := new(model.Quiz)
	if err := c.BodyParser(quiz); err != nil {
		c.Status(400).JSON(fiber.Map{"status": "error", "message": "Could not parse request!", "data": nil})
		return
	}

	db.Save(&quiz)
	if db.Error != nil {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not save quiz!", "data": nil})
		return
	}

	c.JSON(fiber.Map{"status": "success", "message": "Quiz created", "data": quiz.ID})
}

// UpdateQuiz updates a quiz
func UpdateQuiz(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DB

	newQuiz := new(model.Quiz)
	if err := c.BodyParser(newQuiz); err != nil {
		c.Status(400).JSON(fiber.Map{"status": "error", "message": "Could not parse request!", "data": nil})
		return
	}

	var oldQuiz model.Quiz
	db.First(&oldQuiz, id)
	if oldQuiz.ID == 0 {
		c.Status(404).JSON(fiber.Map{"status": "error", "message": "Quiz not found!", "data": nil})
		return
	}

	oldQuiz.Name = newQuiz.Name
	oldQuiz.Description = newQuiz.Description
	oldQuiz.Platform = newQuiz.Platform
	oldQuiz.Questions = newQuiz.Questions
	oldQuiz.Categories = newQuiz.Categories
	db.Save(&oldQuiz)

	if db.Error != nil {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not update quiz!", "data": nil})
		return
	}

	c.JSON(fiber.Map{"status": "success", "message": "Quiz updated", "data": oldQuiz.ID})
}

// DeleteQuiz deletes a quiz
func DeleteQuiz(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DB
	var quiz model.Quiz
	db.First(&quiz, id)
	if quiz.ID == 0 {
		c.Status(404).JSON(fiber.Map{"status": "error", "message": "Quiz not found!", "data": nil})
		return
	}

	db.Delete(&quiz, id)
	if db.Error != nil {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not delete quiz!", "data": nil})
		return
	}

	c.JSON(fiber.Map{"status": "success", "message": "Deleted quiz", "data": nil})
}
