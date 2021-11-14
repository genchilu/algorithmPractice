package numberOfWonderfulSubstrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func wonderfulSubstrings(word string) int64 {
func TestDev1(t *testing.T) {
	word := "aba"
	expect := int64(4)

	actual := wonderfulSubstrings(word)

	assert.Equal(t, expect, actual)
}

func TestDev2(t *testing.T) {
	word := "aabb"
	expect := int64(9)

	actual := wonderfulSubstrings(word)

	assert.Equal(t, expect, actual)
}

func TestDev3(t *testing.T) {
	word := "he"
	expect := int64(2)

	actual := wonderfulSubstrings(word)

	assert.Equal(t, expect, actual)
}
