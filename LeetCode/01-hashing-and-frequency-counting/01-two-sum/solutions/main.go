func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i, num := range nums {
		if index, exists := m[target-num]; exists {
			return []int{i, index}
		}
		m[num] = i
	}

	return nil
}
