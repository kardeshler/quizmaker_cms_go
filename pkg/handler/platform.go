package handler

import (
	"fmt"
	"quizcms/pkg/database"
	"quizcms/pkg/model"

	"github.com/gofiber/fiber"
	"gorm.io/gorm/clause"
)

// GetAllPlatforms return all platform
func GetAllPlatforms(c *fiber.Ctx) {
	db := database.DB
	var platforms []model.Platform

	db.Preload("Languages").Preload("Categories").Find(&platforms)
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
	form := new(model.PlatformCreateUpdate)
	if err := c.BodyParser(form); err != nil {
		c.Status(400).JSON(fiber.Map{"status": "error", "message": "Could not parse request!", "data": nil})
		return
	}

	var oldPlatform model.Platform
	db.Preload(clause.Associations).First(&oldPlatform, id)

	if oldPlatform.ID == 0 {
		c.Status(404).JSON(fiber.Map{"status": "error", "message": "Platform not found!", "data": nil})
		return
	}

	oldPlatform.Name = form.Name

	// we get the categories from the update request body & the actual platform
	formCategories := make([]model.Category, len(form.Categories))
	for index, ctg := range form.Categories {
		formCategories[index] = model.Category{Name: ctg}
	}
	oldCategories := oldPlatform.Categories

	filteredNewCategories := filterCategories(oldCategories, formCategories, true)
	deleteCategories := filterCategories(formCategories, oldCategories, false)

	if len(filteredNewCategories) > 0 {
		fmt.Println("Adding new categories +", filteredNewCategories)
		db.Model(&oldPlatform).Association("Categories").Append(filteredNewCategories)
	}
	if len(deleteCategories) > 0 {
		fmt.Println("Removing filtered out categories from platform +", deleteCategories)
		db.Model(&oldPlatform).Association("Categories").Delete(deleteCategories)
	}

	// we get the languages from the update request body & the actual platform
	formLanguages := make([]model.Language, len(form.Languages))
	for index, lang := range form.Languages {
		formLanguages[index] = model.Language{ShortName: lang}
	}
	oldLanguages := oldPlatform.Languages

	filteredNewLanguages := filterLanguages(oldLanguages, formLanguages, true)
	deleteLanguages := filterLanguages(formLanguages, oldLanguages, false)

	if len(filteredNewLanguages) > 0 {
		fmt.Println("Adding new languages +", filteredNewLanguages)
		db.Model(&oldPlatform).Association("Languages").Append(filteredNewLanguages)
	}
	if len(deleteLanguages) > 0 {
		fmt.Println("Removing filtered out languages from platform +", deleteLanguages)
		db.Model(&oldPlatform).Association("Languages").Delete(deleteLanguages)
	}

	db.Save(&oldPlatform)
	if db.Error != nil {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": db.Error, "data": nil})
		return
	}
	c.JSON(fiber.Map{"status": "success", "message": "Updated platform", "data": oldPlatform.ID})
}

func filterCategories(base []model.Category, target []model.Category, create bool) (out []model.Category) {
	filterMap := make(map[string]struct{}, len(base))
	for _, v := range base {
		filterMap[v.Name] = struct{}{}
	}

	for _, v := range target {
		if _, ok := filterMap[v.Name]; ok == false {
			if create {
				database.DB.Where("name = ?", v.Name).First(&v)
			}
			out = append(out, v)
		}
	}
	return
}

func filterLanguages(base []model.Language, target []model.Language, create bool) (out []model.Language) {
	filterMap := make(map[string]struct{}, len(base))
	for _, v := range base {
		filterMap[v.ShortName] = struct{}{}
	}

	for _, v := range target {
		if _, ok := filterMap[v.ShortName]; ok == false {
			if create {
				database.DB.Where("short_name = ?", v.ShortName).First(&v)
			}
			out = append(out, v)
		}
	}
	return
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
