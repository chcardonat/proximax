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
	privateKey = "B0835BB2375A30D81F232C4ED369A9F1337A04CFD5BA4A5758D29E67037D2AB9"
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

	// nonce of mosaic for which account is owner
	nonce := uint32(0)

	// create mosaic identifier
	mosaicId, err := sdk.NewMosaicIdFromNonceAndOwner(nonce, account.PublicAccount.PublicKey)
	if err != nil {
		fmt.Printf("NewMosaicIdFromNonceAndOwner returned error: %s", err)
		return
	}
	fmt.Printf("Content: \t\t%v", mosaicId)

	// Create a modify metadata transaction
	transaction, err := sdk.NewModifyMetadataMosaicTransaction(
		// The maximum amount of time to include the transaction in the blockchain.
		sdk.NewDeadline(time.Hour),
		// Id of mosaic where metadata should be added
		mosaicId,
		// Actual data which will be added/removed
		[]*sdk.MetadataModification{
			{
				// add or remove metadata
				sdk.AddMetadata,
				// key which should be used to store data
				"my_key_name",
				// actual data which will be stored in associated key
				"my_data",
			},
		},
		networkType,
	)
	if err != nil {
		fmt.Printf("NewModifyMetadataMosaicTransaction returned error: %s", err)
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
