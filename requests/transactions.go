package nemrequests

import (
	"encoding/json"

	nemparams "github.com/wzulfikar/go-nem-client/params"
	nemtypings "github.com/wzulfikar/go-nem-client/typings"

	"github.com/google/go-querystring/query"
)

var TransactionPaths = struct {
	All         string
	Incoming    string
	Unconfirmed string
}{
	"/account/transfers/all",
	"/account/transfers/incoming",
	"/account/unconfirmedTransactions",
}

func (c *Client) GetAllTransactions(addr string, optHash string, optId string) (tx *nemtypings.IncomingTransactions, err error) {
	v, _ := query.Values(nemparams.AllTransactions{
		addr, optHash, optId,
	})

	body, err := c.Get(TransactionPaths.All, v)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(body, &tx); err != nil {
		return nil, err
	}

	return tx, nil
}

func (c *Client) GetUnconfirmedTransactions(addr string) (tx *nemtypings.UnconfirmedTransactionMetaDataPair, err error) {
	v, _ := query.Values(nemparams.UnconfirmedTransactions{
		addr,
	})

	body, err := c.Get(TransactionPaths.Unconfirmed, v)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &tx)
	if err != nil {
		return nil, err
	}

	return tx, nil
}
