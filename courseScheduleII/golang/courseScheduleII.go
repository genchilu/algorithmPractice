package courseScheduleII

func findOrder(numCourses int, prerequisites [][]int) []int {
	m := make(map[int][]int)

	for _, v := range prerequisites {
		if _, ok := m[v[0]]; !ok {
			m[v[0]] = []int{}
		}

		m[v[0]] = append(m[v[0]], v[1])
	}

	res := []int{}
	vistited := make(map[int]bool)

	valid := true

	for i := 0; i < numCourses; i++ {
		mm := make(map[int]bool)
		if !dfs(m, i, &res, vistited, mm) {
			valid = false
			break
		}
	}

	if valid {
		return res
	}

	return []int{}
}

func dfs(m map[int][]int, v int, res *[]int, vistited map[int]bool, curvistited map[int]bool) bool {
	if _, ok := curvistited[v]; ok {
		return false
	}
	curvistited[v] = true

	valid := true
	if _, ok := m[v]; ok {

		for _, vv := range m[v] {
			if !dfs(m, vv, res, vistited, curvistited) {
				valid = false
				break
			}
		}
	}

	if _, ok := vistited[v]; !ok {
		(*res) = append(*res, v)
	}
	vistited[v] = true

	delete(curvistited, v)

	return valid

}
