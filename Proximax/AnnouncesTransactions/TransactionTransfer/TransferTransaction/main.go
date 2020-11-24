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
	receiverKey = "4f61116b8dccc8b340fb25013b768f5980c785bf193754ca100cef4d6c9aa32c"
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
	account, err := sdk.NewAccountFromPrivateKey(privateKey, networkType, client.GenerationHash())
	if err != nil {
		fmt.Printf("NewAccountFromPrivateKey returned error: %s", err)
		return
	}

	// Create an account from a private key
	receiver, err := sdk.NewAccountFromPrivateKey(receiverKey, networkType, client.GenerationHash())
	if err != nil {
		fmt.Printf("NewAccountFromPrivateKey returned error: %s", err)
		return
	}


	// Create a new transfer type transaction
	transaction, err := client.NewTransferTransaction(
		// The maximum amount of time to include the transaction in the blockchain.
		sdk.NewDeadline(time.Hour * 1),
		// The address of the recipient account.
		receiver.Address,
		// The array of mosaic to be sent.
		[]*sdk.Mosaic{sdk.Xpx(81000000)},
		// The transaction message of 1024 characters.
		sdk.NewPlainMessage("Enviado 81xpx a de test1 a test 2"),
	)
	if err != nil {
		fmt.Printf("NewTransferTransaction returned error: %s", err)
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
	time.Sleep(30 * time.Second)

}
