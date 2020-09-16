package handler

import (
	"quizcms/pkg/database"
	"quizcms/pkg/model"

	"github.com/gofiber/fiber"
)

func GetAllCategories(c *fiber.Ctx) {
	db := database.DB
	var categories []model.Category
	db.Find(&categories)
	categoryList := make([]model.CategoryGet, len(categories))
	for index, category := range categories {
		categoryList[index] = category.CategoryGet()
	}
	c.JSON(fiber.Map{"status": "success", "message": "All categories", "data": categoryList})
}

func GetCategory(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DB
	var category model.Category
	db.Find(&category, id)
	if category.Name == "" {
		c.Status(404).JSON(fiber.Map{"status": "error", "message": "Category not found", "data": nil})
		return
	}
	c.JSON(fiber.Map{"status": "success", "message": "Category found", "data": category.CategoryGet()})
}

func CreateCategory(c *fiber.Ctx) {
	db := database.DB
	category := new(model.Category)
	if err := c.BodyParser(&category); err != nil {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create category", "data": nil})
		return
	}
	db.Create(&category)
	if db.Error != nil {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": db.Error, "data": nil})
		return
	}
	c.JSON(fiber.Map{"status": "success", "message": "Created category", "data": category})
}

func UpdateCategory(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DB
	form := new(model.Category)
	if err := c.BodyParser(&form); err != nil {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't update category", "data": nil})
		return
	}
	var category model.Category
	db.First(&category, id)
	if category.ID == 0 {
		c.Status(404).JSON(fiber.Map{"status": "error", "message": "No category found with ID", "data": nil})
		return
	}
	category.Name = form.Name
	category.Description = form.Description
	db.Save(&category)
	if db.Error != nil {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": db.Error, "data": nil})
		return
	}
	c.JSON(fiber.Map{"status": "success", "message": "Updated category", "data": category})
}

func DeleteCategory(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DB

	var category model.Category
	db.First(&category, id)
	if category.Name == "" {
		c.Status(404).JSON(fiber.Map{"status": "error", "message": "No category found with ID", "data": nil})
		return
	}
	db.Delete(&category)
	c.JSON(fiber.Map{"status": "success", "message": "Category successfully deleted", "data": nil})
}
