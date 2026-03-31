package main

import "fmt"

func isPerfectSquare(num int) bool {

	if num == 1 {
		return true
	}

	var start int = 1
	var end int = num
	var mid int = num / 2

	for start < mid {
		if mid * mid == num {
			return true
		}

		if mid * mid > num {
			end = mid
		}

		if mid * mid < num {
			start = mid
		}

		mid = (start + end) / 2
	}
	
	return  false
}