package route

import (
	"context"
	firebase "firebase.google.com/go/v4"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/crashdump/netcp/internal/config"
	"github.com/crashdump/netcp/internal/handler"
	fileStore "github.com/crashdump/netcp/internal/repository/firebase/files"
	metadataStore "github.com/crashdump/netcp/internal/repository/firebase/metadata"
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
)

func Setup(cfg *config.Config) *fiber.App {
	fbc, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing firebase app: %v", err)
	}

	// Init File store
	br, err := fileStore.NewBlobRepo(fbc, cfg.GetString("bucket.name"))
	if err != nil {
		log.Fatalf("Unable to open blob repository")
	}

	// Init Metadata store
	mr, err := metadataStore.NewMetadataRepo(fbc)
	if err != nil {
		log.Fatalf("Unable to open blob repository")
	}

	bs := handler.NewService(br, mr)

	f := fiber.New()

	// Routes
	api := f.Group("/api/v1")
	api.Get("/status", status(bs))

	//api.Use(middlewares.AuthMiddleware())

	api.Post("/blob", addBlob(bs))
	api.Get("/blob/:id", getBlobByShortID(bs))
	api.Delete("/blob/:id", removeBlob(bs))

	f.Get("/docs/*", swagger.Handler)

	f.Static("/", "./ui/dist",
		fiber.Static{
			Compress:      true,
			Browse:        false,
			Index:         "index.html",
			CacheDuration: 30 * time.Second,
			MaxAge:        3600,
		})

	return f
}