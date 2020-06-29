package countingInversion

//import "fmt"

func CountInversions(input []uint) int {
	if (input == nil || len(input) <2) {
		return 0
	}
	_, count:=count(input)
	return count
}

func count(input []uint) ([]uint, int) {
	if (len(input) < 2) {
		return input, 0
	}
	mid := len(input)/2
	linput, lcount := count(input[0:mid])
	rinput, rcount := count(input[mid:])

	count :=0
	lidx :=0
	ridx :=0
	sortedInput := []uint{}
	for lidx <len(linput) || ridx < len(rinput) {
		if (lidx == len(linput)) {
			sortedInput = append(sortedInput, rinput[ridx:]...)
			ridx = len(rinput)
		} else if (ridx == len(rinput)) {
			sortedInput = append(sortedInput, linput[lidx:]...)
			lidx = len(linput)
		} else if (linput[lidx] <= rinput[ridx]) {
			sortedInput = append(sortedInput, linput[lidx])
			lidx++
		} else {
			sortedInput = append(sortedInput, rinput[ridx])
			ridx++
			count+=len(linput) - lidx
		}
	}
	return sortedInput, count+lcount+rcount
}

