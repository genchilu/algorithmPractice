package groupAnagrams

import "sort"

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)

	for _, v := range strs {
		bs := []byte(v)

		sort.SliceStable(bs, func(i, j int) bool { return bs[i] < bs[j] })

		if _, ok := m[string(bs)]; !ok {
			m[string(bs)] = []string{}
		}
		m[string(bs)] = append(m[string(bs)], v)
	}

	result := [][]string{}

	for _, v := range m {
		result = append(result, v)
	}
	return result
}
