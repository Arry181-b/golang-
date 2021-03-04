package main

import (
	"AnderChain/chain"
	"fmt"
)
func main() {
	fmt.Println("hello world")
	blockChain := chain.CreateChainWithGenesis([]byte("helloworld"))
	blockChain.CreateNewBlock([]byte("hello"))
	fmt.Println("区块链的长度",len(blockChain.Blocks))
	fmt.Println("区块0的哈希值：", blockChain.Blocks[0])
	fmt.Println("区块1的哈希值：", blockChain.Blocks[1])



}