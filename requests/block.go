// Wrapper for nemclient and nemparams.
// Method name convention:
// {HTTP Method}{endpoint name}. e,g. `GetAllTransactions`
package nemrequests

import (
	"encoding/json"

	nemparams "github.com/wzulfikar/go-nem-client/params"
	nemtypings "github.com/wzulfikar/go-nem-client/typings"

	"github.com/google/go-querystring/query"
)

var BlockPaths = struct {
	Public string
}{
	"/block/at/public",
}

// Gets a block from the chain that has the given height
func (c *Client) PostBlockOfHeight(blockHeight string) (block *nemtypings.Block, err error) {
	v, _ := query.Values(nemparams.BlockHeight{
		&blockHeight,
	})

	body, err := c.Post(BlockPaths.Public, v)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(body, &block); err != nil {
		return nil, err
	}
	return
}
