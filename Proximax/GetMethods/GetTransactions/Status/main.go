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

// Simple Transaction API request
func main() {

	conf, err := sdk.NewConfig(context.Background(), []string{baseUrl})
	if err != nil {
		fmt.Printf("NewConfig returned error: %s", err)
		return
	}

	// Use the default http client
	client := sdk.NewClient(nil, conf)

	transactionStatus, err := client.Transaction.GetTransactionStatus(context.Background(), "53E6D261C3D23AE5DF9EB60D600899822B9D77997E48DD5DED1A01AE03350676")
	if err != nil {
		fmt.Printf("Transaction.GetTransactionStatus returned error: %s", err)
		return
	}
	fmt.Printf("%s\n\n", transactionStatus.String())
}
