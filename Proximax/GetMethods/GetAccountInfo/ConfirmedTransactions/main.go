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
	// A valid public keys
	publicKey = "760D91BD60806DC77234D12107E136D3B4CC2C603343442012F6EA0F957DA379"
)

func main() {

	conf, err := sdk.NewConfig(context.Background(), []string{baseUrl})
	if err != nil {
		fmt.Printf("NewConfig returned error: %s", err)
		return
	}

	// Use the default http client
	client := sdk.NewClient(nil, conf)

	account, err := sdk.NewAccountFromPublicKey(publicKey, networkType)
	if err != nil {
		fmt.Printf("NewAccountFromPublicKey returned error: %s", err)
		return
	}

	// Get confirmed transactions information.
	transactions, err := client.Account.Transactions(context.Background(), account, nil)
	if err != nil {
		fmt.Printf("Account.Transactions returned error: %s", err)
		return
	}

	for _, transaction := range transactions {
		fmt.Printf("%s\n", transaction.String())
	}
}
