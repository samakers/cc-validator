package luhn

func IsValidLuhn(digits []int) bool {
	checkDigit := 0
	for i := len(digits) - 2; i >= 0; i-- {
		if i%2 == 0 {
			doubled := digits[i] * 2
			if doubled > 9 {
				doubled -= 9
			}
			checkDigit += doubled
		} else {
			checkDigit += digits[i]
		}
	}

	return (10-(checkDigit%10))%10 == digits[len(digits)-1]
}
