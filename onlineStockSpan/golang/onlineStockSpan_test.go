package onlineStockSpan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOnlineStockSpan(t *testing.T) {
	stockSpanner := Constructor()

	testCases := []struct {
		price  int
		expect int
	}{
		{
			100,
			1,
		},
		{
			80,
			1,
		},
		{
			60,
			1,
		},
		{
			70,
			2,
		},
		{
			60,
			1,
		},
		{
			75,
			4,
		},
		{
			85,
			6,
		},
	}

	for _, testCase := range testCases {
		actual := stockSpanner.Next(testCase.price)
		assert.Equal(t, testCase.expect, actual)
	}
}
