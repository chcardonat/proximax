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
	account, err := sdk.NewAccountFromPrivateKey(privateKey, networkType, client.GenerationHash())
	if err != nil {
		fmt.Printf("NewAccountFromPrivateKey returned error: %s", err)
		return
	}

	random := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	nonce := random.Uint32()
	// Create mosaic id of just published mosaic
	mosaicId, err := sdk.NewMosaicIdFromNonceAndOwner(nonce, account.PublicAccount.PublicKey)
	if err != nil {
		fmt.Printf("NewMosaicIdFromNonceAndOwner returned error: %s", err)
		return
	}

	// Create a new account properties mosaic type transaction
	transaction, err := sdk.NewAccountPropertiesMosaicTransaction(
		// The maximum amount of time to include the transaction in the blockchain.
		sdk.NewDeadline(time.Hour * 1),
		// Allow transactions with passed mosaic
		sdk.AllowMosaic,
		// Account properties to update
		[]*sdk.AccountPropertiesMosaicModification{
			{
				sdk.AddProperty,
				mosaicId,
			},
		},
		networkType,
	)
	if err != nil {
		fmt.Printf("NewAccountPropertiesMosaicTransaction returned error: %s", err)
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

	// wait for the transaction to be confirmed! (very important)
	// you can use websockets to wait explicitly for transaction
	// to be in certain state, instead of hard waiting
	time.Sleep(time.Second * 30)
	fmt.Printf("Content: \t\t%v", signedTransaction.Hash)

}