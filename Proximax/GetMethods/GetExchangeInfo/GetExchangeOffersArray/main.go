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

func main() {
	conf, err := sdk.NewConfig(context.Background(), []string{baseUrl})
	if err != nil {
		fmt.Printf("NewConfig returned error: %s", err)
		return
	}

	// Use the default http client
	client := sdk.NewClient(nil, conf)

	offerType := sdk.SellOffer //or sdk.BuyOffer

	// Get offers info by namespaceId (assetId)
	offersInfo, err := client.Exchange.GetExchangeOfferByAssetId(context.Background(), sdk.StorageNamespaceId, offerType)
	if err != nil {
		fmt.Printf("Exchange.GetExchangeOfferByAssetId returned error: %s", err)
		return
	}

	fmt.Println(offersInfo)
}
