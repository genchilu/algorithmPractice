package quicksort

//import "fmt"

func Sort(input []int){
	if input != nil && len(input) >0 {
		quick(input, 0, len(input))
	}
} 

func quick(input []int, begin int, end int) {
	if (end - begin) > 1 {
		finalPivotIdx := begin

		// 這邊可以自由更換選 pivot index 的方法
		curPivotIdx := begin

		// 強制把 pivot 放到 index begin
		swap(input, curPivotIdx, begin)

		for i:=(begin+1);i<end;i++ {
			if input[i] < input[begin] {
				swap(input, i, finalPivotIdx+1)
				finalPivotIdx++
			}
		}
		swap(input, begin, finalPivotIdx)

		quick(input, begin, finalPivotIdx)
		quick(input, finalPivotIdx+1, end)
	}
}

func swap(input []int, fromIdx int, toIdx int) {
	if fromIdx != toIdx {
		tmp := input[fromIdx]
		input[fromIdx] = input[toIdx]
		input[toIdx] = tmp
	}
}