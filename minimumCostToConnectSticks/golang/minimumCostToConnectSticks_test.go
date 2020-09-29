package minimumCostToConnectSticks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectSticks(t *testing.T) {
	testCases := []struct {
		sticks []int
		expect int
	}{
		{
			[]int{2, 4, 3},
			14,
		},
		{
			[]int{1, 8, 3, 5},
			30,
		},
		{
			[]int{5},
			0,
		},
		{
			[]int{3354, 4316, 3259, 4904, 4598, 474, 3166, 6322, 8080, 9009},
			151646,
		},
	}

	for _, testCase := range testCases {
		actual := connectSticks(testCase.sticks)
		assert.Equal(t, testCase.expect, actual)
	}
}
