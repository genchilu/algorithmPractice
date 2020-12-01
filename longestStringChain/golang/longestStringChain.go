package longestStringChain

func longestStrChain(words []string) int {
	lengthMap := make(map[int][]string)
	predecessorMap := make(map[string][]string)

	for _, w := range words {
		if _, ok := lengthMap[len(w)]; !ok {
			lengthMap[len(w)] = []string{}
		}
		lengthMap[len(w)] = append(lengthMap[len(w)], w)
	}

	//fmt.Printf("%v\n", lengthMap)
	for _, w1 := range words {
		if w2s, ok := lengthMap[len(w1)+1]; ok {
			for _, w2 := range w2s {
				if isPredecessor(w1, w2) {
					if _, ok := predecessorMap[w1]; !ok {
						predecessorMap[w1] = []string{}
					}
					predecessorMap[w1] = append(predecessorMap[w1], w2)
				}
			}
		}
	}

	//fmt.Printf("%v\n", predecessorMap)

	max := 0
	isvistited := make(map[string]bool)
	for _, w := range words {
		if _, ok := isvistited[w]; !ok {
			d := dfs(w, predecessorMap, isvistited)
			if d > max {
				max = d
			}
		}
	}
	return max
}

func dfs(w string, predecessorMap map[string][]string, isvistited map[string]bool) int {
	isvistited[w] = true
	if predecessors, ok := predecessorMap[w]; !ok {
		return 1
	} else {
		max := 0
		for _, predecessor := range predecessors {
			d := dfs(predecessor, predecessorMap, isvistited)
			if d > max {
				max = d
			}
		}

		return max + 1
	}
}
func isPredecessor(w1, w2 string) bool {
	for i := 0; i < len(w1); i++ {
		if w1[i] != w2[i] {
			if w1[i:] == w2[i+1:] {
				return true
			} else {
				return false
			}
		}
	}

	return true
}
