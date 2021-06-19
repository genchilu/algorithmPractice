package theMaze

func hasPath(maze [][]int, start []int, destination []int) bool {
	vistited := make([][]bool, len(maze))
	for i := range maze {
		vistited[i] = make([]bool, len(maze[i]))
	}

	return findPath(maze, start, destination, &vistited)
}

func findPath(maze [][]int, start []int, destination []int, vistited *[][]bool) bool {
	if start[0] == destination[0] && start[1] == destination[1] {
		return true
	} else if (*vistited)[start[0]][start[1]] {
		return false
	} else {

		(*vistited)[start[0]][start[1]] = true
		l := findLeft(maze, start)
		r := findRight(maze, start)
		u := findUp(maze, start)
		d := findDown(maze, start)

		return (l != start[1] && findPath(maze, []int{start[0], l}, destination, vistited)) ||
			(r != start[1] && findPath(maze, []int{start[0], r}, destination, vistited)) ||
			(u != start[0] && findPath(maze, []int{u, start[1]}, destination, vistited)) ||
			(d != start[0] && findPath(maze, []int{d, start[1]}, destination, vistited))
	}
}

func findLeft(maze [][]int, cur []int) int {

	l := cur[1]
	for ; l >= 0; l-- {
		if maze[cur[0]][l] == 1 {
			break
		}
	}

	l++
	return l
}

func findRight(maze [][]int, cur []int) int {

	r := cur[1]
	for ; r < len(maze[cur[0]]); r++ {
		if maze[cur[0]][r] == 1 {
			break
		}
	}

	r--
	return r
}

func findUp(maze [][]int, cur []int) int {

	u := cur[0]
	for ; u >= 0; u-- {
		if maze[u][cur[1]] == 1 {
			break
		}
	}

	u++
	return u
}

func findDown(maze [][]int, cur []int) int {

	d := cur[0]
	for ; d < len(maze); d++ {
		if maze[d][cur[1]] == 1 {
			break
		}
	}

	d--
	return d
}
