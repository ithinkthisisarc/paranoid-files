package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func mode_e(target string, export string) {
	data, err := ioutil.ReadFile(target)
	if err != nil {
		fmt.Println("File reading error:\n\t", err, "\n\nMake sure you're in the correct folder and that you spelled everything correctly.")
		return
	}
	fmt.Println("Contents of file:", string(data))
	fmt.Print("\n\tEncrypt? (y/N): ")
	inp := read()
	if strings.HasPrefix(strings.ToLower(inp), "y") {
		cont := string(data)
		encryption(cont, export)
	}
}

func mode_d(target string, export string) {
	_, err := ioutil.ReadFile(target)
	if err != nil {
		fmt.Println("File reading error:\n\t", err, "\n\nMake sure you're in the correct folder and that you spelled everything correctly.")
		return
	}
	fmt.Print("\n\tDecrypt? (y/N): ")
	inp := read()
	if strings.HasPrefix(strings.ToLower(inp), "y") {
		err := ioutil.WriteFile(export, []byte(decryption(target)), 0644)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		fmt.Printf("\nExported to file: %s", export)
	}
}

func mode_b() { // TODO
	var buffer bytes.Buffer

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
		fmt.Print("\n\tByteify? (y/N)")
		inp := read()
		if strings.HasPrefix(strings.ToLower(inp), "y") {
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
		fmt.Print("\n\tUnbyteify? (y/N)")
		inp := read()
		if strings.HasPrefix(strings.ToLower(inp), "y") {
			fmt.Print("\nEnter the name of the file to be exported: ")
			wfilename := read()
			err = ioutil.WriteFile(wfilename, []byte(cont), 0644)
			if err != nil {
				log.Fatalf("error: %v", err)
			}
		}
	}
}
