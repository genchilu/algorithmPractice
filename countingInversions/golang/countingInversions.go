package countingInversion

//import "fmt"

func CountInversions(input []uint) int {

	if (input == nil || len(input) <=1) {
		return 0
	}

	count, _ := counting(input)
	return count
}

func counting(input []uint) (int, []uint) {

	if (len(input) <= 1) {
		return 0, input
	}

	mid := len(input)/2
	lcount, lresult := counting(input[0: mid])
	rcount, rresult := counting(input[mid:])

	lidx :=0
	ridx :=0
	count := 0
	result := []uint{}
	for lidx < len(lresult) || ridx < len(rresult) {
		if lidx == len(lresult) {
			result = append(result, rresult[ridx:]...)
			ridx = len(rresult)
		} else if ridx== len(rresult) {
			result = append(result, lresult[lidx:]...)
			lidx = len(lresult)
		} else if lresult[lidx] <= rresult[ridx] {
			result = append(result, lresult[lidx])
			lidx++
		} else {
			result = append(result, rresult[ridx])
			ridx++
			count += len(lresult) - lidx
		}
	}

	return (count+lcount+rcount), result
}