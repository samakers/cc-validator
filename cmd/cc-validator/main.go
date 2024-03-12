package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/samakers/cc-validator/internal/luhn"
)

func main() {

	cc := os.Args[1]
	num, err := strconv.Atoi(cc)
	if err != nil {
		fmt.Println("Invalid input:", err)
		return
	}

	// Split the integer into individual digits
	numArray := splitInteger(num)

	fmt.Println(numArray)
	fmt.Println(luhn.IsValidLuhn(numArray))
}

func splitInteger(num int) []int {
	// Convert the integer to a string
	numStr := strconv.Itoa(num)

	// Create an array to store the digits
	numArray := make([]int, len(numStr))

	// Split the string into individual characters and convert them to integers
	for i, digit := range numStr {
		numArray[i], _ = strconv.Atoi(string(digit))
	}

	return numArray
}
