package designTicTacToe

type TicTacToe struct {
	players []*Player
	size    int
}

type Player struct {
	mc  map[int]int
	mr  map[int]int
	mld int
	mrd int
}

/** Initialize your data structure here. */
func Constructor(n int) TicTacToe {
	p1 := Player{
		make(map[int]int),
		make(map[int]int),
		0,
		0,
	}
	p2 := Player{
		make(map[int]int),
		make(map[int]int),
		0,
		0,
	}

	players := []*Player{&p1, &p2}
	return TicTacToe{players, n}
}

/** Player {player} makes a move at ({row}, {col}).
  @param row The row of the board.
  @param col The column of the board.
  @param player The player, can be either 1 or 2.
  @return The current winning condition, can be either:
          0: No one wins.
          1: Player 1 wins.
          2: Player 2 wins. */
func (this *TicTacToe) Move(row int, col int, player int) int {
	oppenont := this.players[(player+1)&1]
	me := this.players[(player & 1)]

	if _, ok := oppenont.mc[col]; !ok {
		if _, ok := me.mc[col]; !ok {
			me.mc[col] = 0
		}
		me.mc[col]++

		if me.mc[col] == this.size {
			return player
		}

	}

	if _, ok := oppenont.mr[row]; !ok {
		if _, ok := me.mr[row]; !ok {
			me.mr[row] = 0
		}
		me.mr[row]++

		if me.mr[row] == this.size {
			return player
		}

	}

	if row == col {
		if oppenont.mld == 0 {

			me.mld++

			if me.mld == this.size {
				return player
			}

		}
	}

	if row == this.size-col-1 {
		if oppenont.mrd == 0 {
			me.mrd++

			if me.mrd == this.size {
				return player
			}

		}
	}

	return 0
}
