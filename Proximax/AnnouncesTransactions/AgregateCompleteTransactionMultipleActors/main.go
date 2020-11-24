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
	// Valid private keys
	firstActorPrivateKey = "B0835BB2375A30D81F232C4ED369A9F1337A04CFD5BA4A5758D29E67037D2AB9"
	secondActorPrivateKey = "04DD376196603C44A19FD500492E5DE12DE9ED353DE070A788CB21F210645613"
	// Addresses for transfer
	addressToTransfer  = "VAKCGTVIHHZL3WWPAW3XIBRIQEGR4Y3UXZ7IURYH"
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
	firstActor, err := client.NewAccountFromPrivateKey(firstActorPrivateKey)
	if err != nil {
		fmt.Printf("NewAccountFromPrivateKey returned error: %s", err)
		return
	}
	secondActor, err := client.NewAccountFromPrivateKey(secondActorPrivateKey)
	if err != nil {
		fmt.Printf("NewAccountFromPrivateKey returned error: %s", err)
		return
	}

	// Create a new transfer type transaction
	firstTransferTransaction, err := client.NewTransferTransaction(
		// The maximum amount of time to include the transaction in the blockchain.
		sdk.NewDeadline(time.Hour),
		// The address of the recipient account.
		sdk.NewAddress(addressToTransfer, client.NetworkType()),
		// The array of mosaic to be sent.
		[]*sdk.Mosaic{sdk.Xpx(10)},
		// The transaction message of 1024 characters.
		sdk.NewPlainMessage(""),
	)
	if err != nil {
		fmt.Printf("NewTransferTransaction returned error: %s", err)
		return
	}

	// Create a new transfer type transaction
	secondTransferTransaction, err := client.NewTransferTransaction(
		// The maximum amount of time to include the transaction in the blockchain.
		sdk.NewDeadline(time.Hour),
		// The address of the recipient account.
		sdk.NewAddress(addressToTransfer, client.NetworkType()),
		// The array of mosaic to be sent.
		[]*sdk.Mosaic{sdk.Xpx(10)},
		// The transaction message of 1024 characters.
		sdk.NewPlainMessage(""),
	)
	if err != nil {
		fmt.Printf("NewTransferTransaction returned error: %s", err)
		return
	}

	// Convert an aggregate transaction to an inner transaction including transaction signer.
	firstTransferTransaction.ToAggregate(firstActor.PublicAccount)
	secondTransferTransaction.ToAggregate(secondActor.PublicAccount)

	// Create an aggregate complete transaction
	aggregateTransaction, err := client.NewCompleteAggregateTransaction(
		// The maximum amount of time to include the transaction in the blockchain.
		sdk.NewDeadline(time.Hour),
		// Inner transactions
		[]sdk.Transaction{firstTransferTransaction, secondTransferTransaction},
	)
	if err != nil {
		fmt.Printf("NewCompleteAggregateTransaction returned error: %s", err)
		return
	}

	// Sign transaction
	signedTransaction, err := firstActor.SignWithCosignatures(aggregateTransaction, []*sdk.Account{secondActor})
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

	// wait for the transaction to be confirmed! (very important)
	// you can use websockets to wait explicitly for transaction
	// to be in certain state, instead of hard waiting
	time.Sleep(time.Second * 30)

}