package myCalendarI

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTmp(t *testing.T) {
	a := []int{10, 20}
	b := []int{15, 25}

	actual := isConflict(a, b)
	assert.Equal(t, false, actual)
}
