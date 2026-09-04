func singleNumber(nums []int) int {
    var counter map[int]int = make(map[int]int, len(nums))

	for _, num := range nums {
		counter[num]++
	}

	for k, v := range counter {
		if v == 1 {
			return k
		}
	}

	return 0
}



