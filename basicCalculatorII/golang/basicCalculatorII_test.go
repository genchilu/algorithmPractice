package basicCalculatorII

import (
	"testing"
	"github.com/stretchr/testify/assert"
	//"fmt"
)

func TestSearchRange(t *testing.T) {
	testCases := []struct{
		calculate string
		expectResult int
	} {
		{
			"3 +2*2 ",
			7,
		},
		{
			"300+ 2048*100",
			205100,
		},
		{
			"3/2 ",
			1,
		},
		{
			" 3+5 / 2 ",
			5,
		},
		{
			"1-1-1",
			-1,
		},
		{
			"2*5/3",
			3,
		},
		{
			"6/4*2",
			2,
		},
		{
			"1+2*5/3+6/4*2",
			6,
		},

	}

	for _, testCase := range testCases {
		actualResult := calculate(testCase.calculate)
		
		assert.Equal(t, testCase.expectResult, actualResult)
	}
}