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
	publicKeyOne = "760D91BD60806DC77234D12107E136D3B4CC2C603343442012F6EA0F957DA379"
	publicKeyTwo = "FFF228306B2E14F4F48F031DE8CFF5E6F1D4B87C348E6B9D811D9AA3ACFE94A5"
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

	addressOne, err := sdk.NewAddressFromPublicKey(publicKeyOne, networkType)
	if err != nil {
		fmt.Printf("NewAddressFromPublicKey returned error: %s", err)
		return
	}

	addressTwo, err := sdk.NewAddressFromPublicKey(publicKeyTwo, networkType)
	if err != nil {
		fmt.Printf("NewAddressFromPublicKey returned error: %s", err)
		return
	}

	// Get AccountsInfo for several accounts.
	accountsInfo, err := client.Account.GetAccountsInfo(context.Background(), addressOne, addressTwo)
	if err != nil {
		fmt.Printf("Account.GetAccountsInfo returned error: %s", err)
		return
	}
	for _, accountInfo := range accountsInfo {
		fmt.Printf("%s\n", accountInfo.String())
	}
}