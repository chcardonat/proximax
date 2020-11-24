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

	mosaicId, err := sdk.NewMosaicId(1059995153615339986)
	if err != nil {
		fmt.Printf("NewMosaicId returned error: %s", err)
		return
	}

	// Get mosaic information.
	mosaic, err := client.Mosaic.GetMosaicInfo(context.Background(), mosaicId)
	if err != nil {
		fmt.Printf("Mosaic.GetMosaicInfo returned error: %s", err)
		return
	}
	fmt.Printf("%s\n", mosaic.String())

}
