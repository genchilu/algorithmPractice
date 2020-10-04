package groupAnagrams

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupAnagrams(t *testing.T) {
	testCases := []struct {
		strs   []string
		expect [][]string
	}{
		{
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{
				[]string{"bat"},
				[]string{"nat", "tan"},
				[]string{"ate", "eat", "tea"},
			},
		},
	}

	for _, testCase := range testCases {
		actual := groupAnagrams(testCase.strs)
		assert.Equal(t, len(testCase.expect), len(actual))

		c := 0
		for _, v := range actual {
			macth := false
			for _, vv := range testCase.expect {
				if len(v) == len(vv) {
					sort.Strings(v)
					sort.Strings(vv)

					for i, _ := range v {
						if v[i] != vv[i] {
							break
						}
					}
				}
				macth = true
			}

			if macth {
				c++
			}
		}

		assert.Equal(t, c, len(testCase.expect))
	}
}
