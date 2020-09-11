package quicksort

//import "fmt"

func Sort(input []int){
	if input != nil && len(input) !=0 {
		doQuockSort(input, 0, len(input))
	}
}

func doQuockSort(input []int, begin, end int) {

	if(end-begin) > 0 {
		pivotIdx := begin
		swap(input, pivotIdx, begin)

		pivotIdx = begin

		for i:=pivotIdx+1;i<end;i++ {
			if (input[i] < input[begin]) {
				swap(input, i, pivotIdx)
			}
		}
		swap(input, begin, pivotIdx)

		doQuockSort(input, begin, pivotIdx)
		doQuockSort(input, pivotIdx+1, end)
	}
}

func swap(input []int, from, to int) {
	tmp := input[from]
	input[from] = input[to]
	input[to] = tmp
}



