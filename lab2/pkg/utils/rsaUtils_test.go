package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateAndReadKeys(t *testing.T) {
	privateKeyFile := "private_key_test.pem"
	publicKeyFile := "public_key_test.pem"

	privateKey, publicKey := GenerateRSAKeyPair(2048)

	err := SavePrivateKeyToFile(privateKey, privateKeyFile)
	assert.NoError(t, err)

	err = SavePublicKeyToFile(publicKey, publicKeyFile)
	assert.NoError(t, err)

	loadedPrivateKey, err := ReadPrivateKeyFromFile(privateKeyFile)
	assert.NoError(t, err)
	assert.Equal(t, privateKey.N, loadedPrivateKey.N)

	loadedPublicKey, err := ReadPublicKeyFromFile(publicKeyFile)
	assert.NoError(t, err)
	assert.Equal(t, publicKey.N, loadedPublicKey.N)
}

func TestEncryptAndDecryptData(t *testing.T) {
	privateKey, publicKey := GenerateRSAKeyPair(2048)
	data := []byte("Test data")

	encryptedData, err := EncryptData(publicKey, data)
	assert.NoError(t, err)

	decryptedData, err := DecryptData(privateKey, encryptedData)
	assert.NoError(t, err)
	assert.Equal(t, data, decryptedData)
}
