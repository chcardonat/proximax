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
	// Valid private key
	privateKey = "B0835BB2375A30D81F232C4ED369A9F1337A04CFD5BA4A5758D29E67037D2AB9"
	// Addresses for transfer
	firstAddressToTransfer  = "VBLEPWBVKDLPWSITUBX3MLPNWXIPMED2T57BRWRX"
	secondAddressToTransfer = "VAKCGTVIHHZL3WWPAW3XIBRIQEGR4Y3UXZ7IURYH"
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

	// Create a new transfer type transaction
	firstTransferTransaction, err := client.NewTransferTransaction(
		// The maximum amount of time to include the transaction in the blockchain.
		sdk.NewDeadline(time.Hour*1),
		// The address of the recipient account.
		sdk.NewAddress(firstAddressToTransfer, client.NetworkType()),
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
		sdk.NewDeadline(time.Hour*1),
		// The address of the recipient account.
		sdk.NewAddress(secondAddressToTransfer, client.NetworkType()),
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
	firstTransferTransaction.ToAggregate(account.PublicAccount)
	secondTransferTransaction.ToAggregate(account.PublicAccount)

	// Create an aggregate complete transaction
	aggregateTransaction, err := client.NewCompleteAggregateTransaction(
		// The maximum amount of time to include the transaction in the blockchain.
		sdk.NewDeadline(time.Hour*1),
		// Inner transactions
		[]sdk.Transaction{firstTransferTransaction, secondTransferTransaction},
	)
	if err != nil {
		fmt.Printf("NewCompleteAggregateTransaction returned error: %s", err)
		return
	}

	// Sign transaction
	signedTransaction, err := account.Sign(aggregateTransaction)
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
