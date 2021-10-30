package palindromicSubstrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//countSubstrings(s string)
func TestDev1(t *testing.T) {
	s := "abc"
	expect := 3

	actual := countSubstrings(s)

	assert.Equal(t, expect, actual)
}

func TestDev2(t *testing.T) {
	s := "aaa"
	expect := 6

	actual := countSubstrings(s)

	assert.Equal(t, expect, actual)
}
