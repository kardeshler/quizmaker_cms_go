package main

import (
	"quizcms/pkg/database"
	"quizcms/pkg/router"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	database.ConnectDB()
	defer database.DB.Close()

	router.SetupRoutes(app)
	app.Listen(3000)
}
