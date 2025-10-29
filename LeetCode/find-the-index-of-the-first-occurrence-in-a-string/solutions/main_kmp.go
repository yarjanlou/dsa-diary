package main

func buildLPS(pattern string) []int {
	lps := make([]int, len(pattern))

	var length int = 0
	var i int = 1

	for i < len(pattern) {
		if pattern[i] == pattern[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length > 0 {
				length = lps[length-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}

	return lps
}

func strStrKMP(heystask string, needle string) int {
	if len(needle) == 0 {
		return -1
	}

	// Build LPS (Longest Prefix Suffix) array

	var lps []int = buildLPS(needle)

	var i, j int = 0, 0

	for i < len(heystask) {
		if heystask[i] == needle[j] {
			i++
			j++
			if j == len(needle) {
				return i - j
			}
		} else {
			if j > 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}
	return -1
}
