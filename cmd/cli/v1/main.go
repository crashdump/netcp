package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/alexflint/go-arg"
	"github.com/crashdump/netcp/internal/config"
	"github.com/crashdump/netcp/pkg/entity"
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

var args struct {
	Upload   *UploadCmd   `arg:"subcommand:upload"`
	Download *DownloadCmd `arg:"subcommand:download"`
	List     *ListCmd     `arg:"subcommand:list"`
	Quiet    bool         `arg:"-q"` // this flag is global to all subcommands
}

type UploadCmd struct {
	File string `arg:"positional"`
}

type ListCmd struct {
}

type DownloadCmd struct {
	Code string `arg:"positional"`
	Path string `arg:"positional"`
}

func main() {
	log.Printf("%s (%s)", Name, Version)

	cfg, err := loadConfig()

	p := arg.MustParse(&args)
	if p.Subcommand() == nil {
		p.Fail("missing subcommand, expected 'upload', 'list', or 'download'.")
	}

	url := fmt.Sprintf("http://%s:%s", cfg.GetString("server.host"), cfg.GetString("server.port"))
	o, err := netcp.New(url)
	if err != nil {
		log.Fatal(err.Error())
	}

	switch {
	case args.Upload != nil:
		fileData, err := ioutil.ReadFile(args.Upload.File)
		if err != nil {
			log.Fatal(err)
		}

		_, fileName := path.Split(args.Upload.File)

		err = o.Upload(&entity.APIBlob{
			Filename: fileName,
			Content:  base64.StdEncoding.EncodeToString(fileData),
		})
		if err != nil {
			log.Fatal(err.Error())
		}

		os.Exit(0)

	case args.Download != nil:
		blob, err := o.DownloadByShortID(args.Download.Code)
		if err != nil {
			log.Fatal(err.Error())
		}

		blobBytes, err := base64.StdEncoding.DecodeString(blob.Content)
		if err != nil {
			log.Fatal(err.Error())
		}

		var filePathOut string
		isd, err := isDirectory(args.Download.Path)
		if isd {
			filePathOut = args.Download.Path + "/" + blob.Filename
		} else {
			filePathOut = args.Download.Path
		}

		err = ioutil.WriteFile(filePathOut, blobBytes, 0644)
		if err != nil {
			log.Fatal(err.Error())
		}
		os.Exit(0)

	case args.List != nil:
		log.Fatal("not implemented yet.")
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

func isDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return fileInfo.IsDir(), err
}