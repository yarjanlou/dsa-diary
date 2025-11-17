func isPalindrome(x int) bool {

	var number int = x
	var reverseNumber int = 0

	if x < 0 {
		return false
	}

	for x > 0 {

		reverseNumber = reverseNumber*10 + (x % 10)
		x /= 10
	}

	return reverseNumber == number

}
