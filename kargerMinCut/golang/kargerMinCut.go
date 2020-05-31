package kargerMinCut

import (
	"math/rand"
	"time"
)

func doMinCut(edges [][]int) int {
	if (edges == nil || len(edges) == 0) {
		return 0;
	}

	minCut := len(edges)
	vertextCount := getVertexCount(edges)

	for i:=0;i<vertextCount;i++{
		for j:=0;j<vertextCount;j++{
			curMinCut := doKargerMinCut(edges, vertextCount)
			if (curMinCut < minCut) {
				minCut = curMinCut
			}
		}
	}
	
	return minCut;
}

func doKargerMinCut(edges [][]int, vertextCount int) int {
	rand.Seed(time.Now().UnixNano())
	
	tmpEdges := make([][]int, len(edges));
	copy(edges, tmpEdges);

	vertexMap := initVertexMap(tmpEdges);
	newVertex := findMaxVertex(vertexMap) + 1

	mergeCount := 0
	for mergeCount < vertextCount -2 {
		x := rand.Intn(len(tmpEdges))
		pickedEdge := tmpEdges[x]
		v1 := pickedEdge[0]
		v2 := pickedEdge[1]

		// remove edge
		tmpEdges[x] = tmpEdges[len(tmpEdges)-1]
		tmpEdges = tmpEdges[:len(tmpEdges)-1]

		if (v1 != v2) {
			vertexMap[v1] = newVertex
			vertexMap[v2] = newVertex
			newVertex++
			mergeCount++
		}
	}

	cutCount := 0
	for _, edge := range tmpEdges {
		if (edge[0]!=edge[1]) {
			cutCount++
		}
	}

	return cutCount;
}

func initVertexMap(edges [][]int) map[int]int {

	vertextMap := make(map[int]int)
	for _, v := range edges {
		vertextMap[v[0]] = v[0]
		vertextMap[v[1]] = v[1]
	}

	return vertextMap;
}

func findMaxVertex(vertexMergeMap map[int]int) int {
	max := 0
	for key := range vertexMergeMap {
		if key > max {
			max = key
		}
	}

	return max;
}

func getVertexCount(edges [][]int) int {
	vertexMap := make(map[int]bool)
	for _, v := range edges {
		vertexMap[v[0]] = true
		vertexMap[v[1]] = true
	}

	return len(vertexMap);
}
