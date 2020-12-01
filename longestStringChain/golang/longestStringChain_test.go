package longestStringChain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestStrChain(t *testing.T) {
	testCases := []struct {
		words  []string
		expect int
	}{
		{
			[]string{"a", "b", "ba", "bca", "bda", "bdca"},
			4,
		},
		{
			[]string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"},
			5,
		},
	}

	for _, testCase := range testCases {
		actual := longestStrChain(testCase.words)
		assert.Equal(t, testCase.expect, actual)
	}
}
