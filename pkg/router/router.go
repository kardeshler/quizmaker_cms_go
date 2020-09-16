package router

import (
	"quizcms/pkg/handler"
	"quizcms/pkg/middleware"

	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1", logger.New())
	api.Get("/", handler.Hello)

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", handler.Login)

	// Categories
	category := api.Group("/categories")
	category.Get("/", middleware.Protected(), handler.GetAllCategories)
	category.Get("/:id", middleware.Protected(), handler.GetCategory)
	category.Post("/", middleware.Protected(), handler.CreateCategory)
	category.Delete("/:id", middleware.Protected(), handler.DeleteCategory)

	// Languages
	language := api.Group("/languages")
	language.Get("/", handler.Hello)
	language.Get("/:id", handler.Hello)
	language.Post("/", middleware.Protected(), handler.Hello)
	language.Delete("/:id", middleware.Protected(), handler.Hello)

	// Platforms
	platform := api.Group("/platforms")
	platform.Get("/", handler.Hello)
	platform.Get("/:id", handler.Hello)
	platform.Post("/", middleware.Protected(), handler.Hello)
	platform.Delete("/:id", middleware.Protected(), handler.Hello)

	// Questions
	questions := api.Group("/questions")
	questions.Get("/", handler.Hello)
	questions.Get("/:id", handler.Hello)
	questions.Post("/", middleware.Protected(), handler.Hello)
	questions.Delete("/:id", middleware.Protected(), handler.Hello)

	// Quizzes
	quizzes := api.Group("/quizzes")
	quizzes.Get("/", handler.Hello)
	quizzes.Get("/:id", handler.Hello)
	quizzes.Post("/", middleware.Protected(), handler.Hello)
	quizzes.Delete("/:id", middleware.Protected(), handler.Hello)

}
