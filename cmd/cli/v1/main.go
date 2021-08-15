package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/crashdump/netcp/internal/config"
	//"github.com/crashdump/netcp/pkg/sdk/v1"
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
	flag.StringVar(&flagUrl, "url", "http://127.0.0.1:3000", "URL of Netcp API")
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

	// ...
}
