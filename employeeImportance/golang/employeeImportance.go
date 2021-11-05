package employeeImportance

import "sort"

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {

	sort.SliceStable(employees, func(a, b int) bool {

		return employees[a].Id < employees[b].Id
	})

	return dfs(employees, id)
}

func dfs(employees []*Employee, id int) int {
	idx := findIdx(employees, id)
	if idx == -1 {
		return 0
	}
	val := employees[idx].Importance
	for _, subIdx := range employees[idx].Subordinates {
		val += dfs(employees, subIdx)
	}

	return val
}

func findIdx(employees []*Employee, id int) int {

	l, r := 0, len(employees)-1
	for l <= r {
		m := (l + r) / 2
		if id == employees[m].Id {
			return m
		} else if id < employees[m].Id {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return -1
}
