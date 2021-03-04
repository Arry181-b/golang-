package consensus

import (
	"AnderChain/chain"
	"AnderChain/utils"
	"bytes"
	"crypto/sha256"
	"math/big"
)
const DIFFICULTY = 10 //难度值系数


type Pow struct {
	Block chain.Block
	Target  *big.Int
}



func (pow Pow) FindNonce()int64{
	//1.给定一个nonce值
	var nonce int64
	nonce = 0
	//无限循环
	for {
		hash := CalculateHash(pow.Block, nonce)

		//32 -> 256
		//2.拿到系统目标值
		target := pow.Target

		//3.比较大小
		result := bytes.Compare(hash[:], target.Bytes())
		if result == -1 {
			return nonce
		}
		//4.否则自增
		nonce++
	}
	return 0
}

/**
 *根据区块已有的信息和nonce已有的赋值，计算区块的哈希
 */
func CalculateHash(block chain.Block, nonce int64) [32]byte {
	heightByte, _ := utils.Int2Byte(block.Height)
	versionByte, _ := utils.Int2Byte(block.Version)
	timeByte, _ := utils.Int2Byte(block.TimeStamp)
	nonceByte, _ := utils.Int2Byte(nonce)
	blockByte := bytes.Join([][]byte{heightByte, versionByte, block.PrevHash[:], timeByte, nonceByte, block.Data}, []byte{})
	//计算区块的哈希
	hash := sha256.Sum256(blockByte)
	return hash
}
