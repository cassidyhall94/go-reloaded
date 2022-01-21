package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("File name missing\n")
	}
	fmt.Print("\n") // print new line before results for clarity

	file = []byte(checkUpper(string(file)))
	file = []byte(checkLower(string(file)))
	file = []byte(checkCap(string(file)))

	filehex, _ := checkHex(string(file))
	file = []byte(filehex)

	filebin, _ := checkBin(string(file))
	file = []byte(filebin)

	// fileacon := []rune(string(file))
	// filetransform := acute(fileacon)
	// file = []byte(string(filetransform))

	// print file after transforming functions applied with a new line
	fmt.Print(string(file))
	fmt.Print("\n")
}
