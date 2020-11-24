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
	privateKey = "3B9670B5CB19C893694FC49B461CE489BF9588BE16DBE8DC29CF06338133DEE6"
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

	// Create namespace from it's name
	namespaceId, err := sdk.NewNamespaceIdFromName("mynamespace")
	if err != nil {
		fmt.Printf("NewNamespaceIdFromName returned error: %s", err)
		return

		// Create a new transfer type transaction
		transaction, err := client.NewTransferTransactionWithNamespace(
			// The maximum amount of time to include the transaction in the blockchain.
			sdk.NewDeadline(time.Hour * 1),
			// The address of the recipient account.
			namespaceId,
			// The array of mosaic to be sent.
			[]*sdk.Mosaic{sdk.Xpx(10000000)},
			// The transaction message of 1024 characters.
			sdk.NewPlainMessage("Here you go"),
		)
		if err != nil {
			fmt.Printf("NewTransferTransactionWithNamespace returned error: %s", err)
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
	}
