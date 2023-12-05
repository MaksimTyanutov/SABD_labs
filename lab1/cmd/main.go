package main

import (
	"flag"
	"fmt"
	"os"

	"sabd1/pkg/utils"
)

func main() {
	filePath := flag.String("file", "", "Path to the file")
	mode := flag.Int("mode", 0, "Mode: 1 for obfuscation, 2 for deobfuscation")

	flag.Parse()

	if *filePath == "" {
		fmt.Println("Specify the file path")
		os.Exit(1)
	}

	shift := 3

	err := utils.ProcessFile(*filePath, *mode, shift)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Operation completed successfully.")
}
