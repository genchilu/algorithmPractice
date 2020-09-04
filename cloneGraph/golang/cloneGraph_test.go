package cloneGraph


import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSearchRange(t *testing.T) {
	testCases := []struct{
		node *Node
	} {
		{
			prepareNode([][]int{[]int{2,4}, []int{1,3}, []int{2,4}, []int{1,3}}),
		},
		{
			prepareNode([][]int{[]int{}}),
		},
		{
			prepareNode([][]int{}),
		},
		{
			nil,
		},
	}

	for _, testCase := range testCases {
		actualNode := cloneGraph(testCase.node)

		if (testCase.node == nil) {
			assert.Equal(t, testCase.node, actualNode)
		} else {
			assert.NotSame(t, &testCase.node, &actualNode)
			assert.Equal(t, testCase.node, actualNode)
		}
	}
}

func prepareNode(adjList [][]int) *Node{
	nodeMap := make(map[int]*Node)
	for i, _ := range adjList {
		noedVal := i+1
		node := Node{noedVal, []*Node{}}
		nodeMap[noedVal] = &node
	}

	for i, neighbors := range adjList {
		noedVal := i+1
		node := nodeMap[noedVal]
		for _, neighbor := range neighbors {
			neighborNode := nodeMap[neighbor]
			node.Neighbors = append(node.Neighbors, neighborNode)
		}
	}

	firstNode := nodeMap[1]
	
	return firstNode
}