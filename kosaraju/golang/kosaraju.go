package kosaraju

import(
	//"fmt"
	"sort"
)

type Graph struct {
	edges   map[int][]int
}

func (g *Graph) AddEdge(from int, to int) {
	toVertics, ok := g.edges[from]
	if !ok {
		g.edges[from] = append([]int{}, to)
	} else {
		g.edges[from] = append(toVertics, to)
	}

	_, ok = g.edges[to]
	if !ok {
		g.edges[to] = []int{}
	}
}

func DoKosarajuScc(g Graph) map[int][]int {
	visitRecord := make(map[int]bool)
	finishTimeStack := []int{}

	// do first dfs
	for vertex := range g.edges {
		dfs1st(g, vertex, visitRecord, &finishTimeStack)
	}

	g2 := reverseGraph(g)
	
	componmentMap := make(map[int][]int)
	visitRecord2rd := make(map[int]bool)
	for len(finishTimeStack) > 0 {
		n := len(finishTimeStack) -1
		currVertex := finishTimeStack[n]
		finishTimeStack = finishTimeStack[:n]
		componment := []int{}

		dfs2nd(g2, currVertex, visitRecord2rd, &componment)

		if(len(componment) > 0 ) {
			//fmt.Printf("rund2, len()>0 start vertex:%d componment: %v\n", currVertex, componment)
			sort.Ints(componment)
			componmentMap[componment[0]] = componment
		}
	}

	return componmentMap
}

func dfs1st(g Graph, currVertex int, visitRecord map[int]bool, finishTimeStack *[]int){

	isVisited, ok := visitRecord[currVertex]
	if !ok || !isVisited{
		visitRecord[currVertex] = true
		neighborVertices := g.edges[currVertex]
		for _, neighborVertex := range neighborVertices {
			dfs1st(g, neighborVertex, visitRecord, finishTimeStack)
		}

		*finishTimeStack = append(*finishTimeStack, currVertex)
	}
}


func dfs2nd(g Graph, currVertex int, visitRecord map[int]bool, componment *[]int){

	isVisited, ok := visitRecord[currVertex]
	if !ok || !isVisited{
		*componment = append(*componment, currVertex)
		visitRecord[currVertex] = true
		neighborVertices := g.edges[currVertex]
		for _, neighborVertex := range neighborVertices {
			dfs2nd(g, neighborVertex, visitRecord, componment)
		}
	}
}

func reverseGraph(g Graph) Graph {
	g2 := Graph{make(map[int][]int)}

	for vertex, neighbors := range g.edges {
		to := vertex
		for _, neighbor := range neighbors {
			from := neighbor
			g2.AddEdge(from, to)
		}
	}

	return g2
}