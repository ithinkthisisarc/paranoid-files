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
	fmt.Print("Welcome to Paranoid Files\nPlease select a mode:\n e: encrypt\n d: decrypt\n ")
	mode := read()
	fmt.Print("\n")
	switch strings.ToLower(mode) {
	case "e":
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
			fmt.Println("File reading error:\n\t", err, "\n\nMake sure you're in the correct folder and that you spelled everything correctly.")
			return
		}
		fmt.Println("Contents of file:", string(data))
		fmt.Print("\n\tEncrypt? (Y/N): ")
		inp := read()
		if strings.ToLower(inp) == "y" {
			cont := string(data)
			encryption(cont, name+".para")
		}
	case "d":
		fmt.Println("Selected decrypt mode\n")
		fmt.Print("Enter the file to decrypt (no extension): ")
		name := read()
		buffer.WriteString(name)
		buffer.WriteString(".para")
		rfilename := buffer.String()
		buffer.Reset()
		fmt.Print("Enter the name of the file you want exported: ")
		wfilename := read()
		_, err := ioutil.ReadFile(rfilename)
		if err != nil {
			fmt.Println("File reading error:\n\t", err, "\n\nMake sure you're in the correct folder and that you spelled everything correctly.")
			return
		}
		fmt.Print("\n\tDecrypt? (Y/N): ")
		inp := read()
		if strings.ToLower(inp) == "y" {
			err := ioutil.WriteFile(wfilename, []byte(decryption(rfilename)), 0644)
			if err != nil {
				log.Fatalf("error: %v", err)
			}
		}
	case "b": // TODO
		fmt.Print("Select mode:\n b: byteify\n u: unbyteify\n ")
		m := read()
		switch strings.ToLower(m) {
		case "b":
			fmt.Print("Type the name of the file to byteify (no extension): ")
			name := read()
			fmt.Print("\nType the file extension: ")
			ext := read()
			buffer.Reset()
			buffer.WriteString(name)
			buffer.WriteString(".")
			buffer.WriteString(ext)
			rfilename := buffer.String()
			cont, err := ioutil.ReadFile(rfilename)
			if err != nil {
				fmt.Println("File reading error:\n\t", err, "\n\nMake sure you're in the correct folder and that you spelled everything correctly.")
				return
			}
			fmt.Printf("\tContent of file is:\n\n%s\n", cont)
			fmt.Print("\n\tByteify? (y/n)")
			inp := read()
			if strings.ToLower(inp) == "y" {
				wfilename := name + ".bara"
				bcont := []byte(cont)
				err = ioutil.WriteFile(wfilename, bcont, 0644)
				if err != nil {
					log.Fatalf("error: %v", err)
				}
			}
		case "u":
			fmt.Print("Type the name of the file to unbyteify (no extension): ")
			name := read()
			fmt.Print("\nType the file extension: ")
			ext := read()
			buffer.Reset()
			buffer.WriteString(name)
			buffer.WriteString(".")
			buffer.WriteString(ext)
			rfilename := buffer.String()
			cont, err := ioutil.ReadFile(rfilename)
			if err != nil {
				fmt.Println("File reading error:\n\t", err, "\n\nMake sure you're in the correct folder and that you spelled everything correctly.")
				return
			}
			fmt.Print("\n\tUnbyteify? (y/n)")
			inp := read()
			if strings.ToLower(inp) == "y" {
				fmt.Print("\nEnter the name of the file to be exported: ")
				wfilename := read()
				err = ioutil.WriteFile(wfilename, []byte(cont), 0644)
				if err != nil {
					log.Fatalf("error: %v", err)
				}
			}
		}
	case "exit":
		return
	default:
		fmt.Printf("%s is not a valid mode...\n----------------------\n\n", mode)
		main()
	}
}
