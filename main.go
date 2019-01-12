package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func main() {
	var buffer bytes.Buffer
	fmt.Print("Welcome to Paranoid Files\nPlease select a mode (e, d): ")
	mode := read()
	fmt.Print("\n")
	if mode == "e" {
		fmt.Print("Selected encrypt mode\n\n")
		fmt.Print("Enter the name of the file to be encrypted (no extension): ")
		name := read()
		fmt.Print("\nEnter the extension of the file: ")
		ext := read()
		buffer.WriteString(name)
		buffer.WriteString(".")
		buffer.WriteString(ext)
		rfilename := buffer.String()
		buffer.Reset()
		data, err := ioutil.ReadFile(rfilename)
		if err != nil {
			fmt.Println("File reading error:\n", err)
			return
		}
		fmt.Print("Contents of file:", string(data), "\n\n\n")
		fmt.Print("\tEncrypt? (Y/N): ")
		inp := read()
		if strings.ToLower(inp) == "y" {
			cont := string(data)
			encryption(cont, name+".para")
		}
	} else if mode == "d" {
		fmt.Println("Selected decrypt mode\n")
		fmt.Print("Enter the file to decrypt (no extension): ")
		name := read()
		fmt.Print("\nEnter the file extension: ")
		ext := read()
		buffer.WriteString(name)
		buffer.WriteString(".")
		buffer.WriteString(ext)
		rfilename := buffer.String()
		buffer.Reset()
		fmt.Print("Enter the name of the file you want exported: ")
		wfilename := read()
		_, err := ioutil.ReadFile(rfilename)
		if err != nil {
			fmt.Println("File reading error:\n", err)
			return
		}
		fmt.Print("\tDecrypt? (Y/N): ")
		inp := read()
		if strings.ToLower(inp) == "y" {
			err := ioutil.WriteFile(wfilename, []byte(decryption(rfilename)), 0644)
			if err != nil {
				log.Fatalf("error: %v", err)
			}
		}
	} else if mode == "exit" {
		return
	} else {
		fmt.Printf("%s is not a valid mode...\n----------------------\n\n", mode)
		main()
	}
}
