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
	// Future multisig private key
	multisigPrivateKey = "C8BB4994000422DE988A6FA78E5AE140E397B61C6CC0FCAA7D24D493D6CB743E"
	// Cosignature public keys
	cosignatoryOnePublicKey   = "B28D03B13863FF57B3F81581C02CBB01F0E4253CE6C7A9E02F8EBE4CD556BC52"
	cosignatoryTwoPublicKey   = "2F8B5A224A1B32775D49B3A26B58706158660852C26B07244A0B61DF45FAC8DD"
	cosignatoryThreePublicKey = "FFF228306B2E14F4F48F031DE8CFF5E6F1D4B87C348E6B9D811D9AA3ACFE94A5"
	// Minimal approval count
	minimalApproval = 3
	// Minimal removal count
	minimalRemoval = 2
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
	multisig, err := client.NewAccountFromPrivateKey(multisigPrivateKey)
	if err != nil {
		fmt.Printf("NewAccountFromPublicKey returned error: %s", err)
		return
	}

	cosignerOne, err := client.NewAccountFromPublicKey(cosignatoryOnePublicKey)
	if err != nil {
		fmt.Printf("NewAccountFromPrivateKey returned error: %s", err)
		return
	}
	cosignerTwo, err := client.NewAccountFromPublicKey(cosignatoryTwoPublicKey)
	if err != nil {
		fmt.Printf("NewAccountFromPrivateKey returned error: %s", err)
		return
	}
	cosignerThree, err := client.NewAccountFromPublicKey(cosignatoryThreePublicKey)
	if err != nil {
		fmt.Printf("NewAccountFromPrivateKey returned error: %s", err)
		return
	}

	transaction, err := client.NewModifyMultisigAccountTransaction(
		// The maximum amount of time to include the transaction in the blockchain.
		sdk.NewDeadline(time.Hour*1),
		// The number of signatures needed to approve a transaction.
		minimalApproval,
		// The number of signatures needed to remove a cosignatory.
		minimalRemoval,
		// Array of cosigner accounts added or removed from the multisignature account.
		[]*sdk.MultisigCosignatoryModification{
			{sdk.Add, cosignerOne},
			{sdk.Add, cosignerTwo},
			{sdk.Add, cosignerThree},
		},
	)
	if err != nil {
		fmt.Printf("NewModifyMultisigAccountTransaction returned error: %s", err)
		return
	}

	// Convert transactions to inner for an aggregate transaction
	transaction.ToAggregate(multisig.PublicAccount)

	aggregateBondedTransaction, err := client.NewBondedAggregateTransaction(sdk.NewDeadline(time.Hour), []sdk.Transaction{transaction})

	// Sign transaction
	signedAggregateBoundedTransaction, err := multisig.Sign(aggregateBondedTransaction)
	if err != nil {
		fmt.Printf("Sign returned error: %s", err)
		return
	}
	fmt.Printf("Content: \t\t%v", signedAggregateBoundedTransaction.Hash)

	// Create lock funds transaction for aggregate bounded
	lockFundsTransaction, err := client.NewLockFundsTransaction(
		// The maximum amount of time to include the transaction in the blockchain.
		sdk.NewDeadline(time.Hour),
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
	signedLockFundsTransaction, err := multisig.Sign(lockFundsTransaction)
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

}