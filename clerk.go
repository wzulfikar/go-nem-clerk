package nemclerk

import (
	"fmt"
	nemrequests "local-dev/go-nem-clerk/requests"
	"log"
	"net/http"
	"os"

	"github.com/davecgh/go-spew/spew"
)

// Sample code:
// `nemclerk.NewClient("http://23.228.67.85:7890")`
func NewClient(endpoint string) (*nemrequests.Client, error) {
	return &nemrequests.Client{
		Client:   &http.Client{},
		Endpoint: endpoint,
	}, nil
}

func DefaultClient() (*nemrequests.Client, error) {
	return NewClient(os.Getenv("NEM_CLERK_ENDPOINT"))
}

func SignUnconfirmedTx(c *nemrequests.Client, addr string) {
	tx, err := c.GetUnconfirmedTransactions(addr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Unconfirmed tx:")
	spew.Dump(tx)
}
