package netcp

import (
	"fmt"
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

	client.rcli.SetHeader("User-Agent", "netcp-go/v0.1.0")

	return &client, nil
}

func (c *Client) Status() error {
	_, err := c.rcli.R().Get(fmt.Sprintf("%s/api/v1/status", c.url))
	if err != nil {
		return err
	}
	return nil
}
