package fractionToRecurringDecimal

import (
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	b := []byte{}
	m := make(map[int]map[int]int)

	neg := false
	if numerator < 0 {
		neg = !neg
		numerator = -numerator
	}

	if denominator < 0 {
		neg = !neg
		denominator = -denominator
	}

	//fmt.Printf("neg: %v, %d, %d\n", neg, numerator, denominator)
	q, r := numerator/denominator, numerator%denominator
	b = append(b, strconv.Itoa(q)...)
	if r != 0 {

		numerator = r

		b = append(b, '.')
		for {
			numerator *= 10
			q, r := numerator/denominator, numerator%denominator

			if _, ok := m[q]; !ok {
				m[q] = make(map[int]int)
			}
			if _, ok := m[q][r]; !ok {
				m[q][r] = len(b)
			} else {
				idx := m[q][r]
				b = append(b, []byte{')', ')'}...)
				copy(b[idx+1:], b[idx:])
				b[idx] = '('
				break
			}

			b = append(b, []byte(strconv.Itoa(q))...)

			if r == 0 {
				break
			}

			numerator = r
		}
	}

	result := string(b)
	if neg && result != "0" {
		return "-" + result
	}

	return result
}
