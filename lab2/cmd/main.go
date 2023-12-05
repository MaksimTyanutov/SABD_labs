// main.go
package main

import (
	"crypto/rsa"
	"fmt"
	"os"
	"sabd2/pkg/utils"
)

func main() {
	privateKeyFile := "private_key.pem"
	publicKeyFile := "public_key.pem"
	signatureFile := "signature.txt"

	// Read or generate keys
	var privateKey *rsa.PrivateKey
	var publicKey *rsa.PublicKey

	if _, err := os.Stat(privateKeyFile); os.IsNotExist(err) {
		privateKey, publicKey = utils.GenerateRSAKeyPair(2048)
		utils.SavePrivateKeyToFile(privateKey, privateKeyFile)
		utils.SavePublicKeyToFile(publicKey, publicKeyFile)
	} else {
		privateKey, _ = utils.ReadPrivateKeyFromFile(privateKeyFile)
		publicKey, _ = utils.ReadPublicKeyFromFile(publicKeyFile)
	}

	data := []byte("Hello, World!")

	encryptedData, err := utils.EncryptData(publicKey, data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Encrypted Data: %s\n", encryptedData)

	decryptedData, err := utils.DecryptData(privateKey, encryptedData)
	if err != nil {
		panic(err)
	}

	fmt.Println("Decrypted Data:", string(decryptedData))

	hashedData := utils.HashData(data)

	signature, err := utils.SignData(privateKey, hashedData)
	if err != nil {
		panic(err)
	}

	fmt.Println("Signature:", signature)

	err = utils.SaveSignatureToFile(signature, signatureFile)
	if err != nil {
		panic(err)
	}

	loadedSignature, err := utils.ReadSignatureFromFile(signatureFile)
	if err != nil {
		panic(err)
	}

	err = utils.VerifySignature(publicKey, hashedData, loadedSignature)
	if err != nil {
		panic(err)
	}

	fmt.Println("Signature Verified!")
}
