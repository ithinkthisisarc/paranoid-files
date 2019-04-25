package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	keyFile       = "aes.key"
	encryptedFile = "aes.enc"
)

var IV = []byte("ragARJaaNGaiHjaH")

func readKey(filename string) ([]byte, error) {
	key, err := ioutil.ReadFile(filename)
	if err != nil {
		return key, err
	}
	block, _ := pem.Decode(key)
	return block.Bytes, nil
}

func createKey() []byte {
	genkey := make([]byte, 16)
	_, err := rand.Read(genkey)
	if err != nil {
		log.Fatalf("Failed to read new random key: %s", err)
	}
	return genkey
}

func saveKey(filename string, key []byte) {
	block := &pem.Block{
		Type:  "AES KEY",
		Bytes: key,
	}
	err := ioutil.WriteFile(filename, pem.EncodeToMemory(block), 0644)
	if err != nil {
		log.Fatalf("Failed in saving key to %s: %s", filename, err)
	}
}

func aesKey() []byte {
	file := fmt.Sprintf(keyFile)
	key, err := readKey(file)
	if err != nil {
		log.Println("Creating a new AES key")
		key = createKey()
		saveKey(file, key)
	}
	return key
}

func createCipher() cipher.Block {
	c, err := aes.NewCipher(aesKey())
	if err != nil {
		log.Fatalf("Failed to create the AES cipher: %s", err)
	}
	return c
}

func encryption(plainText string, n string) {
	bytes := []byte(plainText)
	blockCipher := createCipher()
	stream := cipher.NewCTR(blockCipher, IV)
	stream.XORKeyStream(bytes, bytes)
	err := ioutil.WriteFile(n, bytes, 0644)
	if err != nil {
		log.Fatalf("Writing encryption file: %s", err)
	} else {
		fmt.Printf("Message encrypted in file: %s\n\n", n)
	}
}

func decryption(n string) []byte {
	bytes, err := ioutil.ReadFile(n)
	if err != nil {
		log.Fatalf("Reading encrypted file: %s", err)
	}
	blockCipher := createCipher()
	stream := cipher.NewCTR(blockCipher, IV)
	stream.XORKeyStream(bytes, bytes)
	return bytes
}

func mix(text string, x string) {
	a := strings.Split(text, "")
	for j, i := range a {
		switch strings.ToLower(i) {
		case "a":
			a[j] = "t"
		case "b":
			a[j] = "g"
		case "c":
			a[j] = "e"
		case "d":
			a[j] = "q"
		case "e":
			a[j] = "u"
		case "f":
			a[j] = "i"
		case "g":
			a[j] = "c"
		case "h":
			a[j] = "k"
		case "i":
			a[j] = "b"
		case "j":
			a[j] = "r"
		case "k":
			a[j] = "o"
		case "l":
			a[j] = "w"
		case "m":
			a[j] = "n"
		case "n":
			a[j] = "f"
		case "o":
			a[j] = "x"
		case "p":
			a[j] = "j"
		case "q":
			a[j] = "m"
		case "r":
			a[j] = "p"
		case "s":
			a[j] = "s"
		case "t":
			a[j] = "v"
		case "u":
			a[j] = "h"
		case "v":
			a[j] = "l"
		case "w":
			a[j] = "a"
		case "x":
			a[j] = "z"
		case "y":
			a[j] = "y"
		case "z":
			a[j] = "d"
		}
		b := strings.Join(a, "")
		err := ioutil.WriteFile(x, []byte(b), 0644)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
	}
}

func unmix(file string) []byte {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Error reading file:\n ===> %v", err)
	}
	a := strings.Split(string(data), "")
	for j, i := range a {
		switch strings.ToLower(i) {
		case "t":
			a[j] = "a"
		case "g":
			a[j] = "b"
		case "e":
			a[j] = "c"
		case "q":
			a[j] = "d"
		case "u":
			a[j] = "e"
		case "i":
			a[j] = "f"
		case "c":
			a[j] = "g"
		case "k":
			a[j] = "h"
		case "b":
			a[j] = "i"
		case "r":
			a[j] = "j"
		case "o":
			a[j] = "k"
		case "w":
			a[j] = "l"
		case "n":
			a[j] = "m"
		case "f":
			a[j] = "n"
		case "x":
			a[j] = "o"
		case "j":
			a[j] = "p"
		case "m":
			a[j] = "q"
		case "p":
			a[j] = "r"
		case "s":
			a[j] = "s"
		case "v":
			a[j] = "t"
		case "h":
			a[j] = "u"
		case "l":
			a[j] = "v"
		case "a":
			a[j] = "w"
		case "z":
			a[j] = "x"
		case "y":
			a[j] = "y"
		case "d":
			a[j] = "z"
		}
	}
	b := strings.Join(a, "")
	return []byte(b)
}

func read() string {
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	return reader.Text()
}
