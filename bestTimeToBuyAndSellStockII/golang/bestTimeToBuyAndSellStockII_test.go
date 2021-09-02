package bestTimeToBuyAndSellStockII

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDev(t *testing.T) {
	testCases := []struct {
		prices []int
		expect int
	}{
		{
			[]int{7, 1, 5, 3, 6, 4},
			7,
		},
		{
			[]int{1, 2, 3, 4, 5},
			4,
		},
		{
			[]int{7, 6, 4, 3, 1},
			0,
		},
	}

	for _, testCase := range testCases {
		actual := maxProfit(testCase.prices)
		assert.Equal(t, testCase.expect, actual)
	}
}
