package nemparams

// The block height describes the position of the block within the block chain.
// The height is an integer greater than zero.
type BlockHeight struct {
	Height *string `url:"height"`
}
