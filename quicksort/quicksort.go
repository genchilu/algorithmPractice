package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readInputToSlice(filename string) []int {
	var slice []int
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		slice = append(slice, x)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return slice
}

func swap(slice []int, i, j int) {
	tmp := slice[i]
	slice[i] = slice[j]
	slice[j] = tmp
}

func quicksort(slice []int, start, end int, pivotType string) int {
	//fmt.Println("start ", slice)
	//fmt.Printf("start %d, end %d\n", start, end)
	if (end - start) <= 0 {
		return 0
	} else if (end - start) == 1 {
		if slice[start] > slice[end] {
			swap(slice, start, end)
		}
		return 1
	} else {
		var pivotPos int
		if pivotType == "first" {
			pivotPos = start
		} else if pivotType == "end" {
			pivotPos = end
		} else if pivotType == "mid" {
			var midPos int
			if (end-start+1)%2 == 0 {
				midPos = ((end - start + 1) / 2) - 1 + start
			} else {
				midPos = (end-start)/2 + start
			}
			if slice[start] > slice[midPos] && slice[midPos] > slice[end] {
				pivotPos = midPos
			} else if slice[midPos] > slice[start] && slice[start] > slice[end] {
				pivotPos = start
			} else {
				pivotPos = end
			}
		}

		if pivotPos != start {
			swap(slice, start, pivotPos)
		}

		i := start + 1
		for j := (start + 1); j < (end + 1); j++ {
			if slice[j] < slice[start] {
				swap(slice, i, j)
				i++
			}
		}
		swap(slice, (i - 1), start)
		//fmt.Println("after ", slice)
		leftCount := quicksort(slice, start, (i - 2), pivotType)
		rightCount := quicksort(slice, i, end, pivotType)
		comparCount := end - start
		return comparCount + leftCount + rightCount
	}
}

func main() {
	slice := readInputToSlice("input")
	//slice := []int{6, 10, 1, 11, 9, 4, 2, 12, 8, 3, 13, 7}
	count1 := quicksort(slice, 0, len(slice)-1, "first")
	fmt.Println(count1)

	slice = readInputToSlice("input")
	count2 := quicksort(slice, 0, len(slice)-1, "end")
	fmt.Println(count2)

	slice = readInputToSlice("input")
	count := quicksort(slice, 0, len(slice)-1, "mid")
	//fmt.Println(slice)
	fmt.Println(count)
}
