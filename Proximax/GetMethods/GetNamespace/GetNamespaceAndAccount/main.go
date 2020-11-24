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

	// Generate Address struct
	address, err := sdk.NewAddressFromPublicKey(publicKey, networkType)
	if err != nil {
		fmt.Printf("NewAddressFromPublicKey returned error: %s", err)
		return
	}

	namespaces, err := client.Namespace.GetNamespaceInfosFromAccount(context.Background(), address, nil, 0)
	if err != nil {
		fmt.Printf("Namespace.GetNamespaceInfosFromAccount returned error: %s", err)
		return
	}
	for _, namespace := range namespaces {
		fmt.Printf("%s\n", namespace.String())
	}
}