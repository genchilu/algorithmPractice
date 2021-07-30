package guessNumberHigherOrLowerII

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func getMoneyAmount(n int) int
func TestGetMoneyAmount(t *testing.T) {
	testCases := []struct {
		n      int
		expect int
	}{
		{
			1,
			0,
		},
		{
			2,
			1,
		},
		{
			3,
			2,
		},
		{
			4,
			4,
		},
		{
			10,
			16,
		},
	}

	for _, testCase := range testCases {
		actual := getMoneyAmount(testCase.n)
		assert.Equal(t, testCase.expect, actual)
	}
}
