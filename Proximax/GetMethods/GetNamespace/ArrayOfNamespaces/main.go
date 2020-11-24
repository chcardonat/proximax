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

	conf, err := sdk.NewConfig(context.Background(), []string{baseUrl})
	if err != nil {
		fmt.Printf("NewConfig returned error: %s", err)
		return
	}

	// Use the default http client
	client := sdk.NewClient(nil, conf)

	// Generate Ids from namespace names
	proximaxId, err := sdk.NewNamespaceIdFromName("pruebacarlos")
	if err != nil {
		fmt.Printf("NewNamespaceIdFromName returned error: %s", err)
		return
	}
	mynamespaceId, err := sdk.NewNamespaceIdFromName("pruebacarlos.prueba2")
	if err != nil {
		fmt.Printf("NewNamespaceIdFromName returned error: %s", err)
		return
	}

	namespaceNames, err := client.Namespace.GetNamespaceNames(context.Background(), []*sdk.NamespaceId{proximaxId, mynamespaceId})
	if err != nil {
		fmt.Printf("Namespace.GetNamespaceNames returned error: %s", err)
		return
	}
	for _, name := range namespaceNames {
		fmt.Printf("%s\n", name.String())
	}
}