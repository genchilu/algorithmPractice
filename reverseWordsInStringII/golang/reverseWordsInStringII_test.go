package reverseWordsInStringII

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDev(t *testing.T) {
	testCases := []struct {
		s      []byte
		expect []byte
	}{
		{
			[]byte{'t', 'h', 'e', ' ', 's', 'k', 'y', ' ', 'i', 's', ' ', 'b', 'l', 'u', 'e'},
			[]byte{'b', 'l', 'u', 'e', ' ', 'i', 's', ' ', 's', 'k', 'y', ' ', 't', 'h', 'e'},
		},
		{
			[]byte{'a'},
			[]byte{'a'},
		},
	}

	for _, testCase := range testCases {
		reverseWords(testCase.s)
		fmt.Printf("%s\n", string(testCase.s))
		assert.Equal(t, testCase.expect, testCase.s)
	}
}
