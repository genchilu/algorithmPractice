package topKFrequentElements

import "sort"

func topKFrequent(nums []int, k int) []int {
	m1 := make(map[int]int)
	m2 := make(map[int][]int)

	for _, v := range nums {
		m1[v] = m1[v] + 1
	}

	for k, v := range m1 {
		if _, ok := m2[v]; !ok {
			m2[v] = []int{}
		}
		m2[v] = append(m2[v], k)
	}

	i := 0
	counts := make([]int, len(m2))
	for k, _ := range m2 {
		counts[i] = k
		i++
	}
	sort.Ints(counts)

	result := []int{}
	for j := len(counts) - 1; j >= 0; j-- {
		if k > 0 {
			result = append(result, m2[counts[j]]...)
			k -= len(m2[counts[j]])
		} else {
			break
		}
	}

	return result
}
