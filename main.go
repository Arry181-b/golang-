package main

import (
	"AnderChain/chain"
	"fmt"
	"github.com/boltdb/bolt"
)
const BLOCKS = "Ander.db"

func main() {

	//打开数据库文件
	db, err := bolt.Open(BLOCKS, 0600, nil)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close() //xxx.db.lock

	blockChain := chain.CreateChain(db)
	//创世区块
	err = blockChain.CreateGenesis([]byte("helloworld"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//新增一个区块
	err = blockChain.CreateNewBlock([]byte("hello"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}