package main

import (
	"fmt"
	"github.com/crashdump/netcp/internal/config"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoutes(t *testing.T) {
	tests := []struct {
		description string

		// Test input
		route string

		// Expected output
		expectedError bool
		expectedCode  int
		expectedBody  string
	}{
		// TODO: web ui
		//{
		//	description:   "index redirects to ui",
		//	route:         "/",
		//	expectedError: false,
		//	expectedCode:  301,
		//	expectedBody:  "",
		//},
		//{
		//	description:   "ui",
		//	route:         "/ui",
		//	expectedError: false,
		//	expectedCode:  200,
		//	expectedBody:  "something",
		//},
		{
			description:   "API status",
			route:         "/api/v1/status",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "OK",
		},
		{
			description:   "non existing route",
			route:         "/i-dont-exist",
			expectedError: false,
			expectedCode:  404,
			expectedBody:  "Cannot GET /i-dont-exist",
		},
	}

	cfg, err := config.New("cli", "unittest", cfgDefaults)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Setup the app as it is done in the main function
	app := setup(cfg)

	// Iterate through test single test cases
	for _, test := range tests {
		// Create a new http request with the route
		// from the test case
		req, _ := http.NewRequest(
			"GET",
			test.route,
			nil,
		)

		// Perform the request plain with the app.
		// The -1 disables request latency.
		res, err := app.Test(req, -1)

		// verify that no error occurred, that is not expected
		assert.Equalf(t, test.expectedError, err != nil, test.description)

		// As expected errors lead to broken responses, the next
		// test case needs to be processed
		if test.expectedError {
			continue
		}

		// Verify if the status code is as expected
		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

		// Read the response body
		body, err := ioutil.ReadAll(res.Body)

		// Reading the response body should work everytime, such that
		// the err variable should be nil
		assert.Nilf(t, err, test.description)

		// Verify, that the response body contains the expected content
		assert.Containsf(t, string(body), test.expectedBody, test.description)
	}
}
