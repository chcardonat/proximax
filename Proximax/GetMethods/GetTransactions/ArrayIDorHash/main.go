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

	// Get transaction informations for transactionIds or transactionHashes
	transactions, err := client.Transaction.GetTransactions(context.Background(), []string{"53E6D261C3D23AE5DF9EB60D600899822B9D77997E48DD5DED1A01AE03350676", "B0FCD6A4FE0039345FB6ED6943736EC62D0B696F313FB1F803008870B1ACFD86"})
	if err != nil {
		fmt.Printf("Transaction.GetTransactions returned error: %s", err)
		return
	}
	for _, transaction := range transactions {
		fmt.Printf("%s\n\n", transaction.String())
	}
}
