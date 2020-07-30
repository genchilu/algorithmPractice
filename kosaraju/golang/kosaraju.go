package kosaraju

import(
	"sort"
)


func DoKosarajuScc(edges [][]int) map[int][]int {
	if edges == nil || len(edges) == 0 {
		return make(map[int][]int)
	}

	g := make(map[int][]int)
	for _, edge := range edges {
		from := edge[0]
		to := edge[1]
		if _, ok :=g[from]; !ok {
			g[from] = []int{}
		}
		g[from] = append(g[from], to)
	}

	isVistited1 := make(map[int]bool)
	finishStack := []int{}
	for v := range g {
		dfs1(g, v, isVistited1, &finishStack)
	}

	rg := revg(g)
	componments := make(map[int][]int)
	isVistited2 := make(map[int]bool)
	for len(finishStack) > 0 {
		v := finishStack[len(finishStack)-1]
		finishStack = finishStack[0:len(finishStack)-1]

		componment := []int{}
		dfs2(rg, v, isVistited2, &componment)

		if len(componment) > 0 {
			sort.Ints(componment)
			componments[componment[0]] = componment
		}
	}

	return componments
}

func dfs1(g map[int][]int, curVertex int, isVistited map[int]bool, finishStack *[]int) {
	if _, ok := isVistited[curVertex]; !ok {
		isVistited[curVertex] = true
		for _, v := range g[curVertex] {
			dfs1(g, v, isVistited, finishStack)
		}
		*finishStack = append(*finishStack, curVertex)
	}
}

func revg(g map[int][]int) map[int][]int {
	revg := make(map[int][]int)

	for fromV, toVs := range g {
		for _, toV := range toVs {
			if _,ok:=revg[toV];!ok {
				revg[toV] = []int{}
			}

			revg[toV] = append(revg[toV], fromV)
		}
	}

	return revg
}

func dfs2(g map[int][]int, curVertex int, isVistited map[int]bool, componment *[]int) {
	if _,ok:=isVistited[curVertex];!ok {
		isVistited[curVertex] = true
		for _, v := range g[curVertex] {
			dfs2(g, v, isVistited, componment)
		}
		*componment = append(*componment, curVertex)
	}
}




