package netcp_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/crashdump/netcp/cmd/srv/v1/route"
	netcpcfg "github.com/crashdump/netcp/internal/config"
	"github.com/crashdump/netcp/internal/helper"
	"github.com/crashdump/netcp/pkg/netcp"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	client *netcp.Client
}

const (
	apiUrl = "http://127.0.0.1:3000/api/v1/"
)

var cfgDefaults = map[string]interface{}{
	"server.port": "3000",
	"server.host": "127.0.0.1",
	"bucket.name": "cloudcopy-it.appspot.com",
}

func TestSDKTestSuite(t *testing.T) {
	cfg, err := netcpcfg.New("srv", "unittest", cfgDefaults)
	if err != nil {
		t.Error(err)
	}
	err = cfg.ValidateServer()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	f := route.Setup(cfg)
	go func() {
		err := f.Listen(":" + cfg.GetString("server.port"))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	err = helper.WaitForAPI(apiUrl + "status")
	if err != nil {
		t.Error("Unable to start Server backend")
	}

	suite.Run(t, new(TestSuite))
}

func (ts *TestSuite) SetupTest() {
	var err error

	ts.client, err = netcp.New("http://127.0.0.1:3000/api/v1/")
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
			client, err := netcp.New(tt.args.url)
			ts.NoError(err)

			err = client.Status()
			if tt.wantErr {
				ts.Error(err)
			} else {
				ts.NoError(err)
			}
		})
	}
}
