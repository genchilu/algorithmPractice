package minimumRemoveToMakeValidParentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
	//"fmt"
)

func TestMinRemoveToMakeValid(t *testing.T) {

	testCases := []struct {
		s      string
		expect string
	}{
		{
			"lee(t(c)o)de)",
			"lee(t(c)o)de",
		},
		{
			"a)b(c)d",
			"ab(c)d",
		},
		{
			"))((",
			"",
		},
		{
			"(a(b(c)d)",
			"a(b(c)d)",
		},
		{
			"(a(b()d)",
			"a(b()d)",
		},
		{
			"())()(((",
			"()()",
		},
	}

	for _, testCase := range testCases {
		actual := minRemoveToMakeValid(testCase.s)
		assert.Equal(t, testCase.expect, actual)
	}
}
