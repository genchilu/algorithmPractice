package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Vertex struct {
	label          int
	neighborsLabel []int
}

func newVertex(label int) Vertex {
	var v Vertex
	v.label = label
	v.neighborsLabel = []int{}
	return v
}

type Graph struct {
	vertices map[int]Vertex
}

func newGraph() Graph {
	var g Graph
	g.vertices = make(map[int]Vertex)
	return g
}

func (g Graph) addVertex(label int) {
	if _, existed := g.vertices[label]; !existed {
		vertex := newVertex(label)
		g.vertices[label] = vertex
	}
}

func (g Graph) addEdge(from, to int) {
	_, existedFrom := g.vertices[from]
	_, existedTo := g.vertices[to]
	if existedFrom && existedTo {
		v := g.vertices[from]
		v.neighborsLabel = append(v.neighborsLabel, to)
		g.vertices[from] = v
	}
}

func (g Graph) neighbors(label int) []int {
	neighbors := g.vertices[label].neighborsLabel
	sort.Ints(neighbors)
	return neighbors
}

func (g Graph) getAllVerticesSortedLabels() []int {
	keys := make([]int, 0, len(g.vertices))
	for k := range g.vertices {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}

func readInputToGraph(filename string) Graph {
	var edges [][]int
	var labels []int
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		vertices := strings.Split(line, " ")
		label, _ := strconv.Atoi(vertices[0])
		labels = append(labels, label)
		for i := 1; i < len(vertices); i++ {
			endpoint, _ := strconv.Atoi(vertices[i])
			edge := []int{label, endpoint}
			edges = append(edges, edge)
		}
	}
	g := newGraph()
	for _, ele := range labels {
		g.addVertex(ele)
	}
	for _, ele := range edges {
		g.addEdge(ele[0], ele[1])
	}
	return g
}

//second dfs
func dfs2(label int, explored map[int]bool, g Graph, rootLabel int, learderMapLabel map[int][]int) {
	explored[label] = true
	labels := learderMapLabel[rootLabel]
	labels = append(labels, label)
	learderMapLabel[rootLabel] = labels
	for _, neighbor := range g.neighbors(label) {
		if _, isExplored := explored[neighbor]; !isExplored {
			dfs2(neighbor, explored, g, rootLabel, learderMapLabel)
		}
	}
}

//first dfs
func dfs1(label int, explored map[int]bool, g Graph,
	labelMapFinishTime map[int]int, finishTimes *[]int) {
	explored[label] = true
	for _, neighbor := range g.neighbors(label) {
		if _, isExplored := explored[neighbor]; !isExplored {
			dfs1(neighbor, explored, g, labelMapFinishTime, finishTimes)
		}
	}
	finishTime++
	labelMapFinishTime[label] = finishTime
	*finishTimes = append(*finishTimes, finishTime)
}

func revGraph(g Graph, labelMapFinishTime map[int]int, finishTimes []int) Graph {
	revg := newGraph()
	for _, ele := range finishTimes {
		revg.addVertex(ele)
	}
	labelsOri := g.getAllVerticesSortedLabels()
	for _, ele := range labelsOri {
		revEdgeTo := labelMapFinishTime[ele]
		neighborsOri := g.neighbors(ele)
		for _, neighbor := range neighborsOri {
			revEdgeFrom := labelMapFinishTime[neighbor]
			revg.addEdge(revEdgeFrom, revEdgeTo)
		}
	}
	return revg
}

var finishTime int

func main() {
	g := readInputToGraph("SCC.txt")
	labels := g.getAllVerticesSortedLabels()
	explored := map[int]bool{}
	labelMapFinishTime := map[int]int{}
	finishTimes := []int{}
	finishTime = 0
	for idx, _ := range labels {
		ele := labels[len(labels)-idx-1]
		if _, isExplored := explored[ele]; !isExplored {
			dfs1(ele, explored, g, labelMapFinishTime, &finishTimes)
		}
	}
	revg := revGraph(g, labelMapFinishTime, finishTimes)
	newExplored := map[int]bool{}
	leaderMapLabel := map[int][]int{}
	newLabels := revg.getAllVerticesSortedLabels()
	for idx, _ := range newLabels {
		ele := newLabels[len(newLabels)-idx-1]
		if _, isExplored := newExplored[ele]; !isExplored {
			leaderMapLabel[ele] = []int{}
			dfs2(ele, newExplored, revg, ele, leaderMapLabel)
		}
	}
	leadersSize := make([]int, 0, len(leaderMapLabel))
	for leader := range leaderMapLabel {
		leadersSize = append(leadersSize, len(leaderMapLabel[leader]))
	}
	sort.Ints(leadersSize)
	for i := 1; i <= 5; i++ {
		idx := len(leadersSize) - i
		fmt.Println(leadersSize[idx])
	}
	//fmt.Println(g.neighbors(1))
}
