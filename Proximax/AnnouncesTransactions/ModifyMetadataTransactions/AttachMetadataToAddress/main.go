package main

import (
	"context"
	"fmt"
	"github.com/proximax-storage/go-xpx-chain-sdk/sdk"
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

	// Create an accounts from a private keys
	account, err := client.NewAccountFromPrivateKey(privateKey)
	if err != nil {
		fmt.Printf("NewAccountFromPrivateKey returned error: %s", err)
		return
	}

	// Create a modify metadata transaction
	transaction, err := client.NewModifyMetadataAddressTransaction(
		// The maximum amount of time to include the transaction in the blockchain.
		sdk.NewDeadline(time.Hour * 1),
		// Address where metadata should be attached
		account.PublicAccount.Address,
		// Actual data which will be added/removed
		[]*sdk.MetadataModification{
			{
				// add or remove metadata
				sdk.AddMetadata,
				// key which should be used to store data
				"relacionlabor.link1",
				// actual data which will be stored in associated key
				"https://drive.google.com/file/d/1M6VHZziZZAkdJtjbZeBbl1QOZwZttquJ/view?usp=sharing",
			},
		},
	)

	if err != nil {
		fmt.Printf("NewModifyMetadataAddressTransaction returned error: %s", err)
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
	fmt.Printf("Content: \t\t%v", signedTransaction.Hash)

	// wait for the transaction to be confirmed! (very important)
	// you can use websockets to wait explicitly for transaction
	// to be in certain state, instead of hard waiting
	time.Sleep(time.Second * 30)

}