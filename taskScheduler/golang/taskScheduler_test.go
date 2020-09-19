package taskScheduler

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLeastInterval(t *testing.T) {
	testCases := []struct{
		tasks []byte
		n int
		expect int
	}{
		{
			[]byte{'A','A','A','B','B','B'},
			2,
			8,
		},
		{
			[]byte{'A','A','A','B','B','B'},
			0,
			6,
		},
		{
			[]byte{'A','A','A','A','A','A','B','C','D','E','F','G'},
			2,
			16,
		},
	}

	for _, testCase:=range testCases {
		actual := leastInterval(testCase.tasks, testCase.n)
		assert.Equal(t, testCase.expect, actual)
	}
}