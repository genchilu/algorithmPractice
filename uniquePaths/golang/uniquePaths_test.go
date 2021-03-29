package uniquePaths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func uniquePaths(m int, n int) int {
func TestUniquePaths(t *testing.T) {
	testCases := []struct {
		m      int
		n      int
		expect int
	}{
		// {
		// 	7,
		// 	3,
		// 	28,
		// },
		// {
		// 	3,
		// 	3,
		// 	6,
		// },
		{
			1,
			10,
			1,
		},
	}

	for _, testCase := range testCases {
		actual := uniquePaths(testCase.m, testCase.n)
		assert.Equal(t, testCase.expect, actual)
	}
}
