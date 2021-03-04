package router

import (
	"quizcms/pkg/handler"
	"quizcms/pkg/middleware"

	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
)

// SetupRoutes sets the routes up in the rest api
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
	category.Put("/:id", middleware.Protected(), handler.UpdateCategory)
	category.Delete("/:id", middleware.Protected(), handler.DeleteCategory)

	// Languages
	language := api.Group("/languages")
	language.Get("/", middleware.Protected(), handler.GetAllLanguages)
	language.Get("/:id", middleware.Protected(), handler.GetLanguage)
	language.Post("/", middleware.Protected(), handler.CreateLanguage)
	language.Put("/:id", middleware.Protected(), handler.UpdateLanguage)
	language.Delete("/:id", middleware.Protected(), handler.DeleteLanguage)

	// Platforms
	platform := api.Group("/platforms")
	platform.Get("/", middleware.Protected(), handler.GetAllPlatforms)
	platform.Get("/:id", middleware.Protected(), handler.GetPlatform)
	platform.Post("/", middleware.Protected(), handler.CreatePlatform)
	platform.Put("/:id", middleware.Protected(), handler.UpdatePlatform)
	platform.Delete("/:id", middleware.Protected(), handler.DeletePlatform)
	platform.Get("/:id/questions", middleware.Protected(), handler.GetQuizzesOfPlatform)
	platform.Get("/:id/quizzes", middleware.Protected(), handler.GetQuizzesOfPlatform)

	// Questions
	questions := api.Group("/questions")
	questions.Get("/", middleware.Protected(), handler.GetAllQuestions)
	questions.Get("/:id", middleware.Protected(), handler.GetQuestion)
	questions.Post("/", middleware.Protected(), handler.CreateQuestion)
	questions.Put("/:id", middleware.Protected(), handler.UpdateQuestion)
	questions.Delete("/:id", middleware.Protected(), handler.DeleteQuestion)
	// todo: add questions of a platform, of a category, of a language, etc

	// Quizzes
	quizzes := api.Group("/quizzes")
	quizzes.Get("/", middleware.Protected(), handler.GetAllQuizzes)
	quizzes.Get("/:id", middleware.Protected(), handler.GetQuiz)
	quizzes.Post("/", middleware.Protected(), handler.CreateQuiz)
	quizzes.Put("/", middleware.Protected(), handler.UpdateQuiz)
	quizzes.Delete("/:id", middleware.Protected(), handler.DeleteQuiz)
	// todo: add quizzes of a platform, of a category, of a language, etc
}
