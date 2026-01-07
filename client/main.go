package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	fmt.Println("-- Client AES-128 Encryption Test --")
	fmt.Println("AES-128 Initial Key:", initialKey)
	fmt.Println()
	fmt.Println("AES-128 Expanded Key:", expandKey(initialKey))

	initialData := flag.String("data", "Hello ! This is working ! Leonid", "Data to encrypt")
	serverProg := flag.String("server", "../server/server.exe", "Path to server executable")
	flag.Parse()

	initialDataBytes := []byte(*initialData)
	fmt.Println()
	fmt.Println("Initial Data:", *initialData)
	fmt.Println("Initial Data In Bytes:", initialDataBytes)
	fmt.Println("Initial Data length:", len(initialDataBytes))
	fmt.Println()

	encryptedData := aesEncryptBlock(&initialDataBytes, expandKey(initialKey))
	fmt.Println("Encrypted Data In Bytes:", encryptedData)
	b64 := base64.StdEncoding.EncodeToString(encryptedData)
	fmt.Println("Encrypted Data (base64):", b64)
	fmt.Println("Encrypted Data length (bytes):", len(encryptedData))

	fmt.Println()
	fmt.Println("-- Calling server to decrypt --")
	cmd := exec.Command(*serverProg, b64)

	output, err := cmd.CombinedOutput()
	if err != nil {
		// Log the error if the command fails
		log.Fatalf("cmd.CombinedOutput failed with %s\n", err)
	}

	// Print the output as a string
	fmt.Printf("Output:\n\n%s\n", string(output))
}
