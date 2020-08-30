package divideTwoIntegers

import (
	"testing"
	"github.com/stretchr/testify/assert"
	//"fmt"
)

func TestDivide(t *testing.T) {

	testCases := []struct{
		dividend int
		divisor int
		expect int
	} {
		// {
		// 	0,
		// 	1,
		// 	0,
		// },
		// {
		// 	10,
		// 	3,
		// 	3,
		// },
		// {
		// 	7,
		// 	-3,
		// 	-2,
		// },
		// {
		// 	-7,
		// 	3,
		// 	-2,
		// },
		// {
		// 	-7,
		// 	-3,
		// 	2,
		// },
		// {
		// 	1,
		// 	1,
		// 	1,
		// },
		// {
		// 	10,
		// 	2,
		// 	5,
		// },
		// {
		// 	2147483648,
		// 	1,
		// 	2147483647,
		// },
		// {
		// 	-2147483648,
		// 	-1,
		// 	2147483647,
		// },
		// {
		// 	-2147483648,
		// 	1,
		// 	-2147483648,
		// },
		// {
		// 	2147483648,
		// 	-1,
		// 	-2147483648,
		// },
		// {
		// 	-102626058,
		// 	-2147483648,
		// 	0,
		// },
		{
			2147483647,
			2,
			1073741823,
		},
	}

	for _, testCase := range testCases {
		actual := divide(testCase.dividend, testCase.divisor)
		assert.Equal(t, testCase.expect, actual)
	}
}