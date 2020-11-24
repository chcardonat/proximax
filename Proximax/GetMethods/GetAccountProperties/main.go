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
	privateKey = "B0835BB2375A30D81F232C4ED369A9F1337A04CFD5BA4A5758D29E67037D2AB9"
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
