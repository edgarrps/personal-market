package routes

import "github.com/gofiber/fiber/v2"

func Setup(c *fiber.App) {
	app.Post("/api/register", controller.Register)
}
