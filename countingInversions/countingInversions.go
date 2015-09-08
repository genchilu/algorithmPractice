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

func countingMerge(sliceL, sliceR []int) (int, []int) {
	i := 0
	j := 0
	var sortSlice []int
	inversionCount := 0
	for k := 0; k < (len(sliceL) + len(sliceR)); k++ {
		if sliceL[i] <= sliceR[j] {
			sortSlice = append(sortSlice, sliceL[i])
			if i == len(sliceL)-1 {
				sortSlice = append(sortSlice, sliceR[j:len(sliceR)]...)
				break
			} else {
				i++
			}
		} else {
			sortSlice = append(sortSlice, sliceR[j])
			inversionCount += len(sliceL) - i
			if j == len(sliceR)-1 {
				sortSlice = append(sortSlice, sliceL[i:len(sliceL)]...)
				break
			} else {
				j++
			}
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
		countleft, sortedL := countingInversions(slice[0:mid])
		countright, sortedR := countingInversions(slice[mid:length])
		countMerge, sorted := countingMerge(sortedL, sortedR)
		return countleft + countright + countMerge, sorted
	}
}

func main() {
	slice := readInputToSlice("input")
	//slice := []int{1, 3, 5, 2, 4, 6}
	inversionCount, _ := countingInversions(slice)
	fmt.Println(inversionCount)
	//fmt.Println(sortedSlice)
}
