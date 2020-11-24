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
	// Private key of some exist account
	privateKey = "04DD376196603C44A19FD500492E5DE12DE9ED353DE070A788CB21F210645613"
)

func main() {

	conf, err := sdk.NewConfig(context.Background(), []string{baseUrl})
	if err != nil {
		fmt.Printf("NewConfig returned error: %s", err)
		return
	}

	// Use the default http client
	client := sdk.NewClient(nil, conf)

	// Create an account from a private key
	account, err := sdk.NewAccountFromPrivateKey(privateKey, networkType, client.GenerationHash())
	if err != nil {
		fmt.Printf("NewAccountFromPrivateKey returned error: %s", err)
		return
	}

	info, err := client.Account.GetAccountProperties(context.Background(), account.Address)
	if err != nil {
		fmt.Printf("Account.GetAccountProperties returned error: %s", err)
		return
	}
	fmt.Printf("Address %s\n", info.Address)
	fmt.Printf("Allowed Addresses %s\n", info.AllowedAddresses)
	fmt.Printf("Allowed MosaicId %s\n", info.AllowedMosaicId)
	fmt.Printf("Allowed EntityTypes %s\n", info.AllowedEntityTypes)
	fmt.Printf("Blocked Addresses %s\n", info.BlockedAddresses)
	fmt.Printf("Blocked MosaicId %s\n", info.BlockedMosaicId)
	fmt.Printf("Blocked EntityTypes %s\n", info.BlockedEntityTypes)
}
