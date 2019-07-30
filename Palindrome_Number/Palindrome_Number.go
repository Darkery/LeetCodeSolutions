package Palindrome_Number

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	temp := 0
	quotient := x

	for {
		remain := quotient % 10

		temp = temp * 10 + remain

		quotient = quotient / 10
		if quotient == 0 {
			break
		}
	}

	if x == temp {
		return true
	}

	return false
}
