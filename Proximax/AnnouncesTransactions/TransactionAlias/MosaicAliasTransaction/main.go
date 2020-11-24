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

	// Create an account from a private key
	account, err := client.NewAccountFromPrivateKey(privateKey)
	if err != nil {
		fmt.Printf("NewAccountFromPrivateKey returned error: %s", err)
		return
	}

	// Create namespace id from it's name
	namespaceId, err := sdk.NewNamespaceIdFromName("mosaicalias")
	if err != nil {
		fmt.Printf("NewNamespaceIdFromName returned error: %s", err)
		return
	}

	// Create mosaic id from it's id
	mosaicId, err := sdk.NewMosaicId(1470042125674711000)
	if err != nil {
		fmt.Printf("NewMosaicId returned error: %s", err)
		return
	}

	// Create a address alias transaction
	transaction, err := client.NewMosaicAliasTransaction(
		// The maximum amount of time to include the transaction in the blockchain.
		sdk.NewDeadline(time.Hour),
		// The mosaic id which we want to link.
		mosaicId,
		// The namespace id which we want to link.
		namespaceId,
		// The type of action ( in this case we want to link entities ).
		sdk.AliasLink,
	)
	if err != nil {
		fmt.Printf("NewMosaicAliasTransaction returned error: %s", err)
		return
	}

	// Sign transaction
	signedTransaction, err := account.Sign(transaction)
	if err != nil {
		fmt.Printf("Sign returned error: %s", err)
		return
	}
	fmt.Printf("Content: \t\t%v", signedTransaction.Hash)

	// Announce transaction
	_, err = client.Transaction.Announce(context.Background(), signedTransaction)
	if err != nil {
		fmt.Printf("Transaction.Announce returned error: %s", err)
		return
	}
}
