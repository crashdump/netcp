package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/crashdump/netcp/internal/config"
	"github.com/crashdump/netcp/pkg/netcp"
)

var (
	Version = "" // set during build e.g. -ldflags "-X main.appVersion=v0.1.0"
	Name    = "netcp"

	cfgDefaults = map[string]interface{}{
		"srv.url": "http://127.0.0.1:3000",
	}

	flagUrl string
)

func init() {
	flag.StringVar(&flagUrl, "url", "http://127.0.0.1:3000", "Server URL")
}

func main() {
	log.Printf("%s (%s)", Name, Version)

	// Getting configuration base on environment
	env := os.Getenv("ENV")

	cfg, err := config.New("cli", env, cfgDefaults)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = cfg.Load()
	if err != nil {
		fmt.Println(err)
	}

	err = cfg.ValidateClient()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	o, err := netcp.New(cfg.GetString("srv.url"))
	if err != nil {
		log.Fatalf("cannot fetch bearer token: %v", err)
	}

	fmt.Printf("Status: %s", o.Status())
}
