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
	multisigPublicKey = "C8BB4994000422DE988A6FA78E5AE140E397B61C6CC0FCAA7D24D493D6CB743E"
)

func main() {

	conf, err := sdk.NewConfig(context.Background(), []string{baseUrl})
	if err != nil {
		fmt.Printf("NewConfig returned error: %s", err)
		return
	}

	// Use the default http client
	client := sdk.NewClient(nil, conf)

	multisig, err := sdk.NewAccountFromPublicKey(multisigPublicKey, networkType)
	if err != nil {
		fmt.Printf("NewAccountFromPublicKey returned error: %s", err)
		return
	}

	// Get confirmed transactions information.
	multisigAccountInfo, err := client.Account.GetMultisigAccountInfo(context.Background(), multisig.Address)
	if err != nil {
		fmt.Printf("Account.GetMultisigAccountInfo returned error: %s", err)
		return
	}
	fmt.Printf("%s\n", multisigAccountInfo.String() )

	//Get info for a multisigAccount and all hierarchy
	infoMultisign2, err := client.Account.GetMultisigAccountGraphInfo(context.Background(), multisig.Address)
	if err != nil {
		fmt.Printf("Transaction.GetTransaction returned error: %s", err)
		return
	}
	fmt.Println(infoMultisign2)
}
