package zigzagConversion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func convert(s string, numRows int) string {
func TestDev1(t *testing.T) {
	s := "PAYPALISHIRING"
	numRows := 3

	expect := "PAHNAPLSIIGYIR"

	actual := convert(s, numRows)

	assert.Equal(t, expect, actual)
}

func TestDev2(t *testing.T) {
	s := "PAYPALISHIRING"
	numRows := 4

	expect := "PINALSIGYAHRPI"

	actual := convert(s, numRows)

	assert.Equal(t, expect, actual)
}

func TestDev3(t *testing.T) {
	s := "A"
	numRows := 1

	expect := "A"

	actual := convert(s, numRows)

	assert.Equal(t, expect, actual)
}

func TestDev4(t *testing.T) {
	s := "AB"
	numRows := 1

	expect := "AB"

	actual := convert(s, numRows)

	assert.Equal(t, expect, actual)
}

func TestDev5(t *testing.T) {
	s := "ABCD"
	numRows := 2

	expect := "ACBD"

	actual := convert(s, numRows)

	assert.Equal(t, expect, actual)
}
