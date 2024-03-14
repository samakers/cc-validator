package server

import (
	"net/http"
	"strconv"

	"github.com/samakers/cc-validator/internal/luhn" // replace with the actual import path
)

func ValidateHandler(w http.ResponseWriter, r *http.Request) {
	// Get the number from the request
	numStr := r.URL.Query().Get("num")

	// Convert the number to an integer
	num, err := strconv.Atoi(numStr)
	if err != nil {
		http.Error(w, "Invalid number", http.StatusBadRequest)
		return
	}

	// Split the integer into individual digits
	numArray := SplitInteger(num)

	// Check if the number is valid according to the Luhn algorithm
	if luhn.IsValidLuhn(numArray) {
		w.Write([]byte("Passed Luhn!"))
	} else {
		http.Error(w, "Failed Luhn!", http.StatusBadRequest)
	}
}

func SplitInteger(num int) []int {
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
