package blockchain

// Blockchain has all the blocks in the Blockchain
var Blockchain []Block

func replaceChain(newBlocks []Block) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}
