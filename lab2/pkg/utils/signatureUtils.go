package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"io/ioutil"
)

func HashData(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

func SignData(privateKey *rsa.PrivateKey, hashedData []byte) ([]byte, error) {
	return rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashedData)
}

func VerifySignature(publicKey *rsa.PublicKey, hashedData []byte, signature []byte) error {
	return rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashedData, signature)
}

func SaveSignatureToFile(signature []byte, filename string) error {
	return ioutil.WriteFile(filename, signature, 0644)
}

func ReadSignatureFromFile(filename string) ([]byte, error) {
	signature, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return signature, nil
}
