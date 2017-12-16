package nemrequests

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func (c *Client) do(req *http.Request) ([]byte, error) {
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		err = fmt.Errorf("Received unexpected status %d while trying to retrieve the server data with \"%s\"", resp.StatusCode, string(body))
		return nil, err
	}
	return body, nil
}

// Send GET request to given path with optional query params
func (c *Client) Get(path string, optParams ...url.Values) ([]byte, error) {
	req, err := http.NewRequest("GET", c.Endpoint+path, nil)
	if err != nil {
		return nil, err
	}

	if len(optParams) > 0 {
		req.URL.RawQuery = optParams[0].Encode()
	}

	logger("GET", req.URL)
	return c.do(req)
}

func (c *Client) Post(path string, data url.Values) ([]byte, error) {
	req, err := http.NewRequest("POST", c.Endpoint+path, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	logger("POST", req.URL)
	return c.do(req)
}
