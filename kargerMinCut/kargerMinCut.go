package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func readInputToSlice(filename string) ([][]int, [][]int) {
	var slice [][]int
	var idxlabel [][]int
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		vertices := strings.Split(line, "	")
		label, _ := strconv.Atoi(vertices[0])
		label--
		//fmt.Println(label)
		idxlabel = append(idxlabel, []int{label})
		for i := 1; i < len(vertices); i++ {
			endpoint, _ := strconv.Atoi(vertices[i])
			endpoint--
			edge := []int{label, endpoint}
			//Deduplication
			isOnly := true
			sort.Ints(edge)
			for _, ele := range slice {
				if ele[0] == edge[0] && ele[1] == edge[1] {
					isOnly = false
				}
			}
			if isOnly {
				slice = append(slice, edge)
			}
		}
	}
	return slice, idxlabel
}

func kargerMinCut(slice [][]int, idxlabel [][]int) int {
	verticesNumber := len(idxlabel)
	for verticesNumber > 2 {
		//fmt.Println("===============")
		rand.Seed(time.Now().UnixNano())
		selectedge := rand.Intn(len(slice))
		//fmt.Println("select edge ", selectedge)
		endpointIdx1 := slice[selectedge][0]
		endpointIdx2 := slice[selectedge][1]
		idxlabel = mergeVertices(idxlabel, endpointIdx1, endpointIdx2)
		//fmt.Println("new vertices ", idxlabel)
		slice = append(slice[:selectedge], slice[selectedge+1:]...)
		var newSlice [][]int
		for i, _ := range slice {
			if slice[i][0] == endpointIdx1 || slice[i][0] == endpointIdx2 {
				slice[i][0] = len(idxlabel) - 1
			}
			if slice[i][1] == endpointIdx1 || slice[i][1] == endpointIdx2 {
				slice[i][1] = len(idxlabel) - 1
			}
			if slice[i][0] != slice[i][1] {
				newSlice = append(newSlice, slice[i])
			}
		}
		slice = newSlice
		//fmt.Println("new slice ", slice)
		//fmt.Println("new idxlabel ", idxlabel)
		verticesNumber--
	}
	//fmt.Println(idxlabel)
	return len(slice)
}

func mergeVertices(idxlabel [][]int, endpointIdx1, endpointIdx2 int) [][]int {
	var newVertex []int
	//fmt.Println(idxlabel)
	//fmt.Println("endpointIdx1: ", endpointIdx1)
	//fmt.Println("endpointIdx2: ", endpointIdx2)
	for _, ele := range idxlabel[endpointIdx1] {
		newVertex = append(newVertex, ele)
	}
	for _, ele := range idxlabel[endpointIdx2] {
		newVertex = append(newVertex, ele)
	}
	//fmt.Println(newVertex)
	idxlabel = append(idxlabel, newVertex)
	idxlabel[endpointIdx1] = []int{}
	idxlabel[endpointIdx2] = []int{}
	return idxlabel
}

func main() {
	slice, idxlabel := readInputToSlice("kargerMinCut.txt")
	//slice, idxlabel := readInputToSlice("test.txt")
	fmt.Println("len slice: ", len(slice))
	//fmt.Println(idxlabel)
	minCut := len(slice)
	fmt.Println("init min cut: ", minCut)
	for i := 0; i < len(idxlabel)*len(idxlabel); i++ {
		if i%100 == 0 {
			fmt.Printf("iterator %dth\n", i)
		}
		tmpslice := [][]int{}
		for _, ele := range slice {
			edge := []int{ele[0], ele[1]}
			tmpslice = append(tmpslice, edge)
		}

		tmpidxlabel := [][]int{}
		for _, ele := range idxlabel {
			label := []int{ele[0]}
			tmpidxlabel = append(tmpidxlabel, label)
		}
		//fmt.Println("start: ", tmpidxlabel)
		tmp := kargerMinCut(tmpslice, tmpidxlabel)
		if tmp < minCut {
			minCut = tmp
			fmt.Println("new min cut: ", minCut)
		}
		//fmt.Println("after idxlabel ", idxlabel)
		//fmt.Printf("i: %d\n", i)
	}
	fmt.Println(minCut)
}
