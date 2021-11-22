package minesweeper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func updateBoard(board [][]byte, click []int) [][]byte {

func TestDev1(t *testing.T) {
	board := [][]byte{
		{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'M', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'},
	}

	click := []int{3, 0}
	expect := [][]byte{
		{'B', '1', 'E', '1', 'B'},
		{'B', '1', 'M', '1', 'B'},
		{'B', '1', '1', '1', 'B'},
		{'B', 'B', 'B', 'B', 'B'},
	}

	actual := updateBoard(board, click)

	assert.Equal(t, expect, actual)
}

func TestDev2(t *testing.T) {
	board := [][]byte{
		{'B', '1', 'E', '1', 'B'},
		{'B', '1', 'M', '1', 'B'},
		{'B', '1', '1', '1', 'B'},
		{'B', 'B', 'B', 'B', 'B'},
	}

	click := []int{1, 2}
	expect := [][]byte{
		{'B', '1', 'E', '1', 'B'},
		{'B', '1', 'X', '1', 'B'},
		{'B', '1', '1', '1', 'B'},
		{'B', 'B', 'B', 'B', 'B'},
	}

	actual := updateBoard(board, click)

	assert.Equal(t, expect, actual)
}
