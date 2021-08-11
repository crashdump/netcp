### netcp-sdk-go

A Golang SDK for Netcp API.

For best compatibility, please use Go >= 1.12.

### API List

Name | Description | Status
------------ | ------------ | ------------
Rest API | Connect over HTTP, authenticate, and fetch results | <input type="checkbox" checked>
Organisations | CRUD | <input type="checkbox" checked>
Agents | CRUD | <input type="checkbox" unchecked>

### Installation

```shell
go get github.com/netcp-hq/netcp-app/pkg/sdk/v1
```

### Importing

```golang
import (
    "github.com/crashdump/netcp/pkg/sdk/v1
)
```

### Use

```golang

netcli := netcp.New()
netcli.AuthWithOauth2(OauthToken)

```

`AuthWithOauth2(OauthToken)` will always try to use the access_token if present. Otherwise, it will
request one with the refresh_token. Finally, if neither exist, it'll request both through a web-browser
based device authentication.