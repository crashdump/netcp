package main

import (
	"context"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	_ "github.com/crashdump/netcp/api"
	"github.com/crashdump/netcp/internal/config"
	blobStore "github.com/crashdump/netcp/internal/repository/firebase/storage"
	//middlewares "github.com/crashdump/netcp/internal/middleware"
	blobService "github.com/crashdump/netcp/pkg/blob"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var (
	Version = "" // set during build e.g. -ldflags "-X main.appVersion=v0.1.0"
	Name    = "netcp-srv"

	cfgDefaults = map[string]interface{}{
		"srv.url":     "http://127.0.0.1:3000",
		"server.port": "3000",
		"bucket.name": "cloudcopy-it.appspot.com",
	}
)

func main() {
	log.Printf("%s (%s)", Name, Version)

	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "production"
	}

	cfg, err := config.New("srv", env, cfgDefaults)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = cfg.Load()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// PORT environment variable is provided by Cloud Run.
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	cfg.Set("server.port", port)

	app := setup(cfg)

	port = cfg.GetString("server.port")
	log.Printf("server listening on :%s", port)
	log.Fatal(app.Listen(":" + port))
}

func setup(cfg *config.Config) *fiber.App {
	fbc, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing firebase app: %v", err)
	}

	br, err := blobStore.NewBlobRepo(fbc, cfg.GetString("bucket.name"))
	if err != nil {
		log.Fatalf("Unable to open blob repository")
	}
	bs := blobService.NewService(br)

	f := fiber.New()

	// CORS
	url := fmt.Sprintf("http://%s:%s",
		cfg.GetString("server.hostname"),
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
