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

	mosaicId1, err := sdk.NewMosaicId(1059995153615339986)
	if err != nil {
		fmt.Printf("NewMosaicId returned error: %s", err)
		return
	}
	mosaicId2, err := sdk.NewMosaicId(2318234416660457292)
	if err != nil {
		fmt.Printf("NewMosaicId returned error: %s", err)
		return
	}

	// Get mosaic information.
	mosaics, err := client.Mosaic.GetMosaicInfos(context.Background(), []*sdk.MosaicId{mosaicId1, mosaicId2})
	if err != nil {
		fmt.Printf("Mosaic.GetMosaicInfos returned error: %s", err)
		return
	}
	for _, mosaic := range mosaics {
		fmt.Printf("%s\n", mosaic.String())
	}
}
