package consensus

import (
	"AnderChain/chain"
	"math/big"
)

type Consensus interface {
	FindNonce() int64
}

func NewPow(block chain.Block) Consensus{
	initTarget := big.NewInt(1)
	initTarget.Lsh(initTarget, 255 - DIFFICULTY)
	return Pow{block, initTarget}
}

func NewPos(block chain.Block) Consensus{
	return Pos{Block:block}
}