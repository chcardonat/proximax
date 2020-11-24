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
	publicKey = "81298A6EBA208AAAE1BD3278B30BB047D33CEABA4E13FF144753825C499ADB98"
)

func main() {
	conf, err := sdk.NewConfig(context.Background(), []string{baseUrl})
	if err != nil {
		fmt.Printf("NewConfig returned error: %s", err)
		return
	}

	// Use the default http client
	client := sdk.NewClient(nil, conf)

	//Account from private key
	accountSeller, err := client.NewAccountFromPublicKey(publicKey)
	if err != nil {
		fmt.Printf("NewAccountFromPrivateKey returned error: %s", err)
		return
	}

	// Get user offers info
	userExchangeInfo, err := client.Exchange.GetAccountExchangeInfo(context.Background(), accountSeller)
	if err != nil {
		fmt.Printf("Exchange.GetAccountExchangeInfo returned error: %s", err)
		return
	}

	println(userExchangeInfo.String())
}
