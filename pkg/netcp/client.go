package netcp

import (
	"fmt"
	"github.com/crashdump/netcp/pkg/entity"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	rcli *resty.Client
	url  string
}

func New(url string) (*Client, error) {
	client := Client{
		rcli: resty.New(),
		url:  url,
	}

	client.rcli.
		SetHeader("User-Agent", "netcp-go/v0.1.0").
		SetHeader("Content-Type", "application/json")

	return &client, nil
}

func (c *Client) Status() error {
	r, err := c.rcli.R().
		Get(fmt.Sprintf("%s/api/v1/status", c.url))
	if err != nil {
		return err
	}
	fmt.Println(r.String())

	return nil
}

func (c *Client) Upload(blob *entity.APIBlob) error {
	r, err := c.rcli.R().
		SetBody(blob).
		Post(fmt.Sprintf("%s/api/v1/blob", c.url))
	if err != nil {
		return err
	}
	fmt.Println(r.String())

	return nil
}

func (c *Client) DownloadByShortID(id string) (*entity.APIBlob, error) {
	var blob entity.APIBlob
	r, err := c.rcli.R().
		SetResult(&blob).
		Get(fmt.Sprintf("%s/api/v1/blob/%s", c.url, id))
	if err != nil {
		return &blob, err
	}

	fmt.Println(r.String())

	return &blob, nil
}
