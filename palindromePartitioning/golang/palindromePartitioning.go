package palindromePartitioning

import "fmt"

func partition(s string) [][]string {
	return dfs(s, 0)
}

func dfs(s string, start int) [][]string {
	cur_result := [][]string{}
	if isPalindrome(s[start:]) {
		cur_result = append(cur_result, []string{s[start:]})
	}

	for end := start + 1; end < len(s); end++ {
		fmt.Printf("%d, %d\n", start, end)
		if isPalindrome(s[start:end]) {
			sub_result := dfs(s, end)
			for _, ss := range sub_result {
				ss = append([]string{s[start:end]}, ss...)
				cur_result = append(cur_result, ss)
			}
		}
	}

	return cur_result
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}

	return true
}
