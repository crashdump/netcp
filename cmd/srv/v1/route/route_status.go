package route

import (
	"github.com/crashdump/netcp/pkg/entity"
	"github.com/gofiber/fiber/v2"
)

func StatusRouter(app fiber.Router) {
	app.Get("/status", func(c *fiber.Ctx) error {
		return c.JSON(&entity.API{
			Success: true,
			Message: "OK",
		})
	})
}
