package kosaraju

import(
	"sort"
)


func DoKosarajuScc(edges [][]int) map[int][]int {
	if edges == nil {
		return make(map[int][]int)
	}
	g := make(map[int][]int)
	for _, edge := range(edges) {
		from := edge[0]
		to:=edge[1]
		if _,ok:=g[from];!ok{
			g[from] = []int{}
		}
		g[from] = append(g[from], to)
	}

	isVistited := make(map[int]bool)
	finishStack := []int{}
	for v,_:=range(g) {
		dfsRount1(g, v, isVistited, &finishStack)
	}

	revG := reverseG(g)
	isVistited2 := make(map[int]bool)
	componments := make(map[int][]int)

	for len(finishStack) > 0 {
		v := finishStack[len(finishStack)-1]
		finishStack = finishStack[0:len(finishStack)-1]
		componment := []int{}
		dfsRound2(revG, v, isVistited2, &componment)

		if len(componment) > 0 {
			sort.Ints(componment)
			componments[componment[0]] = componment
		}
	}
	
	return componments
}

func dfsRount1(g map[int][]int, curV int, isVistited map[int]bool, finishStack *[]int) {
	if _, ok := isVistited[curV];!ok {
		isVistited[curV] = true
		for _, toV := range(g[curV]) {
			dfsRount1(g, toV, isVistited, finishStack)
		}
		*finishStack = append(*finishStack, curV)
	}
}

func reverseG(g map[int][]int) map[int][]int {
	revG := make(map[int][]int)

	for fromV, toVs := range(g) {
		for _, v := range(toVs) {
			if _, ok := revG[v];!ok{
				revG[v] = []int{}
			}

			revG[v] = append(revG[v], fromV)
		}
	}

	return revG
}

func dfsRound2(g map[int][]int, curV int, isVistited map[int]bool, componment *[]int) {
	if _, ok := isVistited[curV];!ok {
		isVistited[curV] = true
		for _, toV := range(g[curV]) {
			dfsRound2(g, toV, isVistited, componment)
		}
		*componment = append(*componment, curV)
	}

}


