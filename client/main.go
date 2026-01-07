package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	// print initial and expanded keys - for testing only
	fmt.Println("-- Client AES-128 Encryption Test --")
	fmt.Println("AES-128 Initial Key:", initialKey)
	fmt.Println()
	fmt.Println("AES-128 Expanded Key:", expandKey(initialKey))

	// program flags parsing
	initialData := flag.String("data", "Hello ! This is working ! Leonid", "Data to encrypt")
	localServerProg := flag.String("local-server", "../server/build/server.exe", "Path to local server executable file")
	remoteServerAddr := flag.String("remote-server", "", "Remote server address (if empty, local server is used)")
	userName := flag.String("user", "", "MQTT username")
	password := flag.String("pass", "", "MQTT password")
	flag.Parse()

	// convert initial data to bytes
	initialDataBytes := []byte(*initialData)
	fmt.Println()
	fmt.Println("Initial Data:", *initialData)
	fmt.Println("Initial Data In Bytes:", initialDataBytes)
	fmt.Println("Initial Data length:", len(initialDataBytes))
	fmt.Println()

	// encrypt data
	encryptedData := aesEncryptBlock(&initialDataBytes, expandKey(initialKey))
	fmt.Println("Encrypted Data In Bytes:", encryptedData)
	b64 := base64.StdEncoding.EncodeToString(encryptedData)
	fmt.Println("Encrypted Data (base64):", b64)
	fmt.Println("Encrypted Data length (bytes):", len(encryptedData))

	// call server to receive and decrypt
	fmt.Println()
	fmt.Println("-- Calling server to decrypt --")

	if *remoteServerAddr == "" {
		cmd := exec.Command(*localServerProg, b64)

		output, err := cmd.CombinedOutput()
		if err != nil {
			// Log the error if the command fails
			log.Fatalf("cmd.CombinedOutput failed with %s\n", err)
		}

		// Print the output as a string
		fmt.Printf("Output:\n\n%s\n", string(output))
		return
	}

	// connect to MQTT broker if remote server address is provided
	mqttBroker := "tcp://" + *remoteServerAddr + ":1883"
	err := mqttConnect(mqttBroker, *userName, *password)
	if err != nil {
		log.Fatalf("Failed to connect to MQTT broker: %v", err)
	}
	defer mqttDisconnect()

	topic := "test/aes/encrypt"
	err = mqttPublish(topic, encryptedData)
}
