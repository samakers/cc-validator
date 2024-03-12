package main

import (
	"fmt"

	"github.com/samakers/cc-validator/internal/luhn"
)

func main() {
	//Test CC Number
	//4417 1234 5678 9113
	fmt.Println(luhn.IsValidLuhn([]int{4, 4, 1, 7, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 1, 3}))
}
