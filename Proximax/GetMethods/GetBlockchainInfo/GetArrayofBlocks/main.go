package main

import (
	"fmt"
	"github.com/proximax-storage/go-xpx-chain-sdk/sdk"
	"context"
)

const (
	// Catapult-api-rest server.
	baseUrl = "http://bctestnet1.brimstone.xpxsirius.io:3000"
	// Types of network.
	networkType = sdk.PublicTest
)

// Simple Account API request
func main() {

	conf, err := sdk.NewConfig(context.Background(), []string{baseUrl})
	if err != nil {
		fmt.Printf("NewConfig returned error: %s", err)
		return
	}

	// Use the default http client
	client := sdk.NewClient(nil, conf)

	// height of block in blockchain
	height := sdk.Height(1)
	// how many BlockInfo's to return
	limit :=  sdk.Amount(100)

	// Get BlockInfo's by height with limit
	blocks, err := client.Blockchain.GetBlocksByHeightWithLimit(context.Background(), height, limit)
	if err != nil {
		fmt.Printf("Blockchain.GetBlocksByHeightWithLimit returned error: %s", err)
		return
	}
	for _, block := range blocks {
		fmt.Printf("%s\n", block.String())
	}
}
