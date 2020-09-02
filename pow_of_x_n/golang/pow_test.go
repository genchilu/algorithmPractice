package pow


import (
	"testing"
	"github.com/stretchr/testify/assert"
	"math"
	//"fmt"
)

func TestSearchRange(t *testing.T) {
	testCases := []struct{
		x float64
		n int
		expect float64
	} {
		{
			2.0,
			10,
			1024.0,
		},
		{
			2.1,
			3,
			9.261,
		},
		{
			2.0,
			-2,
			0.25,
		},
		{
			4.0,
			2,
			16.0,
		},
		{
			-1.0,
			2147483647,
			-1,
		},
	}

	for _, testCase := range testCases {
		actual := myPow(testCase.x, testCase.n)
		actual = math.Floor(actual*100000)/100000
		assert.Equal(t, testCase.expect, actual)
	}
}