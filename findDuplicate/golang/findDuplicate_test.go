package findDuplicate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func findDuplicate(nums []int) int {
func TestFindDuplicate(t *testing.T) {
	testCases := []struct {
		nums   []int
		expect int
	}{
		{
			[]int{1, 3, 4, 2, 2},
			2,
		},
	}

	for _, testCase := range testCases {
		actual := findDuplicate(testCase.nums)
		assert.Equal(t, testCase.expect, actual)
	}
}
