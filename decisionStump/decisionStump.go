package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func rnadomFloat64() float64 {
	rand.Seed(time.Now().UnixNano())
	return (rand.Float64() * 2) - 1
}

func isEleInSlice(input []float64, x_i float64) bool {
	for _, i := range input {
		if i == x_i {
			return true
		}
	}
	return false
}

func generateInput() ([]float64, []int) {
	var input []float64
	var y []int
	for i := 0; i < 20; i++ {
		x_i := rnadomFloat64()
		for isEleInSlice(input, x_i) {
			x_i = rnadomFloat64()
		}
		input = append(input, x_i)
	}
	sort.Float64s(input)
	for _, x_i := range input {
		rand.Seed(time.Now().UnixNano())
		flip := rand.Intn(100) + 1
		var y_i int
		if x_i > 0.0 {
			y_i = 1
		} else {
			y_i = -1
		}
		if flip <= 20 {
			y_i = -y_i
		}
		y = append(y, y_i)
	}
	return input, y
}

func caculateThetas(input []float64) []float64 {
	var thetas []float64
	for i := 0; i <= len(input); i++ {
		if i == 0 {
			thetas = append(thetas, input[i]-0.1)
		} else if i == len(input) {
			thetas = append(thetas, input[i-1]+0.1)
		} else {
			theta := (input[i-1] + input[i]) / 2
			thetas = append(thetas, theta)
		}
	}
	return thetas
}

func caculateErr(x []float64, theat float64, s int, y []int) float64 {
	errCount := 0
	for i := range x {
		perdict := 1
		if (x[i]-theat)*float64(s) < 0 {
			perdict = -1
		}
		if perdict != y[i] {
			errCount++
		}
	}
	return float64(errCount) / float64(len(x))
}

func caculateEinTheatS(x, theats []float64, y []int) (float64, float64, int) {
	var s int
	var theat float64
	ein := 1.0
	for _, ele := range theats {
		newein := caculateErr(x, ele, 1, y)
		if newein < ein {
			ein = newein
			s = 1
			theat = ele
		}
		newein = caculateErr(x, ele, -1, y)
		if newein < ein {
			ein = newein
			s = -1
			theat = ele
		}
	}
	return ein, theat, s
}

func readInputToSlice(filename string) ([][]float64, []int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var x [][]float64
	var y []int
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		tmp := strings.Split(line, " ")
		y_i, _ := strconv.Atoi(tmp[9])
		var x_i []float64
		for i := 0; i < 9; i++ {
			f, _ := strconv.ParseFloat(tmp[i], 64)
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

func main() {
	//Q17 & Q18
	einSum := 0.0
	eoutSum := 0.0
	for i := 0; i < 5000; i++ {
		x, y := generateInput()
		thetas := caculateThetas(x)
		ein, theta, s := caculateEinTheatS(x, thetas, y)
		einSum += ein
		eout := 0.5 + 0.3*float64(s)*(math.Abs(theta)-1.0)
		eoutSum += eout
		//fmt.Printf("ein: %f, eout: %f, theat: %f, s: %d\n", ein, eout, theta, s)
	}
	fmt.Printf("avg ein: %f, avg eout: %f\n", einSum/5000.0, eoutSum/5000.0)

	//Q19 & 20
	minEin := 1.0
	var dim int
	inputTrain, yTrain := readInputToSlice("train.txt")
	inputTest, yTest := readInputToSlice("test.txt")
	var theta, ein float64
	var s int
	for i := 0; i < len(inputTrain[0]); i++ {
		var x []float64
		for j := 0; j < len(inputTrain); j++ {
			x = append(x, inputTrain[j][i])
		}
		mapY := map[float64]int{}
		for i := range x {
			mapY[x[i]] = yTrain[i]
		}
		sort.Float64s(x)
		var newY []int
		for _, ele := range x {
			newY = append(newY, mapY[ele])
		}
		thetas := caculateThetas(x)
		ein, theta, s = caculateEinTheatS(x, thetas, newY)
		//fmt.Printf("theta: %f, s: %d, dim: %d, ein: %f\n", theta, s, i, ein)
		if ein < minEin {
			minEin = ein
			dim = i
		}
	}
	var testX []float64
	for i := 0; i < len(inputTest); i++ {
		testX = append(testX, inputTest[i][dim])
	}
	eout := caculateErr(testX, theta, s, yTest)
	fmt.Printf("best ein: %f\n", minEin)
	fmt.Printf("eout: %f\n", eout)
}
