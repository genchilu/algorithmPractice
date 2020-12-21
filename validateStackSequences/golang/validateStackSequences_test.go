package validateStackSequences

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func validateStackSequences(pushed []int, popped []int) bool {

func TestValidateStackSequences(t *testing.T) {
	testCases := []struct {
		pushed []int
		popped []int
		expect bool
	}{
		{
			[]int{1, 2, 3, 4, 5},
			[]int{4, 5, 3, 2, 1},
			true,
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{4, 3, 5, 1, 2},
			false,
		},
	}

	for _, testCase := range testCases {
		actual := validateStackSequences(testCase.pushed, testCase.popped)
		assert.Equal(t, testCase.expect, actual)

	}
}
