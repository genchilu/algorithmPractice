package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func readInputToSlice(filename string) ([][]float64, []float64) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var x [][]float64
	var y []float64
	for scanner.Scan() {
		tmp := strings.Split(scanner.Text(), "	")
		xStr := strings.Split(tmp[0], " ")
		y_i, _ := strconv.ParseFloat(tmp[1], 64)
		var x_i []float64
		x_i = append(x_i, 1.0)
		for i := 0; i < 4; i++ {
			f, _ := strconv.ParseFloat(xStr[i], 64)
			x_i = append(x_i, f)
		}

		x = append(x, x_i)
		y = append(y, y_i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	//fmt.Println(x)
	return x, y
}

func correctW(w, x []float64, y float64, param float64) []float64 {
	var result []float64
	//fmt.Println("x: ", x)
	//fmt.Printf("y: %f, param: %f\n", y, param)
	//fmt.Println("origin w: ", w)
	for i := 0; i < len(w); i++ {
		result = append(result, w[i]+(x[i]*y*param))
	}
	//fmt.Println("result: ", result)
	return result
}

func sign(w, x []float64) float64 {
	result := 0.0
	for i := 0; i < len(w); i++ {
		result += w[i] * x[i]
	}
	if result > 0.0 {
		//fmt.Printf("result: %f, return %f\n", result, 1.0)
		return 1.0
	} else {
		//fmt.Printf("result: %f, return %f\n", result, -1.0)
		return -1.0
	}
}

func countPlaPocketError(x [][]float64, y []float64, w []float64) int {
	errCount := 0
	for i := 0; i < len(x); i++ {
		if sign(w, x[i]) != y[i] {
			errCount++
		}
	}
	return errCount
}

func plaPocket(x [][]float64, y []float64, notFindBetter bool,
	iteratorLimit int) []float64 {
	var w = []float64{1.0, 0.0, 0.0, 0.0, 0.0}
	//var tmpw = []float64{1.0, 0.0, 0.0, 0.0, 0.0}
	iterator := 0
	index := 0
	minErrCount := len(x)
	betterW := w
	for iterator < iteratorLimit {
		rand.Seed(time.Now().UnixNano())
		newIndex := index
		for index == newIndex {
			newIndex = rand.Intn(len(x))
		}
		index = newIndex
		signValue := sign(w, x[index])
		if iterator == 0 {
			signValue = -1
		}
		//fmt.Printf("index: %d\n", index)
		//fmt.Printf("index: %d, iterator: %d, signValue: %f, yi: %f\n", index, iterator, signValue, y[index])
		if signValue != y[index] {
			w = correctW(w, x[index], y[index], 1.0)
			errCount := countPlaPocketError(x, y, w)
			//fmt.Printf("signal error, original err: %d, new err: %d\n", originalErrCount, newErrCount)
			if errCount < minErrCount || notFindBetter {
				//fmt.Println("origin w: ", w)
				//fmt.Printf("ori min err count: %d, new err count: %d\n", minErrCount, errCount)
				//fmt.Println("new w: ", betterW)
				minErrCount = errCount
				//fmt.Printf("now min err count: %d\n", minErrCount)
				betterW = w
				//fmt.Println("new w: ", w)
			}
			iterator++
		}
	}
	return betterW
}

func pla(x [][]float64, y []float64, param float64, random bool) int {
	var w = []float64{1.0, 0.0, 0.0, 0.0, 0.0}
	iterator := 1

	var idx []int
	for i := 0; i < len(x); i++ {
		idx = append(idx, i)
	}
	if random {
		shuffle(idx)
	}
	if y[0] != -1 {
		w = correctW(w, x[0], y[0], param)
	}
	allPass := false
	i := 1
	current := 0
	for !allPass {
		index := idx[i]
		if sign(w, x[index]) != y[index] {
			w = correctW(w, x[index], y[index], param)
			iterator++
			current = i
		}
		i = (i + 1) % len(idx)
		if i == current {
			allPass = true
		}
	}
	return iterator
}

func shuffle(a []int) {
	rand.Seed(time.Now().UnixNano())
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}

func main() {
	x, y := readInputToSlice("input1")
	/**Q15**/
	iterator15 := pla(x, y, 1.0, false)
	fmt.Printf("ans for q15: %d\n", iterator15)
	/**/

	/**Q16**/
	iteratorSum := 0
	for i := 0; i < 2000; i++ {
		iteratorSum += pla(x, y, 1.0, true)
		//fmt.Printf("i = %d\n, iteratorSum = %d", i, iteratorSum)
	}
	iteratorAvg := iteratorSum / 2000
	fmt.Printf("ans for q16: %d\n", iteratorAvg)
	/**/

	/**Q17**/
	iteratorSum = 0
	for i := 0; i < 2000; i++ {
		iteratorSum += pla(x, y, 0.5, true)
		//fmt.Printf("i = %d\n, iteratorSum = %d", i, iteratorSum)
	}
	iteratorAvg = iteratorSum / 2000
	fmt.Printf("ans for q17: %d\n", iteratorAvg)
	/**/

	/**Q18**/
	x, y = readInputToSlice("input2")
	xTest, yTest := readInputToSlice("test")
	errPercentSum := 0.0
	errPercent := 0.0
	for i := 0; i < 2000; i++ {
		w := plaPocket(x, y, false, 50)
		errCount := countPlaPocketError(xTest, yTest, w)
		errPercentSum += (float64(errCount) / (float64(len(xTest))))
	}
	errPercent = errPercentSum / 2000.0
	fmt.Printf("ans for q18: %f\n", errPercent)
	/**/

	/**Q19**/
	errPercentSum = 0
	errPercent = 0
	for i := 0; i < 2000; i++ {
		w := plaPocket(x, y, true, 50)
		errCount := countPlaPocketError(xTest, yTest, w)
		errPercentSum += (float64(errCount) / (float64(len(xTest))))
	}
	errPercent = errPercentSum / 2000.0
	fmt.Printf("ans for q19: %f\n", errPercent)
	/**/

	/**Q20**/
	errPercentSum = 0
	errPercent = 0
	for i := 0; i < 2000; i++ {
		w := plaPocket(x, y, false, 100)
		errCount := countPlaPocketError(xTest, yTest, w)
		errPercentSum += (float64(errCount)) / (float64(len(xTest)))
	}
	errPercent = errPercentSum / 2000.0
	fmt.Printf("ans for q20: %f\n", errPercent)
	/**/
}
