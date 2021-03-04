package consensus

import (
	"math/big"
)

type Consensus interface {
	FindNonce() ([32]byte,int64)
}

/**
  * 定义区块结构体的标准  接口
 */
type BlockInterface interface {
	GetHeight() int64
	GetVersion() int64
	GetTimeStamp() int64
	GetPrevHash() [32]byte
	GetData() []byte
}

func NewPow(block BlockInterface) Consensus{
	initTarget := big.NewInt(1)
	initTarget.Lsh(initTarget, 255 - DIFFICULTY)
	return Pow{block, initTarget}
}
