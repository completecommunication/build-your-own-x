package blockchain

import (
	"time"

	"github.com/davecgh/go-spew/spew"
)

const difficulty = 1

// Blockchain is a series of validated Blocks
var Blockchain []Block

func replaceChain(newBlocks []Block) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}

// CreateGenesis chain
func CreateGenesis() {
	t := time.Now()
	genesisBlock := Block{}
	genesisBlock = Block{
		Index:      0,
		Timestamp:  t.String(),
		BPM:        0,
		Hash:       genesisBlock.calculateHash(),
		PrevHash:   "",
		Difficulty: difficulty,
		Nonce:      "",
	}
	spew.Dump(genesisBlock)

	Blockchain = append(Blockchain, genesisBlock)
}
