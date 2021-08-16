package main

import (
	"fmt"
	"github.com/crashdump/netcp/cmd/srv/v1/route"
	"log"
	"os"

	_ "github.com/crashdump/netcp/api"
	"github.com/crashdump/netcp/internal/config"
)

var (
	Version = "" // set during build e.g. -ldflags "-X main.appVersion=v0.1.0"
	Name    = "netcp-srv"

	cfgDefaults = map[string]interface{}{
		"server.host": "127.0.0.1",
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

	err = cfg.ValidateServer()
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

	app := route.Setup(cfg)

	port = cfg.GetString("server.port")
	log.Printf("server listening on :%s", port)
	log.Fatal(app.Listen(":" + port))
}
