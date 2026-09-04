func firstUniqChar(s string) int {
	var directory map[rune]int = make(map[rune]int, len(s))

	for _, char := range s {
		directory[char]++
	}

	for index, char := range s {
		if directory[char] == 1 {
			return index
		}
	}

	return -1
}