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
	nonce = 1234
	// Valid public key of mosaic owner
	publicKey = "F5F090D5EC24E572CC9D4EC3B891D8C4435E986620CB38B003EA19EAD13AC724"
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

	// Generate mosaic id
	mosaicId, err := sdk.NewMosaicIdFromNonceAndOwner(nonce, publicKey)
	if err != nil {
		panic(err)
	}

	info, err := client.Metadata.GetMetadataByMosaicId(context.Background(), mosaicId)
	if err != nil {
		fmt.Printf("Metadata.GetMetadataByMosaicId returned error: %s", err)
		return
	}
	fmt.Printf("Mosaic id %s\n", info.MosaicId)
	fmt.Printf("Mosaic Metadata Type %t\n", sdk.MetadataMosaicType == info.MetadataType)
	fmt.Printf("Metadata for %s: %s\n", key, info.Fields[key])

	// you can get info for multiple mosaics at the same time also
	// infos, err := client.Metadata.GetMosaicMetadatasInfo(context.Background(), mosaicId1, mosaicId2)
}
