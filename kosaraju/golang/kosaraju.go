package kosaraju

import(
	//"fmt"
	"sort"
)


func DoKosarajuScc(edges [][]int) map[int][]int {
	g := make(map[int][]int)

	for _, edge := range edges {
		fromVertex := edge[0]
		toVertex := edge[1]

		if toVertices, found := g[fromVertex]; found {
			g[fromVertex] = append(g[fromVertex], toVertex)
		} else {
			g[fromVertex] = []int{toVertex}
		}
	}

	isVistited := []bool{}

	finishStack := []int{}

	for v, _ := range(g) {
		dfsRound1(g, v, isVistited, finishStack)
	}

	revG := reverseG(g)
	

	for i, _ := range g {
		isVistited[i] = false
	}

	componments := make(map[int][]int)
	for len(finishStack) > 0 {
		n := len(finishStack)-1
		v := finishStack[n]
		finishStack = finishStack[:n]
		componment := []int{}
		dfsRound2(revG, v, isVistited, componment)
		sort.Ints(componment)
		if(len(componment) > 0) {
			componments[componment[0]] = componment
		}
	}

	return componments
}

func dfsRound1(g map[int][]int, fromVertex int, isVistited []bool, finishStack []int) {
	if (!isVistited[fromVertex]) {
		isVistited[fromVertex] = true
		for _, toVertex := range(g[fromVertex]) {
			dfsRound1(g, toVertex, isVistited, finishStack)
		}
		finishStack = append(finishStack, fromVertex)
	}
}

func dfsRound2(g map[int][]int, fromVertex int, isVistited []bool, componments []int) {
	if (!isVistited[fromVertex]) {
		isVistited[fromVertex] = true
		for _, toVertex := range(g[fromVertex]) {
			dfsRound1(g, toVertex, isVistited, componments)
		}
		componments = append(componments, fromVertex)
	}	
}

func reverseG(g map[int][]int) map[int][]int {
	revG := make(map[int][]int)

	for v, toVertices := range g {
		for _, toVertex := range(toVertices) {
			if _, found := revG[toVertex]; found {
				revG[toVertex] = append(revG[toVertex], v)
			} else {
				revG[toVertex] = []int{v}
			}
		}
	}

	return revG
}

