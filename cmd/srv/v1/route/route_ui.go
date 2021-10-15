package route

import (
	"github.com/crashdump/netcp/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func uiRedirect(service handler.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.Redirect("/ui/", 301)
	}
}