package main

import (
	"fmt"
	"os"

	"github.com/crashdump/netcp/cmd/api/v1/router"
	"github.com/crashdump/netcp/internal/config"
	"github.com/gin-gonic/gin"
)

var (
	Version = "" // set during build e.g. -ldflags "-X main.appVersion=v0.1.0"
	Name    = "netcp-api"

	cfgDefaults = map[string]interface{}{
		"api.url":                  "http://127.0.0.1:3000",
		"server.postgres.host":     "localhost",
		"server.postgres.user":     "postgres",
		"server.postgres.password": "postgres",
		"server.postgres.dbname":   "netcp-dev",
		"server.postgres.port":     "5432",
		"server.postgres.sslmode":  "disable",
		"server.tls.enabled":       "false",
		"server.port":              "3000",
	}
)

func main() {
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
	}

	// PORT environment variable is provided by Cloud Run.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	cfg.Set("server.port", port)

	router.Server(cfg)
}