package main

func removeDuplicatesSorted(nums []int) int {

	var j int = 0

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[j] {
			nums[j+1] = nums[i]
			j++
		}
	}

	return j + 1

}
