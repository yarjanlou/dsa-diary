func reverseString(s []byte) {

	var size int = len(s)

	for i := 0; i < size/2; i++ {
		s[i], s[size-1-i] = s[size-1-i], s[i]
	}

}
