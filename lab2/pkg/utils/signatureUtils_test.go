package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveAndReadSignature(t *testing.T) {
	signatureFile := "signature_test.txt"
	signature := []byte("test_signature")

	err := SaveSignatureToFile(signature, signatureFile)
	assert.NoError(t, err)

	loadedSignature, err := ReadSignatureFromFile(signatureFile)
	assert.NoError(t, err)
	assert.Equal(t, signature, loadedSignature)
}
