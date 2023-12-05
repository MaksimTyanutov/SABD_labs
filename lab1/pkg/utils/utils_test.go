package utils

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestReadFromFile(t *testing.T) {
	tempFile, err := ioutil.TempFile("", "testfile")
	if err != nil {
		t.Fatalf("Error creating temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	expectedData := []byte("test data")
	err = ioutil.WriteFile(tempFile.Name(), expectedData, 0644)
	if err != nil {
		t.Fatalf("Error writing to temporary file: %v", err)
	}

	data, err := readFromFile(tempFile.Name())
	if err != nil {
		t.Fatalf("Unexpected error reading from file: %v", err)
	}

	if string(data) != string(expectedData) {
		t.Errorf("Expected %v, but got %v", expectedData, data)
	}
}

func TestWriteToFile(t *testing.T) {
	tempFile, err := ioutil.TempFile("", "testfile")
	if err != nil {
		t.Fatalf("Error creating temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	data := []byte("test data")
	err = writeToFile(tempFile.Name(), data)
	if err != nil {
		t.Fatalf("Unexpected error writing to file: %v", err)
	}

	readData, err := ioutil.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatalf("Error reading from temporary file: %v", err)
	}

	if string(readData) != string(data) {
		t.Errorf("Expected %v, but got %v", data, readData)
	}
}

func TestProcessFile(t *testing.T) {
	tempFile, err := ioutil.TempFile("", "testfile")
	if err != nil {
		t.Fatalf("Error creating temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	expectedData := []byte("test data")
	err = ioutil.WriteFile(tempFile.Name(), expectedData, 0644)
	if err != nil {
		t.Fatalf("Error writing to temporary file: %v", err)
	}

	err = ProcessFile(tempFile.Name(), 1, 3)
	if err != nil {
		t.Fatalf("Unexpected error processing file: %v", err)
	}

	obfuscatedData, err := ioutil.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatalf("Error reading from temporary file: %v", err)
	}

	if string(obfuscatedData) == string(expectedData) {
		t.Error("Expected obfuscated data, but got the original data")
	}

	err = ProcessFile(tempFile.Name(), 2, 3)
	if err != nil {
		t.Fatalf("Unexpected error processing file: %v", err)
	}

	deobfuscatedData, err := ioutil.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatalf("Error reading from temporary file: %v", err)
	}

	if string(deobfuscatedData) != string(expectedData) {
		t.Errorf("Expected %v, but got %v", expectedData, deobfuscatedData)
	}
}
