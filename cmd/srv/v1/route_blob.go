package main

import (
	"github.com/crashdump/netcp/pkg/blob"
	"github.com/crashdump/netcp/pkg/entity"
	"github.com/gofiber/fiber/v2"
)


func BlobRouter(app fiber.Router, service blob.Service) {
	app.Get("/blob", getBlob(service))
	app.Post("/blob", addBlob(service))
	app.Delete("/blob", removeBlob(service))
}

func getBlob(service blob.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetched, err := service.Download()
		var result fiber.Map
		if err != nil {
			result = fiber.Map{
				"status": false,
				"error":  err.Error(),
			}
		} else {
			result = fiber.Map{
				"status": true,
				"Blobs":  fetched,
			}
		}

		return c.JSON(&result)
	}
}

func addBlob(service blob.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entity.Blob
		err := c.BodyParser(&requestBody)
		if err != nil {
			_ = c.JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}
		result, dberr := service.Upload(&requestBody)
		return c.JSON(&fiber.Map{
			"status": result,
			"error":  dberr,
		})
	}
}


func removeBlob(service blob.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entity.Blob
		err := c.BodyParser(&requestBody)
		BlobID := requestBody.ID
		if err != nil {
			_ = c.JSON(&fiber.Map{
				"status": false,
				"error":  err,
			})
		}
		dberr := service.RemoveBlob(BlobID)
		if dberr != nil {
			_ = c.JSON(&fiber.Map{
				"status": false,
				"error":  err,
			})
		}
		return c.JSON(&fiber.Map{
			"status":  false,
			"message": "updated successfully",
		})
	}
}