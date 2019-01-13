package main

import (
	"fmt"
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
	fmt.Print("Welcome to Paranoid Files\nPlease select a mode:\n e: encrypt\n d: decrypt\n ")
	mode := read()
	fmt.Print("\n")
	switch strings.ToLower(mode) {
	case "e":
		mode_e()
	case "d":
		mode_d()
	case "b": // TODO
		mode_b()
	case "exit":
		return
	default:
		fmt.Printf("%s is not a valid mode...\n----------------------\n\n", mode)
		main()
	}
}
