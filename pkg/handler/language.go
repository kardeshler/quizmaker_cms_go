package handler

import (
	"quizcms/pkg/database"
	"quizcms/pkg/model"

	"github.com/gofiber/fiber"
)

// GetAllLanguages returns all languages
func GetAllLanguages(c *fiber.Ctx) {
	db := database.DB
	var languages []model.Language
	db.Find(&languages)
	languageList := make([]model.LanguageGet, len(languages))
	for index, language := range languages {
		languageList[index] = language.LanguageGet()
	}
	c.JSON(fiber.Map{"status": "success", "message": "All languages", "data": languageList})
}

// GetLanguage returns singe language
func GetLanguage(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DB
	var language model.Language
	db.Find(&language, id)
	if language.Name == "" {
		c.Status(404).JSON(fiber.Map{"status": "error", "message": "Language not found!", "data": nil})
		return
	}
	c.JSON(fiber.Map{"status": "success", "message": "Language found", "data": language.LanguageGet()})
}

// CreateLanguage creates one language
func CreateLanguage(c *fiber.Ctx) {
	db := database.DB
	language := new(model.Language)
	if err := c.BodyParser(language); err != nil {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create language!", "data": nil})
		return
	}
	db.Create(&language)
	if db.Error != nil {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": db.Error, "data": nil})
		return
	}
	c.JSON(fiber.Map{"status": "success", "message": "Created language", "data": language})
}

// DeleteLanguage deletes one language
func DeleteLanguage(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DB
	var language model.Language
	db.First(&language, id)

	if language.Name == "" {
		c.Status(404).JSON(fiber.Map{"status": "error", "message": "Language not found!", "data": nil})
		return
	}
	db.Delete(&language)
	c.JSON(fiber.Map{"status": "success", "message": "Language successfully deleted", "data": nil})
}

// UpdateLanguage updates an existing language
func UpdateLanguage(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DB
	newLanguage := new(model.Language)
	if err := c.BodyParser(newLanguage); err != nil {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not update language!", "data": nil})
		return
	}
	var language model.Language
	db.First(&language, id)
	if language.ID == 0 {
		c.Status(404).JSON(fiber.Map{"status": "error", "message": "Language not found!", "data": nil})
		return
	}

	language.Name = newLanguage.Name
	language.ShortName = newLanguage.ShortName
	db.Save(&language)
	if db.Error != nil {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": db.Error, "data": nil})
		return
	}
	c.JSON(fiber.Map{"status": "success", "message": "Language successfully updated", "data": language})
}
