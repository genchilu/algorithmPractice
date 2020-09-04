package cloneGraph

//import "fmt"
type Node struct {
    Val int
    Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
    if node == nil {
        return node
    }
    nodeMap := make(map[int]*Node)
    dfsCloneVal(node, nodeMap)
    dfsCloneNeighbors(node, nodeMap)

    return nodeMap[node.Val]
}

func dfsCloneVal(n *Node, nodeMap map[int]*Node) {
    if _, ok := nodeMap[n.Val];!ok {
        cloneN := Node{n.Val, []*Node{}}
        nodeMap[n.Val] = &cloneN
        for _, neighbor := range n.Neighbors {
            dfsCloneVal(neighbor, nodeMap)
        }
    }
}

func dfsCloneNeighbors(n *Node, nodeMap map[int]*Node) {
    cloneNode := nodeMap[n.Val]
    if len(cloneNode.Neighbors) == 0 {
        for _, neighbor := range n.Neighbors {
            cloneNeighbor := nodeMap[neighbor.Val]
            cloneNode.Neighbors = append(cloneNode.Neighbors, cloneNeighbor)
            dfsCloneNeighbors(neighbor, nodeMap)
        }
    }
}
