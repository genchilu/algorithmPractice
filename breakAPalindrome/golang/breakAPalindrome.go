package breakAPalindrome

func breakPalindrome(palindrome string) string {
	if len(palindrome) == 1 {
		return ""
	} else {
		for i := 0; i < len(palindrome); i++ {
			if palindrome[i] != 'a' {
				if len(palindrome)%2 == 1 && i == len(palindrome)/2 {
					return palindrome[0:len(palindrome)-1] + "b"
				} else {
					return palindrome[0:i] + "a" + palindrome[i+1:]
				}
			}
		}

		return palindrome[0:len(palindrome)-1] + "b"
	}
}
