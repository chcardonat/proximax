package main


import (
	"context"
	"fmt"
	"time"
	"github.com/proximax-storage/go-xpx-chain-sdk/sdk"
)

const (
	// Catapult-api-rest server.
	baseUrl = "http://bctestnet1.brimstone.xpxsirius.io:3000"
	// Types of network.
	networkType = sdk.PublicTest
	// Private key of some exist account
	privateKey = "B0835BB2375A30D81F232C4ED369A9F1337A04CFD5BA4A5758D29E67037D2AB9"
	// A valid public key
	publicKey = "2F8B5A224A1B32775D49B3A26B58706158660852C26B07244A0B61DF45FAC8DD"
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

	// Create an account from a public key
	remoteAccount, err := client.NewAccountFromPublicKey(publicKey)
	if err != nil {
		fmt.Printf("NewAccountFromPublicKey returned error: %s", err)
		return
	}

	// Create a address alias transaction
	transaction, err := client.NewAccountLinkTransaction(
		// The maximum amount of time to include the transaction in the blockchain.
		sdk.NewDeadline(time.Hour),
		// The remote account to which signer wants to delegate importance.
		remoteAccount,
		// The type of action ( in this case we want to link entities ).
		sdk.AccountLink,
	)
	if err != nil {
		fmt.Printf("NewAccountLinkTransaction returned error: %s", err)
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
