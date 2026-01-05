package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	fmt.Println("-- Server AES-128 Decryption Test --")
	fmt.Println("AES-128 Initial Key:", initialKey)
	fmt.Println()
	fmt.Println("AES-128 Expanded Key:", expandKey(initialKey))

	var encryptedDataBytes []byte
	if len(os.Args) == 1 {
		encryptedDataBytes = []byte("Hello ! This is working ! Leonid")
	} else {
		decoded, err := base64.StdEncoding.DecodeString(os.Args[1])
		if err != nil {
			fmt.Println("Failed to decode base64:", err)
			return
		}
		encryptedDataBytes = decoded
	}

	fmt.Println()
	if len(os.Args) == 1 {
		fmt.Println("Encrypted Data:", string(encryptedDataBytes))
	} else {
		fmt.Println("Encrypted Data (base64):", os.Args[1])
	}
	fmt.Println("Encrypted Data In Bytes:", encryptedDataBytes)
	fmt.Println("Encrypted Data length:", len(encryptedDataBytes))
	fmt.Println()
	decryptedData := aes128Decrypt(&encryptedDataBytes, expandKey(initialKey))
	fmt.Println("Decrypted Data In Bytes:", decryptedData)
	fmt.Println("Decrypted Data String:", string(decryptedData))
	fmt.Println("Decrypted Data length:", len(string(decryptedData)))
}
