package route

import (
	"encoding/base64"
	"fmt"
	"github.com/crashdump/netcp/internal/handler"
	"github.com/crashdump/netcp/pkg/entity"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func BlobRouter(f fiber.Router, service handler.Service) {
	f.Post("/blob", addBlob(service))
	f.Get("/blob/:id", getBlobByShortID(service))
	f.Delete("/blob/:id", removeBlob(service))
}

func getBlobByShortID(service handler.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		blob, meta, err := service.DownloadByShortID(c.Params("id"))
		if err != nil {
			return c.JSON(&entity.API{
				Success: false,
				Message: err.Error(),
			})
		}

		return c.JSON(&entity.APIBlob{
			Filename: meta.Filename,
			Content:  base64.StdEncoding.EncodeToString(blob.Content),
		})
	}
}

func getBlobByID(service handler.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.JSON(&entity.API{
				Success: false,
				Message: "invalid UUID",
			})
		}

		blob, meta, err := service.DownloadByID(id)
		if err != nil {
			return c.JSON(&entity.API{
				Success: false,
				Message: err.Error(),
			})
		}

		return c.JSON(&entity.APIBlob{
			Filename: meta.Filename,
			Content:  base64.StdEncoding.EncodeToString(blob.Content),
		})
	}
}

func addBlob(service handler.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entity.APIBlob
		err := c.BodyParser(&requestBody)
		if err != nil {
			_ = c.JSON(&entity.API{
				Success: false,
				Message: err.Error(),
			})
		}

		blobdata, err := base64.StdEncoding.DecodeString(requestBody.Content)
		if err != nil {
			return c.JSON(&entity.API{
				Success: false,
				Message: err.Error(),
			})
		}

		blob := &entity.Blob{
			ID:      uuid.New(),
			Content: blobdata,
		}

		shortid, err := service.Upload(requestBody.Filename, blob)
		if err != nil {
			return c.JSON(&entity.API{
				Success: false,
				Message: err.Error(),
			})
		} else {
			return c.JSON(&entity.API{
				Success: true,
				Message: fmt.Sprintf("Successfully uploaded, short code is %s", shortid),
			})
		}
	}
}

func removeBlob(service handler.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.JSON(&entity.API{
				Success: false,
				Message: "Invalid UUID",
			})
		}

		dberr := service.Remove(id)
		if dberr != nil {
			_ = c.JSON(&entity.API{
				Success: false,
				Message: err.Error(),
			})
		}
		return c.JSON(&entity.API{
			Success: true,
			Message: "Successfully removed",
		})
	}
}
