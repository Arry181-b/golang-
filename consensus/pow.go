package consensus

import (

	"AnderChain/utils"
	"bytes"
	"crypto/sha256"
	"math/big"
)
const DIFFICULTY = 10 //难度值系数

//目的：拿到数据的熟悉值
	//1.通过结构体引用，引用block结构体，热不过访问其熟悉，比如block.Height
	//2.接口

type Pow struct {
	Block BlockInterface
	Target  *big.Int
}



func (pow Pow) FindNonce() ([32]byte ,int64){
	//1.给定一个nonce值
	var nonce int64
	nonce = 0
	//无限循环
	hashBig := new(big.Int)
	for {
		hash := CalculateHash(pow.Block, nonce)

		//32 -> 256
		//2.拿到系统目标值
		target := pow.Target

		//3.比较大小
		hashBig = hashBig.SetBytes(hash[:])
		//result := bytes.Compare(hash[:], target.Bytes())
		result := hashBig.Cmp(target)
		if result == -1 {
			return hash, nonce
		}
		//4.否则自增
		nonce++
	}
	//return 0
}

/**
 *根据区块已有的信息和nonce已有的赋值，计算区块的哈希
 */
func CalculateHash(block BlockInterface, nonce int64) [32]byte {
	heightByte, _ := utils.Int2Byte(block.GetHeight())
	versionByte, _ := utils.Int2Byte(block.GetVersion())
	timeByte, _ := utils.Int2Byte(block.GetTimeStamp())
	nonceByte, _ := utils.Int2Byte(nonce)

	prev := block.GetPrevHash()
	blockByte := bytes.Join([][]byte{heightByte,
		versionByte,
		prev[:],
		timeByte,
		nonceByte,
		block.GetData()}, []byte{})
	//计算区块的哈希
	hash := sha256.Sum256(blockByte)
	return hash
}
