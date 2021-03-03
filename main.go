package main

import (
	"DataCerPlatform/AnderChain/chain"
	"fmt"
)
func main() {
	fmt.Println("hello world")

	block0 := chain.CreateGenesis([]byte("helloworld"))
	block1 := chain.NewBlock(block0.Height,block0.Hash,[]byte("helloworld"))

	fmt.Println(block0)
	fmt.Println(block1)

}