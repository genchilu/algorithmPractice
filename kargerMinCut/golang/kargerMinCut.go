package kargerMinCut

import (
	"math/rand"
	"time"
)

func doMinCut(edges [][]int) int {
	verticesCount := countVertices(edges)

	minCut := len(edges)
	for i:=0;i<verticesCount*verticesCount;i++{
		curMinCut := kargerMinCut(edges)

		if (curMinCut < minCut){
			minCut = curMinCut;
		}

	}
	return minCut
}

func countVertices(edges [][]int) int {
	verticesMap := make(map[int]bool)

	for _, edge := range edges{
		if _, ok := verticesMap[edge[0]]; !ok {
			verticesMap[edge[0]] = true
		}

		if _, ok := verticesMap[edge[1]]; !ok {
			verticesMap[edge[1]] = true
		}
	}

	return len(verticesMap)
}

func kargerMinCut(edges [][]int) int {

	cloneEdges := make([][]int, len(edges))
	copy(cloneEdges, edges)


	verticesMergeIdxMap, mergeIdxVerticesMap, maxVertex:= initVerticesMergeIdxMap(edges)
	

	mergeCount := 0
	for mergeCount < len(verticesMergeIdxMap) -2 {
		rand.Seed(time.Now().UnixNano())
		pickedIdx := rand.Intn(len(cloneEdges))
		pickedEdge := cloneEdges[pickedIdx]

		v0 := pickedEdge[0]
		mergeIdx0 := verticesMergeIdxMap[v0]

		v1 := pickedEdge[1]
		mergeIdx1 := verticesMergeIdxMap[v1]

		if (mergeIdx0 != mergeIdx1) {
			curMergeIdx := maxVertex + mergeCount + 1

			mergeIdxGroup0 := mergeIdxVerticesMap[mergeIdx0]
			for _, v := range mergeIdxGroup0 {
				verticesMergeIdxMap[v] = curMergeIdx
			}

			mergeIdxGroup1 := mergeIdxVerticesMap[mergeIdx1]
			for _, v := range mergeIdxGroup1 {
				verticesMergeIdxMap[v] = curMergeIdx
			}
			
			mergeIdxVerticesMap[curMergeIdx] = append(mergeIdxGroup0, mergeIdxGroup1...)

			delete(mergeIdxVerticesMap, mergeIdx0)
			delete(mergeIdxVerticesMap, mergeIdx1)

			mergeCount++

		}

		cloneEdges = append(cloneEdges[0: pickedIdx], cloneEdges[pickedIdx+1:]...)
	}

	count := 0
	for _, edge := range cloneEdges {
		v0 := edge[0]
		mergeIdx0 := verticesMergeIdxMap[v0]

		v1 := edge[1]
		mergeIdx1 := verticesMergeIdxMap[v1]

		if (mergeIdx0 != mergeIdx1) {
			count++
		}
	}
	return count
}

func initVerticesMergeIdxMap(edges [][]int) (map[int]int,  map[int][]int, int) {
	verticesMergeIdxMap := make(map[int]int)
	mergeIdxVerticesMap := make(map[int][]int)
	maxVertex := 0

	for _, edge := range edges {
		v0 := edge[0]
		v1 := edge[1]

		verticesMergeIdxMap[v0] = v0
		verticesMergeIdxMap[v1] = v1

		if _, ok := mergeIdxVerticesMap[v0]; !ok {
			mergeIdxVerticesMap[v0] = []int{v0}
		}
		if _, ok := mergeIdxVerticesMap[v1]; !ok {
			mergeIdxVerticesMap[v1] = []int{v1}
		}

		if(v0>maxVertex) {
			maxVertex = v0
		}

		if(v1>maxVertex) {
			maxVertex = v1
		}
	}

	return verticesMergeIdxMap, mergeIdxVerticesMap, maxVertex;

}
