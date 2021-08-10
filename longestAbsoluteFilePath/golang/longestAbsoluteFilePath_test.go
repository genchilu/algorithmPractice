package longestAbsoluteFilePath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDev(t *testing.T) {
	testCases := []struct {
		input  string
		expect int
	}{
		{
			"dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext",
			20,
		},
		{
			"dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext",
			32,
		},
		{
			"a",
			0,
		},
		{
			"file1.txt\nfile2.txt\nlongfile.txt",
			12,
		},
		{
			"a\n\taa\n\t\taaa\n\t\t\tfile1.txt\naaaaaaaaaaaaaaaaaaaaa\n\tsth.png",
			29,
		},
		{
			"a\n\taa\n\t\taaa\n\t\t\tfile1234567890123.txt\naaaaaaaaaaaaaaaaaaaaa\n\tsth.png",
			30,
		},
		{
			"rzzmf\nv\n\tix\n\t\tiklav\n\t\t\ttqse\n\t\t\t\ttppzf\n\t\t\t\t\tzav\n\t\t\t\t\t\tkktei\n\t\t\t\t\t\t\thhmav\n\t\t\t\t\t\t\t\tbzvwf.txt",
			47,
		},
	}

	for _, testCase := range testCases {
		actual := lengthLongestPath(testCase.input)
		assert.Equal(t, testCase.expect, actual)
	}
}
