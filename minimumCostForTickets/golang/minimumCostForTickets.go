package minimumCostForTickets

func mincostTickets(days []int, costs []int) int {

	alldays := make([]bool, 366)
	spending := make([]int, 366)
	for _, v := range days {
		alldays[v] = true
	}

	caculCose(alldays, costs, &spending, 1)
	//fmt.Printf("%v\n", spending)
	return spending[1]
}

func caculCose(alldays []bool, costs []int, spending *[]int, n int) int {
	if n > 365 {
		return 0
	}

	if (*spending)[n] == 0 {
		if alldays[n] {
			c1 := caculCose(alldays, costs, spending, n+1) + costs[0]
			c2 := caculCose(alldays, costs, spending, n+7) + costs[1]
			c3 := caculCose(alldays, costs, spending, n+30) + costs[2]
			//fmt.Printf("%d, %d, %d, %d\n", n, c1, c2, c3)
			(*spending)[n] = min(c1, c2, c3)
		} else {
			(*spending)[n] = caculCose(alldays, costs, spending, n+1)
		}
	}

	return (*spending)[n]
}

func min(a, b, c int) int {
	min := a
	if b < min {
		min = b
	}

	if c < min {
		min = c
	}

	return min
}
