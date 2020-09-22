package isGraphBipartite

//import "fmt"

type Msg struct {
	idx int
	color bool
}

func isBipartite(graph [][]int) bool {
	if len(graph) == 0 {
		return false
	}

	if len(graph) == 1 {
		return true
	}


	colors := make(map[int]bool)

	for i, _ := range graph {
		if _,ok := colors[i];!ok {
			queue := []Msg{Msg{i, true}}
			result := _bfs(graph, &queue, colors)
			if !result {
				return result
			}
		}
	}

	return true
}

func _bfs(graph [][]int, queue *[]Msg, colors map[int]bool) bool {
	msg := (*queue)[0]
	*queue = (*queue)[1:]

	if c, ok := colors[msg.idx];!ok {
		colors[msg.idx] = msg.color

		for _, v := range graph[msg.idx] {
			if cc,ok:=colors[v];!ok {
				smsg := Msg{v, !msg.color}
				*queue = append(*queue, smsg)
			} else if cc!=!msg.color{
				return false
			}
		}
 	} else if c != msg.color {
		return false
	}

	if len(*queue) > 0 {
		return _bfs(graph, queue, colors)
	}

	return true
}

