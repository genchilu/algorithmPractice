package designSnakeGame

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestSnakeGame1(t *testing.T) {
// 	snakeGame := Constructor(
// 		3,
// 		2,
// 		[][]int{
// 			[]int{1, 2},
// 			[]int{0, 1},
// 		},
// 	)

// 	actual := snakeGame.Move("R")
// 	assert.Equal(t, 0, actual)

// 	actual = snakeGame.Move("D")
// 	assert.Equal(t, 0, actual)

// 	actual = snakeGame.Move("R")
// 	assert.Equal(t, 1, actual)

// 	actual = snakeGame.Move("U")
// 	assert.Equal(t, 1, actual)

// 	actual = snakeGame.Move("L")
// 	assert.Equal(t, 2, actual)

// 	actual = snakeGame.Move("U")
// 	assert.Equal(t, -1, actual)
// }

func TestSnakeGame2(t *testing.T) {
	snakeGame := Constructor(
		3,
		3,
		[][]int{
			[]int{2, 0},
			[]int{0, 0},
			[]int{0, 2},
			[]int{2, 2},
		},
	)

	actual := snakeGame.Move("D")
	assert.Equal(t, 0, actual)

	actual = snakeGame.Move("D")
	assert.Equal(t, 1, actual)

	actual = snakeGame.Move("R")
	assert.Equal(t, 1, actual)

	actual = snakeGame.Move("U")
	assert.Equal(t, 1, actual)

	actual = snakeGame.Move("U")
	assert.Equal(t, 1, actual)

	actual = snakeGame.Move("L")
	assert.Equal(t, 2, actual)

	actual = snakeGame.Move("D")
	assert.Equal(t, 2, actual)

	actual = snakeGame.Move("R")
	assert.Equal(t, 2, actual)

	actual = snakeGame.Move("R")
	assert.Equal(t, 2, actual)

	actual = snakeGame.Move("U")
	assert.Equal(t, 3, actual)

	actual = snakeGame.Move("L")
	assert.Equal(t, 3, actual)

	actual = snakeGame.Move("D")
	assert.Equal(t, 3, actual)
}
