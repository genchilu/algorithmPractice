package findFirstAndLastPositionOfElementInSortedArray

//import "fmt"

const maxNumsLen = 100000
const max = 1000000000
const min = -1000000000

func searchRange(nums []int, target int) []int {
	if (nums == nil || len(nums) == 0 || len(nums) > maxNumsLen) {
		return []int{-1, -1}
	}

	if (target > max || target < min) {
		return []int{-1, -1}
	}

	begin := -1
	end := -1

	for i, v := range(nums) {
		if (v == target) {
			begin = i
			end = i
			break
		}
	}

	if(begin == -1) {
		return []int{-1, -1}
	}

	for ;end<len(nums); end++ {
		if(nums[end] != target) {
			break
		} 
	}
	end--

    return []int{begin, end}
}