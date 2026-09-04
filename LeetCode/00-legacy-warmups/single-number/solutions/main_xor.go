package main

func singleNumber(nums []int) int {
	var result int = 0

	for _, i := range nums {
		result ^= i
	}

	return result
}