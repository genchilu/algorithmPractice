package sellDiminishingValuedColoredBalls

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	testCases := []struct {
		inventory []int
		orders    int
		expect    int
	}{
		{
			[]int{2, 5},
			4,
			14,
		},
		{
			[]int{3, 5},
			6,
			19,
		},
		{
			[]int{2, 8, 4, 10, 6},
			20,
			110,
		},
		{
			[]int{1000000000},
			1000000000,
			21,
		},
		// {
		// 	[]int{497978859, 167261111, 483575207, 591815159},
		// 	836556809,
		// 	1234,
		// },
	}

	for _, testCase := range testCases {
		actual := maxProfit(testCase.inventory, testCase.orders)
		assert.Equal(t, testCase.expect, actual)
	}
}
