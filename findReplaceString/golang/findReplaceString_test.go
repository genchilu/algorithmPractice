package findReplaceString

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindReplaceString(t *testing.T) {
	testCases := []struct {
		S       string
		indexes []int
		sources []string
		targets []string
		expect  string
	}{
		{
			"abcd",
			[]int{0, 2},
			[]string{"a", "cd"},
			[]string{"eee", "ffff"},
			"eeebffff",
		},
		{
			"abcd",
			[]int{0, 2},
			[]string{"ab", "ec"},
			[]string{"eee", "ffff"},
			"eeecd",
		},
		{
			"vmokgggqzp",
			[]int{3, 5, 1},
			[]string{"kg", "ggq", "mo"},
			[]string{"s", "so", "bfr"},
			"vbfrssozp",
		},
	}

	for _, testCase := range testCases {
		actual := findReplaceString(testCase.S, testCase.indexes, testCase.sources, testCase.targets)
		assert.Equal(t, testCase.expect, actual)
	}
}
