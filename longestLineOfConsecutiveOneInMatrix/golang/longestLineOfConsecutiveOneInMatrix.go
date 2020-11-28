package longestLineOfConsecutiveOneInMatrix

func longestLine(M [][]int) int {
	max, m, n := 0, 0, 0
	m = len(M)
	if m > 0 {
		n = len(M[0])
	}

	for i := 0; i < m; i++ {
		s := 0
		for j := 0; j < n; j++ {
			if M[i][j] == 1 {
				s++
			} else {
				if s > max {
					max = s
				}
				s = 0
			}
		}
		if s > max {
			max = s
		}
	}

	for j := 0; j < n; j++ {
		s := 0
		for i := 0; i < m; i++ {
			if M[i][j] == 1 {
				s++
			} else {
				if s > max {
					max = s
				}
				s = 0
			}
		}
		if s > max {
			max = s
		}
	}

	for j := 0; j < n; j++ {
		s := 0
		for k, l := 0, j; k < m && l < n; {
			if M[k][l] == 1 {
				s++
			} else {
				if s > max {
					max = s
				}
				s = 0
			}
			k++
			l++
		}
		if s > max {
			max = s
		}
	}
	for i := 1; i < m; i++ {
		s := 0
		for k, l := i, 0; k < m && l < n; {
			if M[k][l] == 1 {
				s++
			} else {
				if s > max {
					max = s
				}
				s = 0
			}
			k++
			l++
		}
		if s > max {
			max = s
		}
	}

	for j := n - 1; j >= 0; j-- {
		s := 0
		for k, l := 0, j; k < m && l >= 0; {
			if M[k][l] == 1 {
				s++
			} else {
				if s > max {
					max = s
				}
				s = 0
			}
			k++
			l--
		}
		if s > max {
			max = s
		}
	}
	for i := 1; i < m; i++ {
		s := 0
		for k, l := i, n-1; k < m && l >= 0; {
			if M[k][l] == 1 {
				s++
			} else {
				if s > max {
					max = s
				}
				s = 0
			}
			k++
			l--
			if s > max {
				max = s
			}
		}
	}

	return max
}
