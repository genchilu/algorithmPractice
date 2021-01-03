package reverseWords

import (
	"strings"
)

func reverseWords(s string) string {
	tmp := strings.Fields(s)
	for i, j := 0, len(tmp)-1; i < j; i, j = i+1, j-1 {
		tmp[i], tmp[j] = tmp[j], tmp[i]
	}

	return strings.Join(tmp, " ")
}
