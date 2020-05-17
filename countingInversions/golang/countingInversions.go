package countingInversion

//import "fmt"

func CountInversions(input []uint) int {
	if(input == nil || len(input) <= 1) {
		return 0
	}

	count, _ := countWithMergeSort(input)
	return count
}

func countWithMergeSort(input []uint) (int, []uint) {
	if len(input) == 1 {
		return 0, input
	} else {
		lCount, lSlice := countWithMergeSort(input[0: len(input)/2])
		rCount, rSlice := countWithMergeSort(input[len(input)/2: len(input)])

		i:=0
		j:=0
		count := 0
		var sortedSlice []uint
		for i<len(lSlice) || j<len(rSlice){
			if i == len(lSlice) {
				sortedSlice = append(sortedSlice, rSlice[j])
				j++
			} else if j == len(rSlice) || lSlice[i] <= rSlice[j] {
				sortedSlice = append(sortedSlice, lSlice[i])
				i++
			} else {
				sortedSlice = append(sortedSlice, rSlice[j])
				j++
				count += len(lSlice) -i
			}
		}
		//fmt.Printf("input: %v, l: %v, r: %v, lcount: %d, rcount: %d, count: %d\n", input, lSlice, rSlice, lCount, rCount, count)
		return (count + lCount + rCount), sortedSlice
	}
}