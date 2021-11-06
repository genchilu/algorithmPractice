package singleNumberIII

func singleNumber(nums []int) []int {

	m := make(map[int]bool)

	for _, v := range nums {
		if _, ok := m[v]; ok {
			delete(m, v)
		} else {
			m[v] = true
		}
	}

	result := []int{}
	for k, _ := range m {
		result = append(result, k)
	}
	return result
}
