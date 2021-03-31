package reorderRoutesToMakeAllPathsLeadToTheCityZero

func minReorder(n int, connections [][]int) int {

	conn, rev_conn := make([][]int, n), make([][]int, n)

	for _, v := range connections {
		conn[v[0]] = append(conn[v[0]], v[1])
		rev_conn[v[1]] = append(rev_conn[v[1]], v[0])
	}

	return dfs(0, -1, conn, rev_conn)
}

func dfs(n, from int, conn, rev_conn [][]int) int {
	c := 0
	for _, v := range conn[n] {
		if v != from {
			c += dfs(v, n, conn, rev_conn) + 1
		}
	}

	for _, v := range rev_conn[n] {
		if v != from {
			c += dfs(v, n, conn, rev_conn)
		}
	}

	return c
}
