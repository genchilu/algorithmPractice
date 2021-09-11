package simplifyPath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDev(t *testing.T) {
	testCases := []struct {
		path   string
		expect string
	}{
		{
			"///",
			"/",
		},
		{
			"/home/",
			"/home",
		},
		{
			"/../",
			"/",
		},
		{
			"/home//foo/",
			"/home/foo",
		},
		{
			"/a/./b/../../c/",
			"/c",
		},
	}

	for _, testCase := range testCases {
		actual := simplifyPath(testCase.path)
		assert.Equal(t, testCase.expect, actual)
	}
}
