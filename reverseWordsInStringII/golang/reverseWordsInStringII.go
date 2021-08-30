package reverseWordsInStringII

func reverseWords(s []byte) {
	reverse(s)

	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			reverse(s[start:i])
			start = i + 1
		}
	}

	reverse(s[start:])

}

func reverse(s []byte) {

	l := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		s[i], s[l-i] = s[l-i], s[i]
	}

}
