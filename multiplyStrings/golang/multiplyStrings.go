package multiplyStrings

import "fmt"

func multiply(num1 string, num2 string) string {
	m1 := make([]int, len(num1))
	m2 := make([]int, len(num2))

	for i := 0; i < len(num1); i++ {
		m1[i] = int(num1[i]) - 48
	}

	for i := 0; i < len(num2); i++ {
		m2[i] = int(num2[i]) - 48
	}

	results := make([][]int, len(m2))
	for i := 0; i < len(m2); i++ {
		tmp := make([]int, len(m1))
		copy(tmp, m1)

		for j := 0; j < i; j++ {
			tmp = append(tmp, 0)
		}

		tmp = multiplyNum(tmp, m2[len(m2)-i-1])
		results = append(results, tmp)
	}

	sum := []int{}
	for _, v := range results {
		sum = addNum(sum, v)
	}

	sum = removeLeadZero(sum)
	sumstr := ""

	for i := len(sum) - 1; i >= 0; i-- {
		sumstr = fmt.Sprintf("%d%s", sum[i], sumstr)
	}
	return sumstr
}

func multiplyNum(num []int, multiplicand int) []int {
	c := 0
	for i := len(num) - 1; i >= 0; i-- {
		tmp := num[i]*multiplicand + c
		c = tmp / 10
		num[i] = tmp % 10
	}

	for c > 0 {
		tmp := c % 10
		c = c / 10
		num = append([]int{tmp}, num...)
	}

	return num

}

func addNum(a, b []int) []int {
	longer := b
	minl := len(a)
	if len(b) < len(a) {
		minl = len(b)
		longer = a
	}

	c := 0
	result := []int{}
	i := 0
	for i < minl {
		aa := a[len(a)-1-i]
		bb := b[len(b)-1-i]
		sum := aa + bb + c
		c = sum / 10
		result = append([]int{sum % 10}, result...)
		i++
	}

	for i < len(longer) {
		val := longer[len(longer)-1-i]
		sum := val + c
		c = sum / 10
		result = append([]int{sum % 10}, result...)
		i++
	}

	for c > 0 {
		result = append([]int{c % 10}, result...)
		c = c / 10
	}

	return result
}

func removeLeadZero(a []int) []int {
	i := 0
	for i < len(a)-1 {
		if a[i] != 0 {
			break
		}
		i++
	}

	return a[i:]
}
