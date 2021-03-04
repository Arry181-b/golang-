package consensus

import (
	"AnderChain/chain"
	"fmt"
)

type Pos struct {
	Block chain.Block
}

func (pos Pos) FindNonce()int64 {
	fmt.Println("这是共识机制Pos算法的实现")
	return 0
}