package designSnakeGame

type SnakeGame struct {
	w     int
	h     int
	food  *[][]int
	snake *[][]int
	m     map[int]map[int]bool
}

/** Initialize your data structure here.
  @param width - screen width
  @param height - screen height
  @param food - A list of food positions
  E.g food = [[1,1], [1,0]] means the first food is positioned at [1,1], the second is at [1,0]. */
func Constructor(width int, height int, food [][]int) SnakeGame {
	snake := [][]int{[]int{0, 0}}
	m := make(map[int]map[int]bool)
	m[0] = make(map[int]bool)
	m[0][0] = true
	return SnakeGame{width, height, &food, &snake, m}
}

/** Moves the snake.
  @param direction - 'U' = Up, 'L' = Left, 'R' = Right, 'D' = Down
  @return The game's score after the move. Return -1 if game over.
  Game over when snake crosses the screen boundary or bites its body. */
func (this *SnakeGame) Move(direction string) int {
	i, j := this.nextMove(direction)
	if !this.isSave(i, j) {
		return -1
	}

	(*this.snake) = append([][]int{[]int{i, j}}, (*this.snake)...)
	if _, ok := this.m[i]; !ok {
		this.m[i] = make(map[int]bool)
	}

	if !this.eatFood() {
		tail := (*this.snake)[len(*this.snake)-1]
		(*this.snake) = (*this.snake)[:len(*this.snake)-1]
		ti, tj := tail[0], tail[1]
		delete(this.m[ti], tj)
	}
	if this.bitesBody(i, j) {
		return -1
	}

	this.m[i][j] = true

	return len(*this.snake) - 1
}

func (this *SnakeGame) isSave(i, j int) bool {
	if i < 0 || j < 0 || i == this.h || j == this.w {
		return false
	}

	return true
}

func (this *SnakeGame) bitesBody(i, j int) bool {
	if _, ok1 := this.m[i]; ok1 {
		if _, ok2 := this.m[i][j]; ok2 {
			return true
		}
	}
	return false
}

func (this *SnakeGame) nextMove(direction string) (i, j int) {
	i, j = (*this.snake)[0][0], (*this.snake)[0][1]
	switch direction {
	case "U":
		i--
	case "L":
		j--
	case "R":
		j++
	case "D":
		i++
	}

	return i, j
}

func (this *SnakeGame) eatFood() bool {
	if len(*this.food) > 0 {
		if (*this.snake)[0][0] == (*this.food)[0][0] && (*this.snake)[0][1] == (*this.food)[0][1] {
			(*this.food) = (*this.food)[1:]
			return true
		}
	}

	return false
}
