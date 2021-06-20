package robotBoundedInCircle

func isRobotBounded(instructions string) bool {
	direct := 'N'
	postion := []int{0, 0}

	for i := range instructions {
		switch instructions[i] {
		case 'G':
			switch direct {
			case 'N':
				postion[0]++
			case 'E':
				postion[1]++
			case 'S':
				postion[0]--
			case 'W':
				postion[1]--
			}
		case 'L':
			switch direct {
			case 'N':
				direct = 'E'
			case 'E':
				direct = 'S'
			case 'S':
				direct = 'W'
			case 'W':
				direct = 'N'
			}
		case 'R':
			switch direct {
			case 'N':
				direct = 'W'
			case 'E':
				direct = 'N'
			case 'S':
				direct = 'E'
			case 'W':
				direct = 'S'
			}
		}
	}

	return (postion[0] == 0 && postion[1] == 0) || direct != 'N'
}
