package main

import (
	"AnderChain/chain"
	"fmt"
)
func main() {
	blockChain := chain.CreateChainWithGenesis([]byte("helloworld"))
	blockChain.CreateNewBlock([]byte("hello"))
	//fmt.Println("区块链的长度",len(blockChain.Blocks))

	fmt.Println("区块0：", blockChain.Blocks[0])
	//fmt.Println("区块1的哈希值：", blockChain.Blocks[1])

	firstBlock := blockChain.Blocks[0]
	firstBytes, err := firstBlock.Serialize()
	if err != nil {
		panic(err.Error())
	}
	//直接调用反序列化验证逆过程
	deFirstBlock, err := chain.Deserialize(firstBytes)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(deFirstBlock.Data))



}