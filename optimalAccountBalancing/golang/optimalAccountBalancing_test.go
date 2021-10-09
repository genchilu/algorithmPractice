package optimalAccountBalancing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func minTransfers(transactions [][]int) int {

func TestDev(t *testing.T) {
	testCases := []struct {
		transactions [][]int
		expect       int
	}{
		{
			[][]int{
				{0, 1, 10},
				{1, 0, 1},
				{1, 2, 5},
				{2, 0, 5},
			},
			1,
		},
		{
			[][]int{
				{0, 1, 10},
				{2, 0, 5},
			},
			2,
		},
		{
			[][]int{
				{1, 5, 8},
				{8, 9, 8},
				{2, 3, 9},
				{4, 3, 1},
			},
			4,
		},
		{
			[][]int{
				{1, 2, 3},
				{1, 3, 3},
				{6, 4, 1},
				{5, 4, 4},
			},
			4,
		},
	}

	for _, testCase := range testCases {
		actual := minTransfers(testCase.transactions)
		assert.Equal(t, testCase.expect, actual)
	}
}
