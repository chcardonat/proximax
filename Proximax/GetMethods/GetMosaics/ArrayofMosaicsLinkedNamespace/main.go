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

// Simple Account API request
func main() {

	conf, err := sdk.NewConfig(context.Background(),[]string{baseUrl})
	if err != nil {
		fmt.Printf("NewConfig returned error: %s", err)
		return
	}

	// Use the default http client
	client := sdk.NewClient(nil, conf)

	namespaceId1, err := sdk.NewNamespaceIdFromName("pruebacarlos")
	if err != nil {
		fmt.Printf("NewNamespaceIdFromName returned error: %s", err)
		return
	}
	namespaceId2, err := sdk.NewNamespaceIdFromName("pruebacarlos.prueba2")
	if err != nil {
		fmt.Printf("NewNamespaceIdFromName returned error: %s", err)
		return
	}

	mosaics, err := client.Resolve.GetMosaicInfosByAssetIds(context.Background(), namespaceId1, namespaceId2)
	if err != nil {
		fmt.Printf("Resolve.GetMosaicInfosByAssetIds returned error: %s", err)
		return
	}
	for _, mosaic := range mosaics {
		fmt.Printf("%s\n", mosaic.String())
	}
}