package sdk

import (
	"fmt"
	"gopkg.in/resty.v1"
)

type Client struct {
	rcli *resty.Client
	url  string
	auth auth
}

type auth struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
}

var useragent = "netcp-sdk-go/v0.1.0"

func New(url string) (*Client, error) {
	client := Client{
		rcli: resty.New(),
		url:  url,
	}

	client.rcli.SetHeader("User-Agent", useragent)

	am, err := getAuthenticationMethods(url, client)
	if am.Config.Domain == "" {
		return &Client{}, err
	}


	return &client, nil
}

func (n *Client) Authenticate(token OauthToken) error {
	var err error

	return err
}

func (n *Client) IsAuthenticated() bool {
	// TODO: Actually check the token's validity
	return n.auth.ExpiresIn > 1
}


func (n *Client) GetOauthRefreshToken() string {
	return n.auth.RefreshToken
}

func (n *Client) GetOauthAccessToken() string {
	return n.auth.AccessToken
}

func logHttpRequest(rsp *resty.Response, err error) {
	fmt.Println("Request:")
	fmt.Println("  Method     :", rsp.Request.Method)
	fmt.Println("  URL:       :", rsp.Request.URL)
	fmt.Println("  Body:      :", rsp.Request.Body)
	fmt.Println("Response:")
	fmt.Println("  Status     :", rsp.Status())
	fmt.Println("  Time       :", rsp.Time())
	fmt.Println("  Body       :", string(rsp.Body()))
	fmt.Println("  Error      :", err)
	fmt.Println()
}