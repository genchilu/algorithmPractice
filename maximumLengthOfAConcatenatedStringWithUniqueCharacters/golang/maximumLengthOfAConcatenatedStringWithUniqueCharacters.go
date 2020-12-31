package maximumLengthOfAConcatenatedStringWithUniqueCharacters

import (
	"math/bits"
)

func maxLength(arr []string) int {
	c := []uint32{}
	for _, v := range arr {
		var mask uint32
		for i := 0; i < len(v); i++ {
			mask = mask | 1<<(v[i]-'a')
		}
		if len(v) == bits.OnesCount32(mask) {
			c = append(c, mask)
		}
	}

	if len(c) == 0 {
		return 0
	}

	max := dfs(c, 0, 0, 0)

	return max
}

func dfs(c []uint32, idx int, curmask uint32, max int) int {
	tmp := bits.OnesCount32(curmask)
	if tmp > max {
		max = tmp
	}

	for i := idx; i < len(c); i++ {
		if curmask&c[i] == 0 {
			max = dfs(c, i, curmask|c[i], max)
		}
	}

	return max
}
