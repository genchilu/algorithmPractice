package removeComments

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveComments(t *testing.T) {

	testCases := []struct {
		source []string
		expect []string
	}{
		{
			[]string{"/*Test program */", "int main()", "{ ", "  // variable declaration ", "int a, b, c;", "/* This is a test", "   multiline  ", "   comment for ", "   testing */", "a = b + c;", "}"},
			[]string{"int main()", "{ ", "  ", "int a, b, c;", "a = b + c;", "}"},
		},
		{
			[]string{"a/*comment", "line", "more_comment*/b"},
			[]string{"ab"},
		},
		{
			[]string{"struct Node{", "    /*/ declare members;/**/", "    int size;", "    /**/int val;", "};"},
			[]string{"struct Node{", "    ", "    int size;", "    int val;", "};"},
		},
		{
			[]string{"void func(int k) {", "// this function does nothing /*", "   k = k*2/4;", "   k = k/2;*/", "}"},
			[]string{"void func(int k) {", "   k = k*2/4;", "   k = k/2;*/", "}"},
		},
	}

	for _, testCase := range testCases {
		actual := removeComments(testCase.source)
		assert.Equal(t, testCase.expect, actual)
	}
}
