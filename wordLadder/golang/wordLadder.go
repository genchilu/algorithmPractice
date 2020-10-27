package wordLadder

type Msg struct {
	idx int
	d   int
}

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
	msg := Msg{bi, 1}
	q := []Msg{msg}
	d := len(wordList) + 1
	for len(q) > 0 {
		m := q[0]
		q = q[1:]
		//fmt.Printf("m.idx: %d, m.d:%d, wordList[m.idx]: %s\n", m.idx, m.d, wordList[m.idx])
		if wordList[m.idx] == endWord {
			if d >= m.d {
				d = m.d
			}
		} else {
			for _, vv := range g[m.idx] {
				//fmt.Printf("vv: %d, wordList[vv]: %s\n", vv, wordList[vv])
				mm := Msg{vv, m.d + 1}
				q = append(q, mm)
			}
			g[m.idx] = []int{}
		}
		if d != len(wordList)+1 {
			break
		}
	}

	if d == len(wordList)+1 {
		return 0
	} else {
		return d
	}
}

func genGraph(ws []string) [][]int {
	g := make([][]int, len(ws))

	for i := 0; i < len(ws); i++ {
		adj := []int{}
		for j := 0; j < len(ws); j++ {
			if i != j {
				if calculateDestance(ws[i], ws[j]) == 1 {
					adj = append(adj, j)
				}
			}
		}
		g[i] = adj
	}

	return g
}
func calculateDestance(w1, w2 string) int {
	d := 0
	for i := 0; i < len(w1); i++ {
		if w1[i] != w2[i] {
			d++
		}

		if d == 2 {
			break
		}
	}

	return d
}
