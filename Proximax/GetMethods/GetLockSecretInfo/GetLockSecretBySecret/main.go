package main

import (
	"context"
	"crypto/rand"
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

	//create some proof
	proofB := make([]byte, 8)
	_, err = rand.Read(proofB);
	if err != nil {
		fmt.Printf("rand.Read returned error: %s", err)
		return
	}

	//New proof
	proof := sdk.NewProofFromBytes(proofB)
	//New secret from proof
	secret, _ := proof.Secret(sdk.SHA3_256)
	//Get info by secret
	infosBySecret, err := client.Lock.GetSecretLockInfosBySecret(context.Background(), &secret.Hash)
	for _, v := range infosBySecret {
		fmt.Println(v)
	}
}