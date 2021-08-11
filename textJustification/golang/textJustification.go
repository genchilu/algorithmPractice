package textJustification

import (
	"fmt"
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	result := []string{}
	tmp := []string{}
	count := 0
	for i := range words {
		if count+len(words[i]) > maxWidth {

			result = append(result, format(tmp, maxWidth))

			tmp = []string{words[i]}
			count = len(words[i]) + 1
		} else {
			tmp = append(tmp, words[i])
			count += len(words[i]) + 1
		}
	}

	result = append(result, fmt.Sprintf("%-*s", maxWidth, strings.Join(tmp, " ")))

	return result
}

func format(strs []string, maxWidth int) string {

	if len(strs) == 1 {
		strs[0] = fmt.Sprintf("%-*s", maxWidth, strs[0])
	} else {
		for _, str := range strs {
			maxWidth -= len(str)
		}

		latest := len(strs) - 1
		pad := maxWidth / latest
		r := (maxWidth % latest)
		for i := 0; i < latest; i++ {
			if r > 0 {
				strs[i] = fmt.Sprintf("%-*s", pad+len(strs[i])+1, strs[i])
				r--
			} else {
				strs[i] = fmt.Sprintf("%-*s", pad+len(strs[i]), strs[i])
				//fmt.Printf("%s, %d, %v, %d, %d\n", strs[i], pad, more, maxWidth, latest)
			}
		}

	}
	return strings.Join(strs, "")
}
