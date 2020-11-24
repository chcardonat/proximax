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
	account, err := sdk.NewAccountFromPrivateKey(privateKey, networkType, client.GenerationHash())
	if err != nil {
		fmt.Printf("NewAccountFromPrivateKey returned error: %s", err)
		return
	}
	// Create an account for blocking
	accountToBlock, err := sdk.NewAccount(networkType, client.GenerationHash())
	if err != nil {
		fmt.Printf("NewAccount returned error: %s", err)
		return
	}

	// Create a new account properties address type transaction
	transaction, err := sdk.NewAccountPropertiesAddressTransaction(
		// The maximum amount of time to include the transaction in the blockchain.
		sdk.NewDeadline(time.Hour*1),
		// Block transactions from addresses
		sdk.BlockAddress,
		// Account properties to update
		[]*sdk.AccountPropertiesAddressModification{
			{
				sdk.AddProperty,
				accountToBlock.Address,
			},
		},
		networkType,
	)
	if err != nil {
		fmt.Printf("NewAccountPropertiesAddressTransaction returned error: %s", err)
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
	time.Sleep(time.Second * 5)
	fmt.Printf("Content: \t\t%v", signedTransaction.Hash)
}
