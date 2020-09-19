package exclusiveTimeOfFunctions

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestExclusiveTime(t *testing.T) {

	testCases := []struct{
		n int
		logs []string
		expect []int
	} {
		{
			1,
			[]string{"0:start:0", "0:end:2"},
			[]int{3},
		},
		{
			2,
			[]string{"0:start:0","1:start:2","1:end:5","0:end:6"},
			[]int{3,4},
		},
		{
			1,
			[]string{"0:start:0","0:start:2","0:end:5","0:start:6","0:end:6","0:end:7"},
			[]int{8},
		},
		{
			2,
			[]string{"0:start:0","0:start:2","0:end:5","1:start:6","1:end:6","0:end:7"},
			[]int{7,1},
		},
		{
			2,
			[]string{"0:start:0","0:start:2","0:end:5","1:start:7","1:end:7","0:end:8"},
			[]int{8,1},
		},
	}

	for _, testCase := range testCases {
		actual := exclusiveTime(testCase.n, testCase.logs)
		assert.Equal(t, testCase.expect, actual)
	}
}