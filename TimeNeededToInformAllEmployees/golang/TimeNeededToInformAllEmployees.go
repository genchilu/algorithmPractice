package TimeNeededToInformAllEmployees

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	m := make(map[int][]int)
	for i := 0; i < n; i++ {
		mgr := manager[i]
		if mgr != -1 {
			if _, ok := m[mgr]; !ok {
				m[mgr] = []int{}
			}
			m[mgr] = append(m[mgr], i)
		}
	}

	return dfs(headID, 0, m, informTime)
}

func dfs(id, t int, m map[int][]int, informTime []int) int {

	if _, ok := m[id]; !ok {
		return t
	} else {
		t = t + informTime[id]
		max := t

		for _, e := range m[id] {
			tmp := dfs(e, t, m, informTime)
			if tmp > max {
				max = tmp
			}
		}

		return max
	}

}
