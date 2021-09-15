package bitwiseORsOfSubarrays

func subarrayBitwiseORs(arr []int) int {

	m := make(map[int]bool)
	cuur := make(map[int]bool)

	for _, v := range arr {
		next := make(map[int]bool)

		for vv, _ := range cuur {
			m[vv|v] = true
			next[vv|v] = true
		}

		m[v] = true
		next[v] = true

		cuur = next
	}

	return len(m)
}
