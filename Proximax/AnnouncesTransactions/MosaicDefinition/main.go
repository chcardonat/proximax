package main

import (
	"context"
	"fmt"
	"github.com/proximax-storage/go-xpx-chain-sdk/sdk"
	"math/rand"
	"time"
)

const (
	// Catapult-api-rest server.
	baseUrl = "http://bctestnet1.brimstone.xpxsirius.io:3000"
	// Types of network.
	networkType = sdk.PublicTest
	// Private key of some exist account
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

	// Create an account from a private key
	account, err := client.NewAccountFromPrivateKey(privateKey)
	if err != nil {
		fmt.Printf("NewAccountFromPrivateKey returned error: %s", err)
		return
	}

	random := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	nonce := random.Uint32()

	// Create a new mosaic definition type transaction
	transaction, err := client.NewMosaicDefinitionTransaction(
		// The maximum amount of time to include the transaction in the blockchain.
		sdk.NewDeadline(time.Hour * 1),
		// Mosaic nonce
		nonce,
		// Public key of mosaic owner
		account.PublicAccount.PublicKey,
		sdk.NewMosaicProperties(
			// supply mutable
			true,
			// transferability
			true,
			// divisibility
			4,
			// duration
			sdk.Duration(10000),
		),
	)
	if err != nil {
		fmt.Printf("NewMosaicDefinitionTransaction returned error: %s", err)
		return
	}

	// Sign transaction
	signedTransaction, err := account.Sign(transaction)
	if err != nil {
		fmt.Printf("Sign returned error: %s", err)
		return
	}

	// Announce transaction
	_, err = client.Transaction.Announce(context.Background(), signedTransaction)
	if err != nil {
		fmt.Printf("Transaction.Announce returned error: %s", err)
		return
	}
	time.Sleep(time.Second * 5)
	fmt.Printf("Content: \t\t%v", signedTransaction.Hash)
}