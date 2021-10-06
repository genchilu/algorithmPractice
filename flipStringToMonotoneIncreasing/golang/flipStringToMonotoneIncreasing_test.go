package flipStringToMonotoneIncreasing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func minFlipsMonoIncr(s string) int {
func TestDev(t *testing.T) {
	testCases := []struct {
		s      string
		expect int
	}{
		// {
		// 	"00110",
		// 	1,
		// },
		// {
		// 	"010110",
		// 	2,
		// },
		// {
		// 	"00011000",
		// 	2,
		// },
		{
			"11011",
			1,
		},
	}

	for _, testCase := range testCases {
		actual := minFlipsMonoIncr(testCase.s)
		assert.Equal(t, testCase.expect, actual)
	}
}
