package wordLadder

func ladderLength(beginWord string, endWord string, wordList []string) int {
	var bi int
	for bi = 0; bi < len(wordList); bi++ {
		if wordList[bi] == beginWord {
			break
		}
	}

	if bi == len(wordList) {
		wordList = append(wordList, beginWord)
	}

	g := genGraph(wordList)
	q := []int{bi}
	d := 1
	found := false

	for len(q) > 0 {
		nq := []int{}
		for _, v := range q {
			if wordList[v] == endWord {
				found = true
				break
			} else {
				for _, vv := range g[v] {
					//fmt.Printf("vv: %d, wordList[vv]: %s\n", vv, wordList[vv])
					nq = append(nq, vv)
				}
				g[v] = []int{}
			}
		}

		if found {
			break
		}

		q = nq
		d++
	}

	if found {
		return d
	} else {
		return 0
	}

}

func genGraph(ws []string) [][]int {
	g := make([][]int, len(ws))

	for i := 0; i < len(ws); i++ {
		adj := []int{}
		for j := i + 1; j < len(ws); j++ {
			if isAdj(ws[i], ws[j]) {
				adj = append(adj, j)
				g[i] = append(g[i], j)
				g[j] = append(g[j], i)
			}
		}
	}

	return g
}

func isAdj(w1, w2 string) bool {
	diff := false
	for i := 0; i < len(w1); i++ {
		if w1[i] != w2[i] {
			if diff {
				return false
			}
			diff = true
		}
	}

	return true
}
