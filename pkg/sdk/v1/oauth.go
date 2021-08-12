package sdk

import (
	"errors"
	"fmt"
	"github.com/crashdump/netcp/internal/model"
	"log"
	"time"

	"github.com/crashdump/netcp/internal/browser"
	"gopkg.in/resty.v1"
)

// Ensure to have:
// - Login to Auth0
// - Create a Native Application
// - Enable connections.
// - Enable device Grant types. (Application settings -> Advanced -> Grant Types-> Check device )

// Configured Device Flow User Code Format as per:
// https://auth0.com/docs/get-started/dashboard/configure-device-user-code-settings

// Ensure "Allow Offline Access" is enabled in "Access Settings"

type oauthTokenReq struct {
	ClientID     string `json:"client_id,omitempty"`
	GrantType    string `json:"grant_type,omitempty"`
	DeviceCode   string `json:"device_code,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type OauthToken struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	TokenType    string `json:"token_type,omitempty"`
	IDToken      string `json:"id_token,omitempty"`
	Scope        string `json:"scope,omitempty"`
	ExpiresIn    int    `json:"expires_in,omitempty"`
}

var clientID = "9mRCkgSZkDdJjECOcJivRXdTF0crNbDZ"
var audience = "http://127.0.0.1:3000/api/v1/"

func (n *Client) authenticateWithDevice() error {
	token, err := getOauthDeviceCode()
	if err != nil {
		return err
	}

	n.auth.RefreshToken = token.RefreshToken
	n.auth.AccessToken = token.AccessToken
	n.auth.TokenType = token.TokenType
	return nil
}

func (n *Client) authenticateWithAccessToken(accessToken string) error {
	n.auth.AccessToken = accessToken
	return nil
}

func (n *Client) authenticateWithRefreshToken(refreshToken string) error {
	token, err := getOauthToken(oauthTokenReq{
		ClientID:     clientID,
		GrantType:    "refresh_token",
		RefreshToken: refreshToken,
	})
	if err != nil {
		return err
	}

	n.auth.AccessToken = token.AccessToken
	return nil
}

func getOauthDeviceCode() (OauthToken, error) {
	deviceCodeReq := struct {
		ClientID string `json:"client_id,omitempty"`
		Scope    string `json:"scope,omitempty"`
		Audience string `json:"audience,omitempty"`
	}{
		ClientID: clientID,
		Scope:    "openid offline_access",
		Audience: audience,
	}
	var deviceCodeRsp = struct {
		DeviceCode              string `json:"device_code,omitempty"`
		UserCode                string `json:"user_code,omitempty"`
		VerificationURI         string `json:"verification_uri,omitempty"`
		ExpiresIn               int    `json:"expires_in,omitempty"`
		Interval                int    `json:"interval,omitempty"`
		VerificationURIComplete string `json:"verification_uri_complete,omitempty"`
	}{}

	client := resty.New()
	rsp, err := client.R().
		SetBody(deviceCodeReq).
		SetResult(&deviceCodeRsp).
		Post("https://netcp-dev.eu.auth0.com/oauth/device/code")

	if err != nil {
		return OauthToken{}, err
	}

	if rsp.StatusCode() > 399 {
		log.Println(rsp)
		return OauthToken{}, err
	}

	browser.Open(deviceCodeRsp.VerificationURIComplete)

	// Poll here to check if click was confirmed and valid.
	for {
		time.Sleep(5 * time.Second)
		fmt.Println("Checking if authentication is successful..")

		token, err := getOauthToken(oauthTokenReq{
			ClientID:   clientID,
			GrantType:  "urn:ietf:params:oauth:grant-type:device_code",
			DeviceCode: deviceCodeRsp.DeviceCode,
		})
		if err != nil {
			continue
		}
		if rsp.StatusCode() > 399 {
			continue
		}

		return token, err
	}
}

func getOauthToken(tokenReq oauthTokenReq) (OauthToken, error) {
	var tokenRsp = OauthToken{}

	client := resty.New()
	rsp, err := client.R().
		SetBody(tokenReq).
		SetResult(&tokenRsp).
		Post("https://netcp-dev.eu.auth0.com/oauth/token")

	logHttpRequest(rsp, err)

	if err != nil {
		return tokenRsp, err
	}
	if rsp.StatusCode() > 399 {
		return tokenRsp, errors.New("received a non-2xx http code from the server")
	}

	return tokenRsp, nil
}

func getAuthenticationMethods(url string, netcp Client) (model.Auth, error) {
	var amm []model.Auth
	var am model.Auth
	var ar model.APIError

	// Fetch authentication details from backend
	rsp, err := netcp.rcli.R().
		SetError(&model.APIError{}).
		SetResult(&amm).
		SetError(&ar).
		Get(url + "auth")

	// TODO: better management of log levels for debug
	logHttpRequest(rsp, err)

	if err != nil || ar.Code > 0 || rsp.StatusCode() > 399 {
		return am, err
	}

	// TODO: Extract this and implement support for other authentication types
	for i := range amm {
		if amm[i].Type == "oauth2" {
			fmt.Println(amm[i])
			return amm[i], nil
		}
	}

	return am, errors.New("no authentication method(s) advertised by the server")
}
