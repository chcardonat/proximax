package main


import (
	"context"
	"fmt"
	"github.com/proximax-storage/go-xpx-chain-sdk/sdk"
	"time"
	"math/rand"
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

	mosaicId, _ := sdk.NewMosaicIdFromNonceAndOwner(nonce, account.PublicAccount.PublicKey)

	// Create a new mosaic definition type transaction
	transaction, err := client.NewMosaicSupplyChangeTransaction(
		// The maximum amount of time to include the transaction in the blockchain.
		sdk.NewDeadline(time.Hour * 1),
		// Id of mosaic to change supply
		mosaicId,
		// Supply change direction
		sdk.Increase,
		// Delta
		sdk.Duration(100000000000),
	)
	if err != nil {
		fmt.Printf("NewMosaicSupplyChangeTransaction returned error: %s", err)
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