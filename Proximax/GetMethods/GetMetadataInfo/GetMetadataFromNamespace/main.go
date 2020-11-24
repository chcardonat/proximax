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
	// Valid namespace name
	namespaceName = "pruebacarlos"
	// Valid key which is used to store metadata in address
	key = "my_super_key"
)

func main() {

	conf, err := sdk.NewConfig(context.Background(), []string{baseUrl})
	if err != nil {
		fmt.Printf("NewConfig returned error: %s", err)
		return
	}

	if err != nil {
		panic(err)
	}

	// Use the default http client
	client := sdk.NewClient(nil, conf)

	// Generate namespace id from name
	namespaceId, err := sdk.NewNamespaceIdFromName(namespaceName)
	if err != nil {
		panic(err)
	}

	info, err := client.Metadata.GetMetadataByNamespaceId(context.Background(), namespaceId)
	if err != nil {
		fmt.Printf("Metadata.GetMetadataByNamespaceId returned error: %s", err)
		return
	}
	fmt.Printf("Namespace id %s\n", info.NamespaceId)
	fmt.Printf("Namespace Metadata Type %t\n", sdk.MetadataNamespaceType == info.MetadataType)
	fmt.Printf("Metadata for %s: %s\n", key, info.Fields[key])

	// you can get info for multiple mosaics at the same time also
	// infos, err := client.Metadata.GetNamespaceMetadatasInfo(context.Background(), namespaceId1, namespaceId2)
}