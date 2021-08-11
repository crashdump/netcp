package helpers

import (
	"errors"
	"net/http"
	"time"
)

func WaitForAPI(url string) error {
	c := 1
	for c < 10 {
		time.Sleep(2 * time.Second)
		c += c

		client := http.Client{
			Timeout: 1 * time.Second,
		}
		resp, err := client.Get(url)
		if err != nil {
			continue
		}

		if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
			return nil
		}
	}
	return errors.New("unable to connect")
}