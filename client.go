package nemclient

import (
	"net/http"

	nemrequests "github.com/wzulfikar/go-nem-client/requests"
)

// Sample code:
// `nemrequests.NewClient("http://23.228.67.85:7890")`
func NewClient(endpoint string) (*nemrequests.Client, error) {
	return &nemrequests.Client{
		Client:   &http.Client{},
		Endpoint: endpoint,
	}, nil
}
