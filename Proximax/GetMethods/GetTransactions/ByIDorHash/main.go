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

// Simple Transaction API request
func main() {

	conf, err := sdk.NewConfig(context.Background(), []string{baseUrl})
	if err != nil {
		fmt.Printf("NewConfig returned error: %s", err)
		return
	}

	// Use the default http client
	client := sdk.NewClient(nil, conf)

	transaction, err := client.Transaction.GetTransaction(context.Background(), "f5a1cc34639fa87a86dc19ade53a3621a075a4d0de51b088ca76adb7e6a34991")
	if err != nil {
		fmt.Printf("Transaction.GetTransaction returned error: %s", err)
		return
	}
	fmt.Printf("%s\n\n", transaction.String())
}
