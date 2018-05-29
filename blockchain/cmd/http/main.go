package main

import (
	"log"

	"github.com/alexfalkowski/build-your-own-x/blockchain"
)

func main() {
	go blockchain.CreateGenesis()

	log.Fatal(blockchain.HTTPListenAndServe())
}
