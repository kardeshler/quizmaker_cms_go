package handler

import (
	"quizcms/pkg/database"
	"quizcms/pkg/model"

	"github.com/gofiber/fiber"
)

// GetAllPlatforms return all platform
func GetAllPlatforms(c *fiber.Ctx) {
	db := database.DB
	var platforms []model.Platform

	db.Find(&platforms)
	platformList := make([]model.PlatformGet, len(platforms))
	for index, platform := range platforms {
		platformList[index] = platform.PlatformGet()
	}

	c.JSON(fiber.Map{"status": "success", "message": "All platforms", "data": platformList})
}

// GetPlatform returns single platform
func GetPlatform(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DB
	var platform model.Platform
	db.First(&platform, id)
	if platform.ID == 0 {
		c.Status(404).JSON(fiber.Map{"status": "error", "message": "Platform not found!", "data": nil})
		return
	}
	c.JSON(fiber.Map{"status": "success", "message": "Found platform", "data": platform.PlatformGet()})
}

// CreatePlatform creates a platform
func CreatePlatform(c *fiber.Ctx) {
	db := database.DB
	platform := new(model.Platform)

	if err := c.BodyParser(platform); err != nil {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not parse request!", "data": nil})
		return
	}
	db.Create(&platform)
	if db.Error != nil {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": db.Error, "data": nil})
		return
	}
	c.JSON(fiber.Map{"status": "success", "message": "Created platform", "data": platform.PlatformGet()})
}

// UpdatePlatform updates a platform
func UpdatePlatform(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DB
	newPlatform := new(model.Platform)
	if err := c.BodyParser(newPlatform); err != nil {
		c.Status(400).JSON(fiber.Map{"status": "error", "message": "Could not parse request!", "data": nil})
		return
	}

	var oldPlatform model.Platform
	db.First(&oldPlatform, id)

	if oldPlatform.ID == 0 {
		c.Status(404).JSON(fiber.Map{"status": "error", "message": "Platform not found!", "data": nil})
		return
	}

	oldPlatform.Name = newPlatform.Name
	oldPlatform.Categories = newPlatform.Categories
	oldPlatform.Languages = newPlatform.Languages

	db.Save(&oldPlatform)
	if db.Error != nil {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": db.Error, "data": nil})
		return
	}
	c.JSON(fiber.Map{"status": "success", "message": "Updated platform", "data": oldPlatform.ID})
}

// DeletePlatform deletes a platform
func DeletePlatform(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DB
	var platform model.Platform

	db.First(&platform, id)
	if platform.ID == 0 {
		c.Status(404).JSON(fiber.Map{"status": "error", "message": "Platform not found!", "data": nil})
		return
	}

	db.Delete(&platform, id)
	if db.Error != nil {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not delete platform!", "data": nil})
		return
	}

	c.JSON(fiber.Map{"status": "success", "message": "Deleted platform", "data": nil})
}

// GetQuizzesOfPlatform returns all quizzes of a given platform
func GetQuizzesOfPlatform(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DB
	var platform model.Platform

	db.First(&platform, id)
	if platform.ID == 0 {
		c.Status(404).JSON(fiber.Map{"status": "error", "message": "Platform not found!", "data": nil})
		return
	}

	quizzes := make([]model.Quiz, 0)
	db.Preload("Platforms").Find(&quizzes)

	if db.Error != nil {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not get questions of the platform!", "data": nil})
		return
	}

	quizList := make([]model.QuizGet, len(quizzes))
	for index, quiz := range quizzes {
		quizList[index] = quiz.QuizGet()
	}
	c.JSON(fiber.Map{"status": "success", "message": "Quizzes of the platform", "data": quizList})
}
