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

// Simple Account API request
func main() {

	conf, err := sdk.NewConfig(context.Background(), []string{baseUrl})
	if err != nil {
		fmt.Printf("NewConfig returned error: %s", err)
		return
	}

	// Use the default http client
	client := sdk.NewClient(nil, conf)

	namespaceId, err := sdk.NewNamespaceIdFromName("dolar")
	if err != nil {
		fmt.Printf("NewNamespaceIdFromName returned error: %s", err)
		return
	}

	mosaic, err := client.Resolve.GetMosaicInfoByAssetId(context.Background(), namespaceId)
	if err != nil {
		fmt.Printf("Resolve.GetMosaicInfoByAssetId returned error: %s", err)
		return
	}

	fmt.Printf("%s\n", mosaic.String())
}
