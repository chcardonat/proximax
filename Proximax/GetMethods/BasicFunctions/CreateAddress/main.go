package main

import (
	"fmt"
	"github.com/proximax-storage/go-xpx-chain-sdk/sdk"
)

func main() {

	publicKey := "04dd376196603c44a19fd500492e5de12de9ed353de070a788cb21f210645613"


	Address, _ := sdk.NewAddressFromPublicKey(publicKey, sdk.PublicTest )

	fmt.Printf("Address:\t\t%v\n",Address.Address)
	fmt.Printf("NetworkType:\t%v\n",Address.Type)


}