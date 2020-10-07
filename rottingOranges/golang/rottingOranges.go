package rottingOranges

func orangesRotting(grid [][]int) int {
	b := [][]int{}
	fc := 0
	for i, _ := range grid {
		for j, _ := range grid[i] {
			if grid[i][j] == 2 {
				b = append(b, []int{i, j})
			} else if grid[i][j] == 1 {
				fc++
			}
		}
	}

	if fc == 0 {
		return 0
	}
	//fmt.Printf("00000 %d\n", fc)
	//fmt.Printf("%v\n", b)
	c := 0
	for len(b) > 0 {
		bb := [][]int{}
		r := false
		for _, v := range b {
			i, j := v[0], v[1]

			if grid[i][j] == 1 {
				grid[i][j] = 2
				fc--
				r = true
			}
			if i > 0 && grid[i-1][j] == 1 {
				bb = append(bb, []int{i - 1, j})
			}
			if i < len(grid)-1 && grid[i+1][j] == 1 {
				bb = append(bb, []int{i + 1, j})
			}
			if j > 0 && grid[i][j-1] == 1 {
				bb = append(bb, []int{i, j - 1})
			}
			if j < len(grid[0])-1 && grid[i][j+1] == 1 {
				bb = append(bb, []int{i, j + 1})
			}
		}

		b = bb
		if r {
			c++
		}
		//fmt.Printf("1111 %v\n", b)
	}

	//fmt.Printf("1111 %d\n", fc)
	if fc == 0 {
		return c
	}

	return -1
}
