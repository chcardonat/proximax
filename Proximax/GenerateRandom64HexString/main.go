package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/proximax-storage/go-xpx-chain-sdk/sdk"
)

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func main() {

	val, _ := randomHex(32)
	fmt.Println(val)
	//Create account from private key
	aliceAccount, _ := sdk.NewAccountFromPrivateKey(val, sdk.PublicTest, &sdk.Hash{})

	fmt.Printf("Address:\t%v\n", aliceAccount.Address)
	fmt.Printf("PrivateKey:\t%x\n", aliceAccount.KeyPair.PrivateKey.Raw)
	fmt.Printf("PublicKey:\t%x", aliceAccount.KeyPair.PublicKey.Raw)

}
