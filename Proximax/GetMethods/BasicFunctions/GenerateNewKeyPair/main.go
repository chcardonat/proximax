package main

import (
	"fmt"
	crypto "github.com/proximax-storage/go-xpx-crypto"
)

func main() {

	KeyPair, _ := crypto.NewRandomKeyPair()

	fmt.Printf("PublicKey:\t%x\n", KeyPair.PublicKey.Raw)
	fmt.Printf("PrivateKey:\t%x\n", KeyPair.PrivateKey.Raw)
}