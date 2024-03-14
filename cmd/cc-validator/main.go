package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/samakers/cc-validator/internal/luhn"
	"github.com/samakers/cc-validator/server"
)

func main() {
	//Testing with a single credit card number via command line
	cc := os.Args[1]
	//Convert cc number to integer
	num, err := strconv.Atoi(cc)
	if err != nil {
		fmt.Println("Invalid input:", err)
		return
	}

	// Split the integer into individual digits
	numArray := server.SplitInteger(num)

	fmt.Println(numArray)
	fmt.Println(luhn.IsValidLuhn(numArray))

	//Start server
	server.Start()
}
