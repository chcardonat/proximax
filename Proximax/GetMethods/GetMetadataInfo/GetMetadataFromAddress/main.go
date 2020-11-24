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
	// Valid account address for Mijin Test network
	address = "VCZGZ6VIEVVKAZWQ5CMCEYSL2OT2VD4XCVFZ3AU7"
	// Valid key which is used to store metadata in address
	key = "relacionlabor.link1"
)

func main() {

	conf, err := sdk.NewConfig(context.Background(), []string{baseUrl})
	if err != nil {
		fmt.Printf("NewConfig returned error: %s", err)
		return
	}

	// Use the default http client
	client := sdk.NewClient(nil, conf)

	info, err := client.Metadata.GetMetadataByAddress(context.Background(), address)
	if err != nil {
		fmt.Printf("Metadata.GetMetadataByAddress returned error: %s", err)
		return
	}
	fmt.Printf("Address %s\n", info.Address)
	fmt.Printf("Address %s\n", info)
	fmt.Printf("Metadata for %s: %s\n", key, info.Fields[key])


	// you can get info for multiple addresses at the same time also
	// infos, err := client.Metadata.GetAddressMetadatasInfo(context.Background(), address1, address2)
}