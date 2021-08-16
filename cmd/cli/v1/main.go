package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"github.com/crashdump/netcp/pkg/entity"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/crashdump/netcp/internal/config"
	"github.com/crashdump/netcp/pkg/netcp"
)

var (
	Version = "" // set during build e.g. -ldflags "-X main.appVersion=v0.1.0"
	Name    = "netcp"

	cfgDefaults = map[string]interface{}{
		"server.host": "127.0.0.1",
		"server.port": "3000",
	}

	flagHost string
	flagPort string
)

func init() {
	flag.StringVar(&flagHost, "server.host", "127.0.0.1", "Hostname")
	flag.StringVar(&flagPort, "server.port", "3000", "Port")
}

func main() {
	log.Printf("%s (%s)", Name, Version)

	cfg, err := loadConfig()

	if len(os.Args) < 2 {
		fmt.Println("expected 'upload', 'list', or 'download' subcommands")
		os.Exit(1)
	}

	url := fmt.Sprintf("http://%s:%s", cfg.GetString("server.host"), cfg.GetString("server.port"))
	o, err := netcp.New(url)
	if err != nil {
		log.Fatal(err.Error())
	}

	switch os.Args[1] {
	case "upload":
		if os.Args[2] == "" {
			log.Fatal("usage: ./netcp upload FILENAME")
		}
		pathIn := os.Args[2]
		fileData, err := ioutil.ReadFile(pathIn)
		if err != nil {
			log.Fatal(err)
		}

		_, fileName := path.Split(pathIn)

		err = o.Upload(&entity.APIBlob{
			Filename: fileName,
			Content:  base64.StdEncoding.EncodeToString(fileData),
		})
		if err != nil {
			log.Fatal(err.Error())
		}

		os.Exit(0)

	case "download":
		if os.Args[2] == "" || os.Args[3] == "" {
			log.Fatal("usage: ./netcp download CODE FILENAME")
		}

		codeIn := os.Args[2]
		pathOut := os.Args[3]

		blob, err := o.DownloadByShortID(codeIn)
		if err != nil {
			log.Fatal(err.Error())
		}

		blobBytes, err := base64.StdEncoding.DecodeString(blob.Content)
		if err != nil {
			log.Fatal(err.Error())
		}

		err = ioutil.WriteFile(pathOut, blobBytes, 0644)
		if err != nil {
			log.Fatal(err.Error())
		}

		os.Exit(0)

	case "list":
		log.Fatal("not implemented yet.")

	default:
		log.Fatal("expected 'upload', 'list' or 'download' subcommands")
	}
}

func loadConfig() (*config.Config, error) {
	env := os.Getenv("ENV")
	if env == "" {
		env = "production"
	}

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
	return cfg, err
}
