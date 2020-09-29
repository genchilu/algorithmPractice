package partitionLabels

import (
	"testing"

	"github.com/stretchr/testify/assert"
	//"fmt"
)

func TestPartitionLabels(t *testing.T) {
	testCases := []struct {
		s      string
		expect []int
	}{
		{
			"ababcbacadefegdehijhklij",
			[]int{9, 7, 8},
		},
		{
			"a",
			[]int{1},
		},
		{
			"",
			[]int{},
		},
	}

	for _, testCase := range testCases {
		actual := partitionLabels(testCase.s)
		assert.Equal(t, testCase.expect, actual)
	}
}
