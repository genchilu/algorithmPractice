package isGraphBipartite

//import "fmt"

type Msg struct {
	idx int
	color byte
}

func isBipartite(graph [][]int) bool {
	if len(graph) == 0 {
		return false
	}

	if len(graph) == 1 {
		return true
	}


	colors := make([]byte, len(graph))
	for i,_:=range colors {
		colors[i] = 0
	}

	for i, _ := range graph {
		if colors[i] == byte(0) {
			queue := []Msg{Msg{i, 1}}
			result := _bfs(graph, &queue, &colors)
			if !result {
				return result
			}
		}
	}

	return true
}

func _bfs(graph [][]int, queue *[]Msg, colors *[]byte) bool {
	msg := (*queue)[0]
	*queue = (*queue)[1:]
	idx := msg.idx
	color := msg.color

	if (*colors)[idx] == byte(0) {
		(*colors)[idx] = color
		nextColor := byte(1)
		if color == nextColor {
			nextColor = byte(2)
		}

		for _, v := range graph[idx] {
			msg := Msg{v, nextColor}
			*queue = append(*queue, msg)
		}
 	} else if (*colors)[idx] != color {
		return false
	}

	if len(*queue) > 0 {
		return _bfs(graph, queue, colors)
	}

	return true
}

