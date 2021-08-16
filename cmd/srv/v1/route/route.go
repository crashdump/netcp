package route

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go/v4"
	"github.com/crashdump/netcp/internal/config"
	"github.com/crashdump/netcp/internal/handler"
	fileStore "github.com/crashdump/netcp/internal/repository/firebase/files"
	metadataStore "github.com/crashdump/netcp/internal/repository/firebase/metadata"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	// CORS
	url := fmt.Sprintf("http://%s:%s",
		cfg.GetString("server.host"),
		cfg.GetString("server.port"),
	)
	f.Use(cors.New(cors.Config{
		AllowOrigins: url,
		AllowMethods: "GET,POST,DELETE",
	}))

	// Routes
	UIRouter(f)
	SwaggerRouter(f)

	api := f.Group("/api/v1")
	StatusRouter(api)

	//api.Use(middlewares.AuthMiddleware())

	BlobRouter(api, bs)

	return f
}
