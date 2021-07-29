package angleBetweenHandsOfAClock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func angleClock(hour int, minutes int) float64
func TestAngleClock(t *testing.T) {
	testCases := []struct {
		hour    int
		minutes int
		expect  float64
	}{
		{
			12,
			30,
			165,
		},
		{
			3,
			30,
			75,
		},
		{
			3,
			15,
			7.5,
		},
		{
			4,
			50,
			155,
		},
		{
			12,
			0,
			0,
		},
	}

	for _, testCase := range testCases {
		actual := angleClock(testCase.hour, testCase.minutes)
		assert.Equal(t, testCase.expect, actual)
	}
}
