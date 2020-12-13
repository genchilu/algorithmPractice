package reorganizeString

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//reorganizeString(S string) string {
func TestReorganizeString(t *testing.T) {
	testCases := []struct {
		S      string
		expect string
	}{
		{
			"",
			"",
		},
	}

	for _, testCase := range testCases {
		actual := reorganizeString(testCase.S)
		assert.Equal(t, testCase.expect, actual)
	}
}
