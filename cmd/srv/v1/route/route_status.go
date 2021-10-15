package route

import (
	"github.com/crashdump/netcp/internal/handler"
	"github.com/crashdump/netcp/pkg/entity"
	"github.com/gofiber/fiber/v2"
)

func status(service handler.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(&entity.API{
			Success: true,
			Message: "OK",
		})
	}
}