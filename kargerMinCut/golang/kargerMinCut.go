package kargerMinCut

import (
	"math/rand"
	"time"
)

func doMinCut(edges [][]int) int {
	if edges == nil || len(edges) == 0 {
		return 0
	}

	vertexCount := countVertex(edges)

	minCut := len(edges)
	for i:=0;i<vertexCount*vertexCount;i++ {

		tmpMinCut := doKragerMinCut(edges)
		if tmpMinCut < minCut {
			minCut = tmpMinCut
		}
	}
	return minCut
}

func doKragerMinCut(edges [][]int) int {
	cloneEdges := make([][]int, len(edges))
	copy(cloneEdges, edges)

	vertexMergeIdxMap, mergeIdxVertexMap, maxVertex := initVertexMergeIdxMap(edges)


	mergeCount := 0
	for mergeCount< (len(vertexMergeIdxMap) -2 ) {

		rand.Seed(time.Now().UnixNano())
		idx := rand.Intn(len(cloneEdges))
		edge := cloneEdges[idx]
		v0 := edge[0]
		mergeid0 := vertexMergeIdxMap[v0]
		v1 := edge[1]
		mergeid1 := vertexMergeIdxMap[v1]

		if mergeid0 != mergeid1 {
			mergeCount++
			mergeIdx := mergeCount + maxVertex

			for _, v := range mergeIdxVertexMap[mergeid0] {
				vertexMergeIdxMap[v] = mergeIdx
			}

			for _, v := range mergeIdxVertexMap[mergeid1] {
				vertexMergeIdxMap[v] = mergeIdx
			}

			mergeGroup := append(mergeIdxVertexMap[mergeid0], mergeIdxVertexMap[mergeid1]...)
			mergeIdxVertexMap[mergeIdx] = mergeGroup

			delete(mergeIdxVertexMap, mergeid0)
			delete(mergeIdxVertexMap, mergeid1)
		}

		cloneEdges = append(cloneEdges[0:idx], cloneEdges[idx+1:]...)
	}

	minCut := 0
	for _, edge := range cloneEdges {
		v0 := edge[0]
		v1 := edge[1]

		if vertexMergeIdxMap[v0] != vertexMergeIdxMap[v1] {
			minCut++
		}
	} 

	return minCut
}

func initVertexMergeIdxMap(edges [][]int) (map[int]int, map[int][]int, int) {
	maxVertex := 0
	vertexMergeIdxMap := make(map[int]int)
	mergeIdxVertexMap := make(map[int][]int)

	for _, edge := range edges {
		v0 := edge[0]
		v1 := edge[1]

		vertexMergeIdxMap[v0] = v0
		vertexMergeIdxMap[v1] = v1

		mergeIdxVertexMap[v0] = []int{v0}
		mergeIdxVertexMap[v1] = []int{v1}

		if(v0>maxVertex) {
			maxVertex = v0
		}

		if(v1>maxVertex) {
			maxVertex = v1
		}
	}

	return vertexMergeIdxMap, mergeIdxVertexMap, maxVertex
}

func countVertex(edges [][]int) int {
	vertexMap := make(map[int]bool)

	for _, edge := range edges {
		vertexMap[edge[0]] = true
		vertexMap[edge[1]] = true
	}

	return len(vertexMap)
}

