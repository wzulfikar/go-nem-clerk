package nemparams

type AllTransactions struct {
	Address string `url:"address,omitempty"`
	Hash    string `url:"hash,omitempty"`
	Id      string `url:"id,omitempty"`
}

type UnconfirmedTransactions struct {
	Address string `url:"address"`
}

type RequestPrepareAnnounce struct {
	Transaction string `url:"transaction"`
	PrivateKey  string `url:"privateKey"`
}
