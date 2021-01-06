package handler

import (
	"fmt"
	"quizcms/pkg/database"
	"quizcms/pkg/model"

	"github.com/gofiber/fiber"
)

// GetAllCategories to fetch all categories out there
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

// GetCategory to fetch a category if it exists
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

// CreateCategory to create a brand new category
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

// UpdateCategory to update a category if it exists
func UpdateCategory(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DB
	// Get the category with given id
	// Set updated fields
	// Persist the end result

	var category model.Category
	db.Preload("Platforms").First(&category, id)
	switch {
	case db.Error != nil:
		c.Status(503)
		return
	case category.ID == 0:
		c.Status(404).JSON(fiber.Map{"status": "error", "message": "No category found with ID", "data": nil})
		return
	}

	form := new(model.CategoryCreateUpdate)
	if err := c.BodyParser(&form); err != nil {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't update category", "data": nil})
		return
	}

	category.Name = form.Name
	category.Description = form.Description

	// filter the platforms coming from the update body
	// remove association for those which are filtered out, removed with the update operation
	platforms := make([]model.Platform, len(form.PlatformNames))
	for i, p := range form.PlatformNames {
		platforms[i] = model.Platform{Name: p}
	}

	oldPlatforms := category.Platforms

	filteredPlatforms := filterToCreate(oldPlatforms, platforms)
	deletePlatforms := filterToDelete(oldPlatforms, platforms)

	// todo: find a way to not replace existing associations when you're not changing the platforms at all
	// keep in mind that you are creating new platform instances each time update function's triggered -> Platform{Name: p}
	//db.Model(&category).Association("Platforms").Replace(platforms)

	if len(filteredPlatforms) > 0 {
		fmt.Println("Adding new platforms +", filteredPlatforms)
		db.Model(&category).Association("Platforms").Append(filteredPlatforms)
	}
	if len(deletePlatforms) > 0 {
		fmt.Println("Removing platforms -", deletePlatforms)
		db.Model(&category).Association("Platforms").Delete(deletePlatforms)
	}

	db.Save(&category)
	if db.Error != nil {
		c.Status(500).JSON(fiber.Map{"status": "error", "message": db.Error, "data": nil})
		return
	}
	c.JSON(fiber.Map{"status": "success", "message": "Updated category", "data": category})
}

func filterToCreate(old []model.Platform, new []model.Platform) (out []model.Platform) {
	filterMap := make(map[string]struct{}, len(old))
	for _, v := range old {
		filterMap[v.Name] = struct{}{}
	}

	for _, v := range new {
		if _, ok := filterMap[v.Name]; ok == false {
			out = append(out, v)
		}
	}
	return
}

func filterToDelete(old []model.Platform, new []model.Platform) (out []model.Platform) {
	filterMap := make(map[string]struct{}, len(new))
	for _, v := range new {
		filterMap[v.Name] = struct{}{}
	}

	for _, v := range old {
		if _, ok := filterMap[v.Name]; ok == false {
			out = append(out, v)
		}
	}
	return
}

// DeleteCategory to delete a category if it exists
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
