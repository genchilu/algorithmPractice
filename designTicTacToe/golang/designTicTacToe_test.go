package designTicTacToe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTicTacToe3(t *testing.T) {
	ticTacToe := Constructor(3)

	testCases := []struct {
		row    int
		col    int
		player int
		expect int
	}{
		{
			0,
			0,
			1,
			0,
		},
		{
			0,
			2,
			2,
			0,
		},
		{
			2,
			2,
			1,
			0,
		},
		{
			1,
			1,
			2,
			0,
		},
		{
			2,
			0,
			1,
			0,
		},
		{
			1,
			0,
			2,
			0,
		},
		{
			2,
			1,
			1,
			1,
		},
	}

	for _, testCase := range testCases {
		actual := ticTacToe.Move(testCase.row, testCase.col, testCase.player)
		assert.Equal(t, testCase.expect, actual)
	}
}

func TestTicTacToe2_1(t *testing.T) {
	ticTacToe := Constructor(2)

	testCases := []struct {
		row    int
		col    int
		player int
		expect int
	}{
		{
			0,
			0,
			2,
			0,
		},
		{
			1,
			1,
			1,
			0,
		},
		{
			0,
			1,
			2,
			2,
		},
	}

	for _, testCase := range testCases {
		actual := ticTacToe.Move(testCase.row, testCase.col, testCase.player)
		assert.Equal(t, testCase.expect, actual)
	}
}

func TestTicTacToe2_2(t *testing.T) {
	ticTacToe := Constructor(2)

	testCases := []struct {
		row    int
		col    int
		player int
		expect int
	}{
		{
			0,
			0,
			2,
			0,
		},
		{
			0,
			1,
			1,
			0,
		},
		{
			1,
			1,
			2,
			2,
		},
	}

	for _, testCase := range testCases {
		actual := ticTacToe.Move(testCase.row, testCase.col, testCase.player)
		assert.Equal(t, testCase.expect, actual)
	}
}
