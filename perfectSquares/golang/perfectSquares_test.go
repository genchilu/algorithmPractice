package perfectSquares

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func numSquares(n int) int {

func TestDev(t *testing.T) {
	testCases := []struct {
		n      int
		expect int
	}{
		{
			12,
			3,
		},
		{
			13,
			2,
		},
	}

	for _, testCase := range testCases {
		actual := numSquares(testCase.n)
		assert.Equal(t, testCase.expect, actual)
	}
}
