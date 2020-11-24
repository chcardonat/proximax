package main

import (
	"context"
	"fmt"
	"github.com/proximax-storage/go-xpx-chain-sdk/sdk"
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

	// Get storage of blockchain
	storage, err := client.Blockchain.GetBlockchainStorage(context.Background())
	if err != nil {
		fmt.Printf("Blockchain.GetBlockchainStorage returned error: %s", err)
		return
	}
	fmt.Printf("%s\n", storage)
}
