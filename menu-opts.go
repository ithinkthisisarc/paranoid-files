package main

import (
	//"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"github.com/fatih/color"
)

func mode_e() {
	fmt.Print("Selected encrypt mode\n\n")
	fmt.Print("Enter the name of the file to be encrypted: ")
	rfilename := read()
	data, err := ioutil.ReadFile(rfilename)
	if err != nil {
		fmt.Println("File reading error:\n\t", err, "\n\nMake sure you're in the correct folder and that you spelled everything correctly.")
		return
	}
	fmt.Print("Enter the name of the file to be exported: ")
	wfilename := read()
	fmt.Println("Contents of file:", string(data))
	fmt.Print("\n\tEncrypt? (Y/N): ")
	inp := read()
	if strings.HasPrefix(strings.ToLower(inp), "y") {
		cont := string(data)
		encryption(cont, wfilename)
	}
	color.Blue("\nExported to file: %s\n", wfilename)
}

func mode_d() {
	fmt.Println("Selected decrypt mode\n")
	fmt.Print("Enter the file to decrypt: ")
	rfilename := read()
	fmt.Print("Enter the name of the file you want exported: ")
	wfilename := read()
	_, err := ioutil.ReadFile(rfilename)
	if err != nil {
		fmt.Println("File reading error:\n\t", err, "\n\nMake sure you're in the correct folder and that you spelled everything correctly.")
		return
	}
	fmt.Print("\n\tDecrypt? (Y/N): ")
	inp := read()
	if strings.HasPrefix(strings.ToLower(inp), "y") {
		err := ioutil.WriteFile(wfilename, []byte(decryption(rfilename)), 0644)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		color.Blue("\nExported to file: %s\n", wfilename)
	}
}

func mode_u() {
	fmt.Println("Selected unmix mode\n")
	fmt.Print("Enter file to unmix: ")
	rfilename := read()
	fmt.Print("Enter the name of the file you want exported: ")
	wfilename := read()
	_, err := ioutil.ReadFile(rfilename)
	if err != nil {
		fmt.Println("File reading error:\n\t", err, "\n\nMake sure you're in the correct folder and that you spelled everything correctly.")
		return
	}
	fmt.Print("\n\tUnmix? [Y/n]: ")
	inp := read()
	if strings.HasPrefix(strings.ToLower(inp), "y") {
		err := ioutil.WriteFile(wfilename, []byte(unmix(rfilename)), 0644)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		color.Blue("\nExported to file: %s\n", wfilename)
	}
}

func mode_m() {
	fmt.Println("Selected mix mode\n")
	fmt.Print("Enter the file to mix: ")
	rfilename := read()
	fmt.Print("Enter the name of the file to be exported: ")
	wfilename := read()
	data, err := ioutil.ReadFile(rfilename)
	if err != nil {
		fmt.Printf("File reading error:\n\t ===> %v", err)
		return
	}
	fmt.Print("Mix? [Y/n]: ")
	inp := read()
	if strings.HasPrefix(strings.ToLower(inp), "y") {
		mix(string(data), wfilename)
	}
}