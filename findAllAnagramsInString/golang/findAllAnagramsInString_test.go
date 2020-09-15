package findAllAnagramsInString

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFindAnagrams(t *testing.T){
	testCases := []struct{
		s string
		p string
		expect []int
	}{
		{
			"cbaebabacd",
			"cba",
			[]int{0, 6},
		},
		{
			"abab",
			"ab",
			[]int{0,1,2},
		},
	}

	for _, testCase := range testCases {
		actual := findAnagrams(testCase.s, testCase.p)
		assert.Equal(t, testCase.expect, actual)
		//assert.Equal(t, 0, 1)
	}
	

}