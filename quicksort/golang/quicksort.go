package quicksort

//import "fmt"

func Sort(input []int){
	if (input != nil || len(input)> 0) {
		quick(input, 0, len(input))
	}
}

func quick(input []int, begin, end int) {
	if ((end-begin) > 1 ) {
		pivotIdx := begin;

		swap(input, pivotIdx, begin)
		pivotIdx = begin

		for i:=(begin+1);i<end;i++{
			if(input[i] < input[begin]) {
				pivotIdx++
				swap(input, pivotIdx, i)
			}
		}
		swap(input, pivotIdx, begin)

		quick(input, begin, pivotIdx)
		quick(input, pivotIdx+1, end)
	}
}

func swap(input []int, from, to int) {
	if (from != to) {
		tmp := input[from]
		input[from] = input[to]
		input[to] = tmp
	}
} 