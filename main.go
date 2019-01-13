package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var svar string
	flag.StringVar(&svar, "mode", "e", "Select a mode. Valid modes are:\n E: Encrypt\n D: Decrypt")

	var svar2 string
	flag.StringVar(&svar2, "target", "tester.txt", "Select the file to be encrypted")

	var svar3 string
	flag.StringVar(&svar3, "export", "tester.para", "Select the name of the exported file")
	flag.Parse()
	mode := svar
	target := svar2
	export := svar3
	fmt.Print("\n")
	switch strings.ToLower(mode) {
	case "e":
		mode_e(target, export)
	case "d":
		mode_d(target, export)
	case "b": // TODO
		mode_b()
	default:
		fmt.Printf("%s is not a valid mode...\n----------------------\n\n", mode)
		main()
	}
}
