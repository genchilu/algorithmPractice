package reorganizeString

import "sort"

type ByteCount struct {
	b byte
	c int
}

func reorganizeString(S string) string {
	m := make(map[byte]int)
	for i := 0; i < len(S); i++ {
		if _, ok := m[S[i]]; !ok {
			m[S[i]] = 0
		}
		m[S[i]]++
	}

	byteCounts := []*ByteCount{}
	for k, v := range m {
		byteCounts = append(byteCounts, &ByteCount{k, v})
	}

	sort.Slice(byteCounts, func(i, j int) bool { return byteCounts[j].c < byteCounts[i].c })

	if (byteCounts[0].c*2 - 1) > len(S) {
		return ""
	}

	b := make([]byte, len(S))
	c := 0
	byteCount := byteCounts[c]
	for i := 0; i < len(S); i += 2 {
		b[i] = byteCount.b
		byteCount.c--
		if byteCount.c == 0 {
			c++
			byteCount = byteCounts[c]
		}
	}

	for i := 1; i < len(S); i += 2 {
		b[i] = byteCount.b
		byteCount.c--
		if byteCount.c == 0 && c < (len(byteCounts)-1) {
			c++
			byteCount = byteCounts[c]
		}
	}

	return string(b)
}
