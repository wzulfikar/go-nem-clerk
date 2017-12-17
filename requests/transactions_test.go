package nemrequests

import (
	"fmt"
	nemclient "local-dev/go-nem-client"
	"testing"

	. "github.com/franela/goblin"
)

func TestTransactions(t *testing.T) {
	g := Goblin(t)

	c, _ := nemclient.NewClient("http://23.228.67.85:7890")

	g.Describe("Testing transactions", func() {
		g.It("Should fetch all transactions", func() {
			tx, err := c.GetAllTransactions("TC6Z5CY3TX4V3UCTWFZLCSEH6WOSIZ4PU2RURAHN", "", "")
			if err != nil {
				g.Fail(err)
			}

			fmt.Println("  → Tx found: ", len(tx.Data))
		})

		g.It("Should fetch unconfirmed transactions", func() {
			tx, err := c.GetUnconfirmedTransactions("TDJDSDTS2UD2TLEZEO5DEG6BABJ64M6FTZTHQI6E")
			if err != nil {
				g.Fail(err)
			}
			fmt.Println("  → Tx found: ", len(tx.Data))
		})
	})
}
