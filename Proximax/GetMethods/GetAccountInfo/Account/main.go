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
	// A valid public key.
	publicKey  = "FFF228306B2E14F4F48F031DE8CFF5E6F1D4B87C348E6B9D811D9AA3ACFE94A5"

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



	address, err := sdk.NewAddressFromPublicKey(publicKey, networkType)
	if err != nil {
		fmt.Printf("NewAddressFromPublicKey returned error: %s", err)
		return
	}

	// Get AccountInfo for an account.
	account, err := client.Account.GetAccountInfo(context.Background(), address)
	if err != nil {
		fmt.Printf("Account.GetAccountInfo returned error: %s", err)
		return
	}
	fmt.Printf(account.String())
}