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
	multisigPrivateKey = "A84A3EFDDED96B64E6AFA7A7E989D858AA73942D8C4616F38B3C0F26A2F78963"
	// Cosignature private keys
	cosignatoryOnePrivateKey   = "B0835BB2375A30D81F232C4ED369A9F1337A04CFD5BA4A5758D29E67037D2AB9"
	cosignatoryTwoPrivateKey   = "04DD376196603C44A19FD500492E5DE12DE9ED353DE070A788CB21F210645613"
	cosignatoryThreePrivateKey = "0026693B5A949C95C7FF69D700DD45D5C7108C76B8E205E8991CA3C56BC86C90"
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

	cosignerOne, err := client.NewAccountFromPrivateKey(cosignatoryOnePrivateKey)
	if err != nil {
		fmt.Printf("NewAccountFromPrivateKey returned error: %s", err)
		return
	}
	cosignerTwo, err := client.NewAccountFromPrivateKey(cosignatoryTwoPrivateKey)
	if err != nil {
		fmt.Printf("NewAccountFromPrivateKey returned error: %s", err)
		return
	}
	cosignerThree, err := client.NewAccountFromPrivateKey(cosignatoryThreePrivateKey)
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
			{sdk.Add, cosignerOne.PublicAccount},
			{sdk.Add, cosignerTwo.PublicAccount},
			{sdk.Add, cosignerThree.PublicAccount},
		},
	)
	if err != nil {
		fmt.Printf("NewModifyMultisigAccountTransaction returned error: %s", err)
		return
	}

	// Convert transactions to inner for an aggregate transaction
	// De esta cuenta es que salen los fondos
	transaction.ToAggregate(multisig.PublicAccount)

	aggregateCompletedTransaction, err := client.NewCompleteAggregateTransaction(sdk.NewDeadline(time.Hour), []sdk.Transaction{transaction})

	// Sign transaction
	signedAggregateCompletedTransaction, err := multisig.SignWithCosignatures(aggregateCompletedTransaction, []*sdk.Account{cosignerOne, cosignerTwo, cosignerThree})
	if err != nil {
		fmt.Printf("Sign returned error: %s", err)
		return
	}
	fmt.Printf("Content: \t\t%v", signedAggregateCompletedTransaction.Hash)

	// Announce aggregate bounded transaction
	_, _ = client.Transaction.Announce(context.Background(), signedAggregateCompletedTransaction)
	if err != nil {
		fmt.Printf("Transaction.AnnounceAggregateCompleted returned error: %s", err)
		return
	}

	// Wait for aggregate bounded transaction to be harvested
	time.Sleep(30 * time.Second)

}