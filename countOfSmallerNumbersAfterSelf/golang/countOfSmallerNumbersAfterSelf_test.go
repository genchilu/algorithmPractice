package countOfSmallerNumbersAfterSelf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDev(t *testing.T) {
	testCases := []struct {
		nums   []int
		expect []int
	}{
		{
			[]int{-1},
			[]int{0},
		},
		{
			[]int{-1, -1},
			[]int{0, 0},
		},
		{
			[]int{5, 2, 6, 1},
			[]int{2, 1, 1, 0},
		},
	}

	for _, testCase := range testCases {
		actual := countSmaller(testCase.nums)
		assert.Equal(t, testCase.expect, actual)
	}
}

func TestDev2(t *testing.T) {
	s := []int{1, 2, 3}

	e := []int{1, 2, 4, 3}

	assert.Equal(t, e, insert(s, 2, 4))
}
