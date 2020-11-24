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
	firstPrivateKey = "B0835BB2375A30D81F232C4ED369A9F1337A04CFD5BA4A5758D29E67037D2AB9"
	secondPrivateKey = "04DD376196603C44A19FD500492E5DE12DE9ED353DE070A788CB21F210645613"
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
	firstAccount, err := client.NewAccountFromPrivateKey(firstPrivateKey)
	if err != nil {
		fmt.Printf("NewAccountFromPrivateKey returned error: %s", err)
		return
	}
	secondAccount, err := client.NewAccountFromPrivateKey(secondPrivateKey)
	if err != nil {
		fmt.Printf("NewAccountFromPrivateKey returned error: %s", err)
		return
	}

	// Create a new transfer type transaction
	firstTransferTransaction, err := client.NewTransferTransaction(
		// The maximum amount of time to include the transaction in the blockchain.
		sdk.NewDeadline(time.Hour*1),
		// The address of the recipient account.
		secondAccount.Address,
		// The array of mosaic to be sent.
		[]*sdk.Mosaic{sdk.Xpx(10)},
		// The transaction message of 1024 characters.
		sdk.NewPlainMessage("Let's exchange 10 xpx -> 20 xem"),
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
		firstAccount.Address,
		// The array of mosaic to be sent.
		[]*sdk.Mosaic{sdk.Xem(20)},
		// The transaction message of 1024 characters.
		sdk.NewPlainMessage("Okay"),
	)
	if err != nil {
		fmt.Printf("NewTransferTransaction returned error: %s", err)
		return
	}

	// Convert an aggregate transaction to an inner transaction including transaction signer.
	firstTransferTransaction.ToAggregate(firstAccount.PublicAccount)
	secondTransferTransaction.ToAggregate(secondAccount.PublicAccount)

	aggregateBoundedTransaction, err := client.NewBondedAggregateTransaction(
		// The maximum amount of time to include the transaction in the blockchain.
		sdk.NewDeadline(time.Hour*1),
		// Inner transactions
		[]sdk.Transaction{firstTransferTransaction, secondTransferTransaction},
	)
	if err != nil {
		fmt.Printf("NewBondedAggregateTransaction returned error: %s", err)
		return
	}

	// Sign transaction
	signedAggregateBoundedTransaction, err := firstAccount.Sign(aggregateBoundedTransaction)
	if err != nil {
		fmt.Printf("Sign returned error: %s", err)
		return
	}
	fmt.Printf("Content: \t\t%v", signedAggregateBoundedTransaction.Hash)


	// Create lock funds transaction for aggregate bounded
	lockFundsTransaction, err := client.NewLockFundsTransaction(
		// The maximum amount of time to include the transaction in the blockchain.
		sdk.NewDeadline(time.Hour*1),
		// Funds to lock
		sdk.XpxRelative(10),
		// Duration of lock transaction in blocks
		sdk.Duration(1000),
		// Aggregate bounded transaction for lock
		signedAggregateBoundedTransaction,
	)
	if err != nil {
		fmt.Printf("NewLockFundsTransaction returned error: %s", err)
		return
	}

	// Sign transaction
	signedLockFundsTransaction, err := firstAccount.Sign(lockFundsTransaction)
	if err != nil {
		fmt.Printf("Sign returned error: %s", err)
		return
	}
	fmt.Printf("Content: \t\t%v", signedLockFundsTransaction.Hash)

	// Announce transaction
	_, err = client.Transaction.Announce(context.Background(), signedLockFundsTransaction)
	if err != nil {
		fmt.Printf("Transaction.Announce returned error: %s", err)
		return
	}

	// Wait for lock funds transaction to be harvested
	time.Sleep(30 * time.Second)

	// Announce aggregate bounded transaction
	_, _ = client.Transaction.AnnounceAggregateBonded(context.Background(), signedAggregateBoundedTransaction)
	if err != nil {
		fmt.Printf("Transaction.AnnounceAggregateBonded returned error: %s", err)
		return
	}

	// Wait for aggregate bounded transaction to be harvested
	time.Sleep(30 * time.Second)

	// Create cosignature transaction from first account
	firstAccountCosignatureTransaction := sdk.NewCosignatureTransactionFromHash(signedAggregateBoundedTransaction.Hash)
	signedFirstAccountCosignatureTransaction, err := firstAccount.SignCosignatureTransaction(firstAccountCosignatureTransaction)
	if err != nil {
		fmt.Printf("SignCosignatureTransaction returned error: %s", err)
		return
	}

	// Announce transaction
	_, err = client.Transaction.AnnounceAggregateBondedCosignature(context.Background(), signedFirstAccountCosignatureTransaction)
	if err != nil {
		fmt.Printf("AnnounceAggregateBoundedCosignature returned error: %s", err)
		return
	}

	// Create cosignature transaction from second account
	secondAccountCosignatureTransaction := sdk.NewCosignatureTransactionFromHash(signedAggregateBoundedTransaction.Hash)
	signedSecondAccountCosignatureTransaction, err := secondAccount.SignCosignatureTransaction(secondAccountCosignatureTransaction)
	if err != nil {
		fmt.Printf("SignCosignatureTransaction returned error: %s", err)
		return
	}



	// Announce transaction
	_, err = client.Transaction.AnnounceAggregateBondedCosignature(context.Background(), signedSecondAccountCosignatureTransaction)
	if err != nil {
		fmt.Printf("AnnounceAggregateBoundedCosignature returned error: %s", err)
		return
	}
	// wait for the transaction to be confirmed! (very important)
	// you can use websockets to wait explicitly for transaction
	// to be in certain state, instead of hard waiting
	time.Sleep(30 * time.Second)

}
