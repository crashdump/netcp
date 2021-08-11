package sdk_test

import (
	"errors"
	"fmt"
	"github.com/crashdump/netcp/internal/helpers"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"

	"github.com/crashdump/netcp/cmd/api/v1/router"
	apicfg "github.com/crashdump/netcp/internal/config"
	netcp "github.com/crashdump/netcp/pkg/sdk/v1"
)

type TestSuite struct {
	suite.Suite
	client *netcp.Client
}

const (
	apiUrl = "http://127.0.0.1:3000/api/v1/"
	// this is a test secret and a test organisation
	refreshToken = "G5hGNqg2kcWA2hgCuz7Mo7_baLBDAebix4xqcqoqq_XcY" //nolint:gosec
)

var cfgDefaults = map[string]interface{}{
	"server.postgres.host":     "localhost",
	"server.postgres.user":     "postgres",
	"server.postgres.password": "postgres",
	"server.postgres.dbname":   "netcp-test",
	"server.postgres.port":     "5432",
	"server.postgres.sslmode":  "disable",
	"server.port":              "3000",
	"server.hostname":          "127.0.0.1",
}

func TestSDKTestSuite(t *testing.T) {
	cfg, err := apicfg.New("app", "test", cfgDefaults)
	if err != nil {
		t.Error(err)
	}
	cfg.Load()
	go router.Server(cfg)

	err = helpers.WaitForAPI(apiUrl + "status")
	if err != nil {
		t.Error("Unable to start Server backend")
	}

	suite.Run(t, new(TestSuite))
}

func (ts *TestSuite) SetupTest() {
	var err error

	ts.client, err = netcp.New("http://127.0.0.1:3000/api/v1/")
	ts.NoError(err)

	err = ts.client.Authenticate(netcp.OauthToken{
		RefreshToken: refreshToken,
	})
	ts.NoError(err)
}

func (ts *TestSuite) TestOmniclient_New() {
	type args struct {
		url          string
		refreshToken string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Invalid server URL",
			args: args{
				url: "http://127.0.0.1:1/invalid",
			},
			wantErr: true,
		},
		{
			name: "Valid server URL",
			args: args{
				url: apiUrl,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		ts.Run(tt.name, func() {
			var err error

			ts.client, err = netcp.New("http://127.0.0.1:3000/api/v1/")
			if !tt.wantErr {
				ts.NoError(err)
			}
			if tt.wantErr && ts.client.IsAuthenticated() {
				ts.Error(errors.New("shouldn't be authenticated"))
			}
		})
	}
}

func (ts *TestSuite) TestOmniclient_Authenticate() {
	type args struct {
		token netcp.OauthToken
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "No access token and invalid refresh token",
			args: args{
				token: netcp.OauthToken{
					AccessToken:  "",
					RefreshToken: "123abcd",
				},
			},
			wantErr: true,
		},
		{
			name: "Invalid access token and invalid refresh token",
			args: args{
				token: netcp.OauthToken{
					AccessToken:  "123abcd",
					RefreshToken: "efg456",
				},
			},
			wantErr: true,
		},
		//{
		//	name: "Invalid access token and no refresh token",
		//	args: args{
		//		token: netcp.OauthToken{
		//			AccessToken:  "123abcd",
		//			RefreshToken: "",
		//		},
		//	},
		//	wantErr: true,
		//},

	}
	for _, tt := range tests {
		ts.Run(tt.name, func() {
			err := ts.client.Authenticate(tt.args.token)

			fmt.Println(err)
			if tt.wantErr {
				ts.Error(err)
				return
			}
			// check if we're authenticated
			isAuth := ts.Equal(true, ts.client.IsAuthenticated())
			if !isAuth {
				fmt.Println("Abort tests since we won't be able to run the rest without auth")
				os.Exit(1)
			}
		})
	}
}