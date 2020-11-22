package expressiveWords

import "fmt"

type Count struct {
	ch byte
	c  int
}

func expressiveWords(S string, words []string) int {
	c := convertCounts(S)

	fmt.Printf("%v\n", c)
	sum := 0
	for _, word := range words {
		cc := convertCounts(word)
		if len(cc) == len(c) {
			for i, _ := range c {
				if c[i].ch != cc[i].ch {
					break
				} else {
					if c[i].c < cc[i].c {
						break
					} else if c[i].c > cc[i].c && c[i].c < 3 {
						break
					}
				}

				if i == len(c)-1 {
					fmt.Printf("%s\n", word)
					sum++
				}
			}
		}
	}
	return sum
}

func convertCounts(s string) []Count {
	c := []Count{}
	for i := 0; i < len(s); i++ {
		if len(c) > 0 && c[len(c)-1].ch == s[i] {
			c[len(c)-1].c++
		} else {
			c = append(c, Count{s[i], 1})
		}
	}

	return c
}
