package pairsOfSongsWithTotalDurationsDivisibleBy60

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func numPairsDivisibleBy60(time []int) int {

func TestNumPairsDivisibleBy60(t *testing.T) {
	testCases := []struct {
		time   []int
		expect int
	}{
		{
			[]int{30, 20, 150, 100, 40},
			3,
		},
		{
			[]int{60, 60, 60},
			3,
		},
	}

	for _, testCase := range testCases {
		actual := numPairsDivisibleBy60(testCase.time)
		assert.Equal(t, testCase.expect, actual)
	}
}
