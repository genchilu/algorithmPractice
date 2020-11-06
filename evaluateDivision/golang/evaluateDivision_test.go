package evaluateDivision

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcEquation(t *testing.T) {

	testCases := []struct {
		equations [][]string
		values    []float64
		queries   [][]string
		expect    []float64
	}{
		{
			[][]string{
				[]string{"a", "b"},
				[]string{"b", "c"},
			},
			[]float64{2.0, 3.0},
			[][]string{
				[]string{"a", "c"},
				[]string{"b", "a"},
				[]string{"a", "e"},
				[]string{"a", "a"},
				[]string{"x", "x"},
			},
			[]float64{6.0, 0.5, -1.0, 1.0, -1.0},
		},
		{
			[][]string{
				[]string{"a", "b"},
				[]string{"b", "c"},
				[]string{"bc", "cd"},
			},
			[]float64{1.5, 2.5, 5.0},
			[][]string{
				[]string{"a", "c"},
				[]string{"c", "b"},
				[]string{"bc", "cd"},
				[]string{"cd", "bc"},
			},
			[]float64{3.75, 0.4, 5.0, 0.2},
		},
		{
			[][]string{
				[]string{"a", "b"},
			},
			[]float64{0.5},
			[][]string{
				[]string{"a", "b"},
				[]string{"b", "a"},
				[]string{"a", "c"},
				[]string{"x", "y"},
			},
			[]float64{0.5, 2.0, -1.0, -1.0},
		},
	}

	for _, testCase := range testCases {
		actual := calcEquation(testCase.equations, testCase.values, testCase.queries)
		assert.Equal(t, testCase.expect, actual)

	}
}
