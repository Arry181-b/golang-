package chain

import (
	"AnderChain/consensus"
	"time"
)

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



func (block Block) GetHeight() int64 {
	return block.Height
}

func (block Block)GetVersion() int64 {
	return block.Version
}

func (block Block)GetTimeStamp() int64 {
	return block.TimeStamp
}

func (block Block)GetPrevHash() [32]byte {
	return block.PrevHash
}

func (block Block)GetData() []byte {
	return block.Data
}

/**
 * 计算区块的hash值，并赋值
 */
//func (block *Block) CalculateBlockHash() {
//	heightByte, _ := utils.Int2Byte(block.Height)
//	versionByte, _ := utils.Int2Byte(block.Version)
//	timeByte, _ := utils.Int2Byte(block.TimeStamp)
//	nonceByte, _ := utils.Int2Byte(block.Nonce)
//
//	blockByte := bytes.Join([] []byte{heightByte, versionByte,block.PrevHash[:],timeByte,nonceByte,block.Data}, []byte{})
//	//为区块的哈希字段赋值
//	block.Hash = sha256.Sum256(blockByte)
//}

/**
 *生成创世区块的函数
 */
func CreateGenesis(data []byte) Block {
	genesis := Block{
		Height:  0,
		Version: VERSION,
		PrevHash: [32]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},
		TimeStamp: time.Now().Unix(),
		Data:   data,
		}

		//调用Pow，实现hash计算和寻找nonce
		proof := consensus.	NewPow(genesis)
		hash, nonce := proof.FindNonce()
		genesis.Hash = hash
		genesis.Nonce = nonce

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

	proof := consensus.NewPow(newBlock)
	hash,nonce := proof.FindNonce()
	newBlock.Hash = hash
	newBlock.Nonce = nonce

	return newBlock
}