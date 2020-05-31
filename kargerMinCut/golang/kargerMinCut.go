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
	//vertextCount = 1

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
	copy(tmpEdges, edges);

	vertexMergeIdxMap := make(map[int]int)
	revVertexMergeIdxMap := make(map[int][]int)

	for _, edge := range tmpEdges {
		e1 := edge[0]
		e2 := edge[1]

		_, ok := vertexMergeIdxMap[e1]
		if !ok {
			vertexMergeIdxMap[e1] = e1
		}

		_, ok = revVertexMergeIdxMap[e1]
		if !ok {
			revVertexMergeIdxMap[e1] = []int{ e1 }
		}

		_, ok = vertexMergeIdxMap[e2]
		if !ok {
			vertexMergeIdxMap[e2] = e2
		}

		_, ok = revVertexMergeIdxMap[e2]
		if !ok {
			revVertexMergeIdxMap[e2] = []int{ e2 }
		}
	}

	mergeCount := 0
	for mergeCount < (vertextCount-2) {
		pickIdx := rand.Intn(len(tmpEdges))
		pickedEdge := tmpEdges[pickIdx]
		v1 := pickedEdge[0]
		v2 := pickedEdge[1]

		if(vertexMergeIdxMap[v1] != vertexMergeIdxMap[v2]) {
			curVertexMergeIdx := vertextCount + mergeCount

			curV1Idx := vertexMergeIdxMap[v1]
			curV2Idx := vertexMergeIdxMap[v2]

			for _, e := range revVertexMergeIdxMap[curV1Idx]{
				vertexMergeIdxMap[e] = curVertexMergeIdx
			}

			for _, e := range revVertexMergeIdxMap[curV2Idx]{
				vertexMergeIdxMap[e] = curVertexMergeIdx
			}
		
			revVertexMergeIdxMap[curVertexMergeIdx] = append(revVertexMergeIdxMap[curV1Idx], revVertexMergeIdxMap[curV2Idx]...)
			delete(revVertexMergeIdxMap, v1)
			delete(revVertexMergeIdxMap, v2)

			mergeCount++
		}

		tmpEdges = append(tmpEdges[0:pickIdx], tmpEdges[pickIdx+1:]...)
	}

	cutCount := 0
	for _, edge := range tmpEdges {
		e1 := edge[0]
		e2 := edge[1]
		if(vertexMergeIdxMap[e1] != vertexMergeIdxMap[e2]) {
			cutCount++
		}
	}

	return cutCount;
}




func getVertexCount(edges [][]int) int {
	vertexMap := make(map[int]bool)
	for _, v := range edges {
		vertexMap[v[0]] = true
		vertexMap[v[1]] = true
	}

	return len(vertexMap);
}
