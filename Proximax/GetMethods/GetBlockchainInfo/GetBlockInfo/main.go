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
	height := sdk.Height(2035095)

	// Get BlockInfo by height
	block, err := client.Blockchain.GetBlockByHeight(context.Background(), height)
	if err != nil {
		fmt.Printf("Blockchain.GetBlockByHeight returned error: %s", err)
		return
	}
	fmt.Printf(block.String())
}
