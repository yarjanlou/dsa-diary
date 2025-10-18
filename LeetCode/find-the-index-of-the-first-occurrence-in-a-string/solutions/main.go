package main

func strStr(haystack string, needle string) int16 {
	var i int16 = 0
	for ; i <= int16(len(haystack)-len(needle)); i++ {
		if haystack[i:i+int16(len(needle))] == needle {
			return i
		}
	}
	return -1
}
