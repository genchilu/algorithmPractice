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

func quicksort(slice []int, start, end int) int {
	if (end - start) == 1 {
		return 0
	} else {
		pivot := slice[start]
		i := start + 1
		fmt.Println(i)
		for j := (start + 1); j < (end + 1); j++ {
			fmt.Println(j)
			if slice[j] > pivot {
				fmt.Println(i)
				fmt.Println(j)
				fmt.Println("...")
				tmp := slice[i]
				slice[i] = slice[j]
				slice[j] = tmp
				i++
			}
		}
		tmp := slice[i]
		slice[i] = slice[start]
		slice[start] = tmp
		leftCount := quicksort(slice, start, (i - 1))
		rightCount := quicksort(slice, (i + 1), end)
		return end - start + leftCount + rightCount
	}
}

func main() {
	//slice := readInputToSlice("input")
	slice := []int{2, 1, 4, 6, 3, 7}
	count := quicksort(slice, 0, len(slice)-1)
	fmt.Println(slice)
	fmt.Println(count)
}
