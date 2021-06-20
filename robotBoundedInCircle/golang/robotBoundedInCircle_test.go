package robotBoundedInCircle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//isRobotBounded(instructions string) bool

func TestiIRobotBounded(t *testing.T) {
	testCases := []struct {
		instructions string
		expect       bool
	}{
		{
			"GGLLGG",
			true,
		},
		{
			"GG",
			false,
		},
	}

	for _, testCase := range testCases {
		actual := isRobotBounded(testCase.instructions)
		assert.Equal(t, testCase.expect, actual)
	}
}
