package main

import (
	"fmt"
	"os"
	"strconv"
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
	fmt.Print(string(file))
}

func hex() {
	hex_num := "ab1"
	num, err := strconv.ParseInt(hex_num, 16, 64)
   
	if err != nil {
	  panic(err)
	}
   
	fmt.Println("decimal num: ", num)
   
	hex_num = "12"
	num, err = strconv.ParseInt(hex_num, 16, 64)

	if err != nil {
	  panic(err)
	}
	fmt.Println("decimal num: ", num)
   
	hex_num = "ffff"
	num, err = strconv.ParseInt(hex_num, 16, 64)
	if err != nil {
   
	  panic(err)
   
	}
	fmt.Println("decimal num: ", num)
  }
}