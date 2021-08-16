package main

import (
	"encoding/base64"
	"github.com/crashdump/netcp/pkg/blob"
	"github.com/crashdump/netcp/pkg/entity"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func BlobRouter(f fiber.Router, service blob.Service) {
	f.Post("/blob", addBlob(service))
	f.Get("/blob/:id", getBlob(service))
	f.Delete("/blob/:id", removeBlob(service))
}

func getBlob(service blob.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.JSON(&API{
				Success: false,
				Message: "invalid UUID",
			})
		}

		fetched, err := service.DownloadByID(id)
		if err != nil {
			return c.JSON(&API{
				Success: false,
				Message: err.Error(),
			})
		}

		return c.JSON(&API{
			Success: true,
			Content: base64.StdEncoding.EncodeToString(fetched.Content),
		})
	}
}

func addBlob(service blob.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entity.Blob
		err := c.BodyParser(&requestBody)
		if err != nil {
			_ = c.JSON(&API{
				Success: false,
				Message: err.Error(),
			})
		}

		err = service.Upload("filex", &requestBody)
		if err != nil {
			return c.JSON(&API{
				Success: false,
				Message: err.Error(),
			})
		} else {
			return c.JSON(&API{
				Success: true,
				Message: "Successfully uploaded",
			})
		}
	}
}

func removeBlob(service blob.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.JSON(&API{
				Success: false,
				Message: "Invalid UUID",
			})
		}

		dberr := service.Remove(id)
		if dberr != nil {
			_ = c.JSON(&API{
				Success: false,
				Message: err.Error(),
			})
		}
		return c.JSON(&API{
			Success: true,
			Message: "Successfully removed",
		})
	}
}
