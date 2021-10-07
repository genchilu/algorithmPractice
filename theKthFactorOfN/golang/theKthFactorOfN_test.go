package theKthFactorOfN

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func kthFactor(n int, k int) int {
func TestDev(t *testing.T) {
	testCases := []struct {
		n      int
		k      int
		expect int
	}{
		{
			12,
			3,
			3,
		},
		{
			7,
			2,
			7,
		},
		{
			7,
			2,
			7,
		},
		{
			4,
			4,
			-1,
		},
		{
			1,
			1,
			1,
		},
		{
			1000,
			3,
			4,
		},
		{
			70,
			7,
			35,
		},
		{
			420,
			25,
			-1,
		},
	}

	for _, testCase := range testCases {
		actual := kthFactor(testCase.n, testCase.k)
		assert.Equal(t, testCase.expect, actual)
	}
}
