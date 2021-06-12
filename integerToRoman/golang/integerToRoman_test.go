package integerToRoman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntToRoman(t *testing.T) {
	testCases := []struct {
		num    int
		expect string
	}{
		// {
		// 	3,
		// 	"III",
		// },
		// {
		// 	4,
		// 	"IV",
		// },
		// {
		// 	9,
		// 	"IX",
		// },
		// {
		// 	58,
		// 	"LVIII",
		// },
		// {
		// 	1994,
		// 	"MCMXCIV",
		// },
		{
			41,
			"XLI",
		},
	}

	for _, testCase := range testCases {
		actual := intToRoman(testCase.num)
		assert.Equal(t, testCase.expect, actual)
	}
}
