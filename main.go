package main

import (
	"fmt"
	"log"
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

	filehex, err := checkHex(string(file))
	if err != nil {
		log.Fatal(err)
	}
	file = []byte(filehex)

	filebin, err := checkBin(string(file))
	if err != nil {
		log.Fatal(err)
	}
	file = []byte(filebin)

	file = []byte(checkA(string(file)))

	file = []byte(checkPunc(string(file)))

	file = []byte(checkApos(string(file)))

	// print file after transforming functions applied with a new line
	// fmt.Print(string(file))
	// fmt.Print("\n")

	err = os.WriteFile("result.txt", file, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
