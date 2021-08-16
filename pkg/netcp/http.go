package netcp

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

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
