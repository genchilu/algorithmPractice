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

func quicksort(slice []int, start, end int) int {
	if (end - start) == 0 {
		return 0
	} else if (end - start) == 1 {
		if slice[start] > slice[end] {
			swap(slice, start, end)
		}
		return 1
	} else {
		i := start + 1
		for j := (start + 1); j < (end + 1); j++ {
			if slice[j] < slice[start] {
				swap(slice, i, j)
				i++
			}
		}
		swap(slice, (i - 1), start)
		leftCount := quicksort(slice, start, (i - 2))
		rightCount := quicksort(slice, i, end)
		return end - start + leftCount + rightCount
	}
}

func main() {
	slice := readInputToSlice("input")
	//slice := []int{6, 1, 4, 2, 3, 7}
	count := quicksort(slice, 0, len(slice)-1)
	//fmt.Println(slice)
	fmt.Println(count)
}
