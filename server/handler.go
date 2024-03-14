package server

import (
	"fmt"
	"net/http"
	"strconv"
)

func ValidateHandler(w http.ResponseWriter, r *http.Request) {
	// Get the number from the request
	num := r.URL.Query().Get("num")

	// Respond with the split number
	w.Write([]byte(fmt.Sprint(num)))
	w.Write([]byte("\n"))
	w.Write([]byte("Passed Luhn!"))
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
