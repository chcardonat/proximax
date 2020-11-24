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

	config, err := client.Network.GetNetworkConfigAtHeight(context.Background(), 0)
	if err != nil {
		fmt.Printf("Network.GetNetworkConfig returned error: %s", err)
		return
	}

	for sectionName, fields := range config.BlockChainConfig.Sections {
		for fieldName, value := range fields.Fields {
			fmt.Printf("Section name - %s, field name - %s, field value - %s", sectionName, fieldName, value.String())
		}
	}
}

