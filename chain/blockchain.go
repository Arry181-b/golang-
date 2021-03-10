package chain

import (
	"errors"
	"github.com/boltdb/bolt"
)

const BLOCKS = "blocks"
const LASTHASH = "lasthash"
/**
 *定义区块链结构体，该结构体用于管理区块
 */
type BlockChain struct {
	//key -> value
	//切片
	//Blocks []Block
	DB        *bolt.DB
	LastBlock Block //最新最后的区块
	IteratorBlockHash [32]byte //表示当前迭代到了哪个区块，该变量用于记录迭代到的区块hash
}

func CreateChain(db *bolt.DB)BlockChain{
	return BlockChain{
		DB: db,
	}
}

/**
 *创建一个区块链对象，包含一个创世区块
 */
func (chain *BlockChain) CreateGenesis(data []byte) error {
	var err error
	//genesis持久化到db中去
	engine := chain.DB
	engine.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BLOCKS))
		if bucket == nil { //没有桶
			bucket, err = tx.CreateBucket([]byte(BLOCKS))
			if err != nil {
				return err
				//panic("操作区块存储文件失败，请重试")
			}
		}

		//先查看
		lastHash := bucket.Get([]byte(LASTHASH))
		if len(lastHash) == 0 {
			genesis := CreateGenesis(data)
			genSerBytes, _ := genesis.Serialize()
		//bucket已存在
		// key -> value
		//blockhash -> 区块序列化以后的数据
		bucket.Put(genesis.Hash[:], genSerBytes) //把创世区块保存到boltdb中取
		//使用一个标志，用来记录最新区块的hash，以标明当前文件存储到了最新的哪个区块
		bucket.Put([]byte(LASTHASH),genesis.Hash[:])
			//把genesis赋值给chain的lastblock
			chain.LastBlock = genesis
			chain.IteratorBlockHash = genesis.Hash
		}else {
			//从文件中，读取出最新的区块，并赋值给内存中的LastBlock
			lastHash := bucket.Get([]byte(LASTHASH))
			lastBlockBytes := bucket.Get(lastHash)
			//把反序列化的最后最新的区块赋值给chain.LastBlock
			chain.LastBlock, err = Deserialize(lastBlockBytes)
			chain.IteratorBlockHash = chain.LastBlock.Hash
		}
		return nil
	})
	return err
}

/**
  * 生成一个新区块
 */
func (chain *BlockChain) CreateNewBlock(data []byte) error {
	//目的：生成一个新区块，并保存到bolt.Db中去(持久化）
	//手段(步骤)：
	//1、从文件中查到当前存储的最新区块
	lastBlock := chain.LastBlock
	//3、根据获取的最新区块生成一个新区块
	newBlock := NewBlock(lastBlock.Height, lastBlock.Hash, data)
	//4.将最新区块序列化，得到序列化数据
	var err error
	newBlockSerBytes, err := newBlock.Serialize()
	if err != nil {
		return err
	}
	//5、将序列化数据存储到文件、同时更新到最新区块的标记lashhash，更新为最新区块的hash
		db := chain.DB
		db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BLOCKS))
		if bucket == nil {
			err = errors.New("数据库操作失败，请重试2！")
		}
		//将新生成的区块，保存到文件中去
		bucket.Put(newBlock.Hash[:], newBlockSerBytes)
		//更新到最新区块的标记lashhash，更新为最新区块的hash
		bucket.Put([]byte(LASTHASH), newBlock.Hash[:])
		//更新内存中的nlockchain的LastBlock
		chain.LastBlock = newBlock
		chain.IteratorBlockHash = newBlock.Hash
		return nil
	})
	return err
}

//获取最新区块数据
func (chain *BlockChain) GetLastBlock() (Block) {
	return chain.LastBlock
}


//获取所有区块数据
func (chain *BlockChain) GetAllBlocks() ([]Block, error) {
	//目的：获取所有区块
	//手段（步骤）：
		//1、找到最后一个区块
		db := chain.DB
		var err error
		blocks := make([]Block,0)
		db.View(func(tx *bolt.Tx) error {
			bucket := tx.Bucket([]byte(BLOCKS))
			if bucket == nil {
				err = errors.New("数据库操作失败，请重试4！")
				return err
			}
			var currentHash []byte
			currentHash = bucket.Get([]byte(LASTHASH))
			//2.根据最后一个区块依次往前找
			for {
				currentBlockBytes := bucket.Get(currentHash)
				currentBlock, err := Deserialize(currentBlockBytes)
				if err != nil {
					break
				}
				//3、每次找到的区块防区到一个[]Block容器中
				blocks = append(blocks, currentBlock)
				//4、找到最开始的创世区块时，就结束了，不用找了
				if currentBlock.Height == 0 {
					break
				}
				currentHash = currentBlock.PrevHash[:]
			}

			return nil
		})
		return blocks, err

}

/**
 *该方法用于实现迭代器Iterator的HasNext方法，用于判读是否还有数据
 *如果有数据，返回true,否则返回false
 */
func (chain *BlockChain) HasNext() bool {
	//区块0 -> 区块1 -> 区块2
	//最新区块：区块2
	//步骤：当前区块在哪 -> prevHash -> db
	db := chain.DB
	var hasNext bool
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BLOCKS))
		if bucket == nil {
			return errors.New("数据库操作失败，请重试5")
		}
		//获取前一个区块的区块数据
		preBlockBytes := bucket.Get(chain.IteratorBlockHash[:])
		//如果获取不到前一个区块的数据，说明后面没有区块了
		hasNext = len(preBlockBytes) != 0
		return nil
	})
	return hasNext
}

/**
  * 该方法用于实现迭代器Iterator的HasNext方法，用于取出一个数据
  * 此处，因为是区块链数据集合，因此返回的数据类型是Block
 */
func (chain *BlockChain) Next() Block {
	//1.知道当前在哪个区块
	//2.找当前区块的前一个区块
	//3.找到的区块返回
	db := chain.DB
	var iteratorBlock Block
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BLOCKS))
		if bucket == nil {
			return errors.New("数据库操作失败，请重试6")
		}
		//前一个区块的数据

		blockBytes := bucket.Get(chain.IteratorBlockHash[:])
		iteratorBlock, _ = Deserialize(blockBytes)
		//迭代到当前区块后，更新游标的区块内容
		chain.IteratorBlockHash = iteratorBlock.PrevHash
		return nil
	})
	return iteratorBlock
}








