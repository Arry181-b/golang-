package chain

import "time"

const VERSION = 0x00
//区块的结构定义
type Block struct {
	Height int64
	Version int64
	PrevHash [32]byte
	Hash [32]byte
	//默克尔根
	TimeStamp int64
	//Difficulty int64
	Nonce int64
	//区块体
	Data []byte
}
/**
 *生成创世区块的函数
 */
func CreateGenesis(data []byte) Block {
	genesis := Block{
		Height:0,
		Version: VERSION,
		PrevHash: [32]byte{
			0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},
		TimeStamp: time.Now().Unix(),
		Data:   data,

		}
		//todo 计算并设置nonce 计算并设置hash

		return genesis
	}

/**
 *生成新区快的功能函数
 */
func NewBlock(height int64, prev [32]byte, data []byte) Block{
	newBlock := Block{
		Height: height + 1,
		Version: VERSION,
		PrevHash: prev,
		TimeStamp: time.Now().Unix(),
		Data:  data,
	}
	//todo 设置哈希,寻找并设置nonce
	return newBlock
}