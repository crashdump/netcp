package route

import "github.com/gofiber/fiber/v2"

func UIRouter(f fiber.Router) {
	f.Static("/ui/", "./ui/dist")

	f.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Redirect("/ui/", 301)
	})
}
