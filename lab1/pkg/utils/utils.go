package utils

import (
	"fmt"
	"io/ioutil"
	"sabd1/pkg/obfuscation"
)

func readFromFile(filePath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("rror reading file: %v", err)
	}
	return data, nil
}

func writeToFile(filePath string, data []byte) error {
	err := ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}
	return nil
}

func ProcessFile(filePath string, mode int, shift int) error {
	data, err := readFromFile(filePath)
	if err != nil {
		return err
	}

	var result []byte

	switch mode {
	case 1:
		result = obfuscation.Obfuscate(data, shift)
	case 2:
		result = obfuscation.Deobfuscate(data, shift)
	default:
		return fmt.Errorf("invalid mode. Use 1 for obfuscation or 2 for deobfuscation")
	}

	err = writeToFile(filePath, result)
	if err != nil {
		return err
	}

	return nil
}
