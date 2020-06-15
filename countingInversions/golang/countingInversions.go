package countingInversion

//import "fmt"

func CountInversions(input []uint) int {

	if (input == nil) {
		return 0
	}

	count, _ := counting(input)
	return count

}

func counting(input []uint) (int, []uint){
	if (len(input) <=1) {
		return 0, input
	}

	mid := len(input)/2
	lcount, linput := counting(input[0:mid])
	rcount, rinput := counting(input[mid:])

	count :=0
	mergeResult := []uint{}
	lidx := 0
	ridx := 0

	for lidx<len(linput) || ridx<len(rinput) {
		if (lidx == len(linput)) {
			mergeResult = append(mergeResult, rinput[ridx:]...)
			ridx = len(rinput)
		} else if (ridx == len(rinput)) {
			mergeResult = append(mergeResult, linput[lidx:]...)
			lidx = len(linput)
		} else if (linput[lidx] <= rinput[ridx]) {
			mergeResult = append(mergeResult, linput[lidx])
			lidx++
		} else {
			mergeResult = append(mergeResult, rinput[ridx])
			ridx++
			count+= len(linput) - lidx
		}
	}

	return (count+lcount+rcount),  mergeResult
}
