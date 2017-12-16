package nemrequests

import "net/http"

type Client struct {
	Client   *http.Client
	Endpoint string
}
