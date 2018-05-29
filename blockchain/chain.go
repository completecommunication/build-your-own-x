package blockchain

// Blockchain is a series of validated Blocks
var Blockchain []Block

func replaceChain(newBlocks []Block) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}
