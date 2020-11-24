package findReplaceString

import "sort"

func findReplaceString(S string, indexes []int, sources []string, targets []string) string {

	idx := 0
	result := ""
	m := make(map[int]int)
	for i, v := range indexes {
		m[v] = i
	}
	sort.Ints(indexes)

	for i, _ := range indexes {
		s := indexes[i]
		ii := m[s]
		e := s + len(sources[ii])
		if s < len(S) && e <= len(S) {
			if S[s:e] == sources[ii] {
				head := S[idx:s]
				idx = e
				result += head + targets[ii]
			}
		}
	}
	result += S[idx:]
	return result
}
