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
	// valid addresses
	rawAddressOne = "FFF228306B2E14F4F48F031DE8CFF5E6F1D4B87C348E6B9D811D9AA3ACFE94A5"
	rawAddressTwo = "760D91BD60806DC77234D12107E136D3B4CC2C603343442012F6EA0F957DA379"
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

	// Generate Address struct
	addressOne, err := sdk.NewAddressFromPublicKey(rawAddressOne, networkType)
	if err != nil {
		fmt.Printf("NewAddressFromPublicKey returned error: %s", err)
		return
	}
	addressTwo, err := sdk.NewAddressFromPublicKey(rawAddressTwo, networkType)
	if err != nil {
		fmt.Printf("NewAddressFromPublicKey returned error: %s", err)
		return
	}

	namespaces, err := client.Namespace.GetNamespaceInfosFromAccounts(context.Background(), []*sdk.Address{addressOne, addressTwo}, nil, 0)
	if err != nil {
		fmt.Printf("Namespace.GetNamespaceInfosFromAccounts returned error: %s", err)
		return
	}
	for _, namespace := range namespaces {
		fmt.Printf("%s\n", namespace.String())
	}
}
