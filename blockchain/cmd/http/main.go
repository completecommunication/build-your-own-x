package main

import (
	"log"
	"time"

	"github.com/alexfalkowski/build-your-own-x/blockchain"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	go func() {
		t := time.Now()
		genesisBlock := blockchain.Block{
			Index:     0,
			Timestamp: t.String(),
			BPM:       0,
			Hash:      "",
			PrevHash:  "",
		}
		spew.Dump(genesisBlock)
		blockchain.Blockchain = append(blockchain.Blockchain, genesisBlock)
	}()

	log.Fatal(blockchain.ListenAndServe())
}
