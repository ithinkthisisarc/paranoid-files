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
	switch {
	case strings.HasPrefix(strings.ToLower(mode), "e"):
		mode_e()
		
	case strings.HasPrefix(strings.ToLower(mode), "d"):
		mode_d()

	case strings.HasPrefix(mode, "exit"):
		return
	default:
		fmt.Printf("%s is not a valid mode...\n----------------------\n\n", mode)
		main()
	}
}
