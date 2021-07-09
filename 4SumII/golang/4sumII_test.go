package fourSumCount

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int
func TestFourSumCount(t *testing.T) {

	testCases := []struct {
		nums1  []int
		nums2  []int
		nums3  []int
		nums4  []int
		expect int
	}{
		{
			[]int{1, 2},
			[]int{-2, -1},
			[]int{-1, 2},
			[]int{0, 2},
			2,
		},
	}

	for _, testCase := range testCases {
		actual := fourSumCount(testCase.nums1, testCase.nums2, testCase.nums3, testCase.nums4)
		assert.Equal(t, testCase.expect, actual)
	}
}
