package main

import (
	"github.com/OlegChuev/encryptor/pkg/encryption"
	"github.com/OlegChuev/encryptor/pkg/filesystem"

	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	encryptFlag, decryptFlag, sourceFlag, destFlag, keyFlag, helpFlag := parseFlags()

	if *helpFlag {
		flag.PrintDefaults()
		return
	}

	if *encryptFlag && *decryptFlag {
		fmt.Println("Please specify either -encrypt or -decrypt, not both.")
		return
	}

	if *sourceFlag == "" || *destFlag == "" {
		fmt.Println("Please provide both source and destination file paths.")
		return
	}

	data, err := ioutil.ReadFile(*sourceFlag)
	if err != nil {
		fmt.Println("Error reading source file:", err)
		return
	}

	result, err := process(data, *encryptFlag, *keyFlag)
	if err != nil {
		var processName string
		if *encryptFlag {
			processName = "encryption"
		} else {
			processName = "decryption"
		}

		errMsg := fmt.Sprintf("Error during the %s process:", processName)
		fmt.Println(errMsg, err)
		return
	}

	err = filesystem.WriteToFile(*destFlag, result)
	if err != nil {
		fmt.Println("Error writing to destination file:", err)
		return
	}

	fmt.Println("Operation completed successfully.")
}

func parseFlags() (*bool, *bool, *string, *string, *string, *bool) {
	encryptFlag := flag.Bool("encrypt", false, "Encrypts data")
	decryptFlag := flag.Bool("decrypt", false, "Decrypts data")
	sourceFlag := flag.String("src", "", "Source file")
	destFlag := flag.String("dest", "", "Destination file")
	keyFlag := flag.String("key", "", "Private key string")
	helpFlag := flag.Bool("help", false, "Print this help message")

	flag.Parse()

	return encryptFlag, decryptFlag, sourceFlag, destFlag, keyFlag, helpFlag
}

func process(data []byte, isEncryption bool, key string) (result []byte, err error) {
	if isEncryption {
		result, err = encryption.Encrypt(data, key)
	} else {
		result, err = encryption.Decrypt(data, key)
	}

	return
}
