package findAllAnagramsInString

import (
	"reflect"
)

func findAnagrams(s string, p string) []int {
	if len(s) == 0 {
		return []int{}
	}

	if len(p) > len(s) {
		return []int{}
	}

	pCharCountMap := make(map[byte]uint)
	sWindowCharCountMap := make(map[byte]uint)

	for i:=0;i<len(p);i++{
		if _, ok := pCharCountMap[p[i]];!ok {
			pCharCountMap[p[i]] = 0
		}
		pCharCountMap[p[i]]++

		if _, ok := sWindowCharCountMap[s[i]];!ok {
			sWindowCharCountMap[s[i]] = 0
		}
		sWindowCharCountMap[s[i]]++
	}

	result := []int{}
	if reflect.DeepEqual(pCharCountMap, sWindowCharCountMap) {
		result = append(result, 0)
	}

	for i:=1;i<=len(s)-len(p);i++ {
		sWindowCharCountMap[s[i-1]]--
		if sWindowCharCountMap[s[i-1]] == 0 {
			delete(sWindowCharCountMap, s[i-1])
		}

		if _,ok:=sWindowCharCountMap[s[i+len(p)-1]];!ok{
			sWindowCharCountMap[s[i+len(p)-1]] = 0
		}
		sWindowCharCountMap[s[i+len(p)-1]]++

		if reflect.DeepEqual(pCharCountMap, sWindowCharCountMap) {
			result = append(result, i)
		}
	}

	return result
}
