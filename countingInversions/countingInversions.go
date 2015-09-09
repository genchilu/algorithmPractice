package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
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

func countingMerge(sliceL, sliceR []int) (int, []int) {
	i := 0
	j := 0
	var sortSlice []int
	inversionCount := 0
	for i < len(sliceL) || j < len(sliceR) {
		if i == len(sliceL) {
			sortSlice = append(sortSlice, sliceR[j])
			j++
		} else if j == len(sliceR) {
			sortSlice = append(sortSlice, sliceL[i])
			i++
		} else if sliceL[i] <= sliceR[j] {
			sortSlice = append(sortSlice, sliceL[i])
			i++
		} else {
			sortSlice = append(sortSlice, sliceR[j])
			inversionCount += len(sliceL) - i
			j++
		}
	}
	return inversionCount, sortSlice
}

func countingInversions(slice []int) (int, []int) {
	length := len(slice)
	if length == 1 {
		return 0, slice
	} else {
		mid := length / 2
		var countleft, countright, countMerge int
		var sortedL, sortedR, sorted []int
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			countleft, sortedL = countingInversions(slice[0:mid])
		}()
		wg.Add(1)
		go func() {
			defer wg.Done()
			countright, sortedR = countingInversions(slice[mid:length])
		}()
		wg.Wait()
		countMerge, sorted = countingMerge(sortedL, sortedR)
		return countleft + countright + countMerge, sorted
		//		countleft, sortedL := countingInversions(slice[0:mid])
		//		countright, sortedR := countingInversions(slice[mid:length])
		//		countMerge, sorted := countingMerge(sortedL, sortedR)
		//		return countleft + countright + countMerge, sorted
	}
}

func main() {
	slice := readInputToSlice("input")
	//slice := []int{1, 2, 3, 4}
	inversionCount, _ := countingInversions(slice)
	fmt.Println(inversionCount)
	//fmt.Println(sortedSlice)
}
