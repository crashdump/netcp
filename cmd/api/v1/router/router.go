package router

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"

	//"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"time"

	"github.com/crashdump/netcp/internal/config"
	"github.com/crashdump/netcp/internal/controller"
	"github.com/crashdump/netcp/internal/middleware"
	"github.com/crashdump/netcp/internal/repository"
)

func Server(cfg *config.Config) {
	r := gin.New()

	/*
	 * Postgres
	 */
	db := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.GetString("server.postgres.host"),
		cfg.GetString("server.postgres.user"),
		cfg.GetString("server.postgres.password"),
		cfg.GetString("server.postgres.dbname"),
		cfg.GetString("server.postgres.port"),
		cfg.GetString("server.postgres.sslmode"))
	fmt.Println(db)
	repository.OpenGorm(db)
	repository.AutoMigrate()

	/*
	 * CORS
	 */
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
	NewStatus(new(controller.StatusController), v1Route)
	NewAuth(new(controller.AuthController), v1Route)

	v1Route.Use(middlewares.AuthMiddleware()) // Enforce authentication on this namespace
	NewBlob(new(controller.BlobController), v1Route)
	NewUser(new(controller.UserController), v1Route)

	/*
	 * Run Gin
	 */
	r.Run(":" + cfg.GetString("server.port"))
}