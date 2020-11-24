package main

import (
	"fmt"
	"github.com/proximax-storage/go-xpx-chain-sdk/sdk"
)

const alicePrivateKey = "04dd376196603c44a19fd500492e5de12de9ed353de070a788cb21f210645612"

func main() {

	//Create account from private key
	aliceAccount, _ := sdk.NewAccountFromPrivateKey(alicePrivateKey, sdk.PublicTest, &sdk.Hash{})

	fmt.Printf("Address:\t%v\n", aliceAccount.Address)
	fmt.Printf("PrivateKey:\t%x\n", aliceAccount.KeyPair.PrivateKey.Raw)
	fmt.Printf("PublicKey:\t%x", aliceAccount.KeyPair.PublicKey.Raw)
	fmt.Println(len(alicePrivateKey))


}
