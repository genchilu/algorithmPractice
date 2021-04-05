package jumpGame

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func canJump(nums []int) bool {
func TestCanJump(t *testing.T) {
	testCases := []struct {
		nums   []int
		expect bool
	}{
		{
			[]int{2, 3, 1, 1, 4},
			true,
		},
		{
			[]int{3, 2, 1, 0, 4},
			false,
		},
		{
			[]int{0},
			true,
		},
		{
			[]int{2, 1, 2, 2, 1, 2, 2, 2},
			true,
		},
	}

	for _, testCase := range testCases {
		actual := canJump(testCase.nums)
		assert.Equal(t, testCase.expect, actual)
	}
}
