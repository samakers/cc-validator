package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/samakers/cc-validator/internal/luhn"
)

type RequestPayload struct {
	Num string `json:"num"`
}

type ResponsePayload struct {
	IsValid bool `json:"is_valid"`
}

func ValidateHandler(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON payload
	var p RequestPayload
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Convert the number to an integer
	num, err := strconv.Atoi(p.Num)
	if err != nil {
		http.Error(w, "Invalid number", http.StatusBadRequest)
		return
	}

	// Split the integer into individual digits
	numArray := SplitInteger(num)

	// Check if the number is valid according to the Luhn algorithm
	isValid := luhn.IsValidLuhn(numArray)

	// Create the response payload
	resp := ResponsePayload{
		IsValid: isValid,
	}

	// Set the Content-Type header
	w.Header().Set("Content-Type", "application/json")

	// Return the payload back to the user
	json.NewEncoder(w).Encode(resp)
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
