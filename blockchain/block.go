package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Block represents each 'item' in the blockchain
type Block struct {
	Index      int
	Timestamp  string
	BPM        int
	Hash       string
	PrevHash   string
	Difficulty int
	Nonce      string
}

func (block Block) calculateHash() string {
	record := strconv.Itoa(block.Index) + block.Timestamp + strconv.Itoa(block.BPM) + block.PrevHash + block.Nonce
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
	newBlock.Difficulty = difficulty

	for i := 0; ; i++ {
		hex := fmt.Sprintf("%x", i)
		newBlock.Nonce = hex

		hash := newBlock.calculateHash()
		if !isHashValid(hash, newBlock.Difficulty) {
			fmt.Println(hash, " do more work!")
			time.Sleep(time.Second)
			continue
		} else {
			fmt.Println(hash, " work done!")
			newBlock.Hash = hash
			break
		}
	}

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

func isHashValid(hash string, difficulty int) bool {
	prefix := strings.Repeat("0", difficulty)

	return strings.HasPrefix(hash, prefix)
}
