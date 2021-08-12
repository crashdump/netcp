package main

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"fmt"
	"github.com/crashdump/netcp/internal/controller"
	middlewares "github.com/crashdump/netcp/internal/middleware"
	"github.com/crashdump/netcp/internal/repository"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"google.golang.org/appengine/log"
	"os"
	"time"

	"github.com/crashdump/netcp/cmd/api/v1/route"
	"github.com/crashdump/netcp/internal/config"
	"github.com/gin-gonic/gin"
)

var (
	Version = "" // set during build e.g. -ldflags "-X main.appVersion=v0.1.0"
	Name    = "netcp-api"

	cfgDefaults = map[string]interface{}{
		"api.url":     "http://127.0.0.1:3000",
		"server.port": "3000",
	}
)

func main() {
	ctx := context.Background()

	/*
	 * Configuration
	 */
	env := os.Getenv("GO_ENV")

	if env == "" || env == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode("debug")
	}

	cfg, err := config.New("app", env, cfgDefaults)
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
		port = "8080"
	}
	cfg.Set("server.port", port)


	/*
	 * Firebase
	 */
	fbcli, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Criticalf(ctx, "error initializing firebase app: %v", err)
		os.Exit(1)
	}

	err = repository.Open(fbcli)
	if err != nil {
		log.Criticalf(ctx, "error initializing firebase app: %v", err)
		os.Exit(1)
	}

	/*
	 * Gin
	 */
	r := gin.New()

	// CORS
	url := fmt.Sprintf("http://%s:%s",
		cfg.GetString("server.hostname"),
		cfg.GetString("server.port"),
	)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{url},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	/*
	 * Routes
	 */
	r.Use(static.Serve("/ui/", static.LocalFile("ui/dist", true)))
	r.Use(static.Serve("/swagger/", static.LocalFile("cmd/api/docs", false)))

	// Redirect / to /ui/
	r.GET("/", func(c *gin.Context) {
		c.Redirect(301, "/ui/")
	})

	// Server - Authenticated
	v1Route := r.Group("/api/v1")
	route.NewStatus(new(controller.StatusController), v1Route)
	route.NewAuth(new(controller.AuthController), v1Route)

	v1Route.Use(middlewares.AuthMiddleware()) // Enforce authentication on this namespace
	route.NewBlob(new(controller.BlobController), v1Route)
	route.NewUser(new(controller.UserController), v1Route)

	// RUN
	r.Run(":" + cfg.GetString("server.port"))
}