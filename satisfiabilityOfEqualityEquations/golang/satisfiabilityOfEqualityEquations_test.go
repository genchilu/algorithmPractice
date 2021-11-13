package satisfiabilityOfEqualityEquations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func equationsPossible(equations []string) bool {
func TestDev1(t *testing.T) {
	equations := []string{"a==b", "b!=a"}
	expect := false

	actual := equationsPossible(equations)

	assert.Equal(t, expect, actual)
}

func TestDev2(t *testing.T) {
	equations := []string{"b==a", "a==b"}
	expect := true

	actual := equationsPossible(equations)

	assert.Equal(t, expect, actual)
}

func TestDev3(t *testing.T) {
	equations := []string{"a==b", "b==c", "a==c"}
	expect := true

	actual := equationsPossible(equations)

	assert.Equal(t, expect, actual)
}

func TestDev4(t *testing.T) {
	equations := []string{"a==b", "b!=c", "c==a"}
	expect := false

	actual := equationsPossible(equations)

	assert.Equal(t, expect, actual)
}

func TestDev5(t *testing.T) {
	equations := []string{"c==c", "b==d", "x!=z"}
	expect := true

	actual := equationsPossible(equations)

	assert.Equal(t, expect, actual)
}

func TestDev6(t *testing.T) {
	equations := []string{"c==c", "f!=a", "f==b", "b==c"}
	expect := true

	actual := equationsPossible(equations)

	assert.Equal(t, expect, actual)
}

func TestDev7(t *testing.T) {
	equations := []string{"e==d", "e==a", "f!=d", "b!=c", "a==b"}
	expect := true

	actual := equationsPossible(equations)

	assert.Equal(t, expect, actual)
}
