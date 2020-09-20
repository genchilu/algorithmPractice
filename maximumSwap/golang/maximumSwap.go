package maximumSwap

import (
	//"fmt"
	"math"
)

func maximumSwap(num int) int {
	if num <10 {
		return num
	}

	digit := []int{}
	reminder := 0
	for num>0 {
		quotient := num/10
		reminder = num%10
		digit = append([]int{reminder}, digit...)
		num = quotient
	}

	for i, v := range digit {
		max := v
		maxIdx := i
		for j:=i+1;j<len(digit);j++ {
			if digit[j] >= max {
				max = digit[j]
				maxIdx = j
			}
		}

		if (max > v && maxIdx!=i) {
			digit[i], digit[maxIdx] = digit[maxIdx], digit[i]
			break
		}
	}

	num=0
	for i:=0;i<len(digit);i++{
		num += int(math.Pow(10.0, float64(len(digit)-i-1))) * digit[i]
	}
    return num
}