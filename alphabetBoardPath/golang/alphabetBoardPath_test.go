package alphabetBoardPath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlphabetBoardPath(t *testing.T) {
	testCases := []struct {
		target string
		expect string
	}{
		{
			"leet",
			"RDD!UURRR!!DDD!",
		},
		{
			"code",
			"RR!RRDD!UUL!R!",
		},
	}

	for _, testCase := range testCases {
		actual := alphabetBoardPath(testCase.target)
		assert.Equal(t, testCase.expect, actual)
	}
}
