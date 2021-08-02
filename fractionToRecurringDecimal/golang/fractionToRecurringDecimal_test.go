package fractionToRecurringDecimal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func fractionToDecimal(numerator int, denominator int) string {
func TestFractionToDecimal(t *testing.T) {
	testCases := []struct {
		numerator   int
		denominator int
		expect      string
	}{
		{
			1,
			2,
			"0.5",
		},
		{
			1,
			4,
			"0.25",
		},
		{
			2,
			1,
			"2",
		},
		{
			3,
			2,
			"1.5",
		},
		{
			9,
			8,
			"1.125",
		},
		{
			1,
			3,
			"0.(3)",
		},
		{
			4,
			333,
			"0.(012)",
		},
		{
			1,
			17,
			"0.(0588235294117647)",
		},
		{
			-50,
			8,
			"-6.25",
		},
		{
			-22,
			-2,
			"11",
		},
		{
			-2147483648,
			1,
			"-2147483648",
		},
		{
			0,
			-5,
			"0",
		},
	}

	for _, testCase := range testCases {
		actual := fractionToDecimal(testCase.numerator, testCase.denominator)
		assert.Equal(t, testCase.expect, actual)
	}
}
