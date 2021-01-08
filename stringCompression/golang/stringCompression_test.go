package stringCompression

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func compress(chars []byte) int {
func TestCompress(t *testing.T) {
	testCases := []struct {
		chars  []byte
		expect int
	}{
		{
			[]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'},
			3,
		},
		// {
		// 	[]byte{'a'},
		// 	1,
		// },
		// {
		// 	[]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'},
		// 	4,
		// },
		// {
		// 	[]byte{'a', 'a', 'a', 'b', 'b', 'a', 'a'},
		// 	6,
		// },
	}

	for _, testCase := range testCases {
		actual := compress(testCase.chars)
		assert.Equal(t, testCase.expect, actual)
	}
}
