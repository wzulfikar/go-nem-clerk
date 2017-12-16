package nemparams

type AllTransactions struct {
	Address string `url:"address"`
	Hash    string `url:"hash,omitempty"`
	Id      string `url:"hash,omitempty"`
}

type UnconfirmedTransactions struct {
	Address string `url:"address"`
}
