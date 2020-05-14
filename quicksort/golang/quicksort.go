package quicksort

//import "fmt"

func Sort(input []uint){
	if input != nil {
		quick(input, 0, len(input)-1)
	}
} 

func quick(input []uint, begin int, end int) {
	//fmt.Printf("before: %v, begin: %d, end: %d\n", input, begin, end);
	if end > begin {
		pivotIdx := begin
		for i:= begin+1; i<=end; i++ {
			if input[i] < input[begin] {
				pivotIdx++
				swap(input, pivotIdx, i)
			}
		}

		swap(input, begin, pivotIdx)
		//fmt.Printf("after: %v, begin: %d, end: %d, pivotIdx: %d\n", input, begin, end, pivotIdx);
		quick(input, begin, pivotIdx)
		quick(input, pivotIdx+1, end)
	}
}

func swap(input []uint, a, b int) {
	tmp := input[a]
	input[a] = input[b]
	input[b] = tmp
}