func isPerfectSquare(num int) bool {
	if num < 2 {
		return true
	}

	x := num

	for x > num/x {
		x = (x + num/x) / 2
	}

	return x*x == num
}