package goreloaded

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 3 {
		fmt.Printf("Too many arguments\n")
		return
	} else if len(os.Args) < 3 {
		fmt.Printf("File name missing\n")
		return
	}
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("File name missing\n")
	}
	file = []byte(checkUpper(string(file)))
	fmt.Print(string(file))
}
