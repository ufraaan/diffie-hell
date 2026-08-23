package main

import (
	"fmt"
	"math/big"
	"math/rand/v2"
)

func main() {
	g := big.NewInt(100)
	p := big.NewInt(352)

	//generate alice's private key as an int64, then convert to big.Int
	alicePvtKeyInt := rand.Int64()
	alicePvtKey := big.NewInt(alicePvtKeyInt)
	fmt.Println("alice private key:", alicePvtKey)

	alicePublicKey := new(big.Int)

	//calculate (g ^ alicePvtKey) % p safely
	alicePublicKey.Exp(g, alicePvtKey, p)

	fmt.Println("alice public key:", alicePublicKey)
	


	// generate bob's private key as int64, then convert to big.Int
	bobPvtKeyInt := rand.Int64()
	bobPvtKey := big.NewInt(bobPvtKeyInt)
	fmt.Println("bob private key: ", bobPvtKey)

	bobPublicKey := new(big.Int)

	// calculate (g ^ bobPvtKey) % p safely
	bobPublicKey.Exp(g, bobPvtKey, p)

	fmt.Println("bob public key:", bobPublicKey)


}
