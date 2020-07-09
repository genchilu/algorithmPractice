package kargerMinCut

import (
	"math/rand"
	"time"
)

func doMinCut(edges [][]int) int {

	verticesCount := getVerticesCount(edges)

	minCut := len(edges)
	for i:=0; i<verticesCount*verticesCount; i++ {
		tmpMinCut := kargerMinCut(edges)
		if tmpMinCut < minCut {
			minCut = tmpMinCut
		}
	}

	return minCut
}

func kargerMinCut(edges [][]int) int {
	cloneEdges := make([][]int, len(edges))
	copy(cloneEdges, edges)

	vertexMergeIdxMap, mergeIdxVertexMap, maxVertexId := initVertexMergeIdMap(cloneEdges)
	mergeCount := 0
	
	for mergeCount < (len(vertexMergeIdxMap) -2 ) {
		rand.Seed(time.Now().UnixNano())
		pickIdx := rand.Intn(len(cloneEdges))
		pickEdge := cloneEdges[pickIdx]

		mergeIdx0 := vertexMergeIdxMap[pickEdge[0]]
		mergeIdx1 := vertexMergeIdxMap[pickEdge[1]]
		if mergeIdx0 != mergeIdx1 {
			mergeCount++
			newMergeIdx := maxVertexId + mergeCount
			for _, v:=range(mergeIdxVertexMap[mergeIdx0]) {
				vertexMergeIdxMap[v] = newMergeIdx
			}

			for _, v:=range(mergeIdxVertexMap[mergeIdx1]) {
				vertexMergeIdxMap[v] = newMergeIdx
			}

			mergeIdxVertexMap[newMergeIdx] = append(mergeIdxVertexMap[mergeIdx0], mergeIdxVertexMap[mergeIdx1]...)

			delete(mergeIdxVertexMap, mergeIdx0)
			delete(mergeIdxVertexMap, mergeIdx1)
		}

		cloneEdges = append(cloneEdges[0:pickIdx], cloneEdges[pickIdx+1:]...)
	}

	minCut := 0
	for _, edge := range(cloneEdges) {
		if vertexMergeIdxMap[edge[0]] != vertexMergeIdxMap[edge[1]] {
			minCut++
		}
	}

	return minCut
}

func initVertexMergeIdMap(edges [][]int) (map[int]int, map[int][]int, int) {

	vertexMergeIdxMap := make(map[int]int)
	mergeIdxVertexMap := make(map[int][]int)
	maxVertexId := 0

	for _, edge := range(edges) {
		if _, ok := vertexMergeIdxMap[edge[0]]; !ok{
			vertexMergeIdxMap[edge[0]] = edge[0]
		}
		if _, ok := vertexMergeIdxMap[edge[1]]; !ok{
			vertexMergeIdxMap[edge[1]] = edge[1]
		}

		if _, ok := mergeIdxVertexMap[edge[0]]; !ok{
			mergeIdxVertexMap[edge[0]] = []int{edge[0]}
		}
		if _, ok := mergeIdxVertexMap[edge[1]]; !ok{
			mergeIdxVertexMap[edge[1]] = []int{edge[1]}
		}

		if edge[0] > maxVertexId {
			maxVertexId = edge[0]
		}
		if edge[1] > maxVertexId {
			maxVertexId = edge[1]
		}
	}

	return vertexMergeIdxMap, mergeIdxVertexMap, maxVertexId

}

func getVerticesCount(edges [][]int) int {
	s := make(map[int]bool)

	for _, edge := range(edges) {
		if _, ok := s[edge[0]]; !ok {
			s[edge[0]] = true
		}

		if _, ok := s[edge[1]]; !ok {
			s[edge[1]] = true
		}
	}

	return len(s)
}


