package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title Go Superindo API Product
// @description This is a sample server for Lion Superindo API Product Service.
// @BasePath /api/v1

// @schemes  http
// @host 127.0.0.1:3000
// @Version 1.0.0
func main() {
	app := fiber.New()
    app.Use(cors.New())

	app.Get("/swagger/*", fiberSwagger.WrapHandler)
	
	pathApi := app.Group("/api/v1")

	// get health
	pathApi.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "ok",
		})
	})

	log.Fatal(app.Listen(":3000"))
}