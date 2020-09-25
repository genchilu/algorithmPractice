package verticalOrderTraversalOfABinaryTree

import (
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Msg struct {
	treeNode *TreeNode
	vidx     int
	hidx     int
}

func verticalTraversal(root *TreeNode) [][]int {
	queue := []Msg{Msg{root, 0, 0}}
	verticalMap := make(map[int]map[int][]int)

	bfs(&queue, verticalMap)

	verticalIdxs := []int{}
	for k, _ := range verticalMap {
		verticalIdxs = append(verticalIdxs, k)
	}

	sort.Ints(verticalIdxs)

	result := make([][]int, len(verticalIdxs))
	for i, v := range verticalIdxs {

		horizontalIdxs := []int{}
		for k, _ := range verticalMap[v] {
			horizontalIdxs = append(horizontalIdxs, k)
		}
		sort.Ints(horizontalIdxs)

		vresult := make([]int, 0, len(verticalMap[v]))
		for _, h := range horizontalIdxs {
			sort.Ints(verticalMap[v][h])
			vresult = append(vresult, verticalMap[v][h]...)
		}
		result[i] = vresult
	}
	return result
}

func bfs(queue *[]Msg, verticalMap map[int]map[int][]int) {
	if len(*queue) > 0 {
		msg := (*queue)[0]
		*queue = (*queue)[1:len(*queue)]
		if _, ok := verticalMap[msg.vidx]; !ok {
			verticalMap[msg.vidx] = make(map[int][]int)
		}
		if _, ok := verticalMap[msg.vidx][msg.hidx]; !ok {
			verticalMap[msg.vidx][msg.hidx] = []int{}
		}

		verticalMap[msg.vidx][msg.hidx] = append(verticalMap[msg.vidx][msg.hidx], msg.treeNode.Val)
		if msg.treeNode.Left != nil {
			(*queue) = append(*queue, Msg{msg.treeNode.Left, msg.vidx - 1, msg.hidx + 1})
		}
		if msg.treeNode.Right != nil {
			(*queue) = append(*queue, Msg{msg.treeNode.Right, msg.vidx + 1, msg.hidx + 1})
		}

		bfs(queue, verticalMap)
	}
}
