package route

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func SwaggerRouter(f fiber.Router) {
	f.Get("/docs/*", swagger.Handler)
}
