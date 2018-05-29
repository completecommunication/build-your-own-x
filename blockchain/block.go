package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// Block in the blockchain
type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string
}

func (block Block) calculateHash() string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
	h := sha256.New()
	_, _ = h.Write([]byte(record))
	hashed := h.Sum(nil)

	return hex.EncodeToString(hashed)
}

func (block Block) generateBlock(BPM int) (Block, error) {
	var newBlock Block

	t := time.Now()

	newBlock.Index = block.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = block.Hash
	newBlock.Hash = newBlock.calculateHash()

	return newBlock, nil
}

func (block Block) isBlockValid(newBlock Block) bool {
	if block.Index+1 != newBlock.Index {
		return false
	}

	if block.Hash != newBlock.PrevHash {
		return false
	}

	if newBlock.calculateHash() != newBlock.Hash {
		return false
	}

	return true
}
