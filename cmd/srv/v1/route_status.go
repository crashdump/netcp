package main

import "github.com/gofiber/fiber/v2"

func StatusRouter(app fiber.Router) {
	app.Get("/status", func(c *fiber.Ctx) error {
		return c.JSON(&API{
			Success: true,
			Message: "OK",
		})
	})
}