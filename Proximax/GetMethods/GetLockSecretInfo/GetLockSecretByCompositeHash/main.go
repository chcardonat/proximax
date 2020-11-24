package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/proximax-storage/go-xpx-chain-sdk/sdk"
)

const (
	// Catapult-api-rest server.
	baseUrl = "http://bctestnet1.brimstone.xpxsirius.io:3000"
	// Types of network.
	networkType = sdk.PublicTest
	// Private key of some exist account. Change it
	privateKey = "04DD376196603C44A19FD500492E5DE12DE9ED353DE070A788CB21F210645613"
)

func main() {
	conf, err := sdk.NewConfig(context.Background(), []string{baseUrl})
	if err != nil {
		fmt.Printf("NewConfig returned error: %s", err)
		return
	}

	// Use the default http client
	client := sdk.NewClient(nil, conf)

	//Some hex compositeHash
	bytes, err := hex.DecodeString("11a4cc303e91ae150c6c487ec048a2a07298042427094f2ea6701c25aa565b6c")
	compositeHash := &sdk.Hash{}
	copy(compositeHash[:], bytes)

	//Get secret lock info
	infosByCompositeHash, err := client.Lock.GetSecretLockInfo(context.Background(), compositeHash)
	fmt.Println(infosByCompositeHash)
}
