package randomPickIndex

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPick(t *testing.T) {
	testCases := []struct{
		nums []int
		target int
		expect []int
	}{
		{
			[]int{1,2,3,3,3},
			3,
			[]int{2,3,4},
		},
	}

	for _, testCase := range testCases {
		s := Constructor(testCase.nums)
		actual := s.Pick(testCase.target)

		match := false
		for _, v := range testCase.expect {
			if v == actual {
				match = true
				break
			}
		}
		assert.Equal(t, true, match)
	}
}
