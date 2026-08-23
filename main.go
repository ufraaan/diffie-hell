package main

import (
	"fmt"
	"math/big"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"crypto/rand"
)

func main() {
	g := big.NewInt(100)
	p := big.NewInt(352)

	//generate alice's private key as an int64, then convert to big.Int
	alicePvtKey, _ := rand.Int(rand.Reader, p)
	fmt.Println("alice private key:", alicePvtKey)

	alicePublicKey := new(big.Int)

	//calculate (g ^ alicePvtKey) % p safely
	alicePublicKey.Exp(g, alicePvtKey, p)

	fmt.Println("alice public key:", alicePublicKey)

	// generate bob's private key as int64, then convert to big.Int
	bobPvtKey, _ := rand.Int(rand.Reader, p)
	fmt.Println("bob private key: ", bobPvtKey)

	bobPublicKey := new(big.Int)

	// calculate (g ^ bobPvtKey) % p safely
	bobPublicKey.Exp(g, bobPvtKey, p)

	fmt.Println("bob public key:", bobPublicKey)

	// now we can exchange the public keys and calculate the shared secret
	aliceSharedSecret := new(big.Int).Exp(bobPublicKey, alicePvtKey, p)
	bobSharedSecret := new(big.Int).Exp(alicePublicKey, bobPvtKey, p)
	// both have to be the same

	fmt.Println("alice shared secret: ", aliceSharedSecret)
	fmt.Println("bob shared secret: ", bobSharedSecret)
	if (aliceSharedSecret.Cmp(bobSharedSecret) == 0) {
		fmt.Println("secrets matched")
	} else {
		fmt.Println("error")
	}



	// deriving 32byte aes256 key from the shared secret
	key := sha256.Sum256(aliceSharedSecret.Bytes())

	block, err := aes.NewCipher(key[:])
	if err!= nil {
		panic(err)
	}
	gcm, err := cipher.NewGCM(block) // returns the given 128-bit, block cipher wrapped in Galois Counter Mode with the standard nonce length
	if err!= nil {
		panic(err)
	}

}
