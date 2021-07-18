package encodeAndDecodeStrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeDecode(t *testing.T) {
	testCases := []struct {
		input []string
	}{
		{
			[]string{},
		},
		{
			[]string{""},
		},
		{
			[]string{"Hello", "World"},
		},
		{
			[]string{"", ""},
		},
		{
			[]string{"C#", "&", "~Xp|F", "R4QBf9g=_"},
		},
	}

	for _, testCase := range testCases {
		codec := Codec{}
		encode := codec.Encode(testCase.input)
		actual := codec.Decode(encode)

		assert.Equal(t, testCase.input, actual)
	}
}
