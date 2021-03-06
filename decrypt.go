package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}

func main() {
	/* The input needed for ciphertext is byte[]*/
	fmt.Println("Enter your ciphertext")
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		log.Printf("Failed to read: %v", scanner.Err())
		return
	}
	/*Input is bytes[] repsenteded in hexadecimal notation*/
	ciphertext := scanner.Text()

	//TODO
	/*Decoding from hexadecimal notation back to bytes?*/
	str, err := hex.DecodeString(ciphertext)
	if err != nil {
		fmt.Println("Can't decode the string :/")
	}

	//TODO Input is recieved in hexadecimal nottation
	// Need to decode into byte[] for it to make sense
	plaintext := decrypt(str, "rsgchristmas2019")
	fmt.Println("Decyrption successful!")
	fmt.Printf("Decrypted ciphertext: %s\n", plaintext)

}
