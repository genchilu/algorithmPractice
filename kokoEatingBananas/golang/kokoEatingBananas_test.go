package kokoEatingBananas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func minEatingSpeed(piles []int, h int) int
func TestMinEatingSpeed(t *testing.T) {
	testCases := []struct {
		piles  []int
		h      int
		expect int
	}{
		{
			[]int{3, 6, 7, 11},
			8,
			4,
		},
		{
			[]int{30, 11, 23, 4, 20},
			5,
			30,
		},
		{
			[]int{30, 11, 23, 4, 20},
			6,
			23,
		},
		{
			[]int{1, 1, 1, 1, 100},
			5,
			100,
		},
		{
			[]int{1, 1, 1, 100, 100},
			5,
			100,
		},
	}

	for _, testCase := range testCases {
		actual := minEatingSpeed(testCase.piles, testCase.h)
		assert.Equal(t, testCase.expect, actual)
	}
}
