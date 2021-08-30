package reverseWordsInStringII

func reverseWords(s []byte) {
	r := []byte{}

	start := 0
	for i := range s {
		if s[i] == ' ' {
			tmp := make([]byte, i-start)
			copy(tmp, s[start:i])
			r = append(tmp, r...)
			r = append([]byte{' '}, r...)

			// fmt.Printf("%s\n", string(tmp))
			start = i + 1
		}
	}

	tmp := make([]byte, len(s)-start)
	copy(tmp, s[start:])
	// fmt.Printf("%s\n", string(tmp))

	r = append(tmp, r...)
	// fmt.Printf("%s\n", string(s))
	// fmt.Printf("%s\n", string(r))

	// fmt.Printf("9999 %d, %d\n", len(r), len(s))

	for i := range r {
		s[i] = r[i]
	}
	// fmt.Printf("1111 %s\n", string(s))
}
